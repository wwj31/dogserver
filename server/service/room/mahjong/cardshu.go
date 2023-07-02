package mahjong

import "server/common/log"

// ignore 定缺花色, 返回所有能听牌的单牌以及对应的胡牌类型
func (c Cards) ting(ignore ColorType, gang, pong map[Card]bool) (tingCards map[Card]HuType) {
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
		if huType := c.Insert(card).isHu(gang, pong); huType != HuInvalid {
			tingCards[card] = huType
		}
	}
	return
}

// 牌组能否胡牌
func (c Cards) isHu(gang, pong map[Card]bool) (typ HuType) {
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

			// 判断剩余牌是否胡了,(手牌有1、9，就加入带幺九判断)
			typ = RecurCheckHu(spareHandCards, jiangCards.Has1or9())
			if typ > HuInvalid {
				break
			}
		}
	}

	if typ == HuInvalid {
		return
	}

	return c.upgrade(colors, gang, pong, typ)
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

func (c Cards) upgrade(colors map[int]struct{}, gang, pong map[Card]bool, typ HuType) HuType {
	// 判断清一色升级牌型
	typ = c.qingYiSeUpgrade(colors, typ)

	switch typ {
	case DaiYaoJiu: // 如果是带幺九，判断能否升级全幺九
		upgrade := true
		c.Range(func(color ColorType, number int) bool {
			if number != 1 && number != 9 {
				upgrade = false
				return true
			}
			return false
		})
		if upgrade {
			typ = QuanYaoJiu // 全幺九
		}

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
		// TODO 碰杠有19算不算中张？
		// 判断平胡升中张
		if !c.Has1or9() {
			return ZhongZhang
		}

		// 判断平胡升门清

		// 不能碰
		if len(pong) == 0 {
			return typ
		}

		// 不能明杠
		for _, ming := range gang {
			if ming {
				return typ
			}
		}
		return MenQing

	}
	return typ
}
