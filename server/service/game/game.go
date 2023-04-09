package game

import (
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/actor/event"
	"github.com/wwj31/dogactor/expect"

	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
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
	s.System().OnEvent(s.ID(), func(ev event.EvActorSubMqFin) {
		if actortype.IsActorOf(ev.ActorId, actortype.PlayerActor) {
			if respId, ok := s.respIdMap[ev.ActorId]; ok {
				_ = s.Response(respId, &outer.Ok{})
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
	actMsg, _, _, err := common.UnwrappedGateMsg(msg.RawMsg())
	expect.Nil(err)

	switch pbMsg := actMsg.(type) {
	case *inner.PullPlayer:
		playerId, loading := s.checkAndPullPlayer(pbMsg.RID)
		if !loading {
			_ = s.Response(msg.GetRequestId(), &outer.Ok{})
		} else {
			s.respIdMap[playerId] = msg.GetRequestId()
		}
	default:
		log.Warnw("unknown msg:%v", msg.String())
	}
}

func (s *Game) checkAndPullPlayer(rid string) (playerId actortype.ActorId, loading bool) {
	// TODO::检查玩家是否在其他game节点中,并且通知目标下线,需要将玩家所在节点数据存入redis中以便查询
	playerId = actortype.PlayerId(rid)
	if !s.System().HasActor(playerId) {
		err := s.System().NewActor(
			playerId,
			player.New(rid, s),
			actor.SetMailBoxSize(200),
			//actor.SetLocalized(),
		)
		expect.Nil(err)
		return playerId, true
	}

	return playerId, false
}
