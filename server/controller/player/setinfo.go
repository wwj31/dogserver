package player

import (
	"server/common/router"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"
)

// 设置信息
var _ = router.Reg(func(player *player.Player, msg *outer.SetRoleInfoReq) any {
	r := []rune(msg.Name)
	if len(r) > 20 {
		return &outer.FailRsp{Error: outer.ERROR_FAILED}
	}

	if msg.Name == "" {
		msg.Name = player.Role().Name()
	}

	if msg.Icon == "" {
		msg.Name = player.Role().Icon()
	}
	player.Role().SetBaseInfo(msg.Icon, msg.Name, msg.Gender)
	return &outer.SetRoleInfoRsp{Icon: msg.Icon, Name: msg.Name, Gender: msg.Gender}
})
