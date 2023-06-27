package mahjong_test

import (
	"fmt"
	"testing"

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
