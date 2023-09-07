package room

import (
	"fmt"

	"github.com/golang/protobuf/proto"

	"server/rdsop"
	"server/service/room/mahjong"
	"server/service/room/run"

	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/room"

	"server/common/router"
)

// 创建房间
var _ = router.Reg(func(mgr *room.Mgr, msg *inner.CreateRoomReq) any {
	if msg.GameParams == nil {
		return &inner.Error{ErrorInfo: fmt.Errorf("game params is nil").Error()}
	}

	gameParams := &outer.GameParams{}
	if err := proto.Unmarshal(msg.GetGameParams(), gameParams); err != nil {
		return &inner.Error{ErrorInfo: err.Error()}
	}

	roomId := msg.RoomId
	if roomId == 0 {
		var err error
		roomId, err = mgr.GetRoomId()
		if err != nil {
			return &inner.Error{ErrorInfo: err.Error()}
		}
	}

	newRoomInfo := rdsop.NewRoomInfo{
		RoomId:         roomId,
		CreatorShortId: msg.CreatorShortId,
		AllianceId:     msg.AllianceId,
		GameType:       msg.GameType,
		Params:         gameParams,
	}
	newRoom := room.New(&newRoomInfo)

	var gambling room.Gambling
	switch msg.GameType {
	case room.Mahjong:
		gambling = mahjong.New(newRoom)
	case room.RunFaster:
		gambling = run.New(newRoom)
	}
	newRoom.InjectGambling(gambling)

	roomActor := actortype.RoomName(roomId)
	if err := mgr.System().NewActor(roomActor, newRoom); err != nil {
		log.Errorw("create room failed", "msg", msg, "err", err)
		return &inner.Error{ErrorInfo: err.Error()}
	}

	v, err := mgr.RequestWait(roomActor, &inner.RoomInfoReq{})
	if yes, code := common.IsErr(v, err); yes {
		return code
	}

	roomInfoRsp := v.(*inner.RoomInfoRsp)

	_ = mgr.AddRoom(newRoom)

	// 请求创建roomId为0，表示联盟启动初始化已有的房间
	if msg.RoomId == 0 {
		rdsop.AddAllianceRoom(roomId, msg.AllianceId)
		newRoomInfo.SetInfoToRedis()
	}

	log.Infow("create room", "msg", msg.String(), "new room", newRoomInfo)
	return &inner.CreateRoomRsp{RoomInfo: roomInfoRsp.RoomInfo}
})
