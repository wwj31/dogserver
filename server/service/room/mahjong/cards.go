package mahjong

import (
	"fmt"
	"math/rand"
	"sort"

	"server/common"
	"server/common/log"
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

// Remove 移除一组牌,移除的牌必须全部在手牌中
func (c Cards) Remove(cards ...Card) Cards {
	cardMap := make(map[Card]int) // 统计要移除的牌数量
	for _, card := range cards {
		cardMap[card] += 1
	}

	dst := make(Cards, 0, c.Len())
	for _, card := range c {
		if cardMap[card] > 0 {
			cardMap[card] -= 1
			if cardMap[card] == 0 {
				delete(cardMap, card)
			}
			continue
		}

		if cardMap[card] == 0 {
			dst = append(dst, card)
		}
	}

	if len(cardMap) != 0 {
		log.Errorw("cards remove error handCards:%v cards:%v", c, cards)
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
	if index >= c.Len() || c[index] != card {
		err = fmt.Errorf("pong failed cannot find index cards:%v index:%v card:%v", c, index, card)
		return
	}

	// 找到后把左边一张牌和当前牌，作为碰牌
	// 检查左边那张牌必须相同
	if index == 0 || c[index-1] != card {
		err = fmt.Errorf("pong failed card number != 2:%v index:%v card:%v", c, index, card)
		return
	}

	index--
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
	if index >= c.Len() || c[index] != card {
		err = fmt.Errorf("gang failed cannot find index cards:%v index:%v card:%v", c, index, card)
		return
	}

	// 找到后把左边两张牌和当前牌，作为杠
	// 检查左边两张牌必须相同
	if index < 2 || c[index-1] != card || c[index-2] != card {
		err = fmt.Errorf("gang failed card number != 2:%v index:%v card:%v", c, index, card)
		return
	}

	index -= 2
	cards = append(cards, c[:index]...)
	cards = append(cards, c[index+3:]...)
	return
}

func (c Cards) IsTing() bool {
	// TODO
	return false
}

func (c Cards) IsHu() (typ HuType) {
	duiziGroups := c.Duizi()
	// 没有能做将的牌
	if len(duiziGroups) == 0 {
		return HuInvalid
	}

	// 先判断是不是七对
	if len(duiziGroups) == 7 {
		typ = QiDui
	} else {
		// 去个重
		duiziGroups = RemoveDuplicate(duiziGroups)

		// 挨个做将，再分析剩下的牌型
		for _, jiangCards := range duiziGroups {
			var (
				keziCards   []Cards
				shunziCards []Cards
			)

			spareHandCards := c.Remove(jiangCards...)
			for {
				if spareHandCards.Len() == 0 {
					typ = Hu
					break
				}

			}

			// 如果该对子做将没有hu,直接换下一个
			if typ == HuInvalid {
				continue
			}

			// 检查刻子和顺子的情况，判断是否升级

		}
	}

	if typ != HuInvalid {
		// TODO 检查是否升级

	}
	return typ
}

// ColorCount 当前牌有几个花色
func (c Cards) ColorCount() int {
	colorMap := make(map[int]struct{})
	for _, card := range c {
		color := int(card / 10)
		colorMap[color] = struct{}{}
	}

	return len(colorMap)
}

// HighCard 检查是否存在散牌，不能组成顺子、刻子的牌
func (c Cards) HighCard() bool {
	// TODO
	return false
}

// Duizi 找对子
func (c Cards) Duizi() []Cards {
	var result []Cards
	for i := 0; i < len(c)-1; i++ {
		if c[i] == c[i+1] {
			result = append(result, Cards{c[i], c[i+1]})
			i++ // 跳过下一个已经配对的牌
		}
	}
	return result
}

// Kezi 找刻子
func (c Cards) Kezi() []Cards {
	var result []Cards
	for i := 0; i < len(c)-2; i++ {
		if c[i] == c[i+1] && c[i+1] == c[i+2] {
			result = append(result, Cards{c[i], c[i+1], c[i+2]})
			// 跳过所有相同的牌
			for i < len(c)-2 && c[i] == c[i+2] {
				i++
			}
		}
	}
	return result
}

// RemoveDuplicate 去除对子和顺子中重复的牌组
func RemoveDuplicate(cardsGroup []Cards) []Cards {
	// 创建一个 map 来记录已经出现过的牌组
	uniqueMap := make(map[Card]bool)
	var result []Cards

	// 对每个牌组进行处理
	for _, cards := range cardsGroup {
		// 检查是否已经看到过这个牌组
		if !uniqueMap[cards[0]] {
			uniqueMap[cards[0]] = true
			result = append(result, cards)
		}
	}

	return result
}

// Shunzi 找顺子
func (c Cards) Shunzi() []Cards {
	var result []Cards

	// 用于保存每个牌的数量
	cardCount := make(map[Card]int)
	for _, card := range c {
		cardCount[card]++
	}

	// 获取所有可能的顺子
	for card, count := range cardCount {
		if card%10 <= 7 && cardCount[card+1] > 0 && cardCount[card+2] > 0 {
			for i := 0; i < min(count, cardCount[card+1], cardCount[card+2]); i++ {
				result = append(result, Cards{card, card + 1, card + 2})
			}
		}
	}
	return result
}

func min(a, b, c int) int {
	return common.Min(a, common.Min(b, c))
}
