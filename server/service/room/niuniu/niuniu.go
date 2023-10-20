package niuniu

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/logger"

	"server/common"
	"server/proto/outermsg/outer"

	"github.com/wwj31/dogactor/tools"

	"server/proto/innermsg/inner"
	"server/service/room"
)

const (
	ReadyExpiration     = 5 * time.Second  // 倒计时进入发牌时间
	DealExpiration      = 3 * time.Second  // 发牌状态持续时间
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
	_ = niuniu.fsm.Add(&StateSettlement{NiuNiu: niuniu}) // 游戏结束结算界面

	niuniu.SwitchTo(Ready)

	return niuniu
}

type (
	// 牛牛 参与游戏的玩家数据
	niuniuPlayer struct {
		*room.Player
		score         int64
		totalWinScore int64 // 单局的总输赢
		ready         bool
		readyExpireAt time.Time
		handCards     PokerCards
		finalStatsMsg *outer.NiuNiuFinialPlayerInfo
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
		niuniuPlayers    []*niuniuPlayer // 参与游戏的玩家  seat->player
		masterTimesSeats map[int32]int32 // 每个位置抢庄的倍数
		betTimesSeats    map[int32]int32 // 每个位置押注的倍数
		shows            map[int32]int32 // 亮牌的位置
	}
)

func (f *NiuNiu) SwitchTo(state int) {
	if err := f.fsm.SwitchTo(state); err != nil {
		current := f.fsm.CurrentStateHandler().State()
		f.Log().Errorw("NiuNiu switch to next state failed", "room", f.room.RoomId, "current", current)
		return
	}
	f.currentStateEnterAt = tools.Now()
}

func (f *NiuNiu) toRoomPlayers() (players []*room.Player) {
	f.RangePartInPlayer(func(seat int, player *niuniuPlayer) {
		players = append(players, player.Player)
	})

	return players
}

func (f *NiuNiu) Data(shortId int64) proto.Message {
	info := &outer.NiuNiuGameInfo{
		State:        outer.NiuNiuState(f.fsm.State()),
		StateEnterAt: f.currentStateEnterAt.UnixMilli(),
		StateEndAt:   f.currentStateEndAt.UnixMilli(),
		Players:      f.playersToPB(shortId),
		MasterIndex:  int32(f.masterIndex),
		MasterTimes:  f.masterTimesSeats,
		BetTimes:     f.betTimesSeats,
	}
	return info
}
func (f *NiuNiu) RangePartInPlayer(fn func(seat int, player *niuniuPlayer)) {
	for seatIndex, player := range f.niuniuPlayers {
		if player != nil && player.ready {
			fn(seatIndex, player)
		}
	}
}

func (f *NiuNiu) SeatIndex(shortId int64) int {
	for seatIndex, player := range f.niuniuPlayers {
		if player != nil && player.ShortId == shortId {
			return seatIndex
		}
	}
	return -1
}

func (f *NiuNiu) CanEnter(p *inner.PlayerInfo) bool {
	// 还有空位就能进
	for _, player := range f.niuniuPlayers {
		if player == nil {
			return true
		}
	}
	return false
}

func (f *NiuNiu) CanLeave(p *inner.PlayerInfo) bool {
	if player, _ := f.findNiuNiuPlayer(p.ShortId); player != nil {
		c, exist := f.playerGameCount[p.ShortId]
		if !exist {
			return true
		}
		if c < f.gameParams().PlayCountLimit {
			return false
		}

		// 只有准备和结算时可以离开
		switch f.fsm.State() {
		case Settlement, Ready:
			return true
		}
		return false
	}

	return true
}

func (f *NiuNiu) CanSetGold(p *inner.PlayerInfo) bool {
	if f.fsm.State() == Ready {
		return true
	}
	return false
}

func (f *NiuNiu) PlayerEnter(roomPlayer *room.Player) {
	// 找个位置就坐下
	for i, player := range f.niuniuPlayers {
		if player == nil {
			player = f.newNiuNiuPlayer(roomPlayer)
			player.finalStatsMsg = &outer.NiuNiuFinialPlayerInfo{}
			f.niuniuPlayers[i] = player
			if f.onPlayerEnter != nil {
				f.onPlayerEnter(player)
			}
			break
		}
	}
}

func (f *NiuNiu) PlayerLeave(quitPlayer *room.Player) {
	for idx, player := range f.niuniuPlayers {
		if player != nil && player.ShortId == quitPlayer.ShortId {
			delete(f.playerGameCount, player.ShortId)
			f.room.CancelTimer(quitPlayer.RID)
			f.niuniuPlayers[idx] = nil
			f.Log().Infow("player leave mahjong", "shortId", player.ShortId, "seat", idx, "gold", player.Gold)
			if f.onPlayerLeave != nil {
				f.onPlayerLeave(player)
			}
			return
		}
	}
}

// Handle 游戏消息，全部交由当前状态处理
func (f *NiuNiu) Handle(shortId int64, v any) any {
	return f.fsm.CurrentStateHandler().Handle(shortId, v)
}

func (f *NiuNiu) Log() *logger.Logger {
	return f.room.Log()
}

func (f *NiuNiu) findNiuNiuPlayer(shortId int64) (*niuniuPlayer, int) {
	for i, player := range f.niuniuPlayers {
		if player != nil && player.ShortId == shortId {
			return player, i
		}
	}
	return nil, -1
}
func (f *NiuNiu) participantCount() int {
	var c int
	f.RangePartInPlayer(func(seat int, player *niuniuPlayer) { c++ })
	return c
}

func (f *NiuNiu) playersToPB(shortId int64) (players []*outer.NiuNiuPlayerInfo) {
	for _, player := range f.niuniuPlayers {
		if player == nil {
			players = append(players, nil)
		} else {
			var (
				handCards []int32
				state     = f.fsm.State()
			)
			switch {
			case state == Settlement:
				handCards = player.handCards.ToPB() // 结算时，牌全发
			case state == ShowCards:
				// 亮牌状态，是自己或者已经亮牌了，牌全发
				if player.ShortId == shortId || f.shows[int32(f.SeatIndex(player.ShortId))] == 1 {
					handCards = player.handCards.ToPB()
				} else {
					handCards = make([]int32, len(player.handCards))
				}
			case state < ShowCards:
				// 抢庄、押注状态，自己就发前4张，其他人不发
				if player.ShortId == shortId {
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
				Score:         player.score,
			})
		}
	}
	return
}

func (f *NiuNiu) newNiuNiuPlayer(p *room.Player) *niuniuPlayer {
	return &niuniuPlayer{
		score:  p.Gold,
		Player: p,
	}
}

func (f *NiuNiu) gameParams() *outer.NiuNiuParams {
	return f.room.GameParams.NiuNiu
}

func (f *NiuNiu) baseScore() int64 {
	base := f.gameParams().BaseScore
	if base <= 0 {
		base = 1
	}

	return int64(base * common.Gold1000Times)
}

func (f *NiuNiu) bombWinScore() int64 {
	base := f.gameParams().BaseScore
	if base == 0 {
		base = 1
	}

	return int64(float32(base*common.Gold1000Times) * 5)
}

func (f *NiuNiu) allSeats(ignoreSeat ...int) (result []int) {
	seatMap := map[int]struct{}{}
	for _, seat := range ignoreSeat {
		seatMap[seat] = struct{}{}
	}

	for seatIndex := 0; seatIndex < len(f.room.Players); seatIndex++ {
		if _, ignore := seatMap[seatIndex]; !ignore {
			result = append(result, seatIndex)
		}
	}

	return result
}

func (m *niuniuPlayer) updateScore(val int64) {
	m.score += val
	m.totalWinScore += val // 单局总输赢
	m.finalStatsMsg.TotalScore += val
}
