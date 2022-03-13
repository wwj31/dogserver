package controller

import (
	"reflect"
	"server/common/log"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
)

var _ = regist(MsgName(&outer.ChatReq{}), func(player iface.Player, v interface{}) {
	msg, ok := v.(*outer.ChatReq)
	assert(ok, "chat req msg convert failed", "type", reflect.TypeOf(v).String())

	player.Chat().SendToChannel("world", &outer.ChatNotify{
		SenderId:    player.Role().RoleId(),
		Name:        player.Role().Name(),
		Content:     msg.Content,
		ChannelType: msg.ChannelType,
	})

	log.Infow("chat success ", "player", player.ID(), "msg", msg.String())
})
