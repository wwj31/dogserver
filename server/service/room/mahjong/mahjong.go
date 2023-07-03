package mahjong

import (
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/proto/innermsg/inner"
	"server/service/room"
)

const ReadyTimeout = 20 * time.Second

func New(r *room.Room) *Mahjong {
	mahjong := &Mahjong{
		room: r,
		fsm:  room.NewFSM(),
	}
	_ = mahjong.fsm.Add(&StateReady{Mahjong: mahjong})        // 准备中
	_ = mahjong.fsm.Add(&StateDeal{Mahjong: mahjong})         // 发牌中
	_ = mahjong.fsm.Add(&StateDecideMaster{Mahjong: mahjong}) // 定庄
	_ = mahjong.fsm.Add(&StateDecideMaster{Mahjong: mahjong}) //
	_ = mahjong.fsm.Add(&StateDecideMaster{Mahjong: mahjong})

	return mahjong
}

// 麻将-血战到底 参与游戏的玩家数据
type mahjongPlayer struct {
	*room.Player

	ignoreColor ColorType       // 定缺花色
	handCards   Cards           // 手牌
	lightGang   map[int32]int64 // map[杠牌]ShortId 明杠
	darkGang    map[int32]int64 // map[杠牌]ShortId 暗杠
	pong        map[int32]int64 // map[碰牌]ShortId
}

type Mahjong struct {
	room                *room.Room
	fsm                 *room.FSM
	currentStateEnterAt time.Time // 当前状态的进入时间

	masterIndex int // 庄家位置 0,1,2,3
	gameCount   int // 游戏的连续局数 结算后，有玩家退出，重置0

	cards          Cards // 剩余牌组
	mahjongPlayers [4]*mahjongPlayer
}

func (m *Mahjong) SwitchTo(state int) {
	if err := m.fsm.SwitchTo(state); err != nil {
		current := m.fsm.CurrentStateHandler().State()
		log.Errorw("Mahjong switch to next state failed",
			append([]any{"current", current, "next", state}, m.room.LogInfo())...)
		return
	}
	m.currentStateEnterAt = tools.Now()
}

func (m *Mahjong) SeatIndex(shortId int64) int32 {
	for i, player := range m.mahjongPlayers {
		if player != nil && player.ShortId == shortId {
			return int32(i)
		}
	}
	return -1
}

func (m *Mahjong) CanEnter(p *inner.PlayerInfo) bool {
	if m.fsm.State() == Ready {
		return true
	}
	return false
}

func (m *Mahjong) CanLeave(p *inner.PlayerInfo) bool {
	// 只有准备和结算时可以离开
	switch m.fsm.State() {
	case Ready, Settlement:
		return true
	}
	return false
}
func (m *Mahjong) CanReady(p *inner.PlayerInfo) bool {
	if m.fsm.State() == Ready {
		return true
	}
	return false
}

func (m *Mahjong) PlayerEnter(p *room.Player) {
	var seatIdx = -1
	for i, player := range m.mahjongPlayers {
		if player == nil {
			seatIdx = i
			m.mahjongPlayers[i] = &mahjongPlayer{
				Player:    p,
				lightGang: map[int32]int64{},
				darkGang:  map[int32]int64{},
				pong:      map[int32]int64{},
			}
			break
		}
	}

	if seatIdx >= 0 {
		m.room.AddTimer(p.RID, time.Now().Add(ReadyTimeout), func(dt time.Duration) {
			m.room.PlayerLeave(p.ShortId)
		})
	}
}

func (m *Mahjong) PlayerLeave(p *room.Player) {
	for idx, player := range m.mahjongPlayers {
		if player != nil && player.ShortId == player.ShortId {
			m.room.CancelTimer(p.RID)
			m.mahjongPlayers[idx] = nil
			return
		}
	}
}

func (m *Mahjong) PlayerReady(p *room.Player) {
	if p.Ready {
		m.room.CancelTimer(p.RID)
	} else {
		m.room.AddTimer(p.RID, time.Now().Add(ReadyTimeout), func(dt time.Duration) {
			m.room.PlayerLeave(p.ShortId)
		})
	}
}

// Handle 麻将游戏消息，全部交由当前状态处理
func (m *Mahjong) Handle(v any, shortId int64) any {
	return m.fsm.CurrentStateHandler().Handle(m.fsm, v, shortId)
}

func (m *Mahjong) findMahjongPlayer(shortId int64) *mahjongPlayer {
	for _, player := range m.mahjongPlayers {
		if player != nil && player.ShortId == shortId {
			return player
		}
	}
	return nil
}
