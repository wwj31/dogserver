package handler

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/log"
	"server/common"
	"server/proto/message"
	"server/service/game/player"
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
	_player.Login()
	return &message.EnterGameRsp{}
}
