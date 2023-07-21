package mahjong

import (
	"time"

	"server/proto/outermsg/outer"

	"github.com/wwj31/dogactor/tools"
)

// 结算状态

type StateSettlement struct {
	*Mahjong
}

func (s *StateSettlement) State() int {
	return Settlement
}

func (s *StateSettlement) Enter() {
	s.gameCount++
	if s.isNoHu() {
		// TODO 流局，查大叫
	} else {

	}

	// 分都算完了，就可以清理数据了
	s.clear()
	s.nextMasterIndex() // 计算下一局庄家

	// 结算给个短暂的时间
	s.currentStateEndAt = tools.Now().Add(SettlementDuration)
	s.room.AddTimer(tools.XUID(), s.currentStateEndAt, func(dt time.Duration) {
		s.SwitchTo(Ready)
	})

	s.Log().Infow("[Mahjong] enter state settlement",
		"room", s.room.RoomId, "master", s.masterIndex)
}

func (s *StateSettlement) Leave() {
	s.Log().Infow("[Mahjong] leave state settlement", "room", s.room.RoomId)
}

func (s *StateSettlement) Handle(shortId int64, v any) (result any) {
	return outer.ERROR_MAHJONG_STATE_MSG_INVALID
}

func (s *StateSettlement) nextMasterIndex() {
	if s.mutilHuByIndex != -1 {
		s.masterIndex = s.mutilHuByIndex
	} else if len(s.huSeat) > 0 {
		s.masterIndex = s.huSeat[0]
	} else {
		s.masterIndex = s.nextSeatIndexWithoutHu(s.masterIndex)
	}
}

// 检查是否流局
func (s *StateSettlement) isNoHu() bool {
	for _, player := range s.mahjongPlayers {
		if player.hu != HuInvalid {
			return false
		}
	}
	return true
}
