package player

import (
	"server/common/router"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"
)

// 加入联盟通知
var _ = router.Reg(func(player *player.Player, msg *inner.JoinAllianceNtf) any {
	player.Alliance().SetAllianceId(msg.AllianceId)
	player.Alliance().SetPosition(msg.Position)

	player.GateSession().SendToClient(player, &outer.JoinAllianceNtf{
		AllianceId: msg.AllianceId,
		Position:   msg.Position,
	})
	return nil
})
