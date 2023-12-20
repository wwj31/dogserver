package player

import (
	"server/common"
	"server/common/log"
	"server/common/router"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/rdsop"
	"server/service/game/logic/player"
)

// 玩家登录
var _ = router.Reg(func(player *player.Player, msg *outer.EnterGameReq) any {
	if common.EnterGameToken(msg.RID, msg.NewPlayer) != msg.Checksum {
		log.Warnw("checksum md5 check failed", "msg", msg.String())
		return outer.ERROR_FAILED
	}

	enterGameRsp := &outer.EnterGameRsp{}
	player.Login(msg.NewPlayer, enterGameRsp)

	// 用户留存统计
	if msg.NewPlayer {
		rdsop.AddDailyRegistry(player.Role().ShortId())
	}
	rdsop.AddDailyLogin(player.Role().ShortId())

	return enterGameRsp
})

// 玩家离线
var _ = router.Reg(func(player *player.Player, msg *inner.GSessionClosed) any {
	if player.GateSession().String() == msg.GetGateSession() {
		log.Infow("player offline", "rid", player.RID(), "player gSession", player.GateSession().String(), "msg gSession", msg.GetGateSession())
		player.Logout()
	}
	return nil
})
