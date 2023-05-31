package player

import (
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/router"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/rdsop"
	"server/service/alliance"
	"server/service/game/logic/player"
)

// 加入联盟通知
var _ = router.Reg(func(player *player.Player, msg *outer.AllianceInfoNtf) any {
	player.Alliance().SetAllianceId(msg.AllianceId)
	player.Alliance().SetPosition(msg.Position)
	return msg
})

// 请求设置成员职位
var _ = router.Reg(func(player *player.Player, msg *outer.SetMemberPositionReq) any {
	if player.Alliance().AllianceId() == 0 {
		return outer.ERROR_PLAYER_NOT_IN_ALLIANCE
	}

	if msg.ShortId == 0 {
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	// 找不到设置的玩家，不能设置职位
	playerInfo := rdsop.PlayerInfo(msg.GetShortId())
	if playerInfo.RID == "" {
		return outer.ERROR_CAN_NOT_FIND_PLAYER_INFO
	}

	// 没有联盟，不能设置职位
	if playerInfo.AllianceId == 0 {
		return outer.ERROR_PLAYER_NOT_IN_ALLIANCE
	}

	// 对方没有联盟，不能设置职位
	if playerInfo.AllianceId != player.Alliance().AllianceId() {
		return outer.ERROR_PLAYER_NOT_IN_CORRECT_ALLIANCE
	}

	// 不在一个联盟，不能设置职位
	if playerInfo.Position > player.Alliance().Position() {
		return outer.ERROR_CAN_NOT_SET_HIGHER_POSITION
	}

	// 不是直属下级，不能设置职位
	if playerInfo.UpShortId != player.Role().ShortId() {
		return outer.ERROR_CAN_NOT_SET_NOT_IN_DOWN_POSITION
	}

	allianceActor := actortype.AllianceName(player.Alliance().AllianceId())
	rsp, err := player.RequestWait(allianceActor, &inner.SetMemberPositionReq{
		Player:   &playerInfo,
		Position: int32(msg.Position),
	})

	if err != nil {
		log.Warnw("set member position failed", "err", err, "msg", msg.String())
		return outer.ERROR_FAILED
	}

	if failed, ok := rsp.(*outer.FailRsp); ok {
		log.Warnw("alliance rsp fail", "rsp", failed.String())
		return rsp
	}

	return &outer.SetMemberPositionRsp{
		ShortId:  msg.ShortId,
		Position: msg.Position,
	}
})

// 盟主解散联盟
var _ = router.Reg(func(player *player.Player, msg *outer.DisbandAllianceReq) any {
	if player.Alliance().AllianceId() == 0 {
		return outer.ERROR_PLAYER_NOT_IN_ALLIANCE
	}

	if player.Alliance().Position() != alliance.Master.Int32() {
		return outer.ERROR_PLAYER_POSITION_LIMIT
	}

	allianceActor := actortype.AllianceName(player.Alliance().AllianceId())
	v, err := player.RequestWait(allianceActor, &inner.DisbandAllianceReq{RID: player.RID()})
	if yes, code := common.IsErr(v, err); yes {
		return code
	}

	return &outer.DisbandAllianceRsp{}
})

// 通知 联盟解散
var _ = router.Reg(func(player *player.Player, msg *inner.AllianceDisbandedNtf) any {
	player.Alliance().SetAllianceId(0)
	player.Alliance().SetPosition(0)

	player.GateSession().SendToClient(player, &outer.AllianceInfoNtf{
		AllianceId: 0,
		Position:   0,
	})
	return nil
})
