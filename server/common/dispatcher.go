package common

import (
	"server/common/log"

	"github.com/wwj31/dogactor/tools"
)

type EventType = int
type Listener = int
type (
	Handle func(ev interface{})

	Observer struct {
		listeners map[EventType]map[Listener]Handle
	}
)

func New() Observer {
	return Observer{
		listeners: make(map[EventType]map[Listener]Handle, 10),
	}
}

func (s *Observer) Dispatch(typ EventType, event interface{}) {
	list, ok := s.listeners[typ]
	if !ok {
		return
	}

	for _, fun := range list {
		tools.Try(func() {
			fun(event)
		})
	}
}

func (s *Observer) AddListener(obj Listener, typ EventType, callback Handle) {
	if _, ok := s.listeners[typ]; !ok {
		s.listeners[typ] = make(map[Listener]Handle, 1)
	}

	if _, ok := s.listeners[typ][obj]; ok {
		log.Errorw("repeated listen event", "obj", obj, "typ", typ)
		return
	}

	s.listeners[typ][obj] = callback
}

func (s *Observer) RemoveListener(typ EventType, obj Listener) {
	listenerMap, ok := s.listeners[typ]
	if !ok {
		return
	}

	if _, ok := listenerMap[obj]; !ok {
		return
	}

	delete(listenerMap, obj)
}
