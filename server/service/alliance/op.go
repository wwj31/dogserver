package alliance

import (
	"server/common"
	"server/common/log"
	"server/proto/innermsg/inner"
)

// SetMember 添加成员、更新成员信息
func (a *Alliance) SetMember(playerInfo *inner.PlayerInfo, position ...Position) {
	member, ok := a.members[playerInfo.RID]
	if !ok {
		member = &Member{
			RID:      playerInfo.RID,
			ShortId:  playerInfo.ShortId,
			Position: Normal,
		}
	}

	if len(position) > 0 {
		member.Position = position[0]
	}

	member.Alliance = a
	member.GSession = common.GSession(playerInfo.GSession)
	member.Save()

	log.Infow("setMember", "member info", playerInfo.String())
}
