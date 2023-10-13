package niuniu

import (
	"math/rand"
	"sort"
)

// RandomPokerCards 获得洗好的一副新牌
func RandomPokerCards(ignoreCard []PokerCard) PokerCards {
	cards := make(PokerCards, len(pokerCards52))
	copy(cards, pokerCards52)
	tail := len(cards)
	for i := 0; i < len(cards); i++ {
		idx := rand.Intn(len(cards[:tail]))
		cards[idx], cards[tail-1] = cards[tail-1], cards[idx]
		tail--
	}

	if len(ignoreCard) > 0 {
		cards = cards.Remove(ignoreCard...)
	}
	return cards
}

// Remove 移除一组牌,移除的牌必须全部在牌中
func (p PokerCards) Remove(rmCards ...PokerCard) PokerCards {
	cardMap := make(map[PokerCard]int) // 统计要移除的牌数量
	for _, card := range rmCards {
		cardMap[card] += 1
	}

	var newCards PokerCards
	for _, card := range p {
		if cardMap[card] > 0 {
			cardMap[card]--
		} else {
			newCards = append(newCards, card)
		}
	}
	return newCards
}

// PointCards 找出点数的所有牌,n指定数量,如果要所有牌,n不传
func (p PokerCards) PointCards(point int32, n ...int) PokerCards {
	var newCards PokerCards
	var count int
	for _, card := range p {
		if card.Point() == point {
			newCards = append(newCards, card)
			count++
			if len(n) > 0 && count == n[0] {
				return newCards
			}
		}
	}
	return newCards
}

// CheckExist 检查传入牌是否全部在本牌组中
func (p PokerCards) CheckExist(cards ...PokerCard) bool {
	kv := make(map[PokerCard]struct{}, len(cards))
	for _, card := range p {
		kv[card] = struct{}{}
	}

	for _, card := range cards {
		if _, ok := kv[card]; !ok {
			return false
		}
	}
	return true
}

// ConvertStruct 转换成{point:num}
func (p PokerCards) ConvertStruct() (result map[int32]int) {
	result = map[int32]int{}
	for _, card := range p {
		result[card.Point()]++
	}
	return result
}

// ConvertOrderPoint 把点数按照升序排列在结果中
func (p PokerCards) ConvertOrderPoint(stat map[int32]int) []int32 {
	var arr []int32
	for point := range stat {
		arr = append(arr, point)
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	return arr
}

// Combination 算出需要n张牌的所有组合
func (p PokerCards) Combination(n int) (result []PokerCards) {
	if n == len(p) {
		result = append(result, p)
		return
	}
	if n > len(p) {
		return
	}

	c := PokerCards{}
	var combHelper func(cards PokerCards)
	combHelper = func(cards PokerCards) {
		for i := len(cards) - 1; i >= 0; i-- {
			c = append(c, cards[i])
			if len(c) == n {
				newCards := make(PokerCards, n)
				copy(newCards, c)
				result = append(result, newCards)
			} else {
				combHelper(cards[:i])
			}
			c = c[:len(c)-1]
		}
	}

	combHelper(p)
	return
}

func (p PokerCards) isColorSame() (same bool) {
	if len(p) <= 1 {
		return false
	}

	color := p[0].Color()
	for i := 1; i < len(p); i++ {
		if color != p[i].Color() {
			return false
		}
	}
	return true
}

func (p PokerCards) Sort() PokerCards {
	sort.Slice(p, func(i, j int) bool {
		iPoint := p[i].Point()
		jPoint := p[j].Point()
		if iPoint == jPoint {
			return p[i].Color() < p[j].Color()
		}
		return iPoint < jPoint
	})
	return p
}

func (c CardsGroup) CanCompare(group CardsGroup) bool {
	if len(c.Cards) == 0 || len(group.Cards) == 0 || c.Type != group.Type || c.Type == PokerCardsUnknown {
		return false
	}

	return true
}

func (c CardsGroup) GreaterThan(group CardsGroup) bool {
	if c.Type != group.Type {
		return c.Type > group.Type
	}

	// 顺子和同花顺，比较最左的牌点数
	if group.Type == StraightNiuType || group.Type == ColorStraightType {
		return c.Cards[0].Point() > group.Cards[0].Point()
	}

	return c.Cards[len(c.Cards)].Point() > group.Cards[len(group.Cards)].Point()
}
