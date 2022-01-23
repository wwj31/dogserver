package game

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"server/common"
	"server/common/log"
	"server/common/toml"
	"server/db"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
	"server/service/game/logic/player"
	"server/service/game/logic/player/localmsg"
)

func New(serverId uint16) *Game {
	return &Game{sid: serverId}
}

type Game struct {
	actor.Base
	sid     uint16 // serverId
	genUUID common.UID
	iface.SaveLoader
	playerMgr       iface.PlayerManager
	inactivePlayers map[common.ActorId]struct{}
}

func (s *Game) OnInit() {
	s.SaveLoader = db.New(toml.Get("mysql"), toml.Get("database"))
	s.genUUID = common.NewUID(s.sid)
	s.playerMgr = player.NewMgr(s)
	s.inactivePlayers = make(map[common.ActorId]struct{}, 1000)

	log.Debugf("game OnInit")
}

func (s *Game) OnStop() bool {
	log.Infof("stop game")
	return true
}

func (s *Game) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	actMsg, _, gSession, err := common.UnwrapperGateMsg(msg)
	expect.Nil(err)

	switch pbMsg := actMsg.(type) {
	case *outer.EnterGameReq:
		s.enterGameReq(gSession, pbMsg)
	case *inner.GT2GSessionClosed:
		s.logout(pbMsg)
	default:
		s.toPlayer(gSession, actMsg)
	}
}

func (s *Game) OnHandleEvent(event interface{}) {
	switch ev := event.(type) {
	case *actor.EvDelactor:
		if common.IsActorOf(ev.ActorId, common.Player_Actor) {
			s.inactivePlayers[ev.ActorId] = struct{}{}
		}
	}
}

// SID serverId
func (s *Game) SID() uint16 {
	return s.sid
}

func (s *Game) GenUuid() uint64 {
	return s.genUUID.GenUuid()
}

func (s *Game) PlayerMgr() iface.PlayerManager {
	return s.playerMgr
}

func (s *Game) activatePlayer(rid uint64) common.ActorId {
	playerId := common.PlayerId(rid)
	playerActor := actor.New(playerId, player.New(rid, s), actor.SetMailBoxSize(200), actor.SetLocalized())
	err := s.System().Add(playerActor)
	expect.Nil(err)

	delete(s.inactivePlayers, playerId)
	return playerId
}

// player enter game
func (s *Game) enterGameReq(gSession common.GSession, msg *outer.EnterGameReq) {
	log.Debugw("EnterGameReq", "msg", msg)

	// warn:repeated login
	if _, ok := s.PlayerMgr().PlayerBySession(gSession); ok {
		log.Warnw("player repeated enter game", "gSession", gSession, "localmsg", msg.RID)
		return
	}

	// todo .. decrypt
	var playerId = common.PlayerId(msg.RID)

	if oldSession, ok := s.PlayerMgr().GSessionByPlayer(playerId); ok {
		s.PlayerMgr().DelGSession(oldSession)
	} else {
		playerId = s.activatePlayer(msg.RID)
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
	var (
		actorId common.ActorId
		ok      bool
	)
	if gSession != "" {
		actorId, ok = s.PlayerMgr().PlayerBySession(gSession)
		if !ok {
			log.Warnw("msg to player,but can not found player by gSession", "gSession", gSession)
			return
		}
	} else {

	}

	err := s.Send(actorId, msg)
	expect.Nil(err)
}
