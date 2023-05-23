package player

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/tools"

	"server/common/router"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"
)

func InnerToOuter(player *inner.PlayerInfo) *outer.PlayerInfo {
	return &outer.PlayerInfo{
		RID:        player.RID,
		ShortId:    player.ShortId,
		Name:       player.Name,
		Icon:       player.Icon,
		Gender:     player.Gender,
		AllianceId: player.AllianceId,
		Position:   outer.Position(player.Position),
		LoginAt:    player.LoginAt,
		LogoutAt:   player.LogoutAt,
	}
}

// // 获取上、下级信息
//
//	var _ = router.Reg(func(player *player.Player, msg *outer.AgentMembersReq) any {
//		var (
//			upMember    *outer.PlayerInfo
//			downMembers []*outer.PlayerInfo
//		)
//
//		upShortId := rdsop.AgentUp(player.ShortId())
//		if upShortId != 0 {
//			upInfo := rdsop.PlayerInfo(upShortId)
//			upMember = InnerToOuter(&upInfo)
//		}
//
//		downShortIds := rdsop.AgentDown(player.ShortId())
//		for _, shortId := range downShortIds {
//			downInfo := rdsop.PlayerInfo(shortId)
//			downMembers = append(downMembers, InnerToOuter(&downInfo))
//		}
//
//		return &outer.AgentMembersRsp{
//			UpMember:    upMember,
//			DownMembers: downMembers,
//		}
//	})
//
// 获取上、下级信息
var _ = router.Reg(func(player *player.Player, msg *outer.AgentMembersReq) any {
	var (
		upMember    *outer.PlayerInfo
		downMembers []*outer.PlayerInfo
	)
	upMember = &outer.PlayerInfo{
		RID:        tools.XUID(),
		ShortId:    1678594,
		Name:       "你的大爷",
		Icon:       "8",
		Gender:     0,
		AllianceId: 0,
		Position:   4,
		LoginAt:    tools.TimeFormat(time.Now().Add(-(time.Hour * 4))),
		LogoutAt:   tools.TimeFormat(time.Now()),
	}

	for i := 0; i < 20; i++ {
		downMembers = append(downMembers, &outer.PlayerInfo{
			RID:        tools.XUID(),
			ShortId:    2456730 + int64(i),
			Name:       fmt.Sprintf("你的弟弟_%v", i),
			Icon:       cast.ToString(rand.Intn(8) + 1),
			Gender:     0,
			AllianceId: 0,
			Position:   outer.Position(rand.Intn(2)),
			LoginAt:    tools.TimeFormat(time.Now().Add(-(time.Hour * time.Duration(rand.Intn(24))))),
			LogoutAt:   tools.TimeFormat(time.Now()),
		})
	}

	return &outer.AgentMembersRsp{
		UpMember:    upMember,
		DownMembers: downMembers,
	}
})
