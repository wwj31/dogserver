package mgo

import (
	"context"
	"math/rand"
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
		queue      []docData
	}

	docData struct {
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
		data := docData{key: key, document: doc}
		if !replace(proc.queue, data) {
			proc.queue = append(proc.queue, data)
		}
	}
}

func newProcessor(collection string) *processor {
	proc := &processor{
		fn:         make(chan func()),
		collection: collection,
		queue:      make([]docData, 0),
	}

	waitGroup.Add(1)
	randDur := func() time.Duration {
		return time.Duration(rand.Intn(int(30*time.Second))) + (30 * time.Second)
	}
	t := time.NewTimer(randDur())
	go func() {
		for {
			select {
			case <-ctx.Done():
				ok := make(chan struct{})
				proc.update()
				<-ok
				waitGroup.Done()
				return

			case fn := <-proc.fn:
				fn()

			case <-t.C:
				proc.update()
				t.Reset(randDur())
			}
		}
	}()

	return proc
}

func replace(queue []docData, new docData) bool {
	for i, v := range queue {
		if v.key == new.key {
			queue[i] = new
			return true
		}
	}
	return false
}

func (p *processor) update() {
	arr := p.queue
	p.queue = nil
	for _, v := range arr {
		func() {
			defer func() {
				if delta := time.Now().Sub(time.Now()); delta > time.Millisecond*100 {
					log.Warnw("mongo update exec too long", "delta", delta)
				}
			}()

			_, err := mongodb.Ins.Collection(p.collection).UpdateByID(
				context.Background(),
				v.key,
				bson.M{"$set": v.document},
				options.Update().SetUpsert(true),
			)

			if err != nil {
				log.Errorw("mongo update failed", "collection", p.collection, "key", v.key, "err", err)
				return
			}
		}()
	}
}
