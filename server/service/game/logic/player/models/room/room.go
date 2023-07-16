package room

import (
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/proto/convert"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player/models"
)

type Room struct {
	models.Model
	RoomInfo *inner.RoomInfo
}

func New(base models.Model) *Room {
	mod := &Room{Model: base}
	return mod
}

func (s *Room) OnLogin(first bool, enterGameRsp *outer.EnterGameRsp) {
	if first {

	}

	var clear bool
	// 玩家重登，检查房间是否有效
	if s.RoomInfo != nil && s.RoomInfo.RoomId != 0 {
		roomActor := actortype.RoomName(s.RoomInfo.RoomId)
		v, err := s.Player.RequestWait(roomActor, &inner.RoomLoginReq{
			Player: s.Player.PlayerInfo(),
		})
		if yes, _ := common.IsErr(v, err); yes {
			log.Warnw("room invalid",
				"shortId", s.Player.Role().ShortId(), "roomId", roomActor)
			clear = true
		} else {
			loginCheckRsp := v.(*inner.RoomLoginRsp)
			if loginCheckRsp.Err != 0 {
				log.Warnw("room check rsp failed",
					"shortId", s.Player.Role().ShortId(), "roomId", roomActor, "err", loginCheckRsp.Err)
				clear = true
			} else {
				s.RoomInfo = loginCheckRsp.RoomInfo
				enterGameRsp.GamblingData = loginCheckRsp.GetGamblingData()
			}
		}

		if clear {
			s.RoomInfo = nil
		}
	}

	enterGameRsp.RoomInfo = convert.RoomInfoInnerToOuter(s.RoomInfo)
}

func (s *Room) OnLogout() {
	if s.RoomInfo != nil && s.RoomInfo.RoomId != 0 {
		roomActor := actortype.RoomName(s.RoomInfo.RoomId)
		err := s.Player.Send(roomActor, &inner.RoomLogoutReq{ShortId: s.Player.Role().ShortId()})
		if err != nil {
			log.Warnw("logout room rsp failed",
				"shortId", s.Player.Role().ShortId(),
				"roomId", roomActor,
			)
		}
	}
}

func (s *Room) RoomId() int64 {
	if s.RoomInfo == nil {
		return 0
	}
	return s.RoomInfo.RoomId
}
func (s *Room) SetRoomInfo(info *inner.RoomInfo) { s.RoomInfo = info }
