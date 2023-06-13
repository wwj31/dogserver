package player

import (
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/router"
	"server/proto/convert"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/rdsop"
	"server/service/alliance"
	"server/service/game/logic/player"
)

// 联盟信息变化
var _ = router.Reg(func(player *player.Player, msg *inner.AllianceInfoNtf) any {
	player.Alliance().SetAllianceId(msg.AllianceId)
	player.Alliance().SetPosition(msg.Position)
	player.UpdateInfoToRedis()

	player.SendToClient(&outer.AllianceInfoNtf{
		AllianceId: msg.AllianceId,
		Position:   msg.Position,
	})
	return nil
})

// 搜索玩家信息
var _ = router.Reg(func(player *player.Player, msg *outer.SearchPlayerInfoReq) any {
	playerInfo := rdsop.PlayerInfo(msg.ShortId)
	if playerInfo.RID == "" {
		return outer.ERROR_CAN_NOT_FIND_PLAYER_INFO
	}

	outerPlayerInfo := convert.PlayerInnerToOuter(&playerInfo)
	return &outer.SearchPlayerInfoRsp{Player: outerPlayerInfo}
})

// 邀请加入联盟
var _ = router.Reg(func(player *player.Player, msg *outer.InviteAllianceReq) any {
	if player.Alliance().AllianceId() == 0 {
		return outer.ERROR_PLAYER_NOT_IN_ALLIANCE
	}

	if player.Alliance().Position() == alliance.Normal.Int32() {
		return outer.ERROR_PLAYER_POSITION_LIMIT
	}

	if msg.ShortId == 0 {
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	// 找不到设置的玩家，不能设置职位
	playerInfo := rdsop.PlayerInfo(msg.GetShortId())
	if playerInfo.RID == "" {
		return outer.ERROR_CAN_NOT_FIND_PLAYER_INFO
	}

	// 对方有联盟
	if playerInfo.AllianceId != 0 {
		return outer.ERROR_PLAYER_ALREADY_IN_ALLIANCE
	}

	// 玩家已经有上级
	if playerInfo.UpShortId != 0 {
		return outer.ERROR_PLAYER_ALREADY_HAS_UP
	}

	allianceActor := actortype.AllianceName(player.Alliance().AllianceId())
	result, err := player.RequestWait(allianceActor, &inner.AddMemberReq{
		Player: &playerInfo,
		Ntf:    true,
	})
	if yes, err := common.IsErr(result, err); yes {
		return err
	}

	rdsop.BindAgent(player.Role().ShortId(), playerInfo.ShortId) // 邀请加入联盟绑定上下级关系
	playerInfo.UpShortId = player.Role().ShortId()
	playerInfo.AllianceId = player.Alliance().AllianceId()
	log.Infow("player invite success ", "player", player.Role().ShortId(), "msg", msg.String())
	return &outer.InviteAllianceRsp{Player: convert.PlayerInnerToOuter(&playerInfo)}
})

// 请求设置成员职位
var _ = router.Reg(func(player *player.Player, msg *outer.SetMemberPositionReq) any {
	if player.Alliance().AllianceId() == 0 {
		return outer.ERROR_PLAYER_NOT_IN_ALLIANCE
	}

	if msg.ShortId == 0 {
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	if msg.ShortId == player.Role().ShortId() {
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	// 找不到设置的玩家，不能设置职位
	playerInfo := rdsop.PlayerInfo(msg.GetShortId())
	if playerInfo.RID == "" {
		return outer.ERROR_CAN_NOT_FIND_PLAYER_INFO
	}

	// 对方没有联盟，不能设置职位
	if playerInfo.AllianceId == 0 {
		return outer.ERROR_PLAYER_NOT_IN_ALLIANCE
	}

	// 对方联盟不是本联盟
	if playerInfo.AllianceId != player.Alliance().AllianceId() {
		return outer.ERROR_PLAYER_NOT_IN_CORRECT_ALLIANCE
	}

	// 普通成员没有权限
	if player.Alliance().Position() == alliance.Normal.Int32() {
		return outer.ERROR_PLAYER_POSITION_LIMIT
	}

	// 对方职位比设置者大
	if playerInfo.Position > player.Alliance().Position() {
		return outer.ERROR_PLAYER_POSITION_LIMIT
	}

	// 不是直属下级，不能设置职位(这条规则仅限于队长设小队长)
	if player.Alliance().Position() <= alliance.Captain.Int32() &&
		playerInfo.UpShortId != player.Role().ShortId() {
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

	// 如果职位被盟主，管理员，副盟主设置成功,
	// 那么被设置的人将解除上下级关系，并且绑定盟主为最新上级
	if player.Alliance().Position() >= alliance.Manager.Int32() {
		// 获取盟主
		v, err := player.RequestWait(allianceActor, &inner.AllianceInfoReq{})
		if yes, code := common.IsErr(v, err); yes {
			return code
		}
		alliInfoRsp := v.(*inner.AllianceInfoRsp)
		if playerInfo.UpShortId != alliInfoRsp.MasterShortId {
			if playerInfo.UpShortId != 0 {
				rdsop.AgentCancelUp(playerInfo.ShortId, playerInfo.UpShortId)
			}
			rdsop.BindAgent(alliInfoRsp.MasterShortId, playerInfo.ShortId) // 被管理员以上职位设置职位
			playerInfo.UpShortId = alliInfoRsp.MasterShortId
		}
	}

	playerInfo.Position = int32(msg.Position)
	rdsop.SetPlayerInfo(&playerInfo)
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
	v, err := player.RequestWait(allianceActor, &inner.DisbandAllianceReq{
		RID:        player.RID(),
		AllianceId: player.Alliance().AllianceId(),
	})
	if yes, code := common.IsErr(v, err); yes {
		return code
	}

	return &outer.DisbandAllianceRsp{}
})

// 踢人
var _ = router.Reg(func(player *player.Player, msg *outer.KickOutMemberReq) any {
	playerInfo := rdsop.PlayerInfo(msg.ShortId)
	if playerInfo.RID == "" {
		return outer.ERROR_CAN_NOT_FIND_PLAYER_INFO
	}

	if playerInfo.ShortId == player.Role().ShortId() {
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}
	// 对方没有联盟，不能设置职位
	if playerInfo.AllianceId == 0 {
		return outer.ERROR_PLAYER_NOT_IN_ALLIANCE
	}

	// 操作者权限不够
	if player.Alliance().Position() < alliance.Captain.Int32() {
		return outer.ERROR_PLAYER_POSITION_LIMIT
	}

	// 对方联盟不是本联盟
	if playerInfo.AllianceId != player.Alliance().AllianceId() {
		return outer.ERROR_PLAYER_NOT_IN_CORRECT_ALLIANCE
	}

	// 对方职位比设置者大
	if playerInfo.Position > player.Alliance().Position() {
		return outer.ERROR_PLAYER_POSITION_LIMIT
	}

	// 不是自己的下线不能踢
	if playerInfo.UpShortId != player.Role().ShortId() {
		return outer.ERROR_TARGET_IS_NOT_DOWN
	}

	// 解除被踢者的上下级关系
	rdsop.UnbindAgent(playerInfo.ShortId)
	playerInfo.UpShortId = 0
	rdsop.SetPlayerInfo(&playerInfo)

	// 获取被踢者以及所有下级
	downs := rdsop.AgentDown(playerInfo.ShortId)
	allianceActor := actortype.AllianceName(player.Alliance().AllianceId())
	v, err := player.RequestWait(allianceActor, &inner.KickOutMembersReq{ShortIds: downs})
	if yes, code := common.IsErr(v, err); yes {
		return code
	}

	return &outer.KickOutMemberRsp{}
})
