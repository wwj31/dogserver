package controller

import (
	"server/common/log"
	"server/common/router"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"
)

var _ = router.Reg(func(player *player.Player, msg *outer.ChatReq) {
	player.Chat().SendToChannel("world", &outer.ChatNotify{
		SenderId:    player.Role().RoleId(),
		Name:        player.Role().Name(),
		Content:     msg.Content,
		ChannelType: msg.ChannelType,
	})

	log.Infow("chat success ", "player", player.ID(), "msg", msg.String())
})
