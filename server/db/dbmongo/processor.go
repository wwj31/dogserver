package dbmongo

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/common/log"
	"server/common/mongodb"
)

var (
	processors sync.Map
	ctx        context.Context
	cancel     context.CancelFunc
	waitGroup  sync.WaitGroup
)

type (
	processor struct {
		fn chan func()

		collection string
		queue      []*data
	}

	data struct {
		key      string
		document interface{}
	}
)

func init() {
	ctx, cancel = context.WithCancel(context.Background())
}

func Stop() {
	cancel()
	waitGroup.Wait()
}

func Store(collection, key string, doc interface{}) {
	v, ok := processors.Load(collection)
	if !ok {
		v, _ = processors.LoadOrStore(collection, newProcessor(collection))
	}

	proc := v.(*processor)

	proc.fn <- func() {
		d := &data{key: key, document: doc}
		proc.queue = append(proc.queue, d)
	}
}

func newProcessor(collection string) *processor {
	proc := &processor{
		fn:         make(chan func()),
		collection: collection,
		queue:      make([]*data, 0),
	}

	waitGroup.Add(1)
	tick := time.Tick(time.Second)
	go func() {
		for {
			select {
			case <-ctx.Done():
				proc.update()
				waitGroup.Done()
				return

			case fn := <-proc.fn:
				fn()

			case <-tick:
				proc.update()

			}
		}
	}()

	return proc
}

func (p *processor) update() {
	for _, v := range p.queue {
		_, err := mongodb.Ins.Collection(p.collection).UpdateByID(
			context.Background(),
			v.key,
			bson.M{"$set": v.document},
			options.Update().SetUpsert(true),
		)

		if err != nil {
			log.Errorf("mongo update failed", "collection", p.collection, "key", v.key, "err", err)
			return
		}
	}

	p.queue = p.queue[len(p.queue):]
}
