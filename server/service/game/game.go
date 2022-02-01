package game

import (
	"reflect"
	"server/common"
	"server/common/log"
	"server/common/toml"
	"server/db"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
	"server/service/game/logic/player"
	"server/service/game/logic/player/localmsg"

	"github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
)

func New(serverId uint16) *Game {
	return &Game{sid: serverId}
}

type Game struct {
	actor.Base
	sid     uint16 // serverId
	genUUID common.UID
	iface.SaveLoader
	playerMgr iface.PlayerManager
}

func (s *Game) OnInit() {
	s.SaveLoader = db.New(toml.Get("mysql"), toml.Get("database"))
	s.genUUID = common.NewUID(s.sid)
	s.playerMgr = player.NewMgr(s)

	_ = s.System().RegistEvent(s.ID(), actor.EvDelactor{})
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
	case actor.EvDelactor:
		if common.IsActorOf(ev.ActorId, common.Player_Actor) {
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

func (s *Game) activatePlayer(rid uint64, new bool) common.ActorId {
	playerId := common.PlayerId(rid)
	if ok := s.System().Exist(playerId); !ok || new {
		playerActor := actor.New(playerId, player.New(rid, s), actor.SetMailBoxSize(200), actor.SetLocalized())
		err := s.System().Add(playerActor)
		expect.Nil(err)
	}

	return playerId
}

// player enter game
func (s *Game) enterGameReq(gSession common.GSession, msg *outer.EnterGameReq) {
	log.Debugw("EnterGameReq", "msg", msg)
	if common.LoginMD5(msg.UID, msg.RID, msg.NewPlayer) != msg.Checksum {
		log.Warnw("checksum md5 check faild", "msg", msg.String())
		return
	}

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
		playerId = s.activatePlayer(msg.RID, msg.NewPlayer)
	}
	s.PlayerMgr().AssociateSession(playerId, gSession)

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
		gameWrapper, ok := msg.(*inner.GameMsgWrapper)
		if !ok {
			log.Errorw("unknown msg type ", "msgType", reflect.TypeOf(msg).String())
			return
		}
		v, exist := inner.Spawner(gameWrapper.MsgName)
		if !exist {
			v, exist = outer.Spawner(gameWrapper.MsgName, true)
			if !exist {
				log.Errorw("msg is not in inner msg", "msgName", gameWrapper.MsgName)
				return
			}
		}
		if err := proto.Unmarshal(gameWrapper.Data, v.(proto.Message)); err != nil {
			log.Errorw("unmarshal failed", "msgName", gameWrapper.MsgName)
			return
		}
		msg = v
		actorId = common.PlayerId(gameWrapper.RID)
		// 如果玩家actor已经关闭，需要重新激活
		s.activatePlayer(gameWrapper.RID, false)
	}

	err := s.Send(actorId, msg)
	expect.Nil(err)
}
