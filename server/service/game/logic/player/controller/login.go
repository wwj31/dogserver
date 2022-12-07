package controller

import (
	"server/common"
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
)

// 玩家登录
var _ = registry(&outer.EnterGameReq{}, func(player iface.Player, v interface{}) {
	msg := v.(outer.EnterGameReq)

	if common.LoginMD5(msg.UID, msg.RID, msg.NewPlayer) != msg.Checksum {
		log.Warnw("checksum md5 check faild", "msg", msg.String())
		return
	}

	player.Role().SetRoleId(msg.RID)
	player.Role().SetUId(msg.UID)

	player.Login(msg.NewPlayer)
	player.Send2Client(&outer.EnterGameResp{
		NewPlayer: msg.NewPlayer,
	})
})

// 玩家离线
var _ = registry(&inner.GSessionClosed{}, func(player iface.Player, v interface{}) {
	player.Logout()
})
