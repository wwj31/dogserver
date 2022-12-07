package controller

import (
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
)

// 玩家登录
var _ = regist(&outer.EnterGameReq{}, func(player iface.Player, v interface{}) {
	msg := v.(outer.EnterGameReq)
	player.Role().SetRoleId(msg.RID)
	player.Role().SetUId(msg.UID)

	player.Login(msg.NewPlayer)
	player.Send2Client(&outer.EnterGameResp{
		NewPlayer: msg.NewPlayer,
	})
})

// 玩家离线
var _ = regist(&inner.GT2GSessionClosed{}, func(player iface.Player, v interface{}) {
	player.Logout()
	err := player.Send(player.Gamer().ID(), v)
	assert(err == nil)
})
