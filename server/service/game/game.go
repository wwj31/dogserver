package game

import (
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/actor/event"
	"github.com/wwj31/dogactor/expect"

	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/service/game/logic/player"
)

func New(serverId int32) *Game {
	return &Game{sid: serverId}
}

type Game struct {
	actor.Base
	sid       int32 // serverId
	respIdMap map[actor.Id]string
}

func (s *Game) OnInit() {
	s.respIdMap = make(map[actor.Id]string)
	s.System().OnEvent(s.ID(), func(ev event.EvNewActor) {
		if actortype.IsActorOf(ev.ActorId, actortype.PlayerActor) {
			if respId, ok := s.respIdMap[ev.ActorId]; ok {
				//log.Debugf("the player actor startup ok,response to login %v", respId)
				_ = s.Response(respId, &inner.Ok{})
				delete(s.respIdMap, ev.ActorId)
			}
		}
	})

	log.Infow("game OnInit")
}

func (s *Game) OnStop() bool {
	s.System().CancelAll(s.ID())
	log.Infof("stop game")
	return true
}

// SID serverId
func (s *Game) SID() int32 {
	return s.sid
}

func (s *Game) OnHandle(msg actor.Message) {
	actMsg, _, _, err := common.UnwrappedGateMsg(msg.Payload())
	expect.Nil(err)

	switch pbMsg := actMsg.(type) {
	case *inner.PullPlayer:
		//log.Debugf("pull player %v ", pbMsg.RoleInfo.RID)
		playerId, loading := s.checkAndPullPlayer(pbMsg.Account, pbMsg.RoleInfo)
		if !loading {
			_ = s.Response(msg.GetRequestId(), &inner.Ok{})
		} else {
			s.respIdMap[playerId] = msg.GetRequestId()
		}
	default:
		log.Warnw("unknown msg:%v", msg.String())
	}
}

func (s *Game) checkAndPullPlayer(acc *inner.Account, roleInfo *inner.LoginRoleInfo) (playerId actortype.ActorId, loading bool) {
	playerId = actortype.PlayerId(roleInfo.RID)
	if !s.System().HasActor(playerId) {
		err := s.System().NewActor(
			playerId,
			player.New(acc, roleInfo, s),
			actor.SetMailBoxSize(300),
			//actor.SetLocalized(),
		)
		expect.Nil(err)
		return playerId, true
	}

	return playerId, false
}
