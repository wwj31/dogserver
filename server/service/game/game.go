package game

import (
	"errors"
	"github.com/wwj31/dogactor/actor/actorerr"
	"reflect"
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/toml"
	"server/db"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
	"server/service/game/logic/player"
	"server/service/game/logic/player/localmsg"

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
	s.StoreLoader = db.New(toml.Get("mysql"), toml.Get("database"), s.System())
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
		if actortype.IsActorOf(ev.ActorId, actortype.Player_Actor) {
		}
	}
}

// SID serverId
func (s *Game) SID() uint16 {
	return s.sid
}

// MsgToPlayer send msg to player actor
func (s *Game) MsgToPlayer(rid uint64, sid uint16, msg gogo.Message) {
	actorId := actortype.PlayerId(rid)
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
		err := s.Send(actortype.GameName(int32(sid)), wrapper)
		if err != nil {
			log.Errorw("msg to player from other game send failed", "err", err)
			return
		}
	}

	s.toPlayer("", wrapper)
}

func (s *Game) checkAndActivatePlayer(rid uint64, firstLogin bool) actortype.ActorId {
	playerId := actortype.PlayerId(rid)
	if ok := s.System().Exist(playerId); !ok || firstLogin {
		playerActor := actor.New(playerId, player.New(rid, s, firstLogin), actor.SetMailBoxSize(200), actor.SetLocalized())
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
		playerId = s.checkAndActivatePlayer(msg.RID, msg.NewPlayer)
	}
	s.onlineMgr.AssociateSession(playerId, gSession)

	err := s.Send(playerId, localmsg.Login{GSession: gSession, RId: msg.RID, UId: msg.UID})
	if err != nil {
		log.Errorw("login send error", "rid", msg.RID, "err", err, "playerId", playerId)
		return
	}
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
		actorId actortype.ActorId
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
		actorId = actortype.PlayerId(gameWrapper.RID)
		// try reactivate player actor if actor has exited
		s.checkAndActivatePlayer(gameWrapper.RID, false)
	}

	err := s.Send(actorId, msg)
	expect.Nil(err)
}
