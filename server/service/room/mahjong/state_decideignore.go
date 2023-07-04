package mahjong

import (
	"math/rand"
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/proto/outermsg/outer"
)

// 定缺状态

type StateDecideIgnore struct {
	*Mahjong
}

func (s *StateDecideIgnore) State() int {
	return DecideIgnore
}

func (s *StateDecideIgnore) Enter() {
	s.room.Broadcast(&outer.MahjongBTEDecideIgnoreNtf{})
	s.room.AddTimer(tools.XUID(), tools.Now().Add(15*time.Second), func(time.Duration) {
		s.stateEnd()
	})
	log.Infow("[Mahjong] enter state decide ignore", "room", s.room.RoomId)
}

func (s *StateDecideIgnore) Leave() {
	log.Infow("[Mahjong] leave state decide ignore", "room", s.room.RoomId)
}

func (s *StateDecideIgnore) Handle(shortId int64, v any) (result any) {
	switch msg := v.(type) {
	case *outer.MahjongBTEDecideIgnoreReq:
		if msg.Color == 0 {
			return outer.ERROR_MSG_REQ_PARAM_INVALID
		}
		player := s.findMahjongPlayer(shortId)
		if player == nil {
			return outer.ERROR_ROOM_PLAYER_NOT_IN_GAME
		}

		player.ignoreColor = ColorType(msg.Color)
		log.Infow("MahjongBTEDecideIgnoreReq", "roomId", s.room.RoomId,
			"player", player.ShortId, "ignore color", player.ignoreColor)

		if s.isAllDecide() {
			s.stateEnd()
		}
		return &outer.MahjongBTEDecideIgnoreRsp{}
	}
	return nil
}

// 定缺完成
func (s *StateDecideIgnore) stateEnd() {
	colorMap := map[int64]outer.ColorType{}
	for _, player := range s.mahjongPlayers {
		if player.ignoreColor == ColorUnknown {
			player.ignoreColor = ColorType(rand.Int31n(3) + 1)
		}

		colorMap[player.ShortId] = player.ignoreColor.PB()
	}
	s.room.Broadcast(&outer.MahjongBTEDecideIgnoreEndNtf{Colors: colorMap})

	// 定缺结束，给个几秒播动画
	s.room.AddTimer(tools.XUID(), tools.Now().Add(15*time.Second), func(time.Duration) {
		s.SwitchTo(Playing)
	})
}

// 是否所有玩家都定缺完成
func (s *StateDecideIgnore) isAllDecide() bool {
	for _, player := range s.mahjongPlayers {
		if player.ignoreColor == ColorUnknown {
			return false
		}
	}
	return true
}
