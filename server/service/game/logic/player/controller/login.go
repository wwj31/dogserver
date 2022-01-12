package controller

import (
	"server/proto/inner/inner"
	"server/proto/message"
	"server/service/game/iface"
	"server/service/game/logic/player/localmsg"
)

// 玩家登录
var _ = regist(MsgName(localmsg.Login{}), func(player iface.Player, v interface{}) {
	msg := v.(localmsg.Login)
	player.SetGateSession(msg.GSession)

	isNew := player.IsNewRole()
	//// 新号处理
	if isNew {
		// todo ...
		player.Item().Add(map[int64]int64{123: 999})
	}

	player.Login()

	player.Send2Client(&message.EnterGameResp{
		NewPlayer: isNew,
	})
})

var _ = regist(MsgName(&inner.GT2GSessionClosed{}), func(player iface.Player, v interface{}) {
	player.Logout()
})
