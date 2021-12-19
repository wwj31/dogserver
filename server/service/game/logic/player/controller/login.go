package controller

import (
	"server/proto/inner_message/inner"
	"server/proto/message"
	"server/service/game/iface"
	"server/service/game/logic/player/localmsg"
)

// 玩家登录
var _ = regist(MsgName(localmsg.Login{}), func(player iface.Player, v interface{}) {
	msg := v.(localmsg.Login)
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

var _ = regist(MsgName(&inner.GT2GSessionClosed{}), func(player iface.Player, v interface{}) {
	player.Logout()
	player.Exit()
})
