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
