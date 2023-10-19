package niuniu

type State = int

const (
	Ready        State = iota // 准备状态
	Deal                      // 游戏开始发牌状态
	DecideMaster              // 抢庄
	Betting                   // 押注
	ShowCards                 // 搓牌
	Settlement                // 结算
)
