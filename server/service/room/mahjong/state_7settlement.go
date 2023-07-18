package mahjong

import (
	"github.com/wwj31/dogactor/tools"
	"server/common/log"
	"time"
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

	log.Infow("[Mahjong] enter state settlement",
		"room", s.room.RoomId, "master", s.masterIndex)
}

func (s *StateSettlement) Leave() {
	log.Infow("[Mahjong] leave state settlement", "room", s.room.RoomId)
}

func (s *StateSettlement) Handle(shortId int64, v any) (result any) {
	return nil
}

func (s *StateSettlement) nextMasterIndex() {
	if s.mutilHuByIndex != -1 {
		s.masterIndex = s.mutilHuByIndex
	} else if s.firstHuIndex != -1 {
		s.masterIndex = s.firstHuIndex
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
