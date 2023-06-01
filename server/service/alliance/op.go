package alliance

import (
	"context"

	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/rds"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/rdsop"
)

// AddMember 添加成员、更新成员信息
func (a *Alliance) AddMember(playerInfo *inner.PlayerInfo, ntf bool, position ...Position) *Member {
	member, ok := a.members[playerInfo.RID]
	if !ok {
		member = &Member{
			RID:      playerInfo.RID,
			ShortId:  playerInfo.ShortId,
			Position: Normal,
		}
		a.members[playerInfo.RID] = member
	}

	if len(position) > 0 {
		member.Position = position[0]
	}

	member.Alliance = a
	member.GSession = common.GSession(playerInfo.GSession)
	member.Save()

	if !ntf {
		return member
	}

	// 以下逻辑增对某个玩家被动进入联盟的处理
	if playerInfo.GSession != "" {
		// 玩家在线，通知Player actor修改联盟id，
		a.Send(actortype.PlayerId(playerInfo.RID), &outer.AllianceInfoNtf{
			AllianceId: a.allianceId,
			Position:   member.Position.Int32(),
		})
	} else if member.Position == Master {
		// 盟主不在线进入联盟，单独记录，下次登录时会维护player身上的联盟数据,
		// 非盟主成员不在线进去联盟无需处理，下次进入会检查上级联盟跟随进去
		key := rdsop.JoinAllianceKey(playerInfo.ShortId)
		rds.Ins.Set(context.Background(), key, a.allianceId, 0)
	}

	// 更新玩家公共数据
	playerInfo.Position = Master.Int32()
	playerInfo.AllianceId = a.allianceId
	rdsop.SetPlayerInfo(playerInfo)

	log.Infow("setMember", "member info", playerInfo.String())
	return member
}
