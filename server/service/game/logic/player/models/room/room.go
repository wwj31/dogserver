package room

import (
	gogo "github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/tools"
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
	data     inner.GamblingData // 游戏数据
	RoomInfo *inner.RoomInfo
}

func New(base models.Model) *Room {
	mod := &Room{Model: base}
	return mod
}

func (s *Room) Data() gogo.Message {
	return &s.data
}
func (s *Room) OnLogin(first bool, enterGameRsp *outer.EnterGameRsp) {
	var clear bool
	// 玩家重登，检查房间是否有效
	if s.RoomInfo != nil && s.RoomInfo.RoomId != 0 {
		roomActor := actortype.RoomName(s.RoomInfo.RoomId)
		v, err := s.Player.RequestWait(roomActor, &inner.RoomLoginReq{
			Player: s.Player.PlayerInfo(),
		})
		if yes, _ := common.IsErr(v, err); yes {
			log.Warnw("room invalid", "room", roomActor, "shortId", s.Player.Role().ShortId())
			clear = true
		} else {
			loginCheckRsp := v.(*inner.RoomLoginRsp)
			if loginCheckRsp.Err != 0 {
				log.Warnw("room check rsp failed",
					"shortId", s.Player.Role().ShortId(), "roomActor", roomActor, "err", loginCheckRsp.Err)
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
				"roomActor", roomActor,
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
func (s *Room) AddGamblingHistory(info *inner.HistoryInfo) {
	if s.data.History == nil {
		s.data.History = map[int32]*inner.HistoryInfos{}
	}
	if s.data.History[info.GameType] == nil {
		s.data.History[info.GameType] = &inner.HistoryInfos{List: make([]*inner.HistoryInfo, 0, 10)}
	}
	list := s.data.History[info.GameType].List
	list = append(list, info)
	expireAt := tools.NewTimeEx().BeginOfToday().Add(-3 * tools.Day)
	expireIndex := -1
	for i, historyInfo := range list {
		if historyInfo.GameStartAt < expireAt.UnixMilli() {
			expireIndex = i
			continue
		}
		break
	}
	if expireIndex >= 0 {
		list = list[expireIndex+1:]
	}
	s.data.History[info.GameType].List = list
}
