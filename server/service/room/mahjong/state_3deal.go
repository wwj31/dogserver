package mahjong

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
	*Mahjong
}

func (s *StateDeal) State() int {
	return Deal
}

func (s *StateDeal) Enter() {
	s.cards = RandomCards(nil) // 总共108张
	s.cards = Cards{
		11, 11, 11, 28, 28, 28, 13, 13, 27, 28, 29, 15, 15,
		11, 21, 12, 25, 26, 27, 21, 22, 34, 34, 34, 26, 32,
		12, 22, 14, 14, 14, 24, 24, 24, 29, 13, 12, 23, 25,
		35, 36, 37, 38, 39, 21, 23, 33, 25, 26, 27, 29, 31,
		15,
		16, 38, 14, 16, 16, 17, 23, 18, 18, 18, 32, 19, 33,
		19, 19, 31, 31, 32, 16, 33, 37, 18, 35, 36, 17, 35,
		38, 37, 38, 39, 39, 21, 17, 22, 36, 17, 35, 23, 33,
		31, 25, 36, 13, 26, 12, 37, 32, 27, 24, 39, 19, 22, 34, 15, 29}
	testCardsStr := rds.Ins.Get(context.Background(), "testcards").Val()
	if testCardsStr != "" {
		testCards := Cards{}
		_ = json.Unmarshal([]byte(testCardsStr), &testCards)
		if testCards.Len() > 0 {
			s.cards = testCards
		}
	}
	s.Log().Infow("[Mahjong] enter state deal",
		"room", s.room.RoomId, "params", *s.gameParams(), "cards", s.cards)

	var i int
	for _, player := range s.mahjongPlayers {
		player.handCards = append(Cards{}, s.cards[i:i+13]...).Sort()
		i += 13

		s.room.SendToPlayer(player.ShortId, &outer.MahjongBTEDealNtf{
			Cards:      player.handCards.ToSlice(),
			MasterSeat: int32(s.masterIndex),
		})
	}

	// 庄家多发一张
	master := s.mahjongPlayers[s.masterIndex]
	master.handCards = master.handCards.Insert(s.cards[52])
	s.masterCard14 = s.cards[52]

	// 剩下的算本局牌组
	s.cards = s.cards[53:]

	// 发牌动画后，进入下个状态
	s.currentStateEndAt = tools.Now().Add(DealShowDuration)
	s.room.AddTimer(tools.XUID(), s.currentStateEndAt, func(dt time.Duration) {
		var nextState State
		if s.room.GameParams.Mahjong.HuanSanZhang == 2 {
			nextState = DecideIgnore // 不换牌，直接定缺
		} else {
			nextState = DecideIgnore
		}

		s.SwitchTo(nextState)
	})

	s.Log().Infow("deal finished cards", "room", s.room.RoomId, "spare cards", s.cards)
}

func (s *StateDeal) Leave() {
	for seatIndex, player := range s.mahjongPlayers {
		s.Log().Infow("dealing",
			"room", s.room.RoomId,
			"seat", seatIndex, "player", player.ShortId, "score", player.score,
			"cards", player.handCards)
	}
	s.Log().Infow("[Mahjong] leave state deal", "room", s.room.RoomId)
}

func (s *StateDeal) Handle(shortId int64, v any) (result any) {
	s.Log().Warnw("deal not handle any msg", "msg", reflect.TypeOf(v).String())
	return outer.ERROR_MAHJONG_STATE_MSG_INVALID
}
