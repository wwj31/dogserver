package rdsop

import (
	"bytes"
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/tools"

	"server/common"
	"server/common/log"
	"server/common/rds"
	"server/proto/outermsg/outer"
)

func AddAllianceRoom(roomId int64, allianceId int32) {
	rds.Ins.SAdd(context.Background(), RoomsSetKey(allianceId), roomId)
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
	ManifestId      string            `json:"manifest_id"`
	OwnerMgrActorId string            `json:"owner_mgr_actor_id"` // 归属的房间管理器actorId
	RoomId          int64             `json:"room_id"`
	CreatorShortId  int64             `json:"creator_short_id"`
	AllianceId      int32             `json:"alliance_id"`
	GameType        int32             `json:"game_type"`
	Params          *outer.GameParams `json:"params"`
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
	if info.Params == nil {
		info.Params = &outer.GameParams{}
	}
	return info
}

func AddRoomRecording(recordMsg *outer.Recording) {
	data := common.ProtoMarshal(recordMsg)
	rds.Ins.ZAdd(context.Background(), RoomsRecordingDataKey(recordMsg.Room.RoomId), redis.Z{
		Score:  float64(recordMsg.GameStartAt),
		Member: string(data),
	})

	// 删除4天前的数据
	SevenDayAgo := tools.Now().Add(-4 * tools.Day).UnixMilli()
	rds.Ins.ZRemRangeByScore(context.Background(), RoomsRecordingDataKey(recordMsg.Room.RoomId), "0", cast.ToString(SevenDayAgo))
}

func GetRoomRecording(roomId, gameStartAt int64) *outer.Recording {
	cmd := rds.Ins.ZRangeByScore(context.Background(), RoomsRecordingDataKey(roomId), &redis.ZRangeBy{
		Min:    cast.ToString(float64(gameStartAt)),
		Max:    cast.ToString(float64(gameStartAt)),
		Offset: 0,
		Count:  1,
	})

	result, err := cmd.Result()
	if err != nil {
		log.Warnw("get room recording failed ", "err", err)
		return nil
	}

	if len(result) == 0 {
		log.Warnw("cannot find room recording")
		return nil
	}

	buffer := bytes.NewBufferString(result[0])
	var recording outer.Recording
	if err = proto.Unmarshal(buffer.Bytes(), &recording); err != nil {
		log.Warnw("room recording proto unmarshal failed", "err", err)
		return nil
	}
	return &recording
}
