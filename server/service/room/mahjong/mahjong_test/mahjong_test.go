package mahjong_test

import (
	"fmt"
	"server/service/room/mahjong"
	"testing"
)

func TestName(t *testing.T) {
	cards := mahjong.RandomCards()
	cards.Sort()
	fmt.Println(cards)
}
