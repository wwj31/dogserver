package controller

import (
	"server/common/log"
	"server/common/router"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"

	"github.com/gogo/protobuf/proto"
)

// 使用道具
var _ = router.Reg(func(player *player.Player, msg *outer.UseItemReq) {
	result := player.Item().Use(msg.Items)

	var resp proto.Message
	if result != outer.ERROR_SUCCESS {
		resp = &outer.Fail{
			Error: result,
		}
	} else {
		resp = &outer.UseItemRsp{}
	}
	player.Send2Client(resp)
	log.Infow("use item success ", "player", player.ID(), "msg", msg.String())
})
