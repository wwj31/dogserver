package fasterrun

import (
	"math/rand"
	"sort"

	"server/common"
)

// RandomPokerCards 获得洗好的一副新牌
func RandomPokerCards(ignoreCard PokerCards) PokerCards {
	cards := pokerCards52
	tail := len(cards)
	for i := 0; i < len(cards); i++ {
		idx := rand.Intn(len(cards[:tail]))
		cards[idx], cards[tail-1] = cards[tail-1], cards[idx]
	}

	result := PokerCards(cards[:])
	if len(ignoreCard) > 0 {
		result = result.Remove(ignoreCard...)
	}
	return result
}

func (c PokerCards) Sort() PokerCards {
	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })
	return c
}

func (c PokerCards) Len() int {
	return len(c)
}

// Insert 插入一组牌
func (c PokerCards) Insert(cards ...PokerCard) PokerCards {
	src := c
	var dst PokerCards
	for _, card := range cards {
		dst = PokerCards{}
		index := sort.Search(src.Len(), func(i int) bool { return src[i] > card })
		dst = append(dst, src[:index]...)
		dst = append(dst, card)
		dst = append(dst, src[index:]...)
		src = dst
	}
	return dst
}

// Push 尾部追加
func (c PokerCards) Push(PokerCards ...PokerCard) PokerCards {
	return append(c, PokerCards...)
}

// Remove 移除一组牌,移除的牌必须全部在牌中
func (c PokerCards) Remove(cards ...PokerCard) PokerCards {
	cardMap := make(map[PokerCard]int) // 统计要移除的牌数量
	for _, card := range cards {
		cardMap[card] += 1
	}

	dst := make(PokerCards, 0, c.Len())
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

	return dst
}

// ColorCount 当前牌有几个花色
func (c PokerCards) colors() map[PokerColorType]struct{} {
	colorMap := make(map[PokerColorType]struct{})
	c.RangeSplit(func(color PokerColorType, number int) bool {
		colorMap[color] = struct{}{}
		return false
	})
	return colorMap
}

// colorPokerCards 获得某个花色的所有牌
func (c PokerCards) colorPokerCards(color PokerColorType) PokerCards {
	var cards PokerCards
	for _, card := range c {
		if card.Color() == color {
			cards = append(cards, card)
		}
	}

	return cards
}

func (c PokerCards) Random() PokerCard {
	if c.Len() == 0 {
		return PokerCard(0)
	}
	return c[rand.Intn(len(c))]
}

// CardIndex 返回第一张找的牌的位置
func (c PokerCards) CardIndex(dstCard PokerCard) int {
	for index, card := range c {
		if card == dstCard {
			return index
		}
	}
	return -1
}

func (c PokerCards) ToSlice() (result []int32) {
	for _, card := range c {
		result = append(result, card.Int32())
	}
	return
}

func (c PokerCards) RangeSplit(fn func(color PokerColorType, number int) bool) {
	for _, card := range c {
		if fn(card.Color(), int(card.Point())) {
			break
		}
	}
}

func (c PokerCards) HasColorCard(color PokerColorType) (has bool) {
	c.RangeSplit(func(c PokerColorType, n int) bool {
		if color == c {
			has = true
			return true
		}
		return false
	})
	return has
}

func (c PokerCards) Range(fn func(card PokerCard) bool) {
	for _, card := range c {
		if fn(card) {
			break
		}
	}
}

// Duizi 找对子
func (c PokerCards) Duizi() []PokerCards {
	var result []PokerCards
	for i := 0; i < len(c)-1; i++ {
		if c[i] == c[i+1] {
			result = append(result, PokerCards{c[i], c[i+1]})
			i++ // 跳过下一个已经配对的牌
		}
	}
	return result
}

// Kezi 找刻子
func (c PokerCards) Kezi() []PokerCards {
	var result []PokerCards
	for i := 0; i < len(c)-2; i++ {
		if c[i] == c[i+1] && c[i+1] == c[i+2] {
			result = append(result, PokerCards{c[i], c[i+1], c[i+2]})
			// 跳过所有相同的牌
			for i < len(c)-2 && c[i] == c[i+2] {
				i++
			}
		}
	}
	return result
}

// RemoveDuplicate 去除对子和顺子中重复的牌组
func RemoveDuplicate(PokerCardsGroup []PokerCards) []PokerCards {
	// 创建一个 map 来记录已经出现过的牌组
	uniqueMap := make(map[PokerCard]bool, len(PokerCardsGroup))
	var result []PokerCards

	// 对每个牌组进行处理
	for _, cards := range PokerCardsGroup {
		// 检查是否已经看到过这个牌组
		if !uniqueMap[cards[0]] {
			uniqueMap[cards[0]] = true
			result = append(result, cards)
		}
	}

	return result
}

// Shunzi 找顺子
func (c PokerCards) Shunzi() []PokerCards {
	var result []PokerCards

	// 用于保存每个牌的数量
	cardCount := make(map[PokerCard]int)
	for _, card := range c {
		cardCount[card]++
	}

	// 获取所有可能的顺子
	for card, count := range cardCount {
		if card%10 <= 7 && cardCount[card+1] > 0 && cardCount[card+2] > 0 {
			for i := 0; i < min(count, cardCount[card+1], cardCount[card+2]); i++ {
				result = append(result, PokerCards{card, card + 1, card + 2})
			}
		}
	}
	return result
}

func min(a, b, c int) int {
	return common.Min(a, common.Min(b, c))
}
