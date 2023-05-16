package controller

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"server/common/log"
	"server/common/mongodb"
	"server/common/router"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"
	"server/service/login/account"
)

// 玩家登录
var _ = router.Reg(func(player *player.Player, msg *outer.BindPhoneReq) {
	player.Role().SetPhone(msg.Phone)
	result, err := mongodb.Ins.Collection(account.Collection).
		UpdateByID(context.Background(), player.Account().UID, bson.M{"phone": msg.GetPhone()})
	if err != nil {
		log.Errorf("bing phone failed", "err", err, "rid", player.RID())

		player.Send2Client(&outer.FailRsp{
			Error: outer.ERROR_FAILED,
			Info:  err.Error(),
		})
	}
	player.Role().SetPhone(msg.Phone)
	player.Send2Client(&outer.BindPhoneRsp{})
})
