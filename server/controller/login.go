package controller

import (
	"server/common"
	"server/common/log"
	"server/common/router"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"
)

// 玩家登录
var _ = router.Reg(func(player *player.Player, msg *outer.EnterGameReq) {
	if common.EnterGameToken(msg.RID, msg.NewPlayer) != msg.Checksum {
		log.Warnw("checksum md5 check failed", "msg", msg.String())
		return
	}

	player.Role().SetRoleId(msg.RID)

	player.Login(msg.NewPlayer)
	player.Send2Client(&outer.EnterGameRsp{
		NewPlayer: msg.NewPlayer,
	})
})

// 玩家离线
var _ = router.Reg(func(player *player.Player, msg *inner.GSessionClosed) {
	player.Logout()
})
