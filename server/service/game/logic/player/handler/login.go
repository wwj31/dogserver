package handler

import (
	"server/proto/message"
	"server/service/game/iface"
	"server/service/game/logic/player/msg"
)

var _ = regist(MsgName(msg.Login{}), func(player iface.Player, v interface{}) {
	msg := v.(msg.Login)
	player.SetGateSession(msg.GSession)

	//// 新号处理
	if player.IsNewRole() {
		// todo ...
		player.Item().Add(map[int64]int64{123: 999})
	}

	player.Login()

	player.Send2Client(&message.LoginRsp{
		UID:     player.Role().UUId(),
		RID:     player.Role().RoleId(),
		Cryptic: "",
	})
})
