package fasterrun

import (
	"math/rand"
	"server/common"
	"server/common/log"
	"sort"
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

// PointCards 找出点数的所有牌
func (p PokerCards) PointCards(point int32) PokerCards {
	var newCards PokerCards
	for _, card := range p {
		if card.Point() == point {
			newCards = append(newCards, card)
		}
	}
	return newCards
}

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

// FindBigger 找出指定牌型更大的牌
func (p PokerCards) FindBigger(cardsGroup CardsGroup) (bigger []CardsGroup) {
	// TODO

	return
}

// AnalyzeCards 分析牌型
func (p PokerCards) AnalyzeCards(AAAisBomb bool) (cardsGroup CardsGroup) {
	sort.Slice(&p, func(i, j int) bool { return PokerCard(i).Point() < PokerCard(j).Point() })

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
				cardsGroup.Type = Three
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
					cardsGroup.Type = ThreeWithOne
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
				cardsGroup.Type = ThreeWithTwo
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
			beginPoint       = orderPoints[0]
			sequentialMaxNum = 1 // 最大的连续数量(不计相同的牌)
		)

		for i := 1; i < len(orderPoints); i++ {
			if orderPoints[i] == beginPoint+1 {
				sequentialMaxNum++
			} else {
				sequentialMaxNum = 1
			}
			beginPoint = orderPoints[i]
		}

		// 最长连续数等于本组牌的长度，只能是顺子
		if sequentialMaxNum == l {
			cardsGroup.Type = Straight
			cardsGroup.Cards = append(cardsGroup.Cards, p...)
			return
		}

		// 最长连续数等于本组牌长度的一半，肯定数连对
		if sequentialMaxNum*2 == l {
			cardsGroup.Type = StraightPair
			cardsGroup.Cards = append(cardsGroup.Cards, p...)
			return
		}

		// 最长连续数等于本组牌长度的1/3，肯定飞机
		if sequentialMaxNum*3 == l {
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
		if sequentialMaxNum >= 2 {
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
					cardsGroup.Type = PlaneWithTow
					cardsGroup.Cards = planeCards
					cardsGroup.SideCards = sideCards
					return
				}
			}
		}
	}
	return
}

func (c CardsGroup) CanCompare(group CardsGroup) bool {
	if len(c.Cards) == 0 || len(group.Cards) == 0 || c.Type != group.Type || c.Type == CardsTypeUnknown {
		return false
	}

	return true
}
func (c CardsGroup) Bigger(group CardsGroup) bool { return c.Cards[0].Point() > group.Cards[0].Point() }

func min(a, b, c int) int {
	return common.Min(a, common.Min(b, c))
}
