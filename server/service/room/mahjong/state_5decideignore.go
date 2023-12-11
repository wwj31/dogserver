package mahjong

import (
	"reflect"
	"sort"
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/proto/outermsg/outer"
)

// 定缺状态

type StateDecideIgnore struct {
	*Mahjong
	colorMap map[int64]outer.ColorType
	timerId  string
}

func (s *StateDecideIgnore) State() int {
	return DecideIgnore
}

func (s *StateDecideIgnore) Enter() {
	s.colorMap = make(map[int64]outer.ColorType)
	s.currentStateEndAt = tools.Now().Add(DecideIgnoreExpiration)
	s.room.Broadcast(&outer.MahjongBTEDecideIgnoreNtf{EndAt: s.currentStateEndAt.UnixMilli()})
	s.timerId = s.room.AddTimer(tools.XUID(), s.currentStateEndAt, func(time.Duration) {
		s.stateEnd()
	})
	s.Log().Infow("[Mahjong] enter state decide ignore", "room", s.room.RoomId)
}

func (s *StateDecideIgnore) Leave() {
	s.Log().Infow("[Mahjong] leave state decide ignore", "room", s.room.RoomId, "colors", s.colorMap)
}

func (s *StateDecideIgnore) Handle(shortId int64, v any) (result any) {
	switch msg := v.(type) {
	case *outer.MahjongBTEDecideIgnoreReq:
		if msg.Color == 0 {
			return outer.ERROR_MSG_REQ_PARAM_INVALID
		}

		player, _ := s.findMahjongPlayer(shortId)
		if player == nil {
			return outer.ERROR_ROOM_PLAYER_NOT_IN_GAME
		}

		if player.trusteeship {
			return outer.ERROR_ROOM_NEED_CANCEL_TRUSTEESHIP
		}
		player.ignoreColor = ColorType(msg.Color)

		s.Log().Infow("MahjongBTEDecideIgnoreReq",
			"room", s.room.RoomId, "player", player.ShortId, "ignore color", player.ignoreColor)

		// 广播定缺确认通知
		s.room.Broadcast(&outer.MahjongBTEDecideIgnoreReadyNtf{ShortId: player.ShortId})

		// 所有人都定缺完成，就进入游戏状态
		if s.isAllDecide() {
			s.stateEnd()
		}
		return &outer.MahjongBTEDecideIgnoreRsp{}
	default:
		s.Log().Warnw("decide ignore status has received an unknown message", "msg", reflect.TypeOf(msg).String())
	}
	return outer.ERROR_MAHJONG_STATE_MSG_INVALID
}

// 定缺完成
func (s *StateDecideIgnore) stateEnd() {
	s.room.CancelTimer(s.timerId)
	for _, player := range s.mahjongPlayers {
		if player.ignoreColor == ColorUnknown {
			TongCards := player.handCards.colorCards(Tong)
			TiaoCards := player.handCards.colorCards(Tiao)
			WanCards := player.handCards.colorCards(Wan)

			arr := []struct {
				c   ColorType
				len int
			}{
				{c: Tong, len: TongCards.Len()},
				{c: Tiao, len: TiaoCards.Len()},
				{c: Wan, len: WanCards.Len()},
			}

			sort.Slice(arr, func(i, j int) bool { return arr[i].len < arr[j].len })

			player.ignoreColor = arr[0].c
		}

		s.colorMap[player.ShortId] = player.ignoreColor.PB()
	}
	s.room.Broadcast(&outer.MahjongBTEDecideIgnoreEndNtf{Colors: s.colorMap})

	// 定缺结束，给个几秒播动画
	s.room.AddTimer(tools.XUID(), tools.Now().Add(DecideIgnoreDuration), func(time.Duration) {
		s.SwitchTo(Playing)
	})
}

// 是否所有玩家都定缺完成
func (s *StateDecideIgnore) isAllDecide() bool {
	for _, player := range s.mahjongPlayers {
		if player.ignoreColor == ColorUnknown && !player.trusteeship {
			return false
		}
	}
	return true
}
