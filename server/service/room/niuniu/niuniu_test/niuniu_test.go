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
		niuniu.Spades_A,
		niuniu.Hearts_7,
		niuniu.Diamonds_8,
		niuniu.Hearts_Q,
		niuniu.Joker2,
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

func TestBug(t *testing.T) {
	cards := niuniu.PokerCards{
		205, 305, 106, 206, 306, 408, 409, 213, 313,
	}

	dst := cards.AnalyzeCards(&params)
	fmt.Println(dst)

}

func TestInsteadJoker(t *testing.T) {
	cards := niuniu.PokerCards{
		niuniu.Spades_A,
		niuniu.Hearts_7,
		niuniu.Diamonds_8,
		niuniu.Joker1,
		niuniu.Joker2,
	}
	dst := cards.AnalyzeCards(&params)
	fmt.Println(dst)

	ret := insteadJoker(cards, dst)
	fmt.Println(ret)
}

func insteadJoker(handCards niuniu.PokerCards, cardsGroup niuniu.CardsGroup) (result niuniu.CardsGroup) {
	result = cardsGroup
	var jokers niuniu.PokerCards
	for _, card := range handCards {
		if card.Color() == niuniu.Joker {
			jokers = append(jokers, card)
		}
	}

	if len(jokers) == 0 {
		return
	}

	var dmpCards niuniu.PokerCards
	// 找到dmpCards中存在的一张牌，并且删除
	dmpDelete := func(targetCard niuniu.PokerCard) bool {
		for i, card := range dmpCards {
			if card == targetCard {
				dmpCards = append(dmpCards[:i], dmpCards[i+1:]...)
				return true
			}
		}
		return false
	}

	copy(dmpCards, handCards)
	for i, card := range result.Cards {
		if !dmpDelete(card) {
			result.Cards[i] = jokers[0]
			jokers = jokers[1:]
			if len(jokers) == 0 {
				return
			}
		}
	}

	copy(dmpCards, handCards)
	for i, card := range result.SideCards {
		if !dmpDelete(card) {
			result.SideCards[i] = jokers[0]
			jokers = jokers[1:]
			if len(jokers) == 0 {
				return
			}
		}
	}

	return
}
