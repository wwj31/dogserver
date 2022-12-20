package common

import "server/common/log"

type Observer struct {
	watcher map[string][]func(interface{})
}

func NewObserver() *Observer {
	return &Observer{
		watcher: map[string][]func(interface{}){},
	}
}

func WatchEvent[T any](observer *Observer, fn func(ev T)) {
	if observer == nil {
		log.Errorw("observer is nil")
		return
	}

	var t T
	evName := ProtoType(t)
	if _, ok := observer.watcher[evName]; !ok {
		observer.watcher[evName] = []func(interface{}){}
	}

	arr := observer.watcher[evName]
	arr = append(arr, func(v interface{}) {
		fn(v.(T))
	})
	observer.watcher[evName] = arr
}

func EmitEvent(observer *Observer, v interface{}) {
	if observer == nil {
		log.Errorw("observer is nil")
		return
	}
	evName := ProtoType(v)
	arr := observer.watcher[evName]
	for _, fn := range arr {
		fn(v)
	}
}
