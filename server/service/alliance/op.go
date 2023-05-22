package alliance

import (
	"github.com/wwj31/dogactor/tools"

	"server/common"
	"server/common/log"
	"server/proto/innermsg/inner"
)

// SetMember 添加成员、更新成员信息
func (a *Alliance) SetMember(playerInfo *inner.PlayerInfo) {
	member, ok := a.members[playerInfo.RID]
	if !ok {
		member = &Member{
			RID:       playerInfo.RID,
			ShortId:   playerInfo.ShortId,
			Name:      playerInfo.Name,
			Position:  Normal,
			OnlineAt:  tools.TimeParse(playerInfo.LoginAt),
			OfflineAt: tools.TimeParse(playerInfo.LogoutAt),
			GSession:  common.GSession(playerInfo.GSession),
		}
	}

	member.OnlineAt = tools.TimeParse(playerInfo.LoginAt)
	member.OfflineAt = tools.TimeParse(playerInfo.LogoutAt)
	member.GSession = common.GSession(playerInfo.GSession)

	log.Infow("setMember", "member info", playerInfo.String())
}
