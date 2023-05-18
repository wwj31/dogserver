package controller

import (
	"server/common/router"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"
)

// 设置头像
var _ = router.Reg(func(player *player.Player, msg *outer.SetIconReq) {
	player.Role().SetIcon(msg.Icon)
	player.Send2Client(&outer.BindPhoneRsp{Phone: msg.Icon})
})
