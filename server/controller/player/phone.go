package player

import (
	"context"
	"regexp"
	"unicode/utf8"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"

	"server/common/log"
	"server/common/mongodb"
	"server/common/router"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"
	"server/service/login/account"
)

// 修改密码
var _ = router.Reg(func(player *player.Player, msg *outer.ModifyPasswordReq) any {
	if player.Role().Phone() == "" {
		return &outer.FailRsp{Error: outer.ERROR_MODIFY_PASSWORD_NOT_PHONE}
	}

	result := mongodb.Ins.Collection(account.Collection).FindOne(context.Background(), bson.M{account.Phone: player.Role().Phone()})
	if result.Err() == mongo.ErrNoDocuments {
		return &outer.FailRsp{Error: outer.ERROR_PHONE_NOT_FOUND}
	}

	if !validatePassword(msg.NewPassword) {
		return &outer.FailRsp{Error: outer.ERROR_INVALID_PASSWORD_FORMAT}
	}

	_, err := mongodb.Ins.Collection(account.Collection).
		UpdateByID(context.Background(), player.Account().UID, bson.M{"$set": bson.M{
			"phone_password": msg.GetNewPassword(),
		}})
	if err != nil {
		log.Warnw("bing phone failed", "err", err, "rid", player.RID(), "phone", player.Role().Phone())
		return &outer.FailRsp{
			Error: outer.ERROR_FAILED,
			Info:  err.Error(),
		}
	}
	return &outer.ModifyPasswordRsp{}
})

// 绑定手机号
var _ = router.Reg(func(player *player.Player, msg *outer.BindPhoneReq) any {
	result := mongodb.Ins.Collection(account.Collection).FindOne(context.Background(), bson.M{account.Phone: msg.GetPhone()})
	if result.Err() != mongo.ErrNoDocuments {
		return &outer.FailRsp{Error: outer.ERROR_PHONE_WAS_BOUND}
	}

	if msg.Password == "" {
		return &outer.FailRsp{Error: outer.ERROR_PHONE_PASSWORD_IS_EMPTY}
	}

	if !validatePhoneNumber(msg.Phone) {
		log.Warnw("bing phone validate failed", "rid", player.RID(), "phone", msg.Phone)
		return &outer.FailRsp{Error: outer.ERROR_INVALID_PHONE_FORMAT}
	}

	if !validatePassword(msg.Password) {
		return &outer.FailRsp{Error: outer.ERROR_INVALID_PASSWORD_FORMAT}
	}

	_, err := mongodb.Ins.Collection(account.Collection).
		UpdateByID(context.Background(), player.Account().UID, bson.M{"$set": bson.M{
			"phone":          msg.GetPhone(),
			"phone_password": msg.GetPassword(),
		}})
	if err != nil {
		log.Warnw("bing phone failed", "err", err, "rid", player.RID(), "phone", msg.Phone)
		return &outer.FailRsp{
			Error: outer.ERROR_FAILED,
			Info:  err.Error(),
		}
	}

	player.Role().SetPhone(msg.Phone)
	return &outer.BindPhoneRsp{Phone: msg.Phone}
})

func validatePhoneNumber(phoneNumber string) bool {
	// 使用正则表达式匹配电话号码的模式
	pattern := `^1[3-9]\d{9}$`
	regex := regexp.MustCompile(pattern)

	// 检查给定的字符串是否匹配电话号码模式
	return regex.MatchString(phoneNumber)
}

func validatePassword(password string) bool {
	regex := regexp.MustCompile("[\u4e00-\u9fa5]")
	if regex.MatchString(password) {
		return false
	}

	strLen := utf8.RuneCountInString(password)
	if strLen < 1 || strLen > 20 {
		return false
	}
	return true
}
