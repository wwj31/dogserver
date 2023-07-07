package mahjong

import (
	"time"

	"server/proto/outermsg/outer"

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
	_ = mahjong.fsm.Add(&StateExchange3{Mahjong: mahjong})    // 换三张
	_ = mahjong.fsm.Add(&StateDecideIgnore{Mahjong: mahjong}) // 定缺
	_ = mahjong.fsm.Add(&StatePlaying{Mahjong: mahjong})      // 游戏中

	mahjong.SwitchTo(Ready)

	return mahjong
}

const maxNum = 4

type (
	// 麻将-血战到底 参与游戏的玩家数据
	mahjongPlayer struct {
		*room.Player

		ignoreColor ColorType            // 定缺花色
		exchange    *outer.Exchange3Info // 换三张信息
		handCards   Cards                // 手牌
		lightGang   map[int32]int64      // map[杠牌]ShortId 明杠
		darkGang    map[int32]int64      // map[杠牌]ShortId 暗杠
		pong        map[int32]int64      // map[碰牌]ShortId
	}

	action struct {
		currentActions []outer.ActionType // 当前行动者能执行的行为
		currentHus     []outer.HuType     // 当前行动者能胡的牌
		currentGang    []int32            // 当前行动者能杠的牌
	}

	Mahjong struct {
		room                *room.Room
		fsm                 *room.FSM
		currentStateEnterAt time.Time // 当前状态的进入时间

		masterIndex int // 庄家位置 0,1,2,3
		gameCount   int // 游戏的连续局数 结算后，有玩家退出，重置0

		cards              Cards                  // 剩余牌组
		mahjongPlayers     [maxNum]*mahjongPlayer // 参与游戏的玩家
		latestDrawIndex    int                    // 最后一个摸牌的位置
		latestPlayIndex    int                    // 最后一个打牌的位置
		actionMap          map[int]*action        // 行动者们
		currentActionEndAt time.Time              // 当前行动者结束时间
	}
)

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
	return m.fsm.CurrentStateHandler().Handle(shortId, v)
}

func (m *Mahjong) findMahjongPlayer(shortId int64) (*mahjongPlayer, int) {
	for i, player := range m.mahjongPlayers {
		if player != nil && player.ShortId == shortId {
			return player, i
		}
	}
	return nil, -1
}

func (m *Mahjong) nextSeatIndex(index int32) int32 {
	// 0,1,2,3 东南西北
	index--
	if index < 0 {
		index = 3
	}
	return index
}

// 重置玩家上一把游戏数据
func (m *Mahjong) clearMahjongPlayerInfo() {
	for _, player := range m.mahjongPlayers {
		player.ignoreColor = ColorUnknown
		player.exchange = nil
		player.handCards = nil
		player.lightGang = map[int32]int64{}
		player.darkGang = map[int32]int64{}
		player.pong = map[int32]int64{}
	}
}

func (m *mahjongPlayer) allCardsToPB() *outer.CardsOfBTE {
	return &outer.CardsOfBTE{
		Cards:     m.handCards.ToSlice(),
		LightGang: m.lightGang,
		DarkGang:  m.darkGang,
		Pong:      m.pong,
	}
}

// 检查某个行为是否有效
func (m *action) isValidAction(actionType outer.ActionType) bool {
	for _, act := range m.currentActions {
		if act == actionType {
			return true
		}
	}
	return false
}
