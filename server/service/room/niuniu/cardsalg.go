package niuniu

import (
	"sort"

	"server/proto/outermsg/outer"
)

func (p PokerCards) joker() (cards PokerCards) {
	for _, card := range p {
		if card.Point() == 15 || card.Point() == 16 {
			cards = append(cards, card)
			if len(cards) == 2 {
				return
			}
		}
	}
	return
}

func (p PokerCards) comprehensiveAnalysis(jokerCards PokerCards, params *outer.NiuNiuParams) (cardsGroup CardsGroup) {
	cardsWithoutJoker := p.Remove(jokerCards...) // 先获得去除了王的牌
	var cardsGroups []*CardsGroup
	for _, joker1Card := range pokerCards52 {
		newCards := append(cardsWithoutJoker, joker1Card) // 补第一张王
		if len(newCards) == 4 {
			for _, joker2Card := range pokerCards52 {
				newCards = append(cardsWithoutJoker, joker2Card) // 补第二张王
			}
			cg := newCards.AnalyzeCards(params)
			cardsGroups = append(cardsGroups, &cg)
		} else {
			cg := newCards.AnalyzeCards(params)
			cardsGroups = append(cardsGroups, &cg)
		}
	}
	sort.Slice(cardsGroups, func(i, j int) bool {
		return cardsGroups[j].GreaterThan(*cardsGroups[i])
	})

	return *cardsGroups[len(cardsGroups)-1]
}

// AnalyzeCards 分析牌型
func (p PokerCards) AnalyzeCards(params *outer.NiuNiuParams) (cardsGroup CardsGroup) {
	// 牛牛固定5张牌
	if len(p) == 5 {
		return
	}

	if jokerCards := p.joker(); len(jokerCards) > 0 {
		return p.comprehensiveAnalysis(jokerCards, params)
	}

	sort.Slice(p, func(i, j int) bool { return p[i].Point() < p[j].Point() })

	stat := p.ConvertStruct()
	orderPoints := p.ConvertOrderPoint(stat)

	colorSame := p.isColorSame()
	if len(stat) == 5 && isComb(orderPoints) {
		if colorSame && params.SpecialColorStraight {
			cardsGroup.Type = ColorStraightType
			cardsGroup.Cards = p
			return
		}
		if params.SpecialStraightNiu {
			cardsGroup.Type = StraightNiuType
			cardsGroup.Cards = p
			return
		}
	}

	// 同花牛
	if colorSame && params.SpecialSameColorNiu {
		cardsGroup.Type = SameColorNiuType
		cardsGroup.Cards = p
		return
	}

	// 小五牛判断
	if params.SpecialFiveSmallNiu && func() bool {
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

	if params.SpecialBombNiu {
		// 炸弹牛
		for point, num := range stat {
			if num == 4 {
				cardsGroup.Type = BombNiuType
				cardsGroup.Cards = p.PointCards(point)
				cardsGroup.SideCards = p.Remove(cardsGroup.Cards...)
				return
			}
		}
	}

	// 葫芦牛
	if params.SpecialHuluNiu {
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
	}

	// 五花牛
	if params.SpecialFiveColorNiu && func() bool {
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

	// 牛1~牛10
	var (
		cards       = make(PokerCards, 3, 3) // 主牌
		total10Func func(n int) (found bool)
	)

	total10Func = func(n int) bool {
		for i := 0; i < len(p); i++ {
			cards[n] = p[i]
			if n == 2 {
				var totalPoint int32
				for _, card := range cards {
					// 牛牛特殊规则，JQK判断点数，都作为10点
					point := card.Point()
					if point == 11 || point == 12 || point == 13 {
						point = 10
					}
					totalPoint += point
				}

				if totalPoint%10 == 0 {
					return true
				}
			} else if total10Func(n + 1) {
				return true
			}
		}
		return false
	}

	if found := total10Func(0); found {
		cardsGroup.Cards = cards
		cardsGroup.SideCards = p.Remove(cards...)
		var totalPoint int32
		for _, card := range cardsGroup.SideCards {
			// 牛牛特殊规则，JQK判断点数，都作为10点
			point := card.Point()
			if point == 11 || point == 12 || point == 13 {
				point = 10
			}
			totalPoint += point
			niuniuType := totalPoint % 10
			if niuniuType == 0 {
				niuniuType = 10
			}
			cardsGroup.Type = PokerCardsType(niuniuType)
		}

		return

	}
	return
}

// 是否包含JQk
func isJQK(point int32) bool {
	return point == 11 || point == 12 || point == 13
}

// 判断一堆点数是否连续
func isComb(arr []int32) bool {
	if len(arr) < 2 {
		return false
	}

	// 如果有A,先把A去掉,判断剩下的牌是否连续
	// 统计joker的数量

	hasA := false
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 1 {
			arr = append(arr[:i], arr[i+1:]...)
			hasA = true
		}
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

	// 剩下的牌是连续的，在把A分别作为1和14带入判断
	if comb && hasA {
		if arr[0] == 2 || arr[len(arr)-1] == 13 {
			return true
		}
		return false
	}
	return comb
}

func isJoker(point int32) bool {
	return point == 15 || point == 16
}
