package fasterrun

import (
	"sort"

	"server/common/log"
)

// FindBigger 找出指定牌型更大的牌
func (p PokerCards) FindBigger(cardsGroup CardsGroup) (bigger []CardsGroup) {
	if len(cardsGroup.Cards) == 0 {
		log.Errorw("find bigger unexpected", "group", cardsGroup)
		return nil
	}

	stat := p.ConvertStruct()
	switch cardsGroup.Type {
	case Single:
		for _, card := range p {
			if card.Point() > cardsGroup.Cards[0].Point() {
				bigger = append(bigger, CardsGroup{Type: cardsGroup.Type, Cards: PokerCards{card}})
			}
		}

	case Pair, Trips:
		n := 2
		if cardsGroup.Type == Trips {
			n = 3
		}

		for point, num := range stat {
			if num >= n && point > cardsGroup.Cards[0].Point() {
				cards := p.PointCards(point)
				comb := cards.Combination(len(cards))
				for _, c := range comb {
					bigger = append(bigger, CardsGroup{Type: cardsGroup.Type, Cards: c})
				}
			}
		}
	case TripsWithOne, TripsWithTwo:
		sideCardsSize := 1 // 用于找几张需要带的牌
		if cardsGroup.Type == TripsWithTwo {
			sideCardsSize = 2
		}

		for point, num := range stat {
			if num >= 3 && point > cardsGroup.Cards[0].Point() {
				cards := p.PointCards(point)
				comb := cards.Combination(len(cards)) // 满足点数更大的牌组合
				for _, c := range comb {
					spareCards := p.Remove(c...)
					sideCards := spareCards.SideCards(sideCardsSize)
					if len(sideCards) == 0 {
						continue
					}
					bigger = append(bigger, CardsGroup{Type: cardsGroup.Type, Cards: c, SideCards: sideCards})
				}
			}
		}

	case Straight, StraightPair, Plane:
		n := int(cardsGroup.Type - 5)
		seq := len(cardsGroup.Cards) / n
		groups := p.StraightGroups(cardsGroup.Type, seq)
		for _, group := range groups {
			if group[0].Point() <= cardsGroup.Cards[0].Point() {
				continue
			}

			bigger = append(bigger, CardsGroup{Type: cardsGroup.Type, Cards: group})
		}

	case PlaneWithTwo:
		seq := len(cardsGroup.Cards) / 3
		straightGroups := p.StraightGroups(Plane, seq)
		for _, cards := range straightGroups {
			if cards[0].Point() <= cardsGroup.Cards[0].Point() {
				continue
			}

			spareCards := p.Remove(cards...)
			sideCards := spareCards.SideCards(seq * 2)
			if len(sideCards) == seq*2 {
				bigger = append(bigger, CardsGroup{Type: PlaneWithTwo, Cards: cards, SideCards: sideCards})
			}
		}

	case FourWithTwo, FourWithThree:
		sideCardsSize := 2 // 要找几张副牌
		if cardsGroup.Type == FourWithThree {
			sideCardsSize = 3
		}

		for point, num := range stat {
			if num < 4 || point <= cardsGroup.Cards[0].Point() {
				continue
			}
			cards := p.PointCards(point)
			spareCards := p.Remove(cards...)
			sideCards := spareCards.SideCards(sideCardsSize)
			if len(sideCards) == sideCardsSize {
				bigger = append(bigger, CardsGroup{Type: cardsGroup.Type, Cards: cards, SideCards: sideCards})
			}
		}
	}

	for point, num := range stat {
		if num == 4 {
			bigger = append(bigger, CardsGroup{Type: Bombs, Cards: p.PointCards(point)})
		}
	}
	return
}

// AnalyzeCards 分析牌型
func (p PokerCards) AnalyzeCards(AAAisBomb bool) (cardsGroup CardsGroup) {
	sort.Slice(p, func(i, j int) bool { return p[i].Point() < p[j].Point() })

	if len(p) == 0 {
		return
	}

	if len(p) == 1 {
		cardsGroup.Type = Single
		cardsGroup.Cards = append(cardsGroup.Cards, p...)
		return
	}

	stat := p.ConvertStruct()
	l := len(p)
	switch {
	case l == 2: // 2张，只能组成对子
		if len(stat) == 1 {
			cardsGroup.Type = Pair
			cardsGroup.Cards = append(cardsGroup.Cards, p...)
		}

	case l == 3: // 3张，只能组成三张
		if len(stat) == 1 {
			if AAAisBomb && p[0].Point() == 14 {
				cardsGroup.Type = Bombs
			} else {
				cardsGroup.Type = Trips
			}
			cardsGroup.Cards = append(cardsGroup.Cards, p...)
		}

	case l == 4: // 4张，可能组成:炸弹、三带一、连对
		statLen := len(stat)
		if statLen == 1 {
			cardsGroup.Type = Bombs
			cardsGroup.Cards = append(cardsGroup.Cards, p...)
			return
		}

		if statLen == 2 {
			for point, num := range stat {
				if num == 3 {
					cardsGroup.Type = TripsWithOne
					cardsGroup.Cards = p.PointCards(point)
					cardsGroup.SideCards = p.Remove(cardsGroup.Cards...)
					return
				}
				if num == 2 {
					cardsGroup.Type = StraightPair
					cardsGroup.Cards = append(cardsGroup.Cards, p...)
					return
				}
			}
		}
		log.Errorw("AnalyzeCards cards len:4", "cards", p, "stat", stat)

	case l == 5: // 5张 只能组成三带二、顺子
		// 检查三带二
		for point, num := range stat {
			if num == 3 {
				cardsGroup.Type = TripsWithTwo
				cardsGroup.Cards = p.PointCards(point)
				cardsGroup.SideCards = p.Remove(cardsGroup.Cards...)
				return
			}
		}

		// 检查顺子(只能是5张顺子)
		var straight5 = true
		for i := 1; i < l; i++ {
			if p[i].Point() != p[i-1].Point()+1 {
				straight5 = false
				break
			}
		}
		if straight5 {
			cardsGroup.Type = Straight
			cardsGroup.Cards = append(cardsGroup.Cards, p...)
			return
		}

	case l > 5: // 5张及以上，统一判断
		orderPoints := p.ConvertOrderPoint(stat)
		if len(orderPoints) == 0 {
			log.Errorw("AnalyzeCards cards len:5", "cards", p, "stat", stat)
			cardsGroup.Type = CardsTypeUnknown
			return
		}

		var (
			beginPoint         = orderPoints[0]
			sequentialMaxCount = 1 // 最大的连续数量(不计相同的牌)
		)

		for i := 1; i < len(orderPoints); i++ {
			if orderPoints[i] == beginPoint+1 {
				sequentialMaxCount++
			} else {
				sequentialMaxCount = 1
			}
			beginPoint = orderPoints[i]
		}

		// 最长连续数等于本组牌的长度，只能是顺子
		if sequentialMaxCount == l {
			cardsGroup.Type = Straight
			cardsGroup.Cards = append(cardsGroup.Cards, p...)
			return
		}

		// 最长连续数等于本组牌长度的一半，肯定数连对
		if sequentialMaxCount*2 == l {
			cardsGroup.Type = StraightPair
			cardsGroup.Cards = append(cardsGroup.Cards, p...)
			return
		}

		// 最长连续数等于本组牌长度的1/3，肯定飞机
		if sequentialMaxCount*3 == l {
			cardsGroup.Type = Plane
			cardsGroup.Cards = append(cardsGroup.Cards, p...)
			return
		}

		// 以下牌型用特殊方式检查,四带二、四带三
		for point, num := range stat {
			if num == 4 {
				cards := p.PointCards(point)
				sideCards := p.Remove(cardsGroup.Cards...)
				sideCardsLen := len(sideCards)
				if sideCardsLen == 2 {
					cardsGroup.Type = FourWithTwo
					cardsGroup.Cards = cards
					cardsGroup.SideCards = sideCards
					return
				}

				if sideCardsLen == 3 {
					cardsGroup.Type = FourWithThree
					cardsGroup.Cards = cards
					cardsGroup.SideCards = sideCards
					return
				}
			}
		}

		// 判断飞机带翅膀
		if sequentialMaxCount >= 2 {
			// 先找所有三张
			var planePoints []int32
			for i := 0; i < len(orderPoints); i++ {
				if stat[orderPoints[i]] == 3 {
					if len(planePoints) == 0 || orderPoints[i] == planePoints[len(planePoints)-1]+1 {
						planePoints = append(planePoints, orderPoints[i])
					} else {
						planePoints = []int32{orderPoints[i]}
					}
				}
			}

			if len(planePoints) >= 2 {
				var planeCards PokerCards
				for _, point := range planePoints {
					planeCards = append(planeCards, p.PointCards(point)...)
				}
				sideCards := p.Remove(planeCards...)
				if len(sideCards) == len(planeCards)*2 {
					cardsGroup.Type = PlaneWithTwo
					cardsGroup.Cards = planeCards
					cardsGroup.SideCards = sideCards
					return
				}
			}
		}
	}
	return
}
