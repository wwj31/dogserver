package fasterrun

import (
	"math/rand"
	"sort"

	"server/common/log"
)

// RandomPokerCards 获得洗好的一副新牌
func RandomPokerCards(ignoreCard []PokerCard) PokerCards {
	cards := make(PokerCards, len(pokerCards52))
	copy(cards, pokerCards52[:])
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

// PointCards 找出点数的所有牌 ,n指定数量，如果要所有牌，n不传
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

// StraightGroups 找出所有满足n连顺的牌,只能是顺子，连对，飞机
func (p PokerCards) StraightGroups(t PokerCardsType, n int) (result []PokerCards) {
	if t != Straight && t != StraightPair && t != Plane {
		log.Errorw("StraightGroups args error", "t", t)
		return nil
	}
	if n > len(p) {
		return nil
	}

	same := int(t - 5) // 顺子1张，连队2张，飞机3张
	stat := p.ConvertStruct()
	arr := p.ConvertOrderPoint(stat)
	if len(arr) == 0 {
		return nil
	}

	var (
		seq      = 1
		seqPoint = arr[0]
	)

	for i := 1; i < len(arr); i++ {
		point := arr[i]

		if stat[point] < same {
			continue
		}

		if point == seqPoint+1 {
			seqPoint = point
			seq++
			// 满足n顺的条件，就把这组牌加入结果中
			if seq >= n {
				var sub PokerCards

				for v := point - int32(n-1); v <= point; v++ {
					sub = append(sub, p.PointCards(v, same)...)
				}
				result = append(result, sub)
			}
			continue
		}

		seq = 1
		seqPoint = point
	}

	return
}

// SideCards 主要用于带牌的找牌规则,尽量找落单的，点数小的牌
func (p PokerCards) SideCards(n int) (result PokerCards) {
	if n == len(p) {
		result = make(PokerCards, len(p))
		copy(result, p)
		return result
	}
	stat := p.ConvertStruct()

	var cards PokerCards
	for point, num := range stat {
		if num == 1 {
			cards = append(cards, p.PointCards(point)...)
		}
	}

	var nextCards PokerCards
	if len(cards) <= 4 {
		sort.Slice(cards, func(i, j int) bool { return cards[i].Point() < cards[j].Point() })
		for _, card := range cards {
			result = append(result, card)
		}
		if len(result) >= n {
			result = result[:n]
			return
		} else {
			nextCards = p.Remove(result...)
		}
	}

	addNum := n - len(result)
	for i := 0; i < addNum; i++ {
		if len(nextCards) == 0 {
			return nil
		}

		randIdx := rand.Intn(len(nextCards))
		result = append(result, nextCards[randIdx])
		nextCards = append(nextCards[:randIdx], nextCards[randIdx+1:]...)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Point() < result[j].Point() })
	return result
}

func (c CardsGroup) CanCompare(group CardsGroup) bool {
	if len(c.Cards) == 0 || len(group.Cards) == 0 || c.Type != group.Type || c.Type == CardsTypeUnknown {
		return false
	}

	return true
}
func (c CardsGroup) Bigger(group CardsGroup) bool { return c.Cards[0].Point() > group.Cards[0].Point() }
