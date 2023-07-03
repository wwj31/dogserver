package mahjong

import "server/proto/outermsg/outer"

type HuType int32

const (
	HuInvalid     HuType = iota
	Hu                   // 平胡      0
	QingYiSe             // 清一色    2
	DuiDuiHu             // 对对胡    1
	QingDui              // 清对      3
	JiangDui             // 将对      3
	QiDui                // 七对      2
	LongQiDui            // 龙七对    3
	QingQiDui            // 清七对    4
	JiangQiDui           // 将七对    4
	QingLongQiDui        // 清龙七对   5
	QuanYaoJiu           // 全幺九    3
	MenQing              // 门清      1
	ZhongZhang           // 中张      1
)

type ExtFanType int32

const (
	ExtraInvalid ExtFanType = iota
	Gen                     // 根
	GangShangHua            // 杠上开花
	GangShangPao            // 杠上炮
	QiangGangHu             // 抢杠胡
	ShaoDiHu                // 扫底胡
	JinGouGou               // 金钩胡
	HaiDiPao                // 海底炮
	TianHu                  // 天胡
	Dihu                    // 地胡
)

type ColorType int32

const (
	ColorUnknown ColorType = iota
	Wan                    = 1 // 萬
	Tiao                   = 2 // 条
	Tong                   = 3 // 筒
)

var huStr = map[HuType]string{
	HuInvalid:     "未胡牌",
	Hu:            "平胡",
	DuiDuiHu:      "对对胡",
	QingYiSe:      "清一色",
	QiDui:         "七对",
	LongQiDui:     "龙七对",
	QingDui:       "清对",
	QingQiDui:     "清七对",
	QingLongQiDui: "清龙对",
	QuanYaoJiu:    "全幺九",
	JiangDui:      "将对",
	JiangQiDui:    "将七对",
	MenQing:       "门清",
	ZhongZhang:    "中张",
}

var extraStr = map[ExtFanType]string{
	Gen:          "根",
	GangShangHua: "杠上开花",
	GangShangPao: "杠上炮",
	QiangGangHu:  "抢杠胡",
	ShaoDiHu:     "扫底胡",
	JinGouGou:    "金钩胡",
	HaiDiPao:     "海底炮",
	TianHu:       "天胡",
	Dihu:         "地胡",
}

var colorsStr = map[ColorType]string{
	Wan:  "萬",
	Tiao: "条",
	Tong: "筒",
}

func (h HuType) String() string {
	return huStr[h]
}

func (e ExtFanType) String() string {
	return extraStr[e]
}

func (c ColorType) String() string {
	return colorsStr[c]
}

func (h HuType) PB() outer.HuType {
	return outer.HuType(h)
}

func (e ExtFanType) PB() outer.ExtraType {
	return outer.ExtraType(e)
}

func (c ColorType) PB() outer.ColorType {
	return outer.ColorType(c)
}
