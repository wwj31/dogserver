package game

import (
	"server/common"
	"server/common/log"
	"server/proto/inner_message/inner"
	"server/proto/message"
	"server/service/game/iface"
	"server/service/game/logic/player"
	"server/service/game/logic/player/localmsg"

	"github.com/wwj31/dogactor/expect"

	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
)

func New(s iface.SaveLoader) *Game {
	return &Game{
		SaveLoader: s,
	}
}

type Game struct {
	actor.Base
	common.SendTools
	iface.SaveLoader

	sid int32 // serverId

	playerMgr iface.PlayerManager
}

func (s *Game) OnInit() {
	s.SendTools = common.NewSendTools(s)
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
	}
}

// SID serverId
func (s *Game) SID() int32 {
	return s.sid
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
		playerActor := actor.New(playerId, player.New(msg.RID), actor.SetMailBoxSize(200))
		err := s.System().Add(playerActor)
		if err != nil {
			log.Errorw("add player actor error", "rid", msg.RID, "err", err)
			return
		}
	}

	err := s.Send(playerId, localmsg.Login{GSession: gSession})
	if err != nil {
		log.Errorw("login send error", "rid", msg.RID, "err", err, "playerId", playerId)
		return
	}

	s.PlayerMgr().SetPlayer(gSession, playerId)
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
