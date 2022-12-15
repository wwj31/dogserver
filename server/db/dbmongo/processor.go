package dbmongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/common/log"
	"server/common/mongodb"
	"sync"
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
		kv         map[string]*data
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
		proc.queue = append(proc.queue, &data{
			key:      key,
			document: doc,
		})
	}
}

func newProcessor(collection string) *processor {
	proc := &processor{
		fn:         make(chan func()),
		collection: collection,
		queue:      make([]*data, 0),
		kv:         make(map[string]*data),
	}

	waitGroup.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				waitGroup.Done()
				return

			case fn := <-proc.fn:
				fn()
				proc.update()
				log.Infof("update")
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
}
