package niuniu

import (
	"sort"

	"server/proto/outermsg/outer"
)

func (p PokerCards) joker() (cards PokerCards) {
	for _, card := range p {
		if card.Point() == Joker1.Point() || card.Point() == Joker2.Point() {
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
		newCards := append(cardsWithoutJoker, joker1Card)
		if len(newCards) == 4 {
			for _, joker2Card := range pokerCards52 {
				newCards2 := append(newCards, joker2Card)
				cg := newCards2.AnalyzeCards(params) // 两张王的情况
				if cg.Type != PokerCardsUnknown {
					cardsGroups = append(cardsGroups, &cg)
				}
			}

		} else {
			cg := newCards.AnalyzeCards(params) // 一张王的情况
			cardsGroups = append(cardsGroups, &cg)
			if cg.Type != PokerCardsUnknown {
				cardsGroups = append(cardsGroups, &cg)
			}
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
	if len(p) != 5 {
		return
	}

	if jokerCards := p.joker(); len(jokerCards) > 0 {
		return p.comprehensiveAnalysis(jokerCards, params)
	}

	sort.Slice(p, func(i, j int) bool { return p[i].Point() < p[j].Point() })

	stat := p.ConvertStruct()
	orderPoints := p.ConvertOrderPoint(stat)

	colorSame := p.isColorSame()
	if len(stat) == 5 {
		if comb, Ais14 := isComb(orderPoints); comb {
			var mainCards = p
			// 特殊处理A作为14点的情况
			if Ais14 {
				A := mainCards.PointCards(1)[0]
				mainCards = mainCards.Remove(A)
				mainCards = append(mainCards, PokerCard(A.Color()*100+14))
			}

			if colorSame && params.SpecialColorStraight {
				cardsGroup.Type = ColorStraightType
				cardsGroup.Cards = mainCards.Sort()
				return
			}
			if params.SpecialStraightNiu {
				cardsGroup.Type = StraightNiuType
				cardsGroup.Cards = mainCards.Sort()
				return
			}
		}
	}

	// 五小牛判断
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
		cardsGroup.Cards = p.Sort()
		return
	}

	// 炸弹牛
	if params.SpecialBombNiu {
		for point, num := range stat {
			if num == 4 {
				cards := p
				if point == 1 {
					point = 14
				}

				A := p.PointCards(1)
				cards = p.Remove(A...)
				for _, card := range A {
					cards = append(cards, PokerCard(card.Color()*100+14))
				}
				cardsGroup.Type = BombNiuType
				cardsGroup.Cards = cards.PointCards(point).Sort()
				cardsGroup.SideCards = p.Remove(cardsGroup.Cards...)
				return
			}
		}
	}

	// 葫芦牛
	if params.SpecialHuluNiu {
		cards := p
		A := p.PointCards(1)
		cards = p.Remove(A...)
		for _, card := range A {
			cards = append(cards, PokerCard(card.Color()*100+14))
		}
		if len(A) > 0 {
			stat = cards.ConvertStruct()
		}

		var cards3, cards2 PokerCards
		for point, num := range stat {
			if num == 3 {
				cards3 = cards.PointCards(point)
			}
			if num == 2 {
				cards2 = cards.PointCards(point)
			}
		}

		if len(cards3) == 3 && len(cards2) == 2 {
			cardsGroup.Type = HuluNiuType
			cardsGroup.Cards = cards3.Sort()
			cardsGroup.SideCards = cards2.Sort()
			return
		}
	}

	// 同花牛
	if colorSame && params.SpecialSameColorNiu {
		cardsGroup.Type = SameColorNiuType
		cardsGroup.Cards = p.Sort()
		return
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
		cardsGroup.Cards = p.Sort()
		return
	}

	// 牛1~牛10
	var (
		cards       = make(PokerCards, 3, 3) // 主牌
		total10Func func(n, idx int) bool
		cards10     []PokerCards
	)

	total10Func = func(n, idx int) bool {
		for i := idx; i < len(p); i++ {
			cards[n] = p[i]
			if n == 2 {
				var totalPoint int32
				for _, card := range cards {
					// 特殊规则，JQK判断点数，都作为10点
					point := card.Point()
					if point == 11 || point == 12 || point == 13 {
						point = 10
					}
					totalPoint += point
				}

				if totalPoint%10 == 0 {
					newCards := make(PokerCards, 3, 3)
					copy(newCards, cards.Sort())
					cards10 = append(cards10, newCards)
					return true // 无需区分各种组合成10的情况
				}
				continue
			}
			if total10Func(n+1, i+1) {
				return true
			}
		}
		return false
	}

	total10Func(0, 0)

	if len(cards10) > 0 {
		sort.Slice(cards10, func(i, j int) bool {
			cardi := cards10[i][len(cards10[i])-1]
			cardj := cards10[j][len(cards10[j])-1]
			if cardi.Point() != cardj.Point() {
				return cardi.Point() > cardj.Point()
			}
			return cardi.Color() > cardj.Color()
		})

		bestCards := cards10[0]

		cardsGroup.Cards = bestCards.Sort()
		cardsGroup.SideCards = p.Remove(bestCards...).Sort()
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

	// 没牛
	cardsGroup.Cards = p.Sort()

	return
}

// 是否包含JQk
func isJQK(point int32) bool {
	return point == 11 || point == 12 || point == 13
}

// 判断一堆点数是否连续
func isComb(arr []int32) (comb bool, isAis14 bool) {
	if len(arr) < 2 {
		return false, false
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

	comb = true
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
		if arr[0] == 2 {
			return true, false
		}
		if arr[len(arr)-1] == 13 {
			return true, true
		}

		return false, false
	}
	return comb, false
}
