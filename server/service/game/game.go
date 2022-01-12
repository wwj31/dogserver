package game

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/actor/actorerr"
	"github.com/wwj31/dogactor/expect"
	"server/common"
	"server/common/log"
	"server/common/toml"
	"server/db"
	"server/proto/inner/inner"
	"server/proto/message"
	"server/service/game/iface"
	"server/service/game/logic/player"
	"server/service/game/logic/player/localmsg"
)

func New(serverId uint16) *Game {
	return &Game{sid: serverId}
}

type Game struct {
	actor.Base
	sid uint16 // serverId
	common.UID
	iface.SaveLoader
	playerMgr iface.PlayerManager
}

func (s *Game) OnInit() {
	s.SaveLoader = db.New(toml.Get("mysql"), toml.Get("database"))
	s.UID = common.NewUID(s.sid)
	s.playerMgr = player.NewMgr(s)

	log.Debugf("game OnInit")
}

func (s *Game) OnStop() bool {
	log.Infof("stop game")
	return true
}

func (s *Game) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	actMsg, gSession, err := common.UnwrapperGateMsg(msg)
	expect.Nil(err)

	switch pbMsg := actMsg.(type) {
	case *message.EnterGameReq:
		s.enterGameReq(gSession, pbMsg)
	case *inner.GT2GSessionClosed:
		s.logout(pbMsg)
	default:
		s.toPlayer(gSession, actMsg)
	}
}

// SID serverId
func (s *Game) SID() uint16 {
	return s.sid
}

func (s *Game) GenUuid() uint64 {
	return s.UID.GenUuid()
}

func (s *Game) PlayerMgr() iface.PlayerManager {
	return s.playerMgr
}

// player enter game
func (s *Game) enterGameReq(gSession common.GSession, msg *message.EnterGameReq) {
	log.Debugw("EnterGameReq", "msg", msg)

	// warn:repeated login
	if _, ok := s.PlayerMgr().PlayerBySession(gSession); ok {
		log.Warnw("player repeated enter game", "gSession", gSession, "localmsg", msg.RID)
		return
	}

	// todo .. decrypt
	playerId := common.PlayerId(msg.RID)

	if oldSession, ok := s.PlayerMgr().GSessionByPlayer(playerId); ok {
		s.PlayerMgr().DelGSession(oldSession)
	} else {
		playerActor := actor.New(playerId, player.New(msg.RID, s), actor.SetMailBoxSize(200))
		err := s.System().Add(playerActor)
		if err != nil {
			if !errors.Is(err, actorerr.RegisterActorSameIdErr) {
				log.Errorw("add player actor error", "rid", msg.RID, "err", err)
				return
			} else {
				log.Infow("login activite player actor", "player", playerId)
			}
		}
	}

	s.PlayerMgr().SetPlayer(gSession, playerId)

	err := s.Send(playerId, localmsg.Login{GSession: gSession})
	if err != nil {
		log.Errorw("login send error", "rid", msg.RID, "err", err, "playerId", playerId)
		return
	}

}

// player offline
func (s *Game) logout(msg *inner.GT2GSessionClosed) proto.Message {
	gSession := common.GSession(msg.GateSession)
	playerId, ok := s.PlayerMgr().PlayerBySession(gSession)
	if ok {
		_ = s.Send(playerId, msg)
	}

	s.PlayerMgr().DelGSession(gSession)
	return nil
}

func (s *Game) toPlayer(gSession common.GSession, msg interface{}) {
	actorId, ok := s.PlayerMgr().PlayerBySession(gSession)
	if !ok {
		log.Warnw("msg to player,but can not found player by gSession", "gSession", gSession)
		return
	}
	err := s.Send(actorId, msg)
	expect.Nil(err)
}
