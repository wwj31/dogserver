package mahjong

import (
	"server/common/log"
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

	s.clear()
	s.nextMasterIndex() // 计算下一局庄家
	log.Infow("[Mahjong] enter state settlement",
		"room", s.room.RoomId, "dices", s.room.Dices, "master", s.masterIndex)
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
		s.masterIndex = s.nextSeatIndex(s.masterIndex)
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
