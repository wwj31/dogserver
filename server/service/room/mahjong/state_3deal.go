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
	s.room.GameRecordingStart()
	var (
		initCards    Cards
		baseCardsNum int
	)

	switch s.gameParams().GameMode {
	case 0, 1:
		initCards = cards108[:]
		baseCardsNum = 13
	case 2, 3:
		initCards = cards72[:]
		baseCardsNum = 13
	case 4:
		initCards = cards36[:]
		baseCardsNum = 7 // 两人一房，闲家7张牌，庄家8张牌
	}

	s.cards = initCards.RandomCards(nil)
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

	var start, last int // start 玩家拿牌的起始位置, last 庄家最后一张牌的拿牌位置
	for _, player := range s.mahjongPlayers {
		player.handCards = append(Cards{}, s.cards[start:start+baseCardsNum]...).Sort()
		start += baseCardsNum
		last += baseCardsNum
	}

	// 庄家多发一张
	master := s.mahjongPlayers[s.masterIndex]
	master.handCards = master.handCards.Insert(s.cards[last])
	s.masterCard14 = s.cards[last]

	for _, player := range s.mahjongPlayers {
		s.room.SendToPlayer(player.ShortId, &outer.MahjongBTEDealNtf{
			Cards:      player.handCards.ToSlice(),
			MasterSeat: int32(s.masterIndex),
		})
	}

	// 剩下的算本局牌组
	s.cards = s.cards[last+1:]

	// 发牌动画后，进入下个状态
	s.currentStateEndAt = tools.Now().Add(DealShowDuration)
	s.room.AddTimer(tools.XUID(), s.currentStateEndAt, func(dt time.Duration) {
		var nextState State
		if s.room.GameParams.Mahjong.HuanSanZhang == 0 {
			if s.ignoreState() {
				nextState = DecideIgnore // 不换牌，直接定缺
			} else {
				nextState = Playing // 不换牌，直接游戏
			}
		} else {
			nextState = Exchange3
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
