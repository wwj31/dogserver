package handler

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/log"
	"server/common"
	"server/proto/inner_message/inner"
	"server/proto/message"
	"server/service/game/logic/player"
)

// 玩家请求进入游戏
func (s *Controller) EnterGameReq(sourceId string, gSession common.GSession, pbMsg interface{}) proto.Message {
	msg := pbMsg.(*message.EnterGameReq)
	log.KV("msg", msg).Debug("EnterGameReq")

	// todo .. 解密

	_player, exist := s.PlayerMgr().PlayerByRID(msg.RID)
	if !exist {
		_player = player.New(msg.RID, s)
	}
	_player.SetGateSession(gSession)
	s.PlayerMgr().SetPlayer(_player)

	// 新号处理
	if _player.IsNewRole() {
		// todo ...
		_player.Item().Add(map[int64]int64{123: 999})
	}

	_player.Login()
	log.KVs(log.Fields{"roleId": _player.Role().RoleId(), "gSession": gSession}).Info("player login")
	return &message.EnterGameRsp{}
}

// 玩家离线
func (s *Controller) Logout(sourceId string, gSession common.GSession, pbMsg interface{}) proto.Message {
	msg := pbMsg.(*inner.GT2GSessionClosed)
	gs := common.GSession(msg.GateSession)
	p, ok := s.PlayerMgr().PlayerBySession(gs)
	if ok {
		p.Logout()
	}

	s.PlayerMgr().OfflinePlayer(gs)
	log.KVs(log.Fields{"roleId": p.Role().RoleId(), "gSession": msg.GateSession}).Info("player logout")
	return nil
}
