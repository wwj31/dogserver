package controller

import (
	"github.com/golang/protobuf/proto"
	"reflect"
	"server/common/log"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
)

var _ = regist(MsgName(&outer.UseItemReq{}), func(player iface.Player, v interface{}) {
	msg, ok := v.(*outer.UseItemReq)
	if !ok {
		log.Errorw("use item req msg convert failed", "type", reflect.TypeOf(v).String())
		return
	}

	reslut := player.Item().Use(msg.Items)

	var resp proto.Message
	if reslut != outer.ERROR_SUCCESS {
		resp = &outer.Fail{
			Error: reslut,
		}
	} else {
		resp = &outer.UseItemResp{}
	}
	player.Send2Client(resp)
	log.Infow("use item success ", "player", player.ID(), "msg", msg.String())
})
