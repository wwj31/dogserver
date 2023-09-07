package mahjong

import "server/proto/outermsg/outer"

type HuType int32

func HuTypePB() map[int32]int32 {
	result := make(map[int32]int32)
	for huType, num := range huFan {
		result[int32(huType)] = int32(num)
	}
	return result
}

const (
	HuInvalid     HuType = iota
	Hu                   // 平胡      0
	DuiDuiHu             // 对对胡    1
	QingYiSe             // 清一色    2
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
	JiaXinWu             // 夹心五      1
)

type ExtFanType int32

func ExtFanTypePB() map[int32]int32 {
	result := make(map[int32]int32)
	for extraType, num := range extraFan {
		result[int32(extraType)] = int32(num)
	}
	return result
}

const (
	ExtraInvalid ExtFanType = iota
	GangShangHua            // 杠上开花
	GangShangPao            // 杠上炮
	QiangGangHu             // 抢杠胡
	ShaoDiHu                // 扫底胡
	JinGouGou               // 金钩胡
	HaiDiPao                // 海底炮
	TianHu                  // 天胡
	Dihu                    // 地胡
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

func ExtFanArrToPB(ext []ExtFanType) []outer.ExtraType {
	var result []outer.ExtraType
	for _, v := range ext {
		result = append(result, v.PB())
	}
	return result
}

func ExtraWithin(extras []ExtFanType, v ExtFanType) bool {
	for _, ext := range extras {
		if ext == v {
			return true
		}
	}
	return false
}
