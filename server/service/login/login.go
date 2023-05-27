package login

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/actor/event"
	"github.com/wwj31/dogactor/expect"

	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/rds"
	"server/proto/outermsg/outer"
	"server/service/login/account"
)

const GetAndPopRandInt = `
local result = redis.call('SRANDMEMBER', KEYS[1], 1)  -- 从集合中随机取一个值
if result then
    redis.call('SREM', KEYS[1], result[1])  -- 从集合中删除该值
end
return result
`

type Login struct {
	actor.Base
	sha1         string
	allGameNode  map[string]struct{}
	nextGameNode atomic.Int64
}

func New() *Login {
	return &Login{}
}

func (s *Login) OnInit() {
	log.Infow("login OnInit")
	account.CreateIndex()
	s.sha1 = rds.Ins.ScriptLoad(context.Background(), GetAndPopRandInt).Val()
	s.allGameNode = make(map[string]struct{})

	s.System().OnEvent(s.ID(), func(ev event.EvNewActor) {
		// 维护可用的所有game节点
		log.Infow("new node ", "actor", ev.ActorId)

		if actortype.IsActorOf(ev.ActorId, actortype.GameActor) {
			s.allGameNode[ev.ActorId] = struct{}{}
		}
	})

	s.System().OnEvent(s.ID(), func(ev event.EvDelActor) {
		if actortype.IsActorOf(ev.ActorId, actortype.GameActor) {
			delete(s.allGameNode, ev.ActorId)
		}
	})
}

func (s *Login) OnStop() bool {
	log.Debugw("login stop", "id", s.ID())
	time.Sleep(2 * time.Second)
	return true
}

func (s *Login) OnHandle(m actor.Message) {
	payload := m.Payload()
	v, _, gSession, err := common.UnwrappedGateMsg(payload)

	expect.Nil(err)
	switch msg := v.(type) {
	case *outer.LoginReq:
		err = s.LoginReq(m.GetSourceId(), gSession, msg)
	default:
		err = fmt.Errorf("undefined localmsg type %v", msg)
	}

	if err != nil {
		log.Errorw("handle outer error", "err", err)
	}
}

func (s *Login) LoginReq(sourceId string, gSession common.GSession, msg *outer.LoginReq) error {
	s.Login(gSession, msg)
	return nil
}
