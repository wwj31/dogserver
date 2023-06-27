package mahjong

import (
	"math/rand"
	"sort"
)

// RandomCards 获得洗好的一副新牌
func RandomCards() Cards {
	cards := cards108
	tail := len(cards)
	for i := 0; i < len(cards); i++ {
		idx := rand.Intn(len(cards[:tail]))
		cards[idx], cards[tail-1] = cards[tail-1], cards[idx]
	}
	return cards[:]
}

func (c Cards) Sort() {
	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })
}

func (c Cards) Len() int {
	return len(c)
}

func (c Cards) ToInt32() (result []int32) {
	for _, v := range c {
		result = append(result, v.Int32())
	}
	return
}

func (c Cards) Insert(cards ...Card) Cards {
	src := c
	var dst Cards
	for _, card := range cards {
		dst = Cards{}
		index := sort.Search(src.Len(), func(i int) bool { return src[i] > card })
		dst = append(dst, src[:index]...)
		dst = append(dst, card)
		dst = append(dst, src[index:]...)
		src = dst
	}
	return dst
}

func (c Cards) IsTing() bool {
	// TODO
	return false
}

func (c Cards) IsHu() {
	// TODO
	return
}
