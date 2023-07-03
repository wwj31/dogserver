package mahjong

import (
	"math/rand"
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/proto/outermsg/outer"
	"server/service/room"
)

// 定庄状态

type StateDecideMaster struct {
	*Mahjong
}

func (s *StateDecideMaster) State() int {
	return DecideMaster
}

func (s *StateDecideMaster) Enter(fsm *room.FSM) {
	s.room.Dices[0] = rand.Int31n(6) + 1
	s.room.Dices[1] = rand.Int31n(6) + 1
	// TODO 骰子定庄规则

	s.masterIndex = 0
	s.room.Broadcast(&outer.MahjongBTEDecideMasterNtf{
		Dices:       s.room.Dices,
		MasterIndex: int32(s.masterIndex),
	})

	// 10秒播完动画后，切入换发牌
	s.room.AddTimer(tools.XUID(), tools.Now().Add(10*time.Second), func(dt time.Duration) {
		s.SwitchTo(Deal)
	})

	log.Infow("[Mahjong] leave state  decide master",
		"room", s.room.RoomId, "dices", s.room.Dices, "master", s.masterIndex)
}

func (s *StateDecideMaster) Leave(fsm *room.FSM) {
	log.Infow("[Mahjong] leave state decide master", "room", s.room.RoomId)
}

func (s *StateDecideMaster) Handle(fsm *room.FSM, v any, shortId int64) (result any) {
	return nil
}
