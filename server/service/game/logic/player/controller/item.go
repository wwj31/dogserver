package controller

import (
	"server/common/log"
	"server/proto/message"
	"server/service/game/iface"
)

var _ = regist(MsgName(&message.UseItemReq{}), func(player iface.Player, v interface{}) {
	msg := v.(*message.UseItemReq)
	if !player.Item().Enough(msg.Items) {
		player.Send2Client(&message.Fail{
			Error: message.ERROR_ITEM_NOT_ENOUGH,
			Info:  msg.String(),
		})
		return
	}
	player.Item().Add(msg.Items, true)
	player.Send2Client(&message.UseItemResp{})
	log.Infow("use item success ", "player", player.ID(), "msg", msg.String())
})
