package player

import (
	"github.com/golang/protobuf/proto"
	"sync"
	"sync/atomic"
	"time"

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
)

// 创建房间
var _ = router.Reg(func(p *player.Player, msg *outer.CreateRoomReq) any {
	if msg.GameParams == nil {
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	if p.Alliance().Position() != alliance.Master.Int32() {
		return outer.ERROR_PLAYER_POSITION_LIMIT
	}

	roomMgrId := rdsop.GetRoomMgrId()
	if roomMgrId == -1 {
		return outer.ERROR_FAILED
	}

	gameParamsBytes, _ := proto.Marshal(msg.GetGameParams())
	v, err := p.RequestWait(actortype.RoomMgrName(roomMgrId), &inner.CreateRoomReq{
		GameType:       msg.GameType.Int32(),
		CreatorShortId: p.Role().ShortId(),
		AllianceId:     p.Alliance().AllianceId(),
		GameParams:     gameParamsBytes,
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

	roomActor := actortype.RoomName(msg.RoomId)
	v, err := player.RequestWait(roomActor, &inner.DisbandRoomReq{RoomId: msg.RoomId})
	if yes, code := common.IsErr(v, err); yes {
		return code
	}
	return &outer.DisbandRoomRsp{RoomId: msg.RoomId}
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

	if msg.RoomId == 0 {
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	roomActor := actortype.RoomName(msg.RoomId)
	v, err := p.RequestWait(roomActor, &inner.JoinRoomReq{Player: p.PlayerInfo()})
	if yes, code := common.IsErr(v, err); yes {
		return code
	}
	joinRoomRsp := v.(*inner.JoinRoomRsp)

	p.Room().SetRoomInfo(joinRoomRsp.RoomInfo)
	roomInfo := convert.RoomInfoInnerToOuter(joinRoomRsp.RoomInfo)
	p.UpdateInfoToRedis()
	return &outer.JoinRoomRsp{Room: roomInfo}
})

// 离开房间
var _ = router.Reg(func(p *player.Player, msg *outer.LeaveRoomReq) any {
	if p.Room().RoomId() == 0 {
		return outer.ERROR_PLAYER_NOT_IN_ROOM
	}

	// TODO ...

	roomActor := actortype.RoomName(p.Room().RoomId())
	v, err := p.RequestWait(roomActor, &inner.LeaveRoomReq{ShortId: p.Role().ShortId()})
	if yes, code := common.IsErr(v, err); yes {
		return code
	}

	p.Room().SetRoomInfo(nil)
	p.UpdateInfoToRedis()
	return &outer.LeaveRoomRsp{}
})

// 准备、取消准备
var _ = router.Reg(func(p *player.Player, msg *outer.ReadyReq) any {
	if p.Room().RoomId() == 0 {
		return outer.ERROR_PLAYER_NOT_IN_ROOM
	}

	roomActor := actortype.RoomName(p.Room().RoomId())
	v, err := p.RequestWait(roomActor, &inner.ReadyReq{
		ShortId: p.Role().ShortId(),
		Ready:   msg.Ready,
	})
	if yes, code := common.IsErr(v, err); yes {
		return code
	}

	return &outer.ReadyRsp{Ready: msg.Ready}
})

// 转发所有Client游戏消息至房间
var _ = router.Reg(func(p *player.Player, msg *inner.GamblingMsgToRoomWrapper) any {
	if p.Room().RoomId() == 0 {
		return outer.ERROR_PLAYER_NOT_IN_ROOM
	}

	msgId, ok := p.System().ProtoIndex().MsgNameToId(msg.MsgType)
	if !ok {
		log.Warnw("MsgGamblingMsgToClientWrapper msg name to id failed", "player", p.RID(), "roomId", p.Room().RoomId(), "msg", msg.String())
		return nil
	}
	outerMsg := p.System().ProtoIndex().UnmarshalPbMsg(msgId, msg.Data)

	roomActor := actortype.RoomName(p.Room().RoomId())
	if err := p.Send(roomActor, outerMsg); err != nil {
		log.Warnw("GamblingMsgToRoomWrapper msg send to room failed", "player", p.RID(), "roomId", p.Room().RoomId(), "msg", msg.String())
	}
	return nil
})

// 转发所有房间游戏消息至client
var _ = router.Reg(func(p *player.Player, msg *inner.GamblingMsgToClientWrapper) any {
	if p.GateSession().Invalid() {
		return nil
	}

	msgId, ok := p.System().ProtoIndex().MsgNameToId(msg.MsgType)
	if !ok {
		log.Warnw("MsgGamblingMsgToClientWrapper msg name to id failed", "msg", msg.String())
		return nil
	}
	outerMsg := p.System().ProtoIndex().UnmarshalPbMsg(msgId, msg.Data)
	p.GateSession().SendToClient(p, outerMsg)
	return nil
})
