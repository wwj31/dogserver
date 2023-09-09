package fasterrun

import (
	"github.com/spf13/cast"
)

type (
	PokerCard      int32
	PokerColorType int32
	PokerCards     []PokerCard
	PongGang       map[int32]int64
)

const (
	ColorUnknown PokerColorType = iota
	Clubs                       = 1 // 梅花
	Diamonds                    = 2 // 方块
	Hearts                      = 3 // 红心
	Spades                      = 4 // 黑桃
)

const (
	Clubs_3  PokerCard = 103 // 梅花3
	Clubs_4  PokerCard = 104 // 梅花4
	Clubs_5  PokerCard = 105 // 梅花5
	Clubs_6  PokerCard = 106 // 梅花6
	Clubs_7  PokerCard = 107 // 梅花7
	Clubs_8  PokerCard = 108 // 梅花8
	Clubs_9  PokerCard = 109 // 梅花9
	Clubs_10 PokerCard = 110 // 梅花10
	Clubs_J  PokerCard = 111 // 梅花J
	Clubs_Q  PokerCard = 112 // 梅花Q
	Clubs_K  PokerCard = 113 // 梅花K
	Clubs_A  PokerCard = 114 // 梅花A
	Clubs_2  PokerCard = 115 // 梅花2

	Diamonds_3  PokerCard = 203 // 方块3
	Diamonds_4  PokerCard = 204 // 方块4
	Diamonds_5  PokerCard = 205 // 方块5
	Diamonds_6  PokerCard = 206 // 方块6
	Diamonds_7  PokerCard = 207 // 方块7
	Diamonds_8  PokerCard = 208 // 方块8
	Diamonds_9  PokerCard = 209 // 方块9
	Diamonds_10 PokerCard = 210 // 方块10
	Diamonds_J  PokerCard = 211 // 方块J
	Diamonds_Q  PokerCard = 212 // 方块Q
	Diamonds_K  PokerCard = 213 // 方块K
	Diamonds_A  PokerCard = 214 // 方块A
	Diamonds_2  PokerCard = 215 // 方块2

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
	Hearts_A  PokerCard = 314 // 红心A
	Hearts_2  PokerCard = 315 // 红心2

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
	Spades_A  PokerCard = 414 // 黑桃A
	Spades_2  PokerCard = 415 // 黑桃2
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

func (c PokerCards) String() string {
	var result string
	result = "{"
	for i, card := range c {
		result += card.String()
		if i < len(c)-1 {
			result += ", "
		}
	}
	result += "}"
	return result
}

var pokerCards52 = [52]PokerCard{
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
