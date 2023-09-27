package player

import (
	"server/common/router"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"
)

// 新玩家信息
var _ = router.Reg(func(player *player.Player, msg *inner.NewPlayerInfoReq) any {
	player.Role().SetShortId(msg.ShortId)
	player.Role().SetUID(msg.AccountInfo.UID)
	return &inner.NewPlayerInfoRsp{}
})

// 设置信息
var _ = router.Reg(func(player *player.Player, msg *outer.SetRoleInfoReq) any {
	r := []rune(msg.Name)
	if len(r) > 20 {
		return outer.ERROR_FAILED
	}

	if msg.Name == "" {
		msg.Name = player.Role().Name()
	}

	if msg.Icon == "" {
		msg.Icon = player.Role().Icon()
	}
	player.Role().SetBaseInfo(msg.Icon, msg.Name, msg.Gender)
	return &outer.SetRoleInfoRsp{Icon: msg.Icon, Name: msg.Name, Gender: msg.Gender}
})
