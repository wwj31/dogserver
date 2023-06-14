package convert

import (
	"github.com/wwj31/dogactor/tools"

	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
)

func RoomInfoInnerToOuter(roomInfo *inner.RoomInfo) *outer.RoomInfo {
	if roomInfo == nil {
		return nil
	}
	var players []*outer.PlayerInfo
	for _, v := range roomInfo.Players {
		players = append(players, PlayerInnerToOuter(v))
	}

	return &outer.RoomInfo{
		Id:       roomInfo.RoomId,
		GameType: outer.GameType(roomInfo.GameType),
		Players:  players,
	}
}

func PlayerInnerToOuter(player *inner.PlayerInfo) *outer.PlayerInfo {
	if player == nil {
		return nil
	}
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
