package mahjong

type State = int

const (
	Ready        State = iota // 准备状态
	DecideMaster              // 定庄
	Deal                      // 游戏开始发牌状态
	Exchange3                 // 换三张状态
	DecideIgnore              // 定缺状态
	Playing                   // 游戏状态
	Settlement                // 结算
)
