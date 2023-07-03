package mahjong_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"server/service/room/mahjong"
)

func TestRandomCards(t *testing.T) {
	cards := mahjong.RandomCards()
	cards.Sort()
	fmt.Println(cards)
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
		// Add more test cases for different Hu types...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.IsHu(nil, nil, nil)
			assert.Equal(t, tt.want, tt.c.IsHu(nil, nil, nil))
		})
	}
}
