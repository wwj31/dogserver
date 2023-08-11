package mahjong_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"server/common/rds"
	"server/proto/outermsg/outer"

	"github.com/stretchr/testify/assert"

	"server/service/room/mahjong"
)

func TestRandomCards(t *testing.T) {
	cards := mahjong.RandomCards(nil)
	cards.Sort()
	fmt.Println(cards)
}

func TestIgnoreHandCards(t *testing.T) {
	cards := mahjong.RandomCards(nil)
	handCards := mahjong.Cards{
		11, 11, 11, 28, 28, 28, 13, 13, 27, 28, 29, 15, 15,
		11, 21, 12, 25, 26, 27, 21, 22, 34, 34, 34, 26, 32,
		12, 22, 14, 14, 14, 24, 24, 24, 29, 13, 12, 23, 25,
		35, 36, 37, 38, 39, 21, 23, 33, 25, 26, 27, 29, 31,
		15}

	handCards = handCards.Push(cards.Remove(handCards...)...)
	fmt.Println(handCards, len(handCards))
}

func TestName(t *testing.T) {
	cards := mahjong.Cards{11, 12, 14, 14, 16, 17, 18, 19}
	fmt.Println(cards.Insert(14))
	fmt.Println(cards.Insert(13))
	fmt.Println(cards.Insert(13, 14))
	fmt.Println(cards.Insert(11, 15, 19))
}

func TestDuizi(t *testing.T) {
	cards := mahjong.Cards{11, 11, 11, 12, 12, 12, 18, 19}
	duizi := cards.Duizi()
	fmt.Println(duizi)

	var total []mahjong.Card
	for _, c := range duizi {
		total = append(total, c...)
	}
	cards2 := cards.Remove(total...)
	fmt.Println(cards2)
}

func TestKezi(t *testing.T) {
	cards := mahjong.Cards{11, 11, 11, 12, 12, 12, 18, 19}
	kezi := cards.Kezi()
	fmt.Println(kezi)

	var total []mahjong.Card
	for _, c := range kezi {
		total = append(total, c...)
	}
	cards2 := cards.Remove(total...)
	fmt.Println(cards2)
}

func TestShunzi(t *testing.T) {
	cards := mahjong.Cards{11, 11, 12, 13, 13, 13, 14, 17, 18, 19}
	shunzi := cards.Shunzi()
	fmt.Println(shunzi)

	var total []mahjong.Card
	for _, c := range shunzi {
		total = append(total, c...)
	}
	cards2 := cards.Remove(total...)
	fmt.Println(cards2)
}

func TestHighCard(t *testing.T) {
	cards := mahjong.Cards{11, 12, 13, 14, 15}
	assert.False(t, cards.HighCard(cards.ConvertStruct()))

	cards = mahjong.Cards{11, 13, 14, 15, 17}
	assert.True(t, cards.HighCard(cards.ConvertStruct()))

	cards = mahjong.Cards{11, 11, 11, 12, 12, 13, 13, 14, 14, 15, 15}
	assert.False(t, cards.HighCard(cards.ConvertStruct()))

	cards = mahjong.Cards{11, 11, 11, 12, 12, 13, 14, 14, 15, 15, 17}
	assert.True(t, cards.HighCard(cards.ConvertStruct()))

	cards = mahjong.Cards{11, 11, 11, 11, 12, 12, 12, 13, 13, 13, 14, 14, 14, 15, 15, 15}
	assert.False(t, cards.HighCard(cards.ConvertStruct()))

	cards = mahjong.Cards{11, 11, 11, 11, 12, 12, 12, 13, 13, 13, 14, 14, 14, 15, 15, 17}
	assert.True(t, cards.HighCard(cards.ConvertStruct()))

	cards = mahjong.Cards{21, 23, 24, 25, 27}
	assert.True(t, cards.HighCard(cards.ConvertStruct()))

	cards = mahjong.Cards{21, 21, 21, 22, 22, 23, 23, 24, 24, 25, 25}
	assert.False(t, cards.HighCard(cards.ConvertStruct()))

	cards = mahjong.Cards{21, 21, 21, 22, 22, 23, 24, 24, 25, 25, 27}
	assert.True(t, cards.HighCard(cards.ConvertStruct()))

	cards = mahjong.Cards{21, 21, 21, 21, 22, 22, 22, 23, 23, 23, 24, 24, 24, 25, 25, 25}
	assert.False(t, cards.HighCard(cards.ConvertStruct()))

	cards = mahjong.Cards{21, 21, 21, 21, 22, 22, 22, 23, 23, 23, 24, 24, 24, 25, 25, 27}
	assert.True(t, cards.HighCard(cards.ConvertStruct()))

	cards = mahjong.Cards{11, 12, 13, 14, 15, 17, 18, 19, 31, 32, 33, 33, 34}
	assert.False(t, cards.HighCard(cards.ConvertStruct()))

	cards = mahjong.Cards{31, 32, 33, 34, 35, 37, 38, 39}
	assert.False(t, cards.HighCard(cards.ConvertStruct()))
}

func TestCards(t *testing.T) {
	cards := mahjong.Cards{
		// 发给庄家的13张牌
		11, 28, 29, 21, 12, 12, 22, 13, 13, 23, 25, 15, 22,
		// 发给第二位玩家的13张牌
		11, 11, 11, 25, 26, 27, 21, 22, 34, 34, 34, 26, 32,
		// 发给第三位玩家的13张牌
		12, 13, 14, 14, 14, 24, 24, 24, 28, 28, 27, 28, 29,
		// 发给第四位玩家的13张牌
		35, 36, 37, 38, 39, 21, 23, 33, 25, 26, 27, 28, 31,
		// 给庄家的第14张牌
		15,
	}
	randomCards := mahjong.RandomCards(cards)
	randomCards = append(cards, randomCards...)

	fmt.Println(randomCards)
	fmt.Println(len(randomCards))
}

func TestRecurCheck(t *testing.T) {
	tests := []struct {
		name string
		c    mahjong.Cards
		want mahjong.HuType
	}{
		{
			name: "Hu (平胡1)",
			c:    mahjong.Cards{11, 11, 11, 12, 13, 14, 22, 23, 24, 24, 25, 26, 27, 27},
			want: mahjong.Hu,
		},
		{
			name: "Hu (平胡2)",
			c:    mahjong.Cards{11, 12, 13, 14, 14, 14, 15, 16, 17, 21, 22, 23, 23, 23},
			want: mahjong.Hu,
		},
		{
			name: "Hu (平胡3)",
			c:    mahjong.Cards{11, 11, 11, 12, 12, 12, 13, 13, 13, 21, 22, 23, 23, 23},
			want: mahjong.Hu,
		},
		{
			name: "Hu (平胡4)",
			c:    mahjong.Cards{22, 22},
			want: mahjong.Hu,
		},
		{
			name: "Hu (平胡5)",
			c:    mahjong.Cards{11, 11, 12, 13, 14},
			want: mahjong.Hu,
		},
		{
			name: "Hu (平胡6)",
			c:    mahjong.Cards{11, 12, 13, 14, 14, 14, 24, 24, 24, 28, 28, 27, 28, 29},
			want: mahjong.MenQing,
		},
		// Add more test cases for different Hu types...
	}

	params := &outer.MahjongParams{
		YaoJiuDui:         true,
		MenQingZhongZhang: true,
		TianDiHu:          true,
		DianPaoPingHu:     true,
		JiaXinWu:          true,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.IsHu(nil, nil, nil, 11, params)
			assert.Equal(t, tt.want, tt.c.IsHu(nil, nil, nil, 11, params))
		})
	}
}

func TestCardsTest(t *testing.T) {
	err := rds.NewBuilder().OnConnect(func() {
		fmt.Println("redis connect success")
	}).Connect()
	assert.Nilf(t, err, "connect redis err:%v", err)

	rds.Ins.Set(
		context.Background(),
		"testcards",
		"[11,27,16,22,38,12,32,38,13,14,26,23,14,14,25,33,15,21,14,11,16,15,16,15,17,23,12,23,18,18,18,32,19,33,19,19,31,25,31,31,12,29,32,28,29,16,37,33,26,34,37,34,18,35,26,35,36,21,36,17,34,35,21,11,38,28,37,38,39,39,13,39,25,21,17,28,22,22,36,17,35,23,33,27,31,13,24,24,25,36,13,24,26,28,12,37,32,27,27,24,39,19,22,34,15,29,29,11]",
		-1)
	jsonString := rds.Ins.Get(context.Background(), "testcards").Val()
	jsonCards := mahjong.Cards{}
	_ = json.Unmarshal([]byte(jsonString), &jsonCards)
	cards := jsonCards
	fmt.Println(cards)
	assert.Truef(t, cards.Len() == 108, "error len:%v", len(cards))

	cards.Range(func(card mahjong.Card) bool {
		return !assert.Truef(t, 11 <= card.Int32() && card.Int32() <= 39, "invalid card", card.Int32())
	})

	stats := cards.ConvertStruct()
	for i := 11; i < 40; i++ {
		if stats[i] > 0 {
			if !assert.Truef(t, stats[i] == 4, "invalid i:%v stats:%v", i, stats[i]) {
			}
		}
	}
}
