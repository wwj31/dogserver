package fasterrun

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
	*FasterRun
}

func (s *StateDeal) State() int {
	return Deal
}

func (s *StateDeal) Enter() {
	var ignoreCards PokerCards
	var handCardsNumber int
	if s.gameParams().CardsNumber == 0 { //
		ignoreCards = PokerCards{Hearts_A, Diamonds_A, Clubs_A, Spades_K} // 15张模式，只留黑桃A，去掉黑桃K
		handCardsNumber = 15
	} else if s.gameParams().CardsNumber == 1 {
		ignoreCards = PokerCards{Spades_A} // 16张模式，去掉黑桃A
		handCardsNumber = 16
	}

	cards := RandomPokerCards(ignoreCards)
	testCardsStr := rds.Ins.Get(context.Background(), "fasterrun_testcards").Val()
	if testCardsStr != "" {
		testCards := PokerCards{}
		_ = json.Unmarshal([]byte(testCardsStr), &testCards)
		if len(testCards) > 0 {
			cards = testCards
		}
	}
	s.Log().Infow("[FasterRun] enter state deal",
		"room", s.room.RoomId, "params", *s.gameParams(), "cards", cards)

	var i int
	for _, player := range s.fasterRunPlayers {
		player.handCards = append(PokerCards{}, cards[i:i+handCardsNumber]...).Sort()
		i += handCardsNumber

		s.room.SendToPlayer(player.ShortId, &outer.FasterRunDealNtf{
			Cards:      player.handCards.ToSlice(),
			MasterSeat: int32(s.masterIndex),
		})
	}

	// 发牌动画后，进入下个状态
	s.currentStateEndAt = tools.Now().Add(DealExpiration)
	s.room.AddTimer(tools.XUID(), s.currentStateEndAt, func(dt time.Duration) {
		s.SwitchTo(Playing)
	})
}

func (s *StateDeal) Leave() {
	for seatIndex, player := range s.fasterRunPlayers {
		s.Log().Infow("dealing",
			"room", s.room.RoomId,
			"seat", seatIndex, "player", player.ShortId, "score", player.score,
			"cards", player.handCards)
	}
	s.Log().Infow("[FasterRun] leave state deal", "room", s.room.RoomId)
}

func (s *StateDeal) Handle(shortId int64, v any) (result any) {
	s.Log().Warnw("deal not handle any msg", "msg", reflect.TypeOf(v).String())
	return outer.ERROR_MAHJONG_STATE_MSG_INVALID
}
