package login

import (
	"math"
	"time"

	"server/common/actortype"
	"server/common/log"
)

func (s *Login) getGameNode() string {
	actorId := s.getGameActor()
	_, ok := s.allGameNode[actorId]
	for !ok {
		log.Warnw("game node not exist, redo get", "actorId", actorId)
		time.Sleep(time.Second)
		actorId = s.getGameActor()
	}

	return actorId
}

// round-robin
func (s *Login) getGameActor() string {
	s.nextGameNode.CompareAndSwap(math.MaxInt64-1, 1)
	val := s.nextGameNode.Add(1)
	id := val % int64(len(s.allGameNode))
	actorId := actortype.GameName(int32(id + 1))
	return actorId
}
