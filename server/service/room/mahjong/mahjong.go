package mahjong

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/logger"

	"server/proto/outermsg/outer"

	"github.com/wwj31/dogactor/tools"

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
	pongGangHuGuoExpiration  = 2 * time.Second  // 碰、杠、胡、过持续时间
	playCardExpiration       = 2 * time.Second  // 出牌行为持续时间
	SettlementDuration       = 1 * time.Second  // 结算持续时间
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
		score         int64
		ready         bool
		readyExpireAt time.Time

		ignoreColor ColorType            // 定缺花色
		exchange    *outer.Exchange3Info // 换三张信息

		handCards Cards           // 手牌
		lightGang map[int32]int64 // map[杠牌]ShortId 明杠
		darkGang  map[int32]int64 // map[杠牌]ShortId 暗杠
		gangScore map[int][]int32 // map[peerIndex]seats，杠需要赔付的位置
		pong      map[int32]int64 // map[碰牌]ShortId

		hu          HuType     // 胡牌
		huCard      Card       // 胡的那张牌
		huExtra     ExtFanType // 胡牌额外加番
		huGen       int32      // 胡牌有几根
		huPeerIndex int        // 胡的那次peer下标

	}

	action struct {
		acts    []outer.ActionType // 当前行动者能执行的行为
		hus     []outer.HuType     // 当前行动者能胡的牌
		gang    []int32            // 当前行动者能杠的牌
		newCard Card               // 当前行动者摸到的新牌
	}

	Mahjong struct {
		room                *room.Room
		fsm                 *room.FSM
		currentStateEnterAt time.Time // 当前状态的进入时间
		currentStateEndAt   time.Time // 当前状态的结束时间

		dices          [2]int32 // 两颗骰子数
		masterIndex    int      // 庄家位置 0,1,2,3
		gameCount      int      // 游戏的连续局数
		huSeat         []int32  // 胡牌的位置，依次按顺序加入
		mutilHuByIndex int      // 一炮多响点炮的人

		cards          Cards                  // 剩余牌组
		cardsInDesktop Cards                  // 打出的牌
		cardsPlayOrder []int32                // 出牌座位顺序
		mahjongPlayers [maxNum]*mahjongPlayer // 参与游戏的玩家
		peerRecords    []peerRecords          // 需要关注的操作记录(出牌、摸牌、杠)按触发顺序添加
		actionMap      map[int]*action        // 行动者们

		currentAction      *action   // 当前行动者
		currentActionSeat  int       // 当前行动者座位
		currentActionEndAt time.Time // 当前行动者结束时间
	}
)

func (m *Mahjong) SwitchTo(state int) {
	if err := m.fsm.SwitchTo(state); err != nil {
		current := m.fsm.CurrentStateHandler().State()
		m.Log().Errorw("Mahjong switch to next state failed", "room", m.room.RoomId, "current", current)
		return
	}
	m.currentStateEnterAt = tools.Now()
}

func (m *Mahjong) Data(shortId int64) proto.Message {
	info := &outer.MahjongBTEGameInfo{
		State:            outer.MahjongBTEState(m.fsm.State()),
		StateEnterAt:     m.currentStateEnterAt.UnixMilli(),
		StateEndAt:       m.currentStateEndAt.UnixMilli(),
		Players:          m.playersToPB(shortId, false),
		Dices:            m.dices[:],
		MasterIndex:      int32(m.masterIndex),
		Ex3Info:          m.ex3Info(shortId),
		TotalCardsCount:  int32(m.cards.Len()),
		CardsPlayHistory: m.cardsInDesktop.ToSlice(),
		CardsPlayOrder:   m.cardsPlayOrder,
		ActionEndAt:      m.currentActionEndAt.UnixMilli(),
		ActionShortId:    0,
		ActionType:       nil,
		HuType:           nil,
		GangCards:        nil,
		NewCard:          0,
	}

	// 只有当行动者是出牌状态，才广播行动者
	if m.currentAction != nil && m.currentAction.isValidAction(outer.ActionType_ActionPlayCard) {
		info.ActionShortId = m.mahjongPlayers[m.currentActionSeat].ShortId
	}

	// 判断是当前行动者本人，就发行动数据
	if m.currentActionSeat > 0 {
		p := m.mahjongPlayers[m.currentActionSeat]
		if p.ShortId == shortId {
			info.ActionType = m.currentAction.acts
			info.HuType = m.currentAction.hus
			info.GangCards = m.currentAction.gang
			info.NewCard = m.currentAction.newCard.Int32()
		}
	}

	return info
}

func (m *Mahjong) ex3Info(shortId int64) (info *outer.Exchange3Info) {
	p, _ := m.findMahjongPlayer(shortId)
	if p == nil {
		return nil
	}
	return p.exchange
}

func (m *Mahjong) playersToPB(shortId int64, settlement bool) (players []*outer.MahjongPlayerInfo) {
	for _, player := range m.mahjongPlayers {
		if player == nil {
			players = append(players, nil)
		} else {
			var allCards []int32
			if player.ShortId == shortId || settlement {
				allCards = player.handCards.ToSlice()
			} else {
				handLen := player.handCards.Len()
				allCards = make([]int32, handLen, handLen)
			}

			var huPeer peerRecords
			if player.huPeerIndex != -1 {
				huPeer = m.peerRecords[player.huPeerIndex]
			}

			players = append(players, &outer.MahjongPlayerInfo{
				ShortId:        player.ShortId,
				Ready:          player.ready,
				ReadyExpireAt:  player.readyExpireAt.UnixMilli(),
				Exchange3Ready: player.exchange != nil,
				DecideColor:    outer.ColorType(player.ignoreColor),
				AllCards: &outer.CardsOfBTE{
					Cards:     allCards,
					LightGang: player.lightGang,
					DarkGang:  player.darkGang,
					Pong:      player.pong,
				},
				HuType:      player.hu.PB(),
				HuExtraType: player.huExtra.PB(),
				HuCard:      huPeer.card.Int32(),
				HuGen:       player.huGen,
				Score:       player.score,
			})
		}
	}
	return
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
	for i, player := range m.mahjongPlayers {
		if player == nil {
			m.mahjongPlayers[i] = m.newMahjongPlayer(p)
			m.mahjongPlayers[i].readyExpireAt = time.Now().Add(ReadyExpiration)
			m.readyTimeout(p.RID, p.ShortId, m.mahjongPlayers[i].readyExpireAt)
			break
		}
	}
}

func (m *Mahjong) readyTimeout(rid string, shortId int64, expireAt time.Time) {
	m.room.AddTimer(rid, expireAt, func(dt time.Duration) {
		m.Log().Infow("the player was kicked out of the room due to a timeout in the ready period",
			"room", m.room.RoomId, "player", shortId)
		m.room.PlayerLeave(shortId, true)
	})
}

func (m *Mahjong) PlayerLeave(quitPlayer *room.Player) {
	for idx, player := range m.mahjongPlayers {
		if player != nil && player.ShortId == quitPlayer.ShortId {
			m.room.CancelTimer(quitPlayer.RID)
			m.mahjongPlayers[idx] = nil
			m.Log().Infow("player leave mahjong", "shortId", player.ShortId, "seat", idx)
			return
		}
	}
}

// Handle 麻将游戏消息，全部交由当前状态处理
func (m *Mahjong) Handle(shortId int64, v any) any {
	return m.fsm.CurrentStateHandler().Handle(shortId, v)
}

func (m *Mahjong) Log() *logger.Logger {
	return m.room.Log()
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
		Player:      p,
		score:       p.Gold,
		lightGang:   map[int32]int64{},
		darkGang:    map[int32]int64{},
		pong:        map[int32]int64{},
		gangScore:   map[int][]int32{},
		huPeerIndex: -1,
	}
}

func (m *Mahjong) gameParams() *outer.MahjongParams {
	return m.room.GameParams.Mahjong
}

func (m *Mahjong) clear() {
	// 重置玩家数据
	for i, p := range m.mahjongPlayers {
		m.mahjongPlayers[i] = m.newMahjongPlayer(p.Player)
	}

	m.cards = nil
	m.cardsInDesktop = make(Cards, 0, 8)
	m.cardsPlayOrder = make([]int32, 0, 8)
	m.currentAction = nil
	m.currentActionSeat = -1
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
	for _, act := range a.acts {
		if act == actionType {
			return true
		}
	}
	return false
}

// 判断当前action是否有效
func (a *action) isActivated() bool {
	return len(a.acts) > 0
}

// 删除一个行为
func (a *action) remove(actionType outer.ActionType) {
	for i, act := range a.acts {
		if act == actionType {
			a.acts = append(a.acts[:i], a.acts[i+1:]...)
			return
		}
	}
}

// 当前行为能否杠某张牌
func (a *action) canGang(card Card) bool {
	var valid bool
	for _, c := range a.gang {
		if c == card.Int32() {
			valid = true
			break
		}
	}
	return valid
}

func (a *action) String() string {
	return fmt.Sprintf("actions:%v,hus:%v,gang:%v", a.acts, a.hus, a.gang)
}
