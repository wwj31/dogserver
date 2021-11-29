package handler

import (
	"github.com/golang/protobuf/proto"
	"server/common"
	"server/proto/message"
)

// 玩家请求进入游戏
func (s *Controller) EnterGameReq(sourceId string, gSession common.GSession, pbMsg interface{}) proto.Message {
	_ = pbMsg.(*message.EnterGameReq)

	// 登录成功的消息
	return &message.EnterGameRsp{
		// todo ...
	}
}
