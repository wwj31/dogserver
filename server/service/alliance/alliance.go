package alliance

import (
	"context"
	"fmt"
	"github.com/wwj31/dogactor/actor"
	"go.mongodb.org/mongo-driver/bson"
	"server/common/log"
	"server/common/mongodb"
)

func New(id int32) *Alliance {
	return &Alliance{allianceId: id}
}

type (
	Alliance struct {
		actor.Base
		allianceId int32
		members    map[string]*Member
		masterRID  string
	}
)

func (s *Alliance) OnInit() {
	mongoDBName := fmt.Sprintf("alliance_%v", s.allianceId)
	cur, err := mongodb.Ins.Collection(mongoDBName).Find(context.Background(), bson.M{})
	if err != nil {
		log.Errorw("load all alliance member failed", "err", err)
		return
	}

	var members []*Member
	err = cur.All(context.Background(), &members)
	if err != nil {
		log.Errorw("decode all member failed", "err", err)
		return
	}

	if len(members) == 0 {
		log.Warnw("alliance has no member")
		return
	}

	for _, member := range members {
		s.members[member.RID] = member
		log.Debugf("load member %+v", *member)
	}

	log.Debugf("Alliance OnInit %v members:%v", s.ID(), len(s.members))
}

func (s *Alliance) OnStop() bool {
	log.Infof("stop Alliance %v", s.ID())
	return true
}

func (s *Alliance) OnHandle(msg actor.Message) {

}
