package convert

import (
	"github.com/wwj31/dogactor/tools"

	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/rdsop"
)

func RoomInfoInnerToOuter(roomInfo *inner.RoomInfo) *outer.RoomInfo {
	var players []*outer.PlayerInfo
	for _, v := range roomInfo.Players {
		players = append(players, PlayerInnerToOuter(v))
	}

	creator := rdsop.PlayerInfo(roomInfo.GetCreatorShortId())

	return &outer.RoomInfo{
		Id:      roomInfo.RoomId,
		Creator: PlayerInnerToOuter(&creator),
		Players: players,
	}
}

func PlayerInnerToOuter(player *inner.PlayerInfo) *outer.PlayerInfo {
	return &outer.PlayerInfo{
		RID:        player.RID,
		ShortId:    player.ShortId,
		Name:       player.Name,
		Icon:       player.Icon,
		Gender:     player.Gender,
		AllianceId: player.AllianceId,
		Position:   outer.Position(player.Position),
		LoginAt:    tools.TimeParse(player.LoginAt).Unix(),
		LogoutAt:   tools.TimeParse(player.LogoutAt).Unix(),
		UpShortId:  player.UpShortId,
	}
}
