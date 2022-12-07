package game

import (
	"errors"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/actor/actorerr"
	"github.com/wwj31/dogactor/expect"
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/toml"
	"server/db/dbmysql"
	"server/proto/innermsg/inner"
	"server/service/game/iface"
	"server/service/game/logic/player"
)

func New(serverId uint16) *Game {
	return &Game{sid: serverId}
}

type Game struct {
	actor.Base
	iface.StoreLoader

	sid       uint16 // serverId
	onlineMgr onlineMgr
}

func (s *Game) OnInit() {
	s.StoreLoader = dbmysql.New(toml.Get("mysql"), toml.Get("database"), s.System())
	s.onlineMgr = newMgr(s)

	s.System().OnEvent(s.ID(), func(event actor.EvDelActor) {
		if actortype.IsActorOf(event.ActorId, actortype.Player_Actor) {
		}
	})
	log.Debugf("game OnInit")
}

func (s *Game) OnStop() bool {
	s.System().CancelAll(s.ID())
	log.Infof("stop game")
	return true
}

// SID serverId
func (s *Game) SID() uint16 {
	return s.sid
}

func (s *Game) OnHandle(msg actor.Message) {
	actMsg, _, _, err := common.UnwrapperGateMsg(msg.RawMsg())
	expect.Nil(err)

	switch pbMsg := actMsg.(type) {
	case *inner.PullPlayer:
		s.checkAndPullPlayer(pbMsg.RID)
	default:
		log.Warnw("unknown msg:%v", msg.String())
	}
}

func (s *Game) checkAndPullPlayer(rid string) actortype.ActorId {
	playerId := actortype.PlayerId(rid)
	if act := s.System().LocalActor(playerId); act == nil {
		playerActor := actor.New(
			playerId,
			player.New(rid, s),
			actor.SetMailBoxSize(200),
			//actor.SetLocalized(),
		)

		err := s.System().Add(playerActor)
		if errors.Is(err, actorerr.RegisterActorSameIdErr) {
			log.Errorw("actor add err", "err", err)
			return playerId
		}
		expect.Nil(err)
	}

	return playerId
}
