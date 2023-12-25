package player

import (
	"server/common"
	"server/common/actortype"
	"server/common/router"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/rdsop"
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

// 微信用户信息
var _ = router.Reg(func(player *player.Player, msg *inner.SetWeChatInfoReq) any {
	player.Role().SetBaseInfo(msg.UserInfo.Icon, msg.UserInfo.Name, msg.UserInfo.Gender)
	return &inner.SetWeChatInfoRsp{}
})

// 修改金币
var _ = router.Reg(func(p *player.Player, msg *inner.ModifyGoldReq) any {
	// 仅当修改操作不是来自于房间结算，并且玩家在房间中时执行以下判断
	var delayModify bool
	if msg.RoomId == 0 && p.Room().RoomId() != 0 {
		roomId := actortype.RoomName(p.Room().RoomId())
		v, err := p.RequestWait(roomId, &inner.RoomCanSetGoldReq{ShortId: p.Role().ShortId()})
		if yes, errCode := common.IsErr(v, err); yes {
			return errCode
		}
		rsp := v.(*inner.RoomCanSetGoldRsp)
		if !rsp.Ok {
			delayModify = true
		}
	}

	addGold := msg.Gold
	if msg.Set {
		addGold = msg.Gold - p.Role().Gold()
	}

	// 如果不允许扣为负，就检查
	if !msg.SmallZero && addGold < 0 && p.Role().Gold()+addGold < 0 {
		return &inner.ModifyGoldRsp{Success: false, Info: p.PlayerInfo()}
	}

	// 不能立刻修改，就丢入金币延迟变动中
	if delayModify {
		rdsop.AddDelayModifyGold(p.Role().ShortId(), addGold)
		info := p.PlayerInfo()
		info.Gold += addGold
		return &inner.ModifyGoldRsp{Success: true, Info: info}
	} else {
		p.Role().AddGold(addGold)
		p.Role().SyncDelayGold()
		return &inner.ModifyGoldRsp{Success: true, Info: p.PlayerInfo()}
	}
})
