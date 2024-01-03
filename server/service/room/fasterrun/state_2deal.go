package fasterrun

import (
	"context"
	"encoding/json"
	"math/rand"
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
	s.room.GameRecordingStart()
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
	// 定庄
	s.masterIndex = s.decideMaster(s.gameParams().DecideMasterType)

	s.Log().Infow("[FasterRun] enter state deal", "room", s.room.RoomId,
		"params", *s.gameParams(), "cards", cards, "master", s.masterIndex)

	var i int
	for _, player := range s.fasterRunPlayers {
		player.handCards = append(PokerCards{}, cards[i:i+handCardsNumber]...).Sort()
		if s.gameParams().DoubleHeartsTen {
			player.doubleHearts10 = s.existHeart10(player.handCards)
		}
		i += handCardsNumber

		s.room.SendToPlayer(player.ShortId, &outer.FasterRunDealNtf{
			HandCards:  player.handCards.ToPB(),
			MasterSeat: int32(s.masterIndex),
		})
	}
	s.spareCards = cards[i:]

	// 发牌动画后，进入下个状态
	s.currentStateEndAt = tools.Now().Add(DealExpiration)
	s.room.AddTimer(tools.XUID(), s.currentStateEndAt, func(dt time.Duration) {
		s.SwitchTo(Playing)
	})
}
func (s *StateDeal) existHeart10(cards PokerCards) bool {
	for _, card := range cards {
		if card == Hearts_10 {
			return true
		}
	}
	return false
}

func (s *StateDeal) Leave() {
	for seatIndex, player := range s.fasterRunPlayers {
		s.Log().Infow("dealing", "room", s.room.RoomId,
			"seat", seatIndex, "player", player.ShortId, "score", player.score,
			"cards", player.handCards)
	}
	s.Log().Infow("[FasterRun] leave state deal", "room", s.room.RoomId)
}

func (s *StateDeal) Handle(shortId int64, v any) (result any) {
	s.Log().Warnw("deal not handle any msg", "msg", reflect.TypeOf(v).String())
	return outer.ERROR_MAHJONG_STATE_MSG_INVALID
}

func (s *StateDeal) decideMaster(mode int32) int {
	Spade3Seat := func() int {
		for seat, player := range s.fasterRunPlayers {
			for _, card := range player.handCards {
				if card == Spades_3 {
					return seat
				}
			}
		}
		return 0
	}

	switch mode {
	case 0: // 随机
		return rand.Intn(len(s.fasterRunPlayers))

	case 1: // 黑桃三 为庄 只能在三人游戏模式
		if len(s.fasterRunPlayers) == 2 {
			s.Log().Warnw("player is two can not ues mode 1")
			return s.decideMaster(0)
		}
		return Spade3Seat()

	case 2: // 赢家
		if s.gameCount == 1 {
			return s.decideMaster(0)
		}
		_, seat := s.findFasterRunPlayer(s.lastWinShortId)
		return seat

	case 3: //  首局黑桃三，之后赢家
		if s.gameCount == 1 {
			return Spade3Seat()
		}
		return s.decideMaster(2)

	default:
		s.Log().Warnw("unknown mode", "mode", mode)
		return s.decideMaster(0)
	}
}
