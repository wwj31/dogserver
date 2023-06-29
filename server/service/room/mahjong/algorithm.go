package mahjong

import (
	"fmt"
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

// Insert 插入一组牌
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

// CanPong 判断当前牌组能不能碰
func (c Cards) CanPong(card Card) bool {
	var num int
	for _, handCard := range c {
		if handCard == card {
			num++
			if num == 2 {
				return true
			}
		}
	}
	return false
}

// Pong 碰,返回去除了碰牌后的新手牌，以及碰牌起始下标
func (c Cards) Pong(card Card) (cards Cards, index int, err error) {
	index = sort.Search(c.Len(), func(i int) bool { return c[i] >= card })
	if c[index] != card {
		err = fmt.Errorf("pong failed cannot find index cards:%v index:%v card:%v", c, index, card)
		return
	}

	// 找到后把左边一张牌和当前牌，作为碰牌
	// 检查左边那张牌必须相同
	index--
	if c[index] != card {
		err = fmt.Errorf("pong failed card number != 2:%v index:%v card:%v", c, index, card)
		return
	}

	cards = append(cards, c[:index]...)
	cards = append(cards, c[index+2:]...)
	return
}

// CanGang 判断当前牌组能不能杠
func (c Cards) CanGang(card Card) bool {
	var num int
	for _, handCard := range c {
		if handCard == card {
			num++
			if num == 3 {
				return true
			}
		}
	}
	return false
}

// Gang 杠,返回去除了杠牌后的新手牌，以及杠的起始下标
func (c Cards) Gang(card Card) (cards Cards, index int, err error) {
	index = sort.Search(c.Len(), func(i int) bool { return c[i] >= card })
	if c[index] != card {
		err = fmt.Errorf("gang failed cannot find index cards:%v index:%v card:%v", c, index, card)
		return
	}

	// 找到后把左边两张牌和当前牌，作为杠
	// 检查左边两张牌必须相同
	index -= 2
	if c[index] != card || c[index+1] != card {
		err = fmt.Errorf("gang failed card number != 2:%v index:%v card:%v", c, index, card)
		return
	}

	cards = append(cards, c[:index]...)
	cards = append(cards, c[index+3:]...)
	return
}

func (c Cards) IsTing() bool {
	// TODO
	return false
}

func (c Cards) IsHu() {
	// TODO
	return
}
