package mahjong

import "server/common/log"

// 返回所有能听牌的单牌以及对应的胡牌类型
func (c Cards) ting(ignore ColorType, lightGang, darkGang, pong map[int32]int64) (tingCards map[Card]HuType) {
	tingCards = make(map[Card]HuType)
	if len(c.colors()) == 3 {
		return
	}

	allSingleCards, ok := cardsWithoutIgnore[ignore]
	if !ok {
		log.Errorw("ting ignore error", "ignore", ignore)
		return
	}

	// 检查每一张牌的组合
	for _, card := range allSingleCards {
		if huType := c.Insert(card).IsHu(lightGang, darkGang, pong); huType != HuInvalid {
			tingCards[card] = huType
		}
	}
	return
}

// IsHu 牌组能否胡牌
func (c Cards) IsHu(lightGang, darkGang, pong map[int32]int64) (typ HuType) {
	colors := c.colors()
	if len(colors) == 3 {
		return HuInvalid
	}

	duiziGroups := c.Duizi()
	// 没有能做将的牌
	if len(duiziGroups) == 0 {
		return HuInvalid
	}

	// 先判断七对
	if len(duiziGroups) == 7 {
		typ = QiDui
		// 全是对子，并且还能去重，肯定就是龙七对
		if len(duiziGroups) > len(RemoveDuplicate(duiziGroups)) {
			typ = LongQiDui
		}
	} else {
		// 先检查是否是散牌
		if c.HighCard(c.ConvertStruct()) {
			return HuInvalid
		}

		// 去个重
		duiziGroups = RemoveDuplicate(duiziGroups)

		// 挨个做将，再分析剩下的牌型
		for _, jiangCards := range duiziGroups {
			spareHandCards := c.Remove(jiangCards...)

			// 有散牌
			if spareHandCards.HighCard(spareHandCards.ConvertStruct()) {
				continue
			}

			// 将牌有幺九，并且碰杠也都有幺九，就带上幺九检测
			has1or9 := jiangCards.Has1or9() && pongGangAllHas1or9(lightGang, darkGang, pong)

			// 判断剩余牌是否胡了
			typ = RecurCheckHu(spareHandCards, has1or9)
			if typ > HuInvalid {
				break
			}
		}
	}

	if typ == HuInvalid {
		return
	}

	return c.upgrade(colors, lightGang, darkGang, pong, typ)
}

func (c Cards) qingYiSeUpgrade(colors map[int]struct{}, typ HuType) HuType {
	if typ == HuInvalid {
		return typ
	}
	if len(colors) == 1 {
		switch typ {
		case Hu:
			return QingYiSe
		case DuiDuiHu:
			return QingDui
		case QiDui:
			return QingQiDui
		case LongQiDui:
			return QingLongQiDui
		}
	}
	return typ
}

// 传入花色，传入明杠
func (c Cards) upgrade(colors map[int]struct{}, lightGang, darkGang, pong map[int32]int64, typ HuType) HuType {
	// 判断清一色升级牌型
	typ = c.qingYiSeUpgrade(colors, typ)

	switch typ {
	case DuiDuiHu, QiDui, LongQiDui: // 对对胡->将对对、七对\龙七对->将七对
		upgrade := true
		c.Range(func(color ColorType, number int) bool {
			if number != 2 && number != 5 && number != 8 {
				upgrade = false
				return true
			}
			return false
		})
		if upgrade {
			switch typ {
			case DuiDuiHu:
				typ = JiangDui // 将对
			case QiDui, LongQiDui:
				typ = JiangQiDui // 将七对
			}
		}

	case Hu:
		// 判断平胡升中张 碰杠有幺九不算中张
		if !c.Has1or9() && pongGangAllHasNo1or9(lightGang, darkGang, pong) {
			return ZhongZhang
		}

		// 判断平胡升门清
		if len(pong) == 0 {
			// 不能碰
			return typ
		}

		// 不能明杠
		if len(lightGang) == 0 {
			return MenQing
		}

	}
	return typ
}

// 所有碰杠是否都有幺九牌
func pongGangAllHas1or9(lightGang, darkGang, pong PongGang) bool {
	if len(pong) > 0 && !pong.Has1or9() {
		return false
	}

	if len(lightGang) > 0 && !lightGang.Has1or9() {
		return false
	}

	if len(darkGang) > 0 && !darkGang.Has1or9() {
		return false
	}
	return true
}

// 所有碰杠是否都没有幺九牌
func pongGangAllHasNo1or9(lightGang, darkGang, pong PongGang) bool {
	if len(pong) > 0 && pong.Has1or9() {
		return false
	}

	if len(lightGang) > 0 && lightGang.Has1or9() {
		return false
	}

	if len(darkGang) > 0 && darkGang.Has1or9() {
		return false
	}
	return true
}
