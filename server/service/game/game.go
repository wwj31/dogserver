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
	"time"

	"github.com/wwj31/dogactor/tools"

	gogo "github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
)

func New(serverId uint16) *Game {
	return &Game{sid: serverId}
}

type Game struct {
	actor.Base
	common.UID
	iface.StoreLoader

	sid       uint16 // serverId
	onlineMgr onlineMgr
}

func (s *Game) OnInit() {
	s.StoreLoader = db.New(toml.Get("mysql"), toml.Get("database"))
	s.UID = common.NewUID(s.sid)
	s.onlineMgr = newMgr(s)

	_ = s.System().RegistEvent(s.ID(), actor.EvDelactor{})
	log.Debugf("game OnInit")
}

func (s *Game) OnStop() bool {
	s.System().CancelAll(s.ID())
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

// MsgToPlayer send msg to player actor
func (s *Game) MsgToPlayer(rid uint64, sid uint16, msg gogo.Message) {
	actorId := common.PlayerId(rid)
	gSession, ok := s.onlineMgr.GSessionByPlayer(actorId)
	if ok {
		s.toPlayer(gSession, msg)
		return
	}

	bytes := common.ProtoMarshal(msg)
	wrapper := &inner.GameMsgWrapper{
		RID:     rid,
		MsgName: common.ProtoType(msg),
		Data:    bytes,
	}

	if s.SID() != sid {
		err := s.Send(common.GameName(int32(sid)), wrapper)
		if err != nil {
			log.Errorw("msg to player from other game send failed", "err", err)
			return
		}
	}

	s.toPlayer("", wrapper)
}

func (s *Game) activatePlayer(rid uint64, firstLogin bool) common.ActorId {
	playerId := common.PlayerId(rid)
	if ok := s.System().Exist(playerId); !ok || firstLogin {
		playerActor := actor.New(playerId, player.New(rid, s, firstLogin), actor.SetMailBoxSize(200), actor.SetLocalized())
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
	if _, ok := s.onlineMgr.PlayerBySession(gSession); ok {
		log.Warnw("player repeated enter game", "gSession", gSession, "localmsg", msg.RID)
		return
	}

	// todo .. decrypt
	var playerId = common.PlayerId(msg.RID)

	if oldSession, ok := s.onlineMgr.GSessionByPlayer(playerId); ok {
		s.onlineMgr.DelGSession(oldSession)
	} else {
		playerId = s.activatePlayer(msg.RID, msg.NewPlayer)
	}
	s.onlineMgr.AssociateSession(playerId, gSession)

	err := s.Send(playerId, localmsg.Login{GSession: gSession})
	if err != nil {
		log.Errorw("login send error", "rid", msg.RID, "err", err, "playerId", playerId)
		return
	}
	s.AddTimer("", tools.NowTime()+int64(10*time.Second), func(dt int64) {
		s.MsgToPlayer(msg.RID, s.sid, &outer.UseItemReq{Items: map[int64]int64{123: 1}})
	}, -1)
}

// player offline
func (s *Game) logout(msg *inner.GT2GSessionClosed) gogo.Message {
	gSession := common.GSession(msg.GateSession)
	playerId, ok := s.onlineMgr.PlayerBySession(gSession)
	if ok {
		_ = s.Send(playerId, msg)
	}

	s.onlineMgr.DelGSession(gSession)
	return nil
}

func (s *Game) toPlayer(gSession common.GSession, msg interface{}) {
	var (
		actorId common.ActorId
		ok      bool
	)

	// gSession != "" mean msg from client via gateway,otherwise
	// msg from other actor wrap in inner.GameMsgWrapper
	if gSession != "" {
		actorId, ok = s.onlineMgr.PlayerBySession(gSession)
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
		if err := gogo.Unmarshal(gameWrapper.Data, v.(gogo.Message)); err != nil {
			log.Errorw("unmarshal failed", "msgName", gameWrapper.MsgName)
			return
		}
		msg = v
		actorId = common.PlayerId(gameWrapper.RID)
		// try reactivate player actor if actor has exited
		s.activatePlayer(gameWrapper.RID, false)
	}

	err := s.Send(actorId, msg)
	expect.Nil(err)
}
