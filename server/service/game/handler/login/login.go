package login

import (
	"github.com/golang/protobuf/proto"
	"server/common"
	"server/proto/message"
	"server/service/game/iface"
)

type Login struct {
	game iface.Gamer
}

func Init(g iface.Gamer) {
	handler := &Login{
		game: g,
	}
	g.RegistMsg((*message.EnterGameReq)(nil), handler.EnterGameReq)
}

// 玩家请求进入游戏
func (s *Login) EnterGameReq(sourceId string, gSession common.GSession, pbMsg interface{}) proto.Message {
	_ = pbMsg.(*message.EnterGameReq)

	// 登录成功的消息
	return &message.EnterGameRsp{
		// todo ...
	}
}
