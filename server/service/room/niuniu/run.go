package niuniu

import (
	"fmt"
	"time"

	"server/common"
	"server/common/log"

	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/logger"

	"server/proto/outermsg/outer"

	"github.com/wwj31/dogactor/tools"

	"server/proto/innermsg/inner"
	"server/service/room"
)

const (
	ReadyExpiration    = 20 * time.Second // 准备超时时间
	DealExpiration     = 3 * time.Second  // 发牌状态持续时间
	SettlementDuration = 10 * time.Second // 结算持续时间
)

func New(r *room.Room) *NiuNiu {
	niuniu := &NiuNiu{
		room: r,
		fsm:  room.NewFSM(),
	}
	n := niuniu.playerNumber()
	niuniu.niuniuPlayers = make([]*niuniuPlayer, n, n)
	_ = niuniu.fsm.Add(&StateReady{NiuNiu: niuniu})      // 准备中
	_ = niuniu.fsm.Add(&StateDeal{NiuNiu: niuniu})       // 发牌中
	_ = niuniu.fsm.Add(&StateSettlement{NiuNiu: niuniu}) // 游戏结束结算界面

	niuniu.SwitchTo(Ready)

	return niuniu
}

type (
	// 跑得快 参与游戏的玩家数据
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
		currentStateEnterAt time.Time                         // 当前状态的进入时间
		currentStateEndAt   time.Time                         // 当前状态的结束时间
		playerAutoReady     func(p *niuniuPlayer, ready bool) //
		scoreZeroOver       bool                              // 因为有玩家没分了，而触发的结束

		masterIndex   int             // 庄家位置 0，1，2
		niuniuPlayers []*niuniuPlayer // 参与游戏的玩家
		gameCount     int             // 游戏的连续局数
	}

	PlayCardsRecord struct {
		shortId    int64      // 出牌人
		follow     bool       // true.跟牌出牌，false.有牌权出牌
		cardsGroup CardsGroup // 牌型
		playAt     time.Time  // 出牌时间
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
	for _, p := range f.niuniuPlayers {
		players = append(players, p.Player)
	}
	return players
}

func (f *NiuNiu) playerNumber() int {
	n := f.gameParams().PlayerNumber
	if n == 0 {
		return 2
	} else if n == 1 {
		return 3
	}
	log.Errorw("the player number is unexpected", "number", n)
	return 0
}

func (f *NiuNiu) Data(shortId int64) proto.Message {
	info := &outer.NiuNiuGameInfo{
		State:        outer.NiuNiuState(f.fsm.State()),
		StateEnterAt: f.currentStateEnterAt.UnixMilli(),
		StateEndAt:   f.currentStateEndAt.UnixMilli(),
		GameCount:    int32(f.gameCount),
		Players:      f.playersToPB(shortId, false),
		MasterIndex:  int32(f.masterIndex),
	}
	return info
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
	if f.fsm.State() == Ready {
		return true
	}

	return false
}

func (f *NiuNiu) CanLeave(p *inner.PlayerInfo) bool {
	if p.Gold <= 0 {
		return true
	}

	// 只有准备和结算时可以离开
	switch f.fsm.State() {
	case Settlement, Ready:
		return true
	}

	return false
}

func (f *NiuNiu) CanReady(p *inner.PlayerInfo) bool {
	if f.fsm.State() == Ready {
		return true
	}
	return false
}

func (f *NiuNiu) CanSetGold(p *inner.PlayerInfo) bool {
	if f.fsm.State() == Ready {
		return true
	}
	return false
}

func (f *NiuNiu) PlayerEnter(roomPlayer *room.Player) {
	for i, player := range f.niuniuPlayers {
		if player == nil {
			player = f.newNiuNiuPlayer(roomPlayer)
			player.finalStatsMsg = &outer.NiuNiuFinialPlayerInfo{}
			f.niuniuPlayers[i] = player
			if f.playerAutoReady != nil {
				f.playerAutoReady(player, false)
			}
			break
		}
	}
}

func (f *NiuNiu) readyAfterTimeout(player *niuniuPlayer, expireAt time.Time) {
	f.room.AddTimer(player.RID, expireAt, func(dt time.Duration) {
		f.Log().Infow("the player was kicked out of the room due to a timeout in the ready period",
			"room", f.room.RoomId, "player", player.ShortId)
		f.room.PlayerLeave(player.ShortId, true)
	})
}

func (f *NiuNiu) PlayerLeave(quitPlayer *room.Player) {
	for idx, player := range f.niuniuPlayers {
		if player != nil && player.ShortId == quitPlayer.ShortId {
			f.room.CancelTimer(quitPlayer.RID)
			f.niuniuPlayers[idx] = nil
			f.Log().Infow("player leave mahjong", "shortId", player.ShortId, "seat", idx, "gold", player.Gold)
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

func (f *NiuNiu) playersToPB(shortId int64, settlement bool) (players []*outer.NiuNiuPlayerInfo) {
	for _, player := range f.niuniuPlayers {
		if player == nil {
			players = append(players, nil)
		} else {
			var handCards []int32
			if player.ShortId == shortId || settlement {
				handCards = player.handCards.ToPB()
			} else {
				handCards = make([]int32, len(player.handCards))
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

// 逆时针轮动座位索引,index 当前位置
func (f *NiuNiu) nextSeatIndex(index int) int {
	index--
	if index < 0 {
		index = 1
	}
	return index
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

func (f *NiuNiu) clear() {
	f.scoreZeroOver = false

	// 重置玩家数据
	for i := 0; i < f.playerNumber(); i++ {
		gamer := f.niuniuPlayers[i]
		if gamer != nil {
			f.niuniuPlayers[i] = f.newNiuNiuPlayer(gamer.Player)
			f.niuniuPlayers[i].finalStatsMsg = gamer.finalStatsMsg
		}
	}
}

func (f *NiuNiu) allSeats(ignoreSeat ...int) (result []int) {
	seatMap := map[int]struct{}{}
	for _, seat := range ignoreSeat {
		seatMap[seat] = struct{}{}
	}

	for seatIndex := 0; seatIndex < f.playerNumber(); seatIndex++ {
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

func (p *PlayCardsRecord) String() string {
	if p == nil {
		return ""
	}
	return fmt.Sprintf("{short:%v follow:%v cardsGroup:%v playAt:%v }", p.shortId, p.follow, p.cardsGroup, p.playAt)
}

func (p *PlayCardsRecord) ToPB() *outer.PlayCardsRecord {
	if p == nil {
		return nil
	}

	return &outer.PlayCardsRecord{
		ShortId:    p.shortId,
		Follow:     p.follow,
		CardsGroup: p.cardsGroup.ToPB(),
		PlayAt:     p.playAt.UnixMilli(),
	}
}
