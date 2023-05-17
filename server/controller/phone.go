package controller

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"regexp"
	"server/common/log"
	"server/common/mongodb"
	"server/common/router"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"
	"server/service/login/account"
)

// 绑定手机号
var _ = router.Reg(func(player *player.Player, msg *outer.BindPhoneReq) {
	player.Role().SetPhone(msg.Phone)
	_, err := mongodb.Ins.Collection(account.Collection).
		UpdateByID(context.Background(), player.Account().UID, bson.M{"phone": msg.GetPhone()})
	if err != nil {
		log.Warnw("bing phone failed", "err", err, "rid", player.RID(), "phone", msg.Phone)
		player.Send2Client(&outer.FailRsp{
			Error: outer.ERROR_FAILED,
			Info:  err.Error(),
		})
		return
	}

	if !validatePhoneNumber(msg.Phone) {
		log.Warnw("bing phone validate failed", "rid", player.RID(), "phone", msg.Phone)
		player.Send2Client(&outer.FailRsp{Error: outer.ERROR_INVALID_PHONE})
		return
	}

	player.Role().SetPhone(msg.Phone)
	player.Send2Client(&outer.BindPhoneRsp{})
})

func validatePhoneNumber(phoneNumber string) bool {
	// 使用正则表达式匹配电话号码的模式
	pattern := `^1[3-9]\d{9}$`
	regex := regexp.MustCompile(pattern)

	// 检查给定的字符串是否匹配电话号码模式
	return regex.MatchString(phoneNumber)
}
