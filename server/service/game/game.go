package game

import (
	"errors"
	"github.com/wwj31/dogactor/actor/actorerr"
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/toml"
	"server/db/dbmysql"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
	"server/service/game/logic/player"
	"server/service/game/logic/player/localmsg"

	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
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

func (s *Game) OnHandle(msg actor.Message) {
	actMsg, _, gSession, err := common.UnwrapperGateMsg(msg.RawMsg())
	expect.Nil(err)

	switch pbMsg := actMsg.(type) {
	case *outer.EnterGameReq:
		s.enterGameReq(gSession, pbMsg)
	//case *inner.GT2GSessionClosed:
	//s.logout(pbMsg)
	default:
		log.Warnw("unknown msg:%v", msg.String())
		//s.toPlayer(gSession, actMsg)
	}
}

// SID serverId
func (s *Game) SID() uint16 {
	return s.sid
}

func (s *Game) checkAndActivatePlayer(rid string) actortype.ActorId {
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

// player enter game
func (s *Game) enterGameReq(gSession common.GSession, msg *outer.EnterGameReq) {
	log.Debugw("EnterGameReq", "msg", msg)
	// check sign
	if common.LoginMD5(msg.UID, msg.RID, msg.NewPlayer) != msg.Checksum {
		log.Warnw("checksum md5 check faild", "msg", msg.String())
		return
	}

	// warn:repeated login
	if _, ok := s.onlineMgr.PlayerBySession(gSession); ok {
		log.Warnw("player repeated enter game", "gSession", gSession, "localmsg", msg.RID)
		return
	}

	var playerId = actortype.PlayerId(msg.RID)

	if oldSession, ok := s.onlineMgr.GSessionByPlayer(playerId); ok {
		s.onlineMgr.DelGSession(oldSession)
	} else {
		playerId = s.checkAndActivatePlayer(msg.RID)
	}
	s.onlineMgr.AssociateSession(playerId, gSession)

	err := s.Send(playerId, localmsg.Login{
		GSession: gSession,
		RId:      msg.RID,
		UId:      msg.UID,
		First:    msg.NewPlayer,
	})
	if err != nil {
		log.Errorw("login send error", "rid", msg.RID, "err", err, "playerId", playerId)
		return
	}
}
