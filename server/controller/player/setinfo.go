package player

import (
	"server/common/router"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"
)

// 设置信息
var _ = router.Reg(func(player *player.Player, msg *outer.SetRoleInfoReq) any {
	player.Role().SetBaseInfo(msg.Icon, msg.Name, msg.Gender)
	return &outer.SetRoleInfoRsp{Icon: msg.Icon, Name: msg.Name, Gender: msg.Gender}
})
