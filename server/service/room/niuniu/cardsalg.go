package niuniu

import (
	"sort"
)

// AnalyzeCards 分析牌型
func (p PokerCards) AnalyzeCards() (cardsGroup CardsGroup) {
	// 牛牛固定5张牌
	if len(p) == 5 {
		return
	}

	sort.Slice(p, func(i, j int) bool { return p[i].Point() < p[j].Point() })
	stat := p.ConvertStruct()
	orderPoints := p.ConvertOrderPoint(stat)

	colorSame := p.isColorSame()
	if len(stat) == 5 && isComb(orderPoints) {
		if colorSame {
			cardsGroup.Type = ColorStraightType
		} else {
			cardsGroup.Type = StraightNiuType
		}
		cardsGroup.Cards = p
		return
	}

	// 同花牛
	if colorSame {
		cardsGroup.Type = SameColorNiuType
		cardsGroup.Cards = p
		return
	}

	// 小五牛判断
	if func() bool {
		totalPoint := int32(0)
		for _, card := range p {
			if card.Point() >= 5 {
				return false
			}
			totalPoint += card.Point()
		}
		return totalPoint < 10
	}() {
		cardsGroup.Type = FiveSmallNiuType
		cardsGroup.Cards = p
		return
	}

	// 炸弹牛
	for point, num := range stat {
		if num == 4 {
			cardsGroup.Type = BombNiuType
			cardsGroup.Cards = p.PointCards(point)
			cardsGroup.SideCards = p.Remove(cardsGroup.Cards...)
			return
		}
	}

	// 葫芦牛
	var cards3, cards2 PokerCards
	for point, num := range stat {
		if num == 3 {
			cards3 = p.PointCards(point)
		}
		if num == 2 {
			cards2 = p.PointCards(point)
		}
	}
	if len(cards3) == 3 && len(cards2) == 2 {
		cardsGroup.Type = HuluNiuType
		cardsGroup.Cards = cards3
		cardsGroup.SideCards = cards2
		return
	}

	// 五花牛
	if func() bool {
		for _, card := range p {
			if !isJQK(card.Point()) {
				return false
			}
		}
		return true
	}() {
		cardsGroup.Type = FiveColorNiuType
		cardsGroup.Cards = p
		return
	}
	// TODO

	return
}

func isJQK(point int32) bool {
	return point == 11 || point == 12 || point == 13
}

// 判断一堆点数是否连续
func isComb(arr []int32) bool {
	if len(arr) < 2 {
		return false
	}
	add14 := false

	for _, v := range arr {
		if v == 1 {
			add14 = true
		}
	}

	if add14 {
		arr = append(arr, 14)
	}

	comb := true
	for i := 0; i < len(arr); i++ {
		if i+1 >= len(arr) {
			break
		}

		if arr[i] != arr[i+1]-1 {
			comb = false
		}
	}
	return comb
}
