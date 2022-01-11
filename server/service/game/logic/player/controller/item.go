package controller

import (
	"github.com/golang/protobuf/proto"
	"reflect"
	"server/common/log"
	"server/proto/message"
	"server/service/game/iface"
)

var _ = regist(MsgName(&message.UseItemReq{}), func(player iface.Player, v interface{}) {
	msg, ok := v.(*message.UseItemReq)
	if !ok {
		log.Errorw("use item req msg convert failed", "type", reflect.TypeOf(v).String())
		return
	}

	reslut := player.Item().Use(msg.Items)

	var resp proto.Message
	if reslut != message.ERROR_SUCCESS {
		resp = &message.Fail{
			Error: reslut,
		}
	} else {
		resp = &message.UseItemResp{}
	}
	player.Send2Client(resp)
	log.Infow("use item success ", "player", player.ID(), "msg", msg.String())
})
