package mahjong

import (
	"github.com/wwj31/dogactor/tools"
	"math/rand"
	"server/common/log"
	"server/proto/outermsg/outer"
	"server/service/room"
	"time"
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

	// 10秒播完动画后，切入换三张状态
	s.room.AddTimer(tools.XUID(), tools.Now().Add(10*time.Second), func(dt time.Duration) {
		if err := s.fsm.Switch(Deal); err != nil {
			log.Errorw("enter exchange3 failed on decideMaster", "err", err)
		}
	})

	log.Infow("Mahjong enter decide master",
		"room", s.room.RoomId, "dices", s.room.Dices, "master", s.masterIndex)
}

func (s *StateDecideMaster) Leave(fsm *room.FSM) {
	log.Infow("Mahjong leave decide master", "room", s.room.RoomId)
}

func (s *StateDecideMaster) Handle(fsm *room.FSM, v any, shortId int64) (result any) {
	return nil
}
