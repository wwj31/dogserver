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
	// 没有登录过，是新玩家
	if _player.Role().LoginAt() == 0 {
		// todo ...
	}

	_player.SetGateSession(gSession)
	_player.Login()
	s.PlayerMgr().SetPlayer(_player)

	log.KVs(log.Fields{"roleId": _player.Role().RoleId(), "gSession": gSession}).Info("player login")
	return &message.EnterGameRsp{}
}

// 玩家离线
func (s *Controller) Logout(sourceId string, gSession common.GSession, pbMsg interface{}) proto.Message {
	msg := pbMsg.(*inner.GT2GSessionClosed)
	p, ok := s.PlayerMgr().PlayerBySession(common.GSession(msg.GateSession))
	if ok {
		p.Logout()
	}
	log.KVs(log.Fields{"roleId": p.Role().RoleId(), "gSession": msg.GateSession}).Info("player logout")
	return nil
}
