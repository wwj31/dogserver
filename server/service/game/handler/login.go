package handler

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/log"
	"server/common"
	"server/proto/message"
)

// 玩家请求进入游戏
func (s *Controller) EnterGameReq(sourceId string, gSession common.GSession, pbMsg interface{}) proto.Message {
	msg := pbMsg.(*message.EnterGameReq)
	log.KV("msg", msg).Debug("EnterGameReq")
	// 登录成功的消息
	return &message.EnterGameRsp{
		// todo ...
	}
}
