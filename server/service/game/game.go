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
				log.Debugf("the new player response to the login %v", respId)
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
		log.Debugf("pull player %v ", pbMsg.RID)
		playerId, loading := s.checkAndPullPlayer(pbMsg.RID, pbMsg.ShortId)
		if !loading {
			_ = s.Response(msg.GetRequestId(), &inner.Ok{})
		} else {
			s.respIdMap[playerId] = msg.GetRequestId()
		}
	default:
		log.Warnw("unknown msg:%v", msg.String())
	}
}

func (s *Game) checkAndPullPlayer(rid string, shortId int64) (playerId actortype.ActorId, loading bool) {
	// TODO::检查玩家是否在其他game节点中,并且通知目标下线,需要将玩家所在节点数据存入redis中以便查询
	playerId = actortype.PlayerId(rid)
	if !s.System().HasActor(playerId) {
		err := s.System().NewActor(
			playerId,
			player.New(rid, shortId, s),
			actor.SetMailBoxSize(200),
			//actor.SetLocalized(),
		)
		expect.Nil(err)
		return playerId, true
	}

	return playerId, false
}
