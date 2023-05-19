package controller

import (
	"context"
	"regexp"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"

	"server/common/log"
	"server/common/mongodb"
	"server/common/router"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"
	"server/service/login/account"
)

// 绑定手机号
var _ = router.Reg(func(player *player.Player, msg *outer.BindPhoneReq) {
	result := mongodb.Ins.Collection(account.Collection).FindOne(context.Background(), bson.M{"phone": msg.GetPhone()})
	if result.Err() != mongo.ErrNoDocuments {
		player.Send2Client(&outer.FailRsp{Error: outer.ERROR_PHONE_WAS_BOUND})
		return
	}

	if msg.Password == "" {
		player.Send2Client(&outer.FailRsp{Error: outer.ERROR_PHONE_PASSWORD_IS_EMPTY})
		return
	}

	if !validatePhoneNumber(msg.Phone) {
		log.Warnw("bing phone validate failed", "rid", player.RID(), "phone", msg.Phone)
		player.Send2Client(&outer.FailRsp{Error: outer.ERROR_INVALID_PHONE_FORMAT})
		return
	}

	if !validatePassword(msg.Password) {
		player.Send2Client(&outer.FailRsp{Error: outer.ERROR_INVALID_PASSWORD_FORMAT})
		return
	}

	_, err := mongodb.Ins.Collection(account.Collection).
		UpdateByID(context.Background(), player.Account().UID, bson.M{"$set": bson.M{
			"phone":          msg.GetPhone(),
			"phone_password": msg.GetPassword(),
		}})
	if err != nil {
		log.Warnw("bing phone failed", "err", err, "rid", player.RID(), "phone", msg.Phone)
		player.Send2Client(&outer.FailRsp{
			Error: outer.ERROR_FAILED,
			Info:  err.Error(),
		})
		return
	}

	player.Role().SetPhone(msg.Phone)
	player.Send2Client(&outer.BindPhoneRsp{Phone: msg.Phone})
})

func validatePhoneNumber(phoneNumber string) bool {
	// 使用正则表达式匹配电话号码的模式
	pattern := `^1[3-9]\d{9}$`
	regex := regexp.MustCompile(pattern)

	// 检查给定的字符串是否匹配电话号码模式
	return regex.MatchString(phoneNumber)
}

func validatePassword(password string) bool {
	if len(password) < 6 {
		return false
	}
	return true
}
