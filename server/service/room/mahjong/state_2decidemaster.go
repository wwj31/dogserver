package mahjong

import (
	"math/rand"
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/proto/outermsg/outer"
)

// 定庄状态

type StateDecideMaster struct {
	*Mahjong
}

func (s *StateDecideMaster) State() int {
	return DecideMaster
}

func (s *StateDecideMaster) Enter() {
	s.room.Dices[0] = rand.Int31n(6) + 1
	s.room.Dices[1] = rand.Int31n(6) + 1

	s.masterIndex = rand.Intn(maxNum)

	// 广播定庄 庄家和骰子
	s.room.Broadcast(&outer.MahjongBTEDecideMasterNtf{
		Dices:       s.room.Dices,
		MasterIndex: int32(s.masterIndex),
	})

	// 10秒播完动画后，切入换发牌
	s.room.AddTimer(tools.XUID(), tools.Now().Add(DecideMasterShowDuration), func(dt time.Duration) {
		s.SwitchTo(Deal)
	})

	log.Infow("[Mahjong] enter state  decide master",
		"room", s.room.RoomId, "dices", s.room.Dices, "master", s.masterIndex)
}

func (s *StateDecideMaster) Leave() {
	log.Infow("[Mahjong] leave state decide master", "room", s.room.RoomId)
}

func (s *StateDecideMaster) Handle(shortId int64, v any) (result any) {
	return nil
}
