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
var _ = router.Reg(func(player *player.Player, msg *outer.EnterGameReq) any {
	if common.EnterGameToken(msg.RID, msg.NewPlayer) != msg.Checksum {
		log.Warnw("checksum md5 check failed", "msg", msg.String())
		return &outer.FailRsp{}
	}

	enterGameRsp := &outer.EnterGameRsp{}
	player.Login(msg.NewPlayer, enterGameRsp)

	return enterGameRsp
})

// 玩家离线
var _ = router.Reg(func(player *player.Player, msg *inner.GSessionClosed) any {
	if player.GateSession().String() == msg.GetGateSession() {
		player.Logout()
	}
	return nil
})
