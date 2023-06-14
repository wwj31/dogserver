package rdsop

import (
	"context"
	"github.com/spf13/cast"
	"server/common/log"
	"server/common/rds"
)

func AddRoom(roomId, allianceId int32) {
	rds.Ins.SAdd(context.Background(), RoomsKey(allianceId), roomId)
}

func DelRoom(roomId, allianceId int32) {
	rds.Ins.SRem(context.Background(), RoomsKey(allianceId), roomId)
}

func RoomList(allianceId int32) (roomIds []int32) {
	arr, err := rds.Ins.SMembers(context.Background(), RoomsKey(allianceId)).Result()
	if err != nil {
		log.Errorw("rds op room list failed", "err", err)
		return nil
	}

	for _, v := range arr {
		roomIds = append(roomIds, cast.ToInt32(v))
	}

	return roomIds
}
