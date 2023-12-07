package niuniu

import (
	"context"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/redis/go-redis/v9"
	"github.com/wwj31/dogactor/logger"

	"github.com/wwj31/dogactor/tools"

	"server/common"
	"server/proto/outermsg/outer"
	"server/rdsop"

	"server/proto/innermsg/inner"
	"server/service/room"
)

const (
	ReadyExpiration     = 5 * time.Second  // 倒计时进入发牌时间
	DealExpiration      = 3 * time.Second  // 发牌状态持续时间(参与人数*该值)
	MasterExpiration    = 10 * time.Second // 抢庄状态持续时间
	BettingExpiration   = 10 * time.Second // 押注状态持续时间
	ShowCardsExpiration = 10 * time.Second // 亮牌状态持续时间
	SettlementDuration  = 10 * time.Second // 结算持续时间
)

func New(r *room.Room) *NiuNiu {
	niuniu := &NiuNiu{
		room: r,
		fsm:  room.NewFSM(),
	}
	niuniu.niuniuPlayers = make([]*niuniuPlayer, 10, 10)
	niuniu.playerGameCount = make(map[int64]int32)
	_ = niuniu.fsm.Add(&StateReady{NiuNiu: niuniu})      // 准备中
	_ = niuniu.fsm.Add(&StateDeal{NiuNiu: niuniu})       // 发牌中
	_ = niuniu.fsm.Add(&StateMaster{NiuNiu: niuniu})     // 抢庄
	_ = niuniu.fsm.Add(&StateBetting{NiuNiu: niuniu})    // 押注
	_ = niuniu.fsm.Add(&StateShow{NiuNiu: niuniu})       // 开牌
	_ = niuniu.fsm.Add(&StateSettlement{NiuNiu: niuniu}) // 游戏结束结算界面

	niuniu.SwitchTo(Ready)

	return niuniu
}

type (
	// 牛牛 参与游戏的玩家数据
	niuniuPlayer struct {
		*room.Player
		score         int64 // 本局参与的实时分
		winScore      int64 // 单局的总输赢
		LastWinScore  int64 // 上一把总输赢
		ready         bool
		readyExpireAt time.Time
		handCards     PokerCards
		cardsGroup    CardsGroup
	}

	NiuNiu struct {
		room                *room.Room
		fsm                 *room.FSM
		currentStateEnterAt time.Time // 当前状态的进入时间
		currentStateEndAt   time.Time // 当前状态的结束时间
		onPlayerEnter       func(p *niuniuPlayer)
		onPlayerLeave       func(p *niuniuPlayer)

		playerGameCount  map[int64]int32 // 参与者的游戏次数
		masterIndex      int             // 庄家位置
		lastMasterShort  int64           // 上一把庄家shortId
		pushBetIndex     []int32         // 抢庄后，能推注的玩家位置
		randMasterSeat   []int32         // 参与随机选庄的位置
		niuniuPlayers    []*niuniuPlayer // 参与游戏的玩家  seat->player
		masterTimesSeats map[int32]int32 // 每个位置抢庄的倍数
		betGoldSeats     map[int32]int64 // 每个位置押注的分数
		showCards        map[int32]int32 // 亮牌的位置
		settlementMsg    *outer.NiuNiuSettlementNtf
	}
)

func (n *NiuNiu) SwitchTo(state int) {
	if err := n.fsm.SwitchTo(state); err != nil {
		current := n.fsm.CurrentStateHandler().State()
		n.Log().Errorw("NiuNiu switch to next state failed", "room", n.room.RoomId, "current", current)
		return
	}
	n.currentStateEnterAt = tools.Now()
}

func (n *NiuNiu) toRoomPlayers() (players []*room.Player) {
	n.RangePartInPlayer(func(seat int, player *niuniuPlayer) {
		players = append(players, player.Player)
	})

	return players
}

func (n *NiuNiu) Data(shortId int64) proto.Message {
	info := &outer.NiuNiuGameInfo{
		State:        outer.NiuNiuState(n.fsm.State()),
		StateEnterAt: n.currentStateEnterAt.UnixMilli(),
		StateEndAt:   n.currentStateEndAt.UnixMilli(),
		Players:      n.playersToPB(shortId),
		MasterIndex:  int32(n.masterIndex),
		MasterTimes:  n.masterTimesSeats,
		BetGold:      n.betGoldSeats,
		Settlement:   n.settlementMsg,
	}
	return info
}
func (n *NiuNiu) RangePartInPlayer(fn func(seat int, player *niuniuPlayer)) {
	for seatIndex, player := range n.niuniuPlayers {
		if player != nil && player.ready {
			fn(seatIndex, player)
		}
	}
}

func (n *NiuNiu) SeatIndex(shortId int64) int {
	for seatIndex, player := range n.niuniuPlayers {
		if player != nil && player.ShortId == shortId {
			return seatIndex
		}
	}
	return -1
}

func (n *NiuNiu) CanEnter(p *inner.PlayerInfo) bool {
	if p.Gold < n.baseScore() {
		return false
	}
	// 还有空位就能进
	for _, player := range n.niuniuPlayers {
		if player == nil {
			return true
		}
	}
	return false
}

func (n *NiuNiu) CanLeave(p *inner.PlayerInfo) bool {
	if player, _ := n.findNiuNiuPlayer(p.ShortId); player != nil {
		c, exist := n.playerGameCount[p.ShortId]
		if !exist {
			return true
		}
		if c < n.gameParams().PlayCountLimit && n.playerCount() > 1 {
			return false
		}

		// 只有准备和结算时可以离开
		switch n.fsm.State() {
		case Settlement, Ready:
			return true
		}
		return false
	}

	return true
}

func (n *NiuNiu) CanSetGold(p *inner.PlayerInfo) bool {
	if n.fsm.State() == Ready {
		return true
	}
	return false
}

// RecordingPlayback 当前状态是否需要记录回播内容
func (n *NiuNiu) CanRecordingPlayback() bool {
	if n.fsm.State() == Ready {
		return false
	}
	return true
}

func (n *NiuNiu) PlayerOnline(shortId int64) {}
func (n *NiuNiu) PlayerOffline(shortId int64) {
	if player, _ := n.findNiuNiuPlayer(shortId); player != nil && n.CanLeave(player.PlayerInfo) {
		n.room.PlayerLeave(shortId, true)
	}

}

func (n *NiuNiu) PlayerEnter(roomPlayer *room.Player) {
	// 找个位置就坐下
	for i, player := range n.niuniuPlayers {
		if player == nil {
			player = n.newNiuNiuPlayer(roomPlayer)
			n.niuniuPlayers[i] = player
			if n.onPlayerEnter != nil {
				n.onPlayerEnter(player)
			}
			break
		}
	}
}

func (n *NiuNiu) PlayerLeave(quitPlayer *room.Player) {
	for idx, player := range n.niuniuPlayers {
		if player != nil && player.ShortId == quitPlayer.ShortId {
			delete(n.playerGameCount, player.ShortId)
			n.room.CancelTimer(quitPlayer.RID)
			n.niuniuPlayers[idx] = nil
			n.Log().Infow("player leave niu niu ", "shortId", player.ShortId, "seat", idx, "gold", player.Gold)
			if n.onPlayerLeave != nil {
				n.onPlayerLeave(player)
			}
			return
		}
	}
}

// Handle 游戏消息，全部交由当前状态处理
func (n *NiuNiu) Handle(shortId int64, v any) any {
	return n.fsm.CurrentStateHandler().Handle(shortId, v)
}

func (n *NiuNiu) Log() *logger.Logger {
	return n.room.Log()
}

func (n *NiuNiu) findNiuNiuPlayer(shortId int64) (*niuniuPlayer, int) {
	for i, player := range n.niuniuPlayers {
		if player != nil && player.ShortId == shortId {
			return player, i
		}
	}
	return nil, -1
}
func (n *NiuNiu) participantCount() (c int) {
	n.RangePartInPlayer(func(seat int, player *niuniuPlayer) { c++ })
	return c
}

func (n *NiuNiu) playersToPB(shortId int64) (players []*outer.NiuNiuPlayerInfo) {
	for _, player := range n.niuniuPlayers {
		if player == nil {
			players = append(players, nil)
		} else {
			var (
				handCards []int32
				state     = n.fsm.State()
			)
			switch {
			case state == Settlement:
				handCards = player.handCards.ToPB() // 结算时，牌全发
			case state == ShowCards:
				// 亮牌状态，是自己或者已经亮牌了，牌全发
				if player.ShortId == shortId || n.showCards[int32(n.SeatIndex(player.ShortId))] == 1 {
					handCards = player.handCards.ToPB()
				} else {
					handCards = make([]int32, len(player.handCards))
				}
			case state < ShowCards:
				// 抢庄、押注状态，自己就发前4张，其他人不发
				if player.ShortId == shortId && len(player.handCards) >= 5 {
					handCards = player.handCards[0:4].ToPB()
					handCards = append(handCards, 0)
				} else {
					handCards = make([]int32, 5)
				}
			}

			players = append(players, &outer.NiuNiuPlayerInfo{
				ShortId:       player.ShortId,
				Ready:         player.ready,
				ReadyExpireAt: player.readyExpireAt.UnixMilli(),
				HandCards:     handCards,
				CardsType:     player.cardsGroup.ToPB(),
				Score:         player.score,
				CanPushBet:    n.canPushBet(player.ShortId) == outer.ERROR_OK,
			})
		}
	}
	return
}

func (n *NiuNiu) newNiuNiuPlayer(p *room.Player) *niuniuPlayer {
	return &niuniuPlayer{
		score:  p.Gold,
		Player: p,
	}
}

// 开局扣门票
func (n *NiuNiu) masterRebate() {
	if n.gameParams().MasterRebate <= 0 {
		return
	}

	record := &outer.RebateDetailInfo{
		Type:      outer.GameType_NiuNiu,
		BaseScore: n.gameParams().BaseScore,
		CreateAt:  tools.Now().UnixMilli(),
	}

	var pip redis.Pipeliner
	score := int64(n.gameParams().MasterRebate * common.Gold1000Times) // 门票
	n.RangePartInPlayer(func(seat int, player *niuniuPlayer) {
		player.score = common.Max(player.score-score, 0)
		record.Gold = score
		record.ShortId = player.ShortId
		rdsop.RecordRebateGold(common.JsonMarshal(record), n.room.MasterShortId, score, pip)
	})

	if _, err := pip.Exec(context.Background()); err != nil {
		n.Log().Errorw("master rebate redis failed", "err", err)
	}
}

func (n *NiuNiu) gameParams() *outer.NiuNiuParams {
	return n.room.GameParams.NiuNiu
}

func (n *NiuNiu) baseScore() int64 {
	base := n.gameParams().BaseScore
	if base <= 0 {
		base = 1
	}

	return int64(base * common.Gold1000Times)
}

func (n *NiuNiu) bombWinScore() int64 {
	base := n.gameParams().BaseScore
	if base == 0 {
		base = 1
	}

	return int64(float32(base*common.Gold1000Times) * 5)
}

func (n *NiuNiu) allSeats(ignoreSeat ...int) (result []int) {
	seatMap := map[int]struct{}{}
	for _, seat := range ignoreSeat {
		seatMap[seat] = struct{}{}
	}

	for seatIndex := 0; seatIndex < len(n.room.Players); seatIndex++ {
		if _, ignore := seatMap[seatIndex]; !ignore {
			result = append(result, seatIndex)
		}
	}

	return result
}

func (m *niuniuPlayer) updateScore(val int64) {
	m.score += val
	m.winScore += val // 单局总输赢
	m.LastWinScore += val
}

func (n *NiuNiu) playerCount() int32 {
	var count int32

	for _, p := range n.niuniuPlayers {
		if p != nil {
			count++
		}
	}
	return count
}

func (n *NiuNiu) canPushBet(shortId int64) outer.ERROR {
	if n.fsm.State() == Ready || n.fsm.State() >= ShowCards {
		return outer.ERROR_NIUNIU_DISALLOW_PUSH_WITH_NOT_BETTING
	}

	player, seat := n.findNiuNiuPlayer(shortId)
	// 上把没赢钱，或者赢钱小于5倍底分，不能推注
	if player.LastWinScore <= 0 || player.LastWinScore < n.baseScore()*5 {
		return outer.ERROR_NIUNIU_DISALLOW_PUSH_WITH_NOT_WIN
	}

	// 上把是庄家不能推注
	if player.ShortId == n.lastMasterShort {
		return outer.ERROR_NIUNIU_DISALLOW_PUSH_WITH_LAST_IS_MASTER
	}

	// 发牌和抢庄阶段
	if n.fsm.State() == Deal {
		return outer.ERROR_OK
	}

	// 剩下的判断只可能是在,抢庄和押注阶段
	var can bool
	for _, index := range n.pushBetIndex {
		if seat == int(index) {
			can = true
			break
		}
	}

	//  没有抢过最大倍数庄，不允许推注
	if !can {
		return outer.ERROR_NIUNIU_DISALLOW_PUSH_WITH_NOT_MASTER
	}

	return outer.ERROR_OK
}

func (n *NiuNiu) clear() {
	for seatIndex, oldPlayer := range n.niuniuPlayers {
		if oldPlayer != nil {
			n.niuniuPlayers[seatIndex] = n.newNiuNiuPlayer(oldPlayer.Player)
			n.niuniuPlayers[seatIndex].LastWinScore = oldPlayer.LastWinScore
		}
	}
}
