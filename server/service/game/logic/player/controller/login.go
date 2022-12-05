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
	player.Role().SetRoleId(msg.RId)
	player.Role().SetUId(msg.UId)

	player.Login(msg.First)
	player.Send2Client(&outer.EnterGameResp{
		NewPlayer: msg.First,
	})
})

// 玩家离线
var _ = regist(&inner.GT2GSessionClosed{}, func(player iface.Player, v interface{}) {
	player.Logout()
	err := player.Send(player.Gamer().ID(), v)
	assert(err == nil)
})
