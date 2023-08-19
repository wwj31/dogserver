package player

import (
	"reflect"
	"sync"
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/proto"

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

	if p.Role().Gold() <= p.Role().GoldLine() {
		return outer.ERROR_ROOM_CANNOT_ENTER_WITH_GOLD_LINE
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

	return &outer.JoinRoomRsp{
		Room:         roomInfo,
		GamblingData: joinRoomRsp.GamblingData,
	}
})

// 离开房间
var _ = router.Reg(func(p *player.Player, msg *outer.LeaveRoomReq) any {
	if p.Room().RoomId() == 0 {
		return outer.ERROR_PLAYER_NOT_IN_ROOM
	}

	roomActor := actortype.RoomName(p.Room().RoomId())
	v, err := p.RequestWait(roomActor, &inner.LeaveRoomReq{ShortId: p.Role().ShortId()})
	if yes, code := common.IsErr(v, err); yes {
		return code
	}

	p.Room().SetRoomInfo(nil)
	p.UpdateInfoToRedis()
	return &outer.LeaveRoomRsp{}
})

// 被房间主动踢出
var _ = router.Reg(func(p *player.Player, msg *inner.RoomKickOutNtf) any {
	if p.Room().RoomId() != msg.RoomId {
		log.Warnw("player kick out room unexception", "player room info", p.Room().RoomId(), "msg", msg.String())
		return nil
	}

	p.Room().SetRoomInfo(nil)
	return nil
})

// 房间通知结算后修改金币
var _ = router.Reg(func(p *player.Player, msg *inner.ModifyGoldReq) any {
	addGold := msg.Gold
	if msg.SetOrAdd {
		addGold = msg.Gold - p.Role().Gold()
	}

	p.Role().AddGold(addGold)
	return &inner.ModifyGoldRsp{Info: p.PlayerInfo()}
})

// 转发所有Client游戏消息至房间
var _ = router.Reg(func(p *player.Player, msg *inner.GamblingMsgToRoomWrapper) any {
	if p.Room().RoomId() == 0 {
		return outer.ERROR_PLAYER_NOT_IN_ROOM
	}

	roomActor := actortype.RoomName(p.Room().RoomId())
	v, err := p.RequestWait(roomActor, msg)
	if yes, code := common.IsErr(v, err); yes {
		return code
	}

	toClientWrapper, ok := v.(*inner.GamblingMsgToClientWrapper)
	if !ok {
		log.Errorw("msg should be GamblingMsgToClientWrapper", "type", reflect.TypeOf(v).String())
		return nil
	}

	rsp := handlerGamblingMsgToClientWrapper(p, toClientWrapper)
	return rsp
})

// 转发所有房间游戏消息至client(主要用于处理Ntf类消息)
var _ = router.Reg(func(p *player.Player, msg *inner.GamblingMsgToClientWrapper) any {
	outerMsg := handlerGamblingMsgToClientWrapper(p, msg)
	p.GateSession().SendToClient(p, outerMsg)
	return nil
})

func handlerGamblingMsgToClientWrapper(p *player.Player, msg *inner.GamblingMsgToClientWrapper) proto.Message {
	if p.GateSession().Invalid() {
		return nil
	}

	msgId, ok := p.System().ProtoIndex().MsgNameToId(msg.MsgType)
	if !ok {
		log.Warnw("MsgGamblingMsgToClientWrapper msg name to id failed", "msg", msg.String())
		return nil
	}

	outerMsg := p.System().ProtoIndex().UnmarshalPbMsg(msgId, msg.Data)
	return outerMsg
}
