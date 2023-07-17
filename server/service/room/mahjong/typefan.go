package mahjong

var (
	huFan = map[HuType]int{
		Hu:            0, // 平胡
		QingYiSe:      2, // 清一色
		DuiDuiHu:      1, // 对对胡
		QingDui:       3, // 清对
		JiangDui:      4, // 将对
		QiDui:         2, // 七对
		LongQiDui:     2, // 龙七对 (龙七对3番=七对2番+根1番)
		QingQiDui:     4, // 清七对
		JiangQiDui:    4, // 将七对
		QingLongQiDui: 4, // 清龙七对 (清龙七对5番=清七对4番+根1番)
		QuanYaoJiu:    3, // 全幺九
		MenQing:       1, // 门清
		ZhongZhang:    1, // 中张
		JiaXinWu:      1, // 夹心五
	}

	extraFan = map[ExtFanType]int{
		GangShangHua: 1, // 杠上开花
		GangShangPao: 1, // 杠上炮
		QiangGangHu:  1, // 抢杠胡
		ShaoDiHu:     1, // 扫底胡
		JinGouGou:    1, // 金钩胡
		HaiDiPao:     1, // 海底炮
		TianHu:       3, // 天胡
		Dihu:         2, // 地胡
	}
)
