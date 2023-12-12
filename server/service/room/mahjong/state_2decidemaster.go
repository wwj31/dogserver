package mahjong

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/wwj31/dogactor/tools"

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

	s.dices[0] = rand.Int31n(6) + 1
	s.dices[1] = rand.Int31n(6) + 1

	s.masterIndex = rand.Intn(s.playerNumber())
	//s.masterIndex = 0

	// 广播定庄 庄家和骰子
	s.room.Broadcast(&outer.MahjongBTEDecideMasterNtf{
		Dices:       s.dices[:],
		MasterIndex: int32(s.masterIndex),
	})

	// 10秒播完动画后，切入换发牌
	s.currentStateEndAt = tools.Now().Add(DecideMasterShowDuration)
	s.room.AddTimer(tools.XUID(), s.currentStateEndAt, func(dt time.Duration) {
		s.SwitchTo(Deal)
	})

	s.Log().Infow("[Mahjong] enter state  decide master",
		"room", s.room.RoomId, "dices", s.dices, "master", s.masterIndex)
}

func (s *StateDecideMaster) Leave() {
	s.Log().Infow("[Mahjong] leave state decide master", "room", s.room.RoomId)
}

func (s *StateDecideMaster) Handle(shortId int64, v any) (result any) {
	s.Log().Warnw("decide master not handle any msg", "msg", reflect.TypeOf(v).String())
	return outer.ERROR_MAHJONG_STATE_MSG_INVALID
}
