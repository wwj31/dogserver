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

	enterGameRsp := &outer.EnterGameRsp{}
	player.Login(msg.NewPlayer, enterGameRsp)
	player.Send2Client(enterGameRsp)
})

// 玩家离线
var _ = router.Reg(func(player *player.Player, msg *inner.GSessionClosed) {
	if player.GateSession().String() == msg.GetGateSession() {
		player.Logout()
	}
})
