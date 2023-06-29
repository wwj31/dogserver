package mahjong

import (
	"time"

	"server/common/log"
	"server/proto/innermsg/inner"
	"server/service/room"
)

const ReadyTimeout = 20 * time.Second

func New(r *room.Room) *Mahjong {
	fsm := room.NewFSM()
	_ = fsm.Add(&StateDeal{})

	return &Mahjong{
		room: r,
		FSM:  fsm,
	}
}

type mahjongPlayer struct {
	*room.Player
	masterIndex int

	tickOutTimerId string // 准备倒计时
}
type Mahjong struct {
	room *room.Room
	*room.FSM
	mahjongPlayers []*mahjongPlayer
}

func (m *Mahjong) Init() {
	if err := m.Switch(Ready); err != nil {
		log.Warnw("Mahjong start switch ready failed", m.room.LogInfo()...)
	}
}

func (m *Mahjong) CanEnter(p *inner.PlayerInfo) bool {
	if m.State() == Ready {
		return true
	}
	return false
}
func (m *Mahjong) CanLeave(p *inner.PlayerInfo) bool {
	// 只有准备和结算时可以离开
	switch m.State() {
	case Ready, Settlement:
		return true
	}
	return false
}
func (m *Mahjong) CanReady(p *inner.PlayerInfo) bool {
	if m.State() == Ready {
		return true
	}
	return false
}

func (m *Mahjong) PlayerEnter(p *room.Player) {
	m.room.AddTimer(p.RID, time.Now().Add(ReadyTimeout), func(dt time.Duration) {
		m.room.PlayerLeave(p.ShortId)
	})
}
func (m *Mahjong) PlayerLeave(p *room.Player) {
}
func (m *Mahjong) PlayerReady(p *room.Player) {
}

func (m *Mahjong) Handle(v any) any {
	return v
}

func (m *Mahjong) findMahjongPlayer(shortId int64) *mahjongPlayer {
	for _, p := range m.mahjongPlayers {
		if p.ShortId == shortId {
			return p
		}
	}
	return nil
}
