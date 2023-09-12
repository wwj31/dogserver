package mahjong_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"server/service/room/fasterrun"

	"github.com/stretchr/testify/assert"

	"server/common/rds"

	"server/service/room/mahjong"
)

func TestComb(t *testing.T) {
	cards := fasterrun.PokerCards{fasterrun.Diamonds_3, fasterrun.Diamonds_4, fasterrun.Clubs_5, fasterrun.Diamonds_6, fasterrun.Hearts_7}
	v := cards.Combination(2)
	fmt.Println(v)
}

func TestCardsTest(t *testing.T) {
	err := rds.NewBuilder().DB(1).OnConnect(func() {
		fmt.Println("redis connect success")
	}).Connect()
	assert.Nilf(t, err, "connect redis err:%v", err)

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
