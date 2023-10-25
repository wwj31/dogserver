package niuniu

import (
	"fmt"

	"github.com/spf13/cast"

	"server/proto/outermsg/outer"
)

type (
	PokerCard      int32
	PokerColorType int32
	PokerCardsType int32
	PokerCards     []PokerCard
	CardsGroup     struct {
		Type      PokerCardsType // 牌型
		Cards     PokerCards     // 主牌
		SideCards PokerCards     // 副牌
	}
)

const (
	PokerCardsUnknown PokerCardsType = iota
	Niu1Type                         = 1  // 牛1
	Niu2Type                         = 2  // 牛2
	Niu3Type                         = 3  // 牛3
	Niu4Type                         = 4  // 牛4
	Niu5Type                         = 5  // 牛5
	Niu6Type                         = 6  // 牛6
	Niu7Type                         = 7  // 牛7
	Niu8Type                         = 8  // 牛8
	Niu9Type                         = 9  // 牛9
	NiuNiuType                       = 10 // 牛牛
	StraightNiuType                  = 11 // 顺子牛
	FiveColorNiuType                 = 12 // 五花牛
	SameColorNiuType                 = 13 // 同花牛
	HuluNiuType                      = 14 // 葫芦牛
	BombNiuType                      = 15 // 炸弹牛
	FiveSmallNiuType                 = 16 // 五小牛
	ColorStraightType                = 17 // 同花顺
)

const (
	ColorUnknown PokerColorType = iota
	Diamonds                    = 1 // 方块
	Clubs                       = 2 // 梅花
	Hearts                      = 3 // 红心
	Spades                      = 4 // 黑桃
	Joker                       = 5 // 王
)

var CardsTypeTimes = [2]map[PokerCardsType]int32{
	{
		PokerCardsUnknown: 1,
		Niu1Type:          1,
		Niu2Type:          1,
		Niu3Type:          1,
		Niu4Type:          1,
		Niu5Type:          1,
		Niu6Type:          1,
		Niu7Type:          2,
		Niu8Type:          2,
		Niu9Type:          2,
		NiuNiuType:        3,
		StraightNiuType:   4,
		FiveColorNiuType:  4,
		SameColorNiuType:  4,
		HuluNiuType:       4,
		BombNiuType:       4,
		FiveSmallNiuType:  5,
		ColorStraightType: 5,
	},
	{
		PokerCardsUnknown: 1,
		Niu1Type:          1,
		Niu2Type:          2,
		Niu3Type:          3,
		Niu4Type:          4,
		Niu5Type:          5,
		Niu6Type:          6,
		Niu7Type:          7,
		Niu8Type:          8,
		Niu9Type:          9,
		NiuNiuType:        10,
		StraightNiuType:   11,
		FiveColorNiuType:  11,
		SameColorNiuType:  12,
		HuluNiuType:       13,
		BombNiuType:       14,
		FiveSmallNiuType:  15,
		ColorStraightType: 15,
	},
}

const (
	Diamonds_A  PokerCard = 101 // 方块A
	Diamonds_2  PokerCard = 102 // 方块2
	Diamonds_3  PokerCard = 103 // 方块3
	Diamonds_4  PokerCard = 104 // 方块4
	Diamonds_5  PokerCard = 105 // 方块5
	Diamonds_6  PokerCard = 106 // 方块6
	Diamonds_7  PokerCard = 107 // 方块7
	Diamonds_8  PokerCard = 108 // 方块8
	Diamonds_9  PokerCard = 109 // 方块9
	Diamonds_10 PokerCard = 110 // 方块10
	Diamonds_J  PokerCard = 111 // 方块J
	Diamonds_Q  PokerCard = 112 // 方块Q
	Diamonds_K  PokerCard = 113 // 方块K

	Clubs_A  PokerCard = 201 // 梅花A
	Clubs_2  PokerCard = 202 // 梅花2
	Clubs_3  PokerCard = 203 // 梅花3
	Clubs_4  PokerCard = 204 // 梅花4
	Clubs_5  PokerCard = 205 // 梅花5
	Clubs_6  PokerCard = 206 // 梅花6
	Clubs_7  PokerCard = 207 // 梅花7
	Clubs_8  PokerCard = 208 // 梅花8
	Clubs_9  PokerCard = 209 // 梅花9
	Clubs_10 PokerCard = 210 // 梅花10
	Clubs_J  PokerCard = 211 // 梅花J
	Clubs_Q  PokerCard = 212 // 梅花Q
	Clubs_K  PokerCard = 213 // 梅花K

	Hearts_A  PokerCard = 301 // 红心A
	Hearts_2  PokerCard = 302 // 红心2
	Hearts_3  PokerCard = 303 // 红心3
	Hearts_4  PokerCard = 304 // 红心4
	Hearts_5  PokerCard = 305 // 红心5
	Hearts_6  PokerCard = 306 // 红心6
	Hearts_7  PokerCard = 307 // 红心7
	Hearts_8  PokerCard = 308 // 红心8
	Hearts_9  PokerCard = 309 // 红心9
	Hearts_10 PokerCard = 310 // 红心10
	Hearts_J  PokerCard = 311 // 红心J
	Hearts_Q  PokerCard = 312 // 红心Q
	Hearts_K  PokerCard = 313 // 红心K

	Spades_A  PokerCard = 401 // 黑桃A
	Spades_2  PokerCard = 402 // 黑桃2
	Spades_3  PokerCard = 403 // 黑桃3
	Spades_4  PokerCard = 404 // 黑桃4
	Spades_5  PokerCard = 405 // 黑桃5
	Spades_6  PokerCard = 406 // 黑桃6
	Spades_7  PokerCard = 407 // 黑桃7
	Spades_8  PokerCard = 408 // 黑桃8
	Spades_9  PokerCard = 409 // 黑桃9
	Spades_10 PokerCard = 410 // 黑桃10
	Spades_J  PokerCard = 411 // 黑桃J
	Spades_Q  PokerCard = 412 // 黑桃Q
	Spades_K  PokerCard = 413 // 黑桃K

	Joker1 PokerCard = 516 // 小王
	Joker2 PokerCard = 517 // 大王
)

func (m PokerCard) String() string {
	return cast.ToString(m.Int())
}

func (m PokerCard) Color() PokerColorType {
	return PokerColorType(m / 100)
}

func (m PokerCard) Point() int32 {
	return int32(m % 100)
}

func (m PokerCard) Int32() int32 {
	return int32(m)
}

func (m PokerCard) Int() int {
	return int(m)
}

func (c CardsGroup) String() string {
	return fmt.Sprintf("type:%v cards:%v side cards:%v", c.Type, c.Cards, c.SideCards)
}

func (c CardsGroup) BiggestCard() PokerCard {
	result := c.Cards[0]

	var A PokerCard
	for i := 1; i < len(c.Cards); i++ {
		point := c.Cards[i].Point()
		if point == 16 || point == 17 {
			continue
		}

		if point == 1 {
			A = c.Cards[i]
		}

		if result.Point() < point {
			result = c.Cards[i]
		}
	}

	// 如果A作为顺子中的14，那么A最大
	if result.Point() == 13 && A.Point() == 1 && (c.Type == StraightNiuType || c.Type == ColorStraightType) {
		return A
	}

	return result
}

func (p PokerCards) ToPB() []int32 {
	result := make([]int32, 0, len(p))
	for _, card := range p {
		result = append(result, card.Int32())
	}
	return result
}

func (c PokerColorType) ToPB() outer.NiuNiuPokerColorType {
	return outer.NiuNiuPokerColorType(c)
}

func (c PokerCardsType) ToPB() outer.NiuNiuPokerCardsType {
	return outer.NiuNiuPokerCardsType(c)
}

func (c CardsGroup) ToPB() *outer.NiuNiuCardsGroup {
	if c.Type == PokerCardsUnknown {
		return nil
	}

	return &outer.NiuNiuCardsGroup{
		Type:      c.Type.ToPB(),
		Cards:     c.Cards.ToPB(),
		SideCards: c.SideCards.ToPB(),
	}
}

var pokerCards52 = []PokerCard{
	Clubs_A, Diamonds_A, Hearts_A, Spades_A,
	Clubs_2, Diamonds_2, Hearts_2, Spades_2,
	Clubs_3, Diamonds_3, Hearts_3, Spades_3,
	Clubs_4, Diamonds_4, Hearts_4, Spades_4,
	Clubs_5, Diamonds_5, Hearts_5, Spades_5,
	Clubs_6, Diamonds_6, Hearts_6, Spades_6,
	Clubs_7, Diamonds_7, Hearts_7, Spades_7,
	Clubs_8, Diamonds_8, Hearts_8, Spades_8,
	Clubs_9, Diamonds_9, Hearts_9, Spades_9,
	Clubs_10, Diamonds_10, Hearts_10, Spades_10,
	Clubs_J, Diamonds_J, Hearts_J, Spades_J,
	Clubs_Q, Diamonds_Q, Hearts_Q, Spades_Q,
	Clubs_K, Diamonds_K, Hearts_K, Spades_K,
}
