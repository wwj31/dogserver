package mahjong

type HuType int32
type ExtFanType int32

const (
	HuInvalid  HuType = iota
	Hu                // 平胡
	DuiDuiHu          // 对对胡
	QingYiSe          // 清一色
	QiDui             // 七对
	LongQiDui         // 龙七对
	QingDui           // 清对
	QingQiDui         // 清七对
	QuanYaoJiu        // 全幺九
	JiangDui          // 将对
	JiangQiDui        // 将七对
	MenQing           // 门清
	ZhongZhang        // 中张
	JiaXinWu          // 夹心五
)

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

var hu = map[HuType]string{
	HuInvalid:  "未胡牌",
	Hu:         "平胡",
	DuiDuiHu:   "对对胡",
	QingYiSe:   "清一色",
	QiDui:      "七对",
	LongQiDui:  "龙七对",
	QingDui:    "清对",
	QingQiDui:  "清七对",
	QuanYaoJiu: "全幺九",
	JiangDui:   "将对",
	JiangQiDui: "将七对",
	MenQing:    "门清",
	ZhongZhang: "中张",
	JiaXinWu:   "夹心五",
}

var extra = map[ExtFanType]string{
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

func (h HuType) String() string {
	return hu[h]
}

func (e ExtFanType) String() string {
	return extra[e]
}
