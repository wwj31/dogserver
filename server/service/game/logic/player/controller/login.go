package controller

import (
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
	"server/service/game/logic/player/localmsg"
)

// 玩家登录
var _ = regist(&localmsg.Login{}, func(player iface.Player, v interface{}) {
	msg := v.(localmsg.Login)
	player.SetGateSession(msg.GSession)
	isNew := player.IsNewRole()
	player.Role().SetRoleId(msg.RId)
	player.Role().SetUUId(msg.UId)

	player.Login()
	player.Send2Client(&outer.EnterGameResp{
		NewPlayer: isNew,
	})
})

// 玩家离线
var _ = regist(&inner.GT2GSessionClosed{}, func(player iface.Player, v interface{}) {
	player.Logout()
	err := player.Send(player.Gamer().ID(), v)
	assert(err == nil)
})
