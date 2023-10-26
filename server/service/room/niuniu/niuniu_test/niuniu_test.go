package niuniu_test

import (
	"fmt"
	"testing"

	"server/proto/outermsg/outer"
	"server/service/room/niuniu"
)

var params = outer.NiuNiuParams{
	SpecialStraightNiu:   true,
	SpecialFiveColorNiu:  true,
	SpecialSameColorNiu:  true,
	SpecialHuluNiu:       true,
	SpecialBombNiu:       true,
	SpecialFiveSmallNiu:  true,
	SpecialColorStraight: true,
	LaiZi:                true,
	AllowScoreSmallZero:  true,
	ReBate:               nil,
}

func TestNiuNiuCardsType(t *testing.T) {
	cards := niuniu.PokerCards{
		niuniu.Spades_K,
		niuniu.Spades_Q,
		niuniu.Clubs_10,
		niuniu.Hearts_J,
		niuniu.Spades_A,
	}
	dst := cards.AnalyzeCards(&params)
	fmt.Println(dst)

	cards = niuniu.PokerCards{
		niuniu.Spades_A,
		niuniu.Spades_2,
		niuniu.Spades_3,
		niuniu.Clubs_4,
		niuniu.Hearts_5,
	}
	dst = cards.AnalyzeCards(&params)
	fmt.Println(dst)

	// 1 2 3 5 7  127 35 235 17
	cards = niuniu.PokerCards{
		niuniu.Spades_A,
		niuniu.Spades_5,
		niuniu.Spades_3,
		niuniu.Clubs_7,
		niuniu.Hearts_2,
	}
	dst = cards.AnalyzeCards(&params)
	fmt.Println(dst)

	// 1 2 3 5 7  127 35 235 17
	cards = niuniu.PokerCards{
		niuniu.Spades_2,
		niuniu.Spades_A,
		niuniu.Spades_3,
		niuniu.Clubs_6,
		niuniu.Hearts_2,
	}
	dst = cards.AnalyzeCards(&params)
	fmt.Println(dst)
}
