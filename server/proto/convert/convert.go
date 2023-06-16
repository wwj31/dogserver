package convert

import (
	"github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
)

func RoomPlayerInfoInnerToOuter(roomPlayer *inner.RoomPlayerInfo) *outer.RoomPlayerInfo {
	return &outer.RoomPlayerInfo{
		BaseInfo: PlayerInnerToOuter(roomPlayer.BaseInfo),
		Ready:    roomPlayer.Ready,
	}
}

func RoomInfoInnerToOuter(roomInfo *inner.RoomInfo) *outer.RoomInfo {
	if roomInfo == nil {
		return nil
	}
	var players []*outer.RoomPlayerInfo
	for _, roomPlayer := range roomInfo.Players {
		players = append(players, RoomPlayerInfoInnerToOuter(roomPlayer))
	}

	gameParams := &outer.GameParams{}
	if err := proto.Unmarshal(roomInfo.GetGameParams(), gameParams); err != nil {
		log.Errorw("game params unmarshal failed", "err", err, "game params", gameParams)
		return nil
	}

	return &outer.RoomInfo{
		Id:         roomInfo.RoomId,
		GameType:   outer.GameType(roomInfo.GameType),
		GameParams: gameParams,
		Players:    players,
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
