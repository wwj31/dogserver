package player

import (
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/router"
	"server/proto/convert"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/rdsop"
	"server/service/alliance"
	"server/service/game/logic/player"
	"sync"
	"sync/atomic"
	"time"
)

// 创建房间
var _ = router.Reg(func(player *player.Player, msg *outer.CreateRoomReq) any {
	if player.Alliance().Position() != alliance.Master.Int32() {
		return outer.ERROR_PLAYER_POSITION_LIMIT
	}

	roomMgrId := rdsop.GetRoomMgrId()
	if roomMgrId == -1 {
		return outer.ERROR_FAILED
	}

	roomMgrActor := actortype.RoomMgrName(roomMgrId)
	v, err := player.RequestWait(roomMgrActor, &inner.CreateRoomReq{
		GameType: msg.GameType.Int32(),
		Creator:  player.PlayerInfo(),
	})
	if yes, code := common.IsErr(v, err); yes {
		return code
	}
	createRoomRsp := v.(*inner.CreateRoomRsp)

	//player.Room().SetRoomInfo(createRoomRsp.RoomInfo)
	roomInfo := convert.RoomInfoInnerToOuter(createRoomRsp.RoomInfo)
	return &outer.CreateRoomRsp{Room: roomInfo}
})

// 解散房间
var _ = router.Reg(func(player *player.Player, msg *outer.DisbandRoomReq) any {
	if player.Alliance().Position() != alliance.Master.Int32() {
		return outer.ERROR_PLAYER_POSITION_LIMIT
	}

	roomActor := actortype.RoomName(msg.Id)
	v, err := player.RequestWait(roomActor, &inner.DisbandRoomReq{RoomId: msg.GetId()})
	if yes, code := common.IsErr(v, err); yes {
		return code
	}
	return &outer.DisbandRoomRsp{Id: msg.Id}
})

// 房间列表
var _ = router.Reg(func(player *player.Player, msg *outer.RoomListReq) any {
	if player.Alliance().AllianceId() == 0 {
		return outer.ERROR_PLAYER_NOT_IN_ALLIANCE
	}

	var (
		roomList []*outer.RoomInfo
		lock     sync.Mutex
		count    atomic.Int32
		timeout  atomic.Bool
	)

	roomListIds := rdsop.RoomList(player.Alliance().AllianceId())

	for _, roomId := range roomListIds {
		roomActor := actortype.RoomName(roomId)
		go func() {
			defer func() {
				recover()
				count.Add(1)
			}()

			v, err := player.RequestWait(roomActor, &inner.RoomInfoReq{}, 3*time.Second)
			if yes, code := common.IsErr(v, err); yes {
				log.Warnw("room list info failed", "code", code)
				return
			}

			if timeout.Load() {
				return
			}

			roomInfoRsp := v.(*inner.RoomInfoRsp)
			lock.Lock()
			defer lock.Unlock()
			roomList = append(roomList, convert.RoomInfoInnerToOuter(roomInfoRsp.RoomInfo))
		}()
	}

	for i := 0; i <= 100; i++ {
		time.Sleep(20 * time.Millisecond)
		if count.Load() == int32(len(roomListIds)) {
			break
		}
		if i == 100 {
			timeout.Store(true)
		}
		if i%10 == 0 {
			log.Warnw("room list sleep too long", "i", i)
		}
	}

	lock.Lock()
	defer lock.Unlock()
	return &outer.RoomListRsp{RoomList: roomList}
})

// 加入房间
var _ = router.Reg(func(p *player.Player, msg *outer.JoinRoomReq) any {
	if p.Room().RoomId() != 0 {
		return outer.ERROR_PLAYER_ALREADY_IN_ROOM
	}

	roomActor := actortype.RoomName(p.Room().RoomId())
	v, err := p.RequestWait(roomActor, &inner.JoinRoomReq{Player: p.PlayerInfo()})
	if yes, code := common.IsErr(v, err); yes {
		return code
	}
	joinRoomRsp := v.(*inner.JoinRoomRsp)

	p.Room().SetRoomInfo(joinRoomRsp.RoomInfo)
	roomInfo := convert.RoomInfoInnerToOuter(joinRoomRsp.RoomInfo)
	return &outer.JoinRoomRsp{Room: roomInfo}
})

// 离开房间
var _ = router.Reg(func(player *player.Player, msg *outer.LeaveRoomReq) any {
	// TODO ...
	return &outer.LeaveRoomRsp{}
})
