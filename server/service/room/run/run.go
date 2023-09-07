package run

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/logger"

	"server/proto/outermsg/outer"

	"github.com/wwj31/dogactor/tools"

	"server/proto/innermsg/inner"
	"server/service/room"
)

func New(r *room.Room) *FasterRun {
	fasterRun := &FasterRun{
		room: r,
		fsm:  room.NewFSM(),
	}
	//_ = fasterRun.fsm.Add(&StateReady{Mahjong: fasterRun})      // 准备中
	//_ = fasterRun.fsm.Add(&StateDeal{Mahjong: fasterRun})       // 发牌中
	//_ = fasterRun.fsm.Add(&StatePlaying{Mahjong: fasterRun})    // 游戏中
	//_ = fasterRun.fsm.Add(&StateSettlement{Mahjong: fasterRun}) // 游戏结束结算界面

	fasterRun.SwitchTo(Ready)

	return fasterRun
}

const maxNum = 2

type (
	// 跑得快 参与游戏的玩家数据
	fasterRunPlayer struct {
		*room.Player
		score         int64
		totalWinScore int64 // 单局的总输赢
		ready         bool
		readyExpireAt time.Time
		finalStatsMsg *outer.MahjongBTEFinialPlayerInfo
	}

	FasterRun struct {
		room                *room.Room
		fsm                 *room.FSM
		currentStateEnterAt time.Time                            // 当前状态的进入时间
		currentStateEndAt   time.Time                            // 当前状态的结束时间
		playerAutoReady     func(p *fasterRunPlayer, ready bool) //
		fasterRunPlayers    [maxNum]*fasterRunPlayer             // 参与游戏的玩家

	}
)

func (f *FasterRun) SwitchTo(state int) {
	if err := f.fsm.SwitchTo(state); err != nil {
		current := f.fsm.CurrentStateHandler().State()
		f.Log().Errorw("FasterRun switch to next state failed", "room", f.room.RoomId, "current", current)
		return
	}
	f.currentStateEnterAt = tools.Now()
}

func (f *FasterRun) Data(shortId int64) proto.Message {
	info := &outer.MahjongBTEGameInfo{
		State:        outer.MahjongBTEState(f.fsm.State()),
		StateEnterAt: f.currentStateEnterAt.UnixMilli(),
		StateEndAt:   f.currentStateEndAt.UnixMilli(),
	}
	return info
}

func (f *FasterRun) SeatIndex(shortId int64) int {
	for seatIndex, player := range f.fasterRunPlayers {
		if player != nil && player.ShortId == shortId {
			return seatIndex
		}
	}
	return -1
}

func (f *FasterRun) CanEnter(p *inner.PlayerInfo) bool {
	if f.fsm.State() == Ready {
		return true
	}

	return false
}

func (f *FasterRun) CanLeave(p *inner.PlayerInfo) bool {
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

func (f *FasterRun) CanReady(p *inner.PlayerInfo) bool {
	if f.fsm.State() == Ready {
		return true
	}
	return false
}

func (f *FasterRun) CanSetGold(p *inner.PlayerInfo) bool {
	if f.fsm.State() == Ready {
		return true
	}
	return false
}

func (f *FasterRun) PlayerEnter(roomPlayer *room.Player) {
	for i, player := range f.fasterRunPlayers {
		if player == nil {
			player = f.newFasterRunPlayer(roomPlayer)
			player.finalStatsMsg = &outer.MahjongBTEFinialPlayerInfo{}
			f.fasterRunPlayers[i] = player
			if f.playerAutoReady != nil {
				f.playerAutoReady(player, false)
			}
			break
		}
	}
}

func (f *FasterRun) readyAfterTimeout(player *fasterRunPlayer, expireAt time.Time) {
	f.room.AddTimer(player.RID, expireAt, func(dt time.Duration) {
		f.Log().Infow("the player was kicked out of the room due to a timeout in the ready period",
			"room", f.room.RoomId, "player", player.ShortId)
		f.room.PlayerLeave(player.ShortId, true)
	})
}

func (f *FasterRun) PlayerLeave(quitPlayer *room.Player) {
	for idx, player := range f.fasterRunPlayers {
		if player != nil && player.ShortId == quitPlayer.ShortId {
			f.room.CancelTimer(quitPlayer.RID)
			f.fasterRunPlayers[idx] = nil
			f.Log().Infow("player leave mahjong", "shortId", player.ShortId, "seat", idx, "gold", player.Gold)
			return
		}
	}
}

// Handle 麻将游戏消息，全部交由当前状态处理
func (f *FasterRun) Handle(shortId int64, v any) any {
	return f.fsm.CurrentStateHandler().Handle(shortId, v)
}

func (f *FasterRun) Log() *logger.Logger {
	return f.room.Log()
}

func (f *FasterRun) findFasterRunPlayer(shortId int64) (*fasterRunPlayer, int) {
	for i, player := range f.fasterRunPlayers {
		if player != nil && player.ShortId == shortId {
			return player, i
		}
	}
	return nil, -1
}

// 逆时针轮动座位索引,index 当前位置
func (f *FasterRun) nextSeatIndex(index int) int {
	index--
	if index < 0 {
		index = 1
	}
	return index
}

func (f *FasterRun) newFasterRunPlayer(p *room.Player) *fasterRunPlayer {
	return &fasterRunPlayer{
		score:  p.Gold,
		Player: p,
	}
}

func (f *FasterRun) gameParams() *outer.FasterRunParams {
	return f.room.GameParams.FasterRun
}

func (f *FasterRun) clear() {
	// 重置玩家数据
	for i := 0; i < maxNum; i++ {
		gamer := f.fasterRunPlayers[i]
		if gamer != nil {
			f.fasterRunPlayers[i] = f.newFasterRunPlayer(gamer.Player)
			f.fasterRunPlayers[i].finalStatsMsg = gamer.finalStatsMsg
		}
	}
}

func (f *FasterRun) allSeats(ignoreSeat ...int) (result []int) {
	seatMap := map[int]struct{}{}
	for _, seat := range ignoreSeat {
		seatMap[seat] = struct{}{}
	}

	for seatIndex := 0; seatIndex < maxNum; seatIndex++ {
		if _, ignore := seatMap[seatIndex]; !ignore {
			result = append(result, seatIndex)
		}
	}

	return result
}

func (m *fasterRunPlayer) updateScore(val int64) {
	m.score += val
	m.totalWinScore += val // 单局总输赢
	m.finalStatsMsg.TotalScore += val
}
