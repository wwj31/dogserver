package game

import (
	"server/common"
	"server/proto/inner_message/inner"
	"server/proto/message"
	"server/service/game/iface"
	"server/service/game/logic/player"
	msg2 "server/service/game/logic/player/msg"

	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/log"
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

	sid int32 // 服务器Id

	playerMgr iface.PlayerManager
}

func (s *Game) OnInit() {
	s.SendTools = common.NewSendTools(s)
	s.playerMgr = player.NewMgr(s)

	log.Debug("game OnInit")
}

// 区服id
func (s *Game) OnStop() bool {
	log.Info("stop game")
	return true
}

// 区服id
func (s *Game) SID() int32 {
	return s.sid
}

func (s *Game) PlayerMgr() iface.PlayerManager {
	return s.playerMgr
}

func (s *Game) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	actMsg, gSession, err := common.UnwrapperGateMsg(msg)
	expect.Nil(err)

	switch pbMsg := actMsg.(type) {
	case *message.EnterGameReq:
		s.EnterGameReq(gSession, pbMsg)
	case *inner.GT2GSessionClosed:
		s.Logout(gSession, pbMsg)
	}
}

// 玩家请求进入游戏
func (s *Game) EnterGameReq(gSession common.GSession, msg *message.EnterGameReq) {
	log.KV("msg", msg).Debug("EnterGameReq")

	// 重复登录
	if _, ok := s.PlayerMgr().PlayerBySession(gSession); ok {
		log.KVs(log.Fields{"gSession": gSession, "msg": msg.RID}).Warn("player repeated enter game")
		return
	}

	// todo .. 解密
	playerId := common.PlayerId(msg.RID)
	// 重连删除旧session,否则创建Player Actor
	if oldSession, ok := s.PlayerMgr().GSessionByPlayer(playerId); ok {
		s.PlayerMgr().DelGSession(oldSession)
	} else {
		playerActor := actor.New(playerId, player.New(msg.RID), actor.SetMailBoxSize(500))
		err := s.System().Regist(playerActor)
		if err != nil {
			log.KVs(log.Fields{"rid": msg.RID, "err": err}).Error("regist player actor error")
			return
		}
	}

	err := s.Send(playerId, msg2.MsgLogin{GSession: gSession})
	if err != nil {
		log.KVs(log.Fields{"rid": msg.RID, "err": err, "playerId": playerId}).Error("login send error")
		return
	}

	s.PlayerMgr().SetPlayer(gSession, playerId)
	//
	//if !exist {
	//	_player = player.New(msg.RID, s)
	//}
	//_player.SetGateSession(gSession)
	//s.PlayerMgr().SetPlayer(gSession, playerId)
	//
	//// 新号处理
	//if _player.IsNewRole() {
	//	// todo ...
	//	_player.Item().Add(map[int64]int64{123: 999})
	//}
	//
	//_player.Login()
}

// 玩家离线
func (s *Game) Logout(gs common.GSession, msg *inner.GT2GSessionClosed) proto.Message {
	playerId, ok := s.PlayerMgr().PlayerBySession(gs)
	if ok {
		s.Send(playerId, msg)
	}

	s.PlayerMgr().DelGSession(gs)
	return nil
}
