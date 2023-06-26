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
