package niuniu

import (
	"context"
	"encoding/json"
	"reflect"
	"time"

	"server/common/rds"

	"github.com/wwj31/dogactor/tools"

	"server/proto/outermsg/outer"
)

// 发牌状态

type StateDeal struct {
	*NiuNiu
}

func (s *StateDeal) State() int {
	return Deal
}

const handCardsSize = 5

func (s *StateDeal) Enter() {
	cards := RandomPokerCards(nil)
	testCardsStr := rds.Ins.Get(context.Background(), "niuniu_testcards").Val()
	if testCardsStr != "" {
		testCards := PokerCards{}
		_ = json.Unmarshal([]byte(testCardsStr), &testCards)
		if len(testCards) > 0 {
			cards = testCards
		}
	}

	s.Log().Infow("[NiuNiu] enter state deal", "room", s.room.RoomId,
		"params", *s.gameParams(), "cards", cards, "master", s.masterIndex)

	var i int
	for _, player := range s.niuniuPlayers {
		player.handCards = append(PokerCards{}, cards[i:i+handCardsSize]...).Sort()
		i += handCardsSize

		s.room.SendToPlayer(player.ShortId, &outer.FasterRunDealNtf{
			HandCards:  player.handCards.ToPB(),
			MasterSeat: int32(s.masterIndex),
		})
	}

	// 发牌动画后，进入下个状态
	s.currentStateEndAt = tools.Now().Add(DealExpiration)
	s.room.AddTimer(tools.XUID(), s.currentStateEndAt, func(dt time.Duration) {
		s.SwitchTo(DecideMaster)
	})
}

func (s *StateDeal) Leave() {
	for seatIndex, player := range s.niuniuPlayers {
		s.Log().Infow("dealing", "room", s.room.RoomId,
			"seat", seatIndex, "player", player.ShortId, "score", player.score,
			"cards", player.handCards)
	}
	s.Log().Infow("[NiuNiu] leave state deal", "room", s.room.RoomId)
}

func (s *StateDeal) Handle(shortId int64, v any) (result any) {
	s.Log().Warnw("deal not handle any msg", "msg", reflect.TypeOf(v).String())
	return outer.ERROR_MAHJONG_STATE_MSG_INVALID
}
