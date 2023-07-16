package mahjong

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"time"

	"server/proto/outermsg/outer"

	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/proto/innermsg/inner"
	"server/service/room"
)

const (
	ReadyExpiration          = 20 * time.Second // 准备超时时间
	DecideMasterShowDuration = 3 * time.Second  // 定庄广播后的动画播放时间
	DealShowDuration         = 2 * time.Second  // 发牌广播后的动画播放时间
	Exchange3Expiration      = 5 * time.Second  // 换三张持续时间
	Exchange3ShowDuration    = 1 * time.Second  // 换三张结束后的动画播放时间
	DecideIgnoreExpiration   = 5 * time.Second  // 定缺持续时间
	DecideIgnoreDuration     = 1 * time.Second  // 定缺结束后的动画播放时间
	pongGangHuGuoExpiration  = 6 * time.Second  // 碰、杠、胡、过持续时间
	playCardExpiration       = 3 * time.Second  // 摸牌后的行为持续时间(出牌，杠，胡)
	SettlementDuration       = 2 * time.Second  // 结算持续时间
)

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
	_ = mahjong.fsm.Add(&StateSettlement{Mahjong: mahjong})   // 游戏结束结算界面

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
		hu          HuType               // 胡牌
		huExtra     ExtFanType           // 胡牌额外加番
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

		masterIndex    int // 庄家位置 0,1,2,3
		gameCount      int // 游戏的连续局数
		firstHuIndex   int // 第一个胡牌的人
		mutilHuByIndex int // 一炮多响点炮的人

		cards              Cards                  // 剩余牌组
		cardsInDesktop     Cards                  // 打出的牌
		mahjongPlayers     [maxNum]*mahjongPlayer // 参与游戏的玩家
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

func (m *Mahjong) Data() proto.Message {
	return &outer.MahjongBTEGameInfo{}
}

func (m *Mahjong) SeatIndex(shortId int64) int {
	for seatIndex, player := range m.mahjongPlayers {
		if player != nil && player.ShortId == shortId {
			return seatIndex
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
			m.mahjongPlayers[i] = m.newMahjongPlayer(p)
			break
		}
	}

	if seatIdx >= 0 {
		m.room.AddTimer(p.RID, time.Now().Add(ReadyExpiration), func(dt time.Duration) {
			log.Infow("the player was kicked out of the room due to a timeout in the ready period",
				"roomId", m.room.RoomId, "player", p.ShortId)
			m.room.PlayerLeave(p.ShortId, true)
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
		m.room.AddTimer(p.RID, time.Now().Add(ReadyExpiration), func(dt time.Duration) {
			m.room.PlayerLeave(p.ShortId, true)
		})
	}
}

// Handle 麻将游戏消息，全部交由当前状态处理
func (m *Mahjong) Handle(shortId int64, v any) any {
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

// 逆时针轮动座位索引,index 当前位置
func (m *Mahjong) nextSeatIndex(index int) int {
	// 0,1,2,3 东南西北
	for {
		index--
		if index < 0 {
			index = 3
		}

		player := m.mahjongPlayers[index]
		if player.hu == HuInvalid {
			break
		}
	}
	return index
}

func (m *Mahjong) nextSeatIndexWithoutHu(index int) int {
	// 0,1,2,3 东南西北
	index--
	if index < 0 {
		index = 3
	}
	return index
}

func (m *Mahjong) newMahjongPlayer(p *room.Player) *mahjongPlayer {
	return &mahjongPlayer{
		Player:    p,
		lightGang: map[int32]int64{},
		darkGang:  map[int32]int64{},
		pong:      map[int32]int64{},
	}
}

func (m *Mahjong) clear() {
	// 重置玩家数据
	for i, p := range m.mahjongPlayers {
		m.mahjongPlayers[i] = m.newMahjongPlayer(p.Player)
		m.mahjongPlayers[i].Ready = false
	}

	m.cards = nil
	m.cardsInDesktop = nil
	m.actionMap = make(map[int]*action)
	m.currentActionEndAt = time.Time{}
}

func (m *mahjongPlayer) allCardsToPB() *outer.CardsOfBTE {
	return &outer.CardsOfBTE{
		Cards:     m.handCards.ToSlice(),
		LightGang: m.lightGang,
		DarkGang:  m.darkGang,
		Pong:      m.pong,
	}
}

// 检测该行为是否在当前可操作行为中
func (a *action) isValidAction(actionType outer.ActionType) bool {
	for _, act := range a.currentActions {
		if act == actionType {
			return true
		}
	}
	return false
}

// 判断当前action是否有效
func (a *action) isActivated() bool {
	return len(a.currentActions) > 0
}

// 删除一个行为
func (a *action) remove(actionType outer.ActionType) {
	for i, act := range a.currentActions {
		if act == actionType {
			a.currentActions = append(a.currentActions[:i], a.currentActions[i+1:]...)
			return
		}
	}
}

func (a *action) String() string {
	return fmt.Sprintf("actions:%v,hus:%v,gang:%v", a.currentActions, a.currentHus, a.currentGang)
}
