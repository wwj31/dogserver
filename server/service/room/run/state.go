package run

type State = int

const (
	Ready      State = iota // 准备状态
	Deal                    // 游戏开始发牌状态
	Playing                 // 游戏状态
	Settlement              // 结算
)
