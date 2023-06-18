package rdsop

import (
	"context"
	"github.com/spf13/cast"
	"server/common"
	"server/common/log"
	"server/common/rds"
	"server/proto/outermsg/outer"
)

func AddAllianceRoom(roomId int64, allianceId int32) {
	rds.Ins.SAdd(context.Background(), RoomsSetKey(allianceId), roomId)
	rds.Ins.HSet(context.Background(), RoomsSetKey(allianceId))
}

func SubAllianceRoom(roomId int64, allianceId int32) {
	rds.Ins.SRem(context.Background(), RoomsSetKey(allianceId), roomId)
}

func RoomList(allianceId int32) (roomIds []int64) {
	arr, err := rds.Ins.SMembers(context.Background(), RoomsSetKey(allianceId)).Result()
	if err != nil {
		log.Errorw("rds op room list failed", "err", err)
		return nil
	}

	for _, v := range arr {
		roomIds = append(roomIds, cast.ToInt64(v))
	}

	return roomIds
}

type NewRoomInfo struct {
	RoomId         int64             `json:"room_id"`
	CreatorShortId int64             `json:"creator_short_id""`
	AllianceId     int32             `json:"alliance_id"`
	GameType       int32             `json:"game_type"`
	Params         *outer.GameParams `json:"params"`
}

func (n NewRoomInfo) SetInfoToRedis() {
	rds.Ins.Set(context.Background(), RoomsInfoKey(n.RoomId), common.JsonMarshal(&n), 0)
}

func DelRoomInfoFromRedis(roomId int64) {
	rds.Ins.Del(context.Background(), RoomsInfoKey(roomId))
}

func (n NewRoomInfo) GetInfoFromRedis() NewRoomInfo {
	var info NewRoomInfo
	str, _ := rds.Ins.Get(context.Background(), RoomsInfoKey(n.RoomId)).Result()
	common.JsonUnmarshal(str, &info)
	return info
}
