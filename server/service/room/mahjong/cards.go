package mahjong

import (
	"fmt"
	"math/rand"
	"sort"

	"server/common"
	"server/common/log"
)

// MaxCardNum 创建一个41大小的数组，有效索引为11-39,即为牌，值为牌数量
const MaxCardNum = 41

// RandomCards 获得洗好的一副新牌
func RandomCards(ignoreCard Cards) Cards {
	cards := cards108
	tail := len(cards)
	for i := 0; i < len(cards); i++ {
		idx := rand.Intn(len(cards[:tail]))
		cards[idx], cards[tail-1] = cards[tail-1], cards[idx]
	}

	result := Cards(cards[:])
	if len(ignoreCard) > 0 {
		result = result.Remove(ignoreCard...)
	}
	return result
}

func (c Cards) Sort() Cards {
	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })
	return c
}

func (c Cards) Len() int {
	return len(c)
}

// Insert 插入一组牌
func (c Cards) Insert(cards ...Card) Cards {
	src := c
	var dst Cards
	for _, card := range cards {
		dst = Cards{}
		index := sort.Search(src.Len(), func(i int) bool { return src[i] > card })
		dst = append(dst, src[:index]...)
		dst = append(dst, card)
		dst = append(dst, src[index:]...)
		src = dst
	}
	return dst
}

// Remove 移除一组牌,移除的牌必须全部在牌中
func (c Cards) Remove(cards ...Card) Cards {
	cardMap := make(map[Card]int) // 统计要移除的牌数量
	for _, card := range cards {
		cardMap[card] += 1
	}

	dst := make(Cards, 0, c.Len())
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

	if len(cardMap) != 0 {
		log.Errorw("cards remove error handCards:%v cards:%v", c, cards)
	}

	return dst
}

// CanPongTo 判断当前牌组能不能碰
func (c Cards) CanPongTo(card Card) bool {
	var num int
	for _, handCard := range c {
		if handCard == card {
			num++
			if num == 2 {
				return true
			}
		}
	}
	return false
}

// Pong 碰,返回去除了碰牌后的新手牌，以及碰牌起始下标
func (c Cards) Pong(card Card) (cards Cards, index int, err error) {
	index = sort.Search(c.Len(), func(i int) bool { return c[i] >= card })
	if index >= c.Len() || c[index] != card {
		err = fmt.Errorf("pong failed cannot find index cards:%v index:%v card:%v", c, index, card)
		return
	}

	// 找到后把左边一张牌和当前牌，作为碰牌
	// 检查左边那张牌必须相同
	if c[index+1] != card {
		err = fmt.Errorf("pong failed card number != 2:%v index:%v card:%v", c, index, card)
		return
	}

	cards = append(cards, c[:index]...)
	cards = append(cards, c[index+2:]...)
	return
}

// CanGangTo 能否杠这张牌
func (c Cards) CanGangTo(card Card) bool {
	var num int
	for _, handCard := range c {
		if handCard == card {
			num++
			if num == 3 {
				return true
			}
		}
	}
	return false
}

// HasGang 有没有能杠的牌
func (c Cards) HasGang() (cards Cards) {
	cs := c.ConvertStruct()
	for i := 11; i < MaxCardNum; i++ {
		if cs[i] == 4 {
			cards = append(cards, Card(i))
		}
	}
	return
}

// Gang 杠,返回去除了杠牌后的新手牌，以及杠的起始下标
func (c Cards) Gang(card Card) (cards Cards, index int, err error) {
	index = sort.Search(c.Len(), func(i int) bool { return c[i] >= card })
	if index >= c.Len() || c[index] != card {
		err = fmt.Errorf("gang failed cannot find index cards:%v index:%v card:%v", c, index, card)
		return
	}

	// 找到后把左边两张牌和当前牌，作为杠
	// 检查左边两张牌必须相同
	if c[index+1] != card || c[index+2] != card {
		err = fmt.Errorf("gang failed card number != 2:%v index:%v card:%v", c, index, card)
		return
	}

	cards = append(cards, c[:index]...)
	cards = append(cards, c[index+3:]...)
	return
}

// RecurCheckHu 给一副去除了将牌的牌组，判断有没有一种组合能把所有牌都组成刻子或顺子
func RecurCheckHu(cards Cards, upCardsHas1or9 bool) HuType {
	if cards.Len() == 0 {
		if upCardsHas1or9 {
			return QuanYaoJiu
		} else {
			return Hu
		}
	}

	// 当前牌组是散牌,返回失败
	if cards.HighCard(cards.ConvertStruct()) {
		return HuInvalid
	}

	// 所有牌刚好全部是刻子,返回成功
	allKezi := cards.Kezi()
	if len(allKezi)*3 == cards.Len() {
		return DuiDuiHu
	}

	allShunzi := cards.Shunzi()

	// 所有砍
	allKan := append(RemoveDuplicate(allShunzi), RemoveDuplicate(allKezi)...)
	for _, kan := range allKan {
		tmp := cards.Remove(kan...)
		if h := RecurCheckHu(tmp, kan.Has1or9()); h > HuInvalid {
			return h
		}
	}
	return HuInvalid
}

// Has1or9 是否带有1或9
func (c Cards) Has1or9() bool {
	for _, card := range c {
		num := card.Int() / 10
		if num == 1 || num == 9 {
			return true
		}
	}
	return false
}

// HasJiaXinWu 是否具备夹心五
func (c Cards) HasJiaXinWu(card Card) bool {
	if card.Int32()%10 != 5 {
		return false
	}

	var has4, has6 bool
	c.Range(func(card Card) bool {
		if card.Int() == card.Int()-1 {
			has4 = true
		}

		if card.Int() == card.Int()-1 {
			has6 = true
		}

		if has4 && has6 {
			return true
		}

		return false
	})
	return false
}

// ColorCount 当前牌有几个花色
func (c Cards) colors() map[int]struct{} {
	colorMap := make(map[int]struct{})
	for _, card := range c {
		color := int(card / 10)
		colorMap[color] = struct{}{}
	}

	return colorMap
}

// colorCards 获得某个花色的所有牌
func (c Cards) colorCards(color ColorType) Cards {
	var cards Cards
	for _, card := range c {
		if ColorType(card/10) == color {
			cards = append(cards, card)
		}
	}

	return cards
}

func (c Cards) Random() Card {
	if c.Len() == 0 {
		return Card(0)
	}
	return c[rand.Intn(len(c))]
}

// CardIndex 返回第一张找的牌的位置
func (c Cards) CardIndex(dstCard Card) int {
	for index, card := range c {
		if card == dstCard {
			return index
		}
	}
	return -1
}

func (c Cards) ToSlice() (result []int32) {
	for _, card := range c {
		result = append(result, card.Int32())
	}
	return
}

func (c Cards) RangeSplit(fn func(color ColorType, number int) bool) {
	for _, card := range c {
		if fn(ColorType(card%10), int(card/10)) {
			break
		}
	}
}

func (c Cards) Range(fn func(card Card) bool) {
	for _, card := range c {
		if fn(card) {
			break
		}
	}
}

// HighCard 检查是否存在散牌,存在不能组成顺子、刻子、对子的单牌
func (c Cards) HighCard(cardsStat [MaxCardNum]int) bool {
	if c.Len() == 0 {
		return false
	}

	if c.Len() == 1 {
		return true
	}
	var continuous int

	check := func(i int) bool {
		// 断开连后，如果之前连续是1或者2，那么这两张连着的牌，有一张是单牌，那就构成散牌
		if 0 < continuous && continuous < 3 {
			if cardsStat[i-1] == 1 || cardsStat[i-2] == 1 {
				return true
			}
		}
		return false
	}

	for i := 11; i < MaxCardNum; i++ {
		if cardsStat[i] == 0 {
			if check(i) {
				return true
			}
			continuous = 0
			continue
		}

		continuous++
	}

	return check(MaxCardNum)
}

// ConvertStruct 转换牌型结构为统计结构，槽位下标表示牌，值表示牌的数量，前11个槽位无用
// 例如 [40]int{ [0]=0, [1]=0, ... [11]=2, [12]=3, [13]=2, [14]=0, [15]=2, }
// 表示 一万2张，二万3张，三万2张，四万0张，五万2张
func (c Cards) ConvertStruct() (result [MaxCardNum]int) {
	for _, card := range c {
		if card.Int() >= MaxCardNum {
			log.Errorw("card number out of range ", "card", card)
			return
		}

		result[card.Int()]++
	}
	return result
}

// Duizi 找对子
func (c Cards) Duizi() []Cards {
	var result []Cards
	for i := 0; i < len(c)-1; i++ {
		if c[i] == c[i+1] {
			result = append(result, Cards{c[i], c[i+1]})
			i++ // 跳过下一个已经配对的牌
		}
	}
	return result
}

// Kezi 找刻子
func (c Cards) Kezi() []Cards {
	var result []Cards
	for i := 0; i < len(c)-2; i++ {
		if c[i] == c[i+1] && c[i+1] == c[i+2] {
			result = append(result, Cards{c[i], c[i+1], c[i+2]})
			// 跳过所有相同的牌
			for i < len(c)-2 && c[i] == c[i+2] {
				i++
			}
		}
	}
	return result
}

// RemoveDuplicate 去除对子和顺子中重复的牌组
func RemoveDuplicate(cardsGroup []Cards) []Cards {
	// 创建一个 map 来记录已经出现过的牌组
	uniqueMap := make(map[Card]bool, len(cardsGroup))
	var result []Cards

	// 对每个牌组进行处理
	for _, cards := range cardsGroup {
		// 检查是否已经看到过这个牌组
		if !uniqueMap[cards[0]] {
			uniqueMap[cards[0]] = true
			result = append(result, cards)
		}
	}

	return result
}

// Shunzi 找顺子
func (c Cards) Shunzi() []Cards {
	var result []Cards

	// 用于保存每个牌的数量
	cardCount := make(map[Card]int)
	for _, card := range c {
		cardCount[card]++
	}

	// 获取所有可能的顺子
	for card, count := range cardCount {
		if card%10 <= 7 && cardCount[card+1] > 0 && cardCount[card+2] > 0 {
			for i := 0; i < min(count, cardCount[card+1], cardCount[card+2]); i++ {
				result = append(result, Cards{card, card + 1, card + 2})
			}
		}
	}
	return result
}

func min(a, b, c int) int {
	return common.Min(a, common.Min(b, c))
}

func (p PongGang) Has1or9() bool {
	for card := range p {
		n := card / 10
		if n == 1 || n == 9 {
			return true
		}
	}
	return false
}
