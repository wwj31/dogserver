package mahjong

type (
	Card  int32
	Cards []Card
)

const (
	WAN_1 Card = 11 // 1萬
	WAN_2 Card = 12 // 2萬
	WAN_3 Card = 13 // 3萬
	WAN_4 Card = 14 // 4萬
	WAN_5 Card = 15 // 5萬
	WAN_6 Card = 16 // 6萬
	WAN_7 Card = 17 // 7萬
	WAN_8 Card = 18 // 8萬
	WAN_9 Card = 19 // 9萬

	TIAO_1 Card = 21 // 1条
	TIAO_2 Card = 22 // 2条
	TIAO_3 Card = 23 // 3条
	TIAO_4 Card = 24 // 4条
	TIAO_5 Card = 25 // 5条
	TIAO_6 Card = 26 // 6条
	TIAO_7 Card = 27 // 7条
	TIAO_8 Card = 28 // 8条
	TIAO_9 Card = 29 // 9条

	TONG_1 Card = 31 // 1筒
	TONG_2 Card = 32 // 2筒
	TONG_3 Card = 33 // 3筒
	TONG_4 Card = 34 // 4筒
	TONG_5 Card = 35 // 5筒
	TONG_6 Card = 36 // 6筒
	TONG_7 Card = 37 // 7筒
	TONG_8 Card = 38 // 8筒
	TONG_9 Card = 39 // 9筒

)

var strCard = map[Card]string{
	WAN_1: "一萬",
	WAN_2: "二萬",
	WAN_3: "三萬",
	WAN_4: "四萬",
	WAN_5: "五萬",
	WAN_6: "六萬",
	WAN_7: "七萬",
	WAN_8: "八萬",
	WAN_9: "九萬",

	TIAO_1: "一条",
	TIAO_2: "二条",
	TIAO_3: "三条",
	TIAO_4: "四条",
	TIAO_5: "五条",
	TIAO_6: "六条",
	TIAO_7: "七条",
	TIAO_8: "八条",
	TIAO_9: "九条",

	TONG_1: "一筒",
	TONG_2: "二筒",
	TONG_3: "三筒",
	TONG_4: "四筒",
	TONG_5: "五筒",
	TONG_6: "六筒",
	TONG_7: "七筒",
	TONG_8: "八筒",
	TONG_9: "九筒",
}

func (m Card) String() string {
	return strCard[m]
}

func (m Card) Int32() int32 {
	return int32(m)
}

func (m Card) Int() int {
	return int(m)
}

var cards108 = [108]Card{
	WAN_1, WAN_1, WAN_1, WAN_1,
	WAN_2, WAN_2, WAN_2, WAN_2,
	WAN_3, WAN_3, WAN_3, WAN_3,
	WAN_4, WAN_4, WAN_4, WAN_4,
	WAN_5, WAN_5, WAN_5, WAN_5,
	WAN_6, WAN_6, WAN_6, WAN_6,
	WAN_7, WAN_7, WAN_7, WAN_7,
	WAN_8, WAN_8, WAN_8, WAN_8,
	WAN_9, WAN_9, WAN_9, WAN_9,

	TONG_1, TONG_1, TONG_1, TONG_1,
	TONG_2, TONG_2, TONG_2, TONG_2,
	TONG_3, TONG_3, TONG_3, TONG_3,
	TONG_4, TONG_4, TONG_4, TONG_4,
	TONG_5, TONG_5, TONG_5, TONG_5,
	TONG_6, TONG_6, TONG_6, TONG_6,
	TONG_7, TONG_7, TONG_7, TONG_7,
	TONG_8, TONG_8, TONG_8, TONG_8,
	TONG_9, TONG_9, TONG_9, TONG_9,

	TIAO_1, TIAO_1, TIAO_1, TIAO_1,
	TIAO_2, TIAO_2, TIAO_2, TIAO_2,
	TIAO_3, TIAO_3, TIAO_3, TIAO_3,
	TIAO_4, TIAO_4, TIAO_4, TIAO_4,
	TIAO_5, TIAO_5, TIAO_5, TIAO_5,
	TIAO_6, TIAO_6, TIAO_6, TIAO_6,
	TIAO_7, TIAO_7, TIAO_7, TIAO_7,
	TIAO_8, TIAO_8, TIAO_8, TIAO_8,
	TIAO_9, TIAO_9, TIAO_9, TIAO_9,
}

// 去除了缺一门花色的牌组
var cardsWithoutIgnore = map[ColorType]Cards{
	Wan: {
		TONG_1, TONG_2, TONG_3, TONG_4, TONG_5, TONG_6, TONG_7, TONG_8, TONG_9,
		TIAO_1, TIAO_2, TIAO_3, TIAO_4, TIAO_5, TIAO_6, TIAO_7, TIAO_8, TIAO_9,
	},

	Tiao: {
		TONG_1, TONG_2, TONG_3, TONG_4, TONG_5, TONG_6, TONG_7, TONG_8, TONG_9,
		WAN_1, WAN_2, WAN_3, WAN_4, WAN_5, WAN_6, WAN_7, WAN_8, WAN_9,
	},

	Tong: {
		WAN_1, WAN_2, WAN_3, WAN_4, WAN_5, WAN_6, WAN_7, WAN_8, WAN_9,
		TIAO_1, TIAO_2, TIAO_3, TIAO_4, TIAO_5, TIAO_6, TIAO_7, TIAO_8, TIAO_9,
	},
}
