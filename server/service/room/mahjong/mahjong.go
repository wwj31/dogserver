package mahjong

import (
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/service/room"
	"time"
)

const ReadyTimeout = 20 * time.Second

func New(r *room.Room) *mahjong {
	fsm := room.NewFSM()
	_ = fsm.Add(&StateDeal{})

	return &mahjong{
		room: r,
		FSM:  fsm,
	}
}

type mahjongPlayer struct {
	*room.Player

	tickOutTimerId string // 准备倒计时
}
type mahjong struct {
	room *room.Room
	*room.FSM
	mahjongPlayers []*mahjongPlayer
}

func (m *mahjong) Init() {
	if err := m.Switch(Ready); err != nil {
		log.Warnw("mahjong start switch ready failed", m.room.LogInfo()...)
	}
}

func (m *mahjong) CanEnter(p *inner.PlayerInfo) bool {
	if m.State() == Ready {
		return true
	}
	return false
}
func (m *mahjong) CanLeave(p *inner.PlayerInfo) bool {
	// 只有准备和结算时可以离开
	switch m.State() {
	case Ready, Settlement:
		return true
	}
	return false
}
func (m *mahjong) CanReady(p *inner.PlayerInfo) bool {
	if m.State() == Ready {
		return true
	}
	return false
}

func (m *mahjong) PlayerEnter(p *room.Player) {
	m.room.AddTimer(p.RID, time.Now().Add(ReadyTimeout), func(dt time.Duration) {
		m.room.PlayerLeave(p.ShortId)
	})
}
func (m *mahjong) PlayerLeave(p *room.Player) {
}
func (m *mahjong) PlayerReady(p *room.Player) {
}

func (m *mahjong) Handle(v any) any {
	return v
}

func (m *mahjong) findMahjongPlayer(shortId int64) *mahjongPlayer {
	for _, p := range m.mahjongPlayers {
		if p.ShortId == shortId {
			return p
		}
	}
	return nil
}
