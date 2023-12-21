package player

import (
	"context"

	"server/common"
	"server/common/log"
	"server/common/rds"
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
	player.PlayerInfo()

	enterGameRsp := &outer.EnterGameRsp{}
	player.Login(msg.NewPlayer, enterGameRsp)

	// 用户留存统计
	pip := rds.Ins.Pipeline()
	if msg.NewPlayer {
		rdsop.AddDailyRegistry(player.Role().ShortId(), pip)
	}
	rdsop.AddDailyLogin(player.Role().ShortId(), pip)
	rdsop.SetRealTimeUser(player.Role().ShortId(), pip)

	if _, er := pip.Exec(context.Background()); er != nil {
		log.Errorw("login user stats pip exec failed", "err", er, "short", player.Role().ShortId())
	}

	return enterGameRsp
})

// 玩家离线
var _ = router.Reg(func(player *player.Player, msg *inner.GSessionClosed) any {
	if player.GateSession().String() == msg.GetGateSession() {
		log.Infow("player offline", "rid", player.RID(), "player gSession", player.GateSession().String(), "msg gSession", msg.GetGateSession())
		player.Logout()

		pip := rds.Ins.Pipeline()
		rdsop.UnsetRealTimeUser(player.Role().ShortId(), pip)
		if _, er := pip.Exec(context.Background()); er != nil {
			log.Errorw("logout user stats pip exec failed", "err", er, "short", player.Role().ShortId())
		}
	}
	return nil
})
