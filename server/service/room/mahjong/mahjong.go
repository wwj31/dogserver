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
	Exchange3Expiration      = 20 * time.Second // 换三张持续时间
	Exchange3ShowDuration    = 1 * time.Second  // 换三张结束后的动画播放时间
	DecideIgnoreExpiration   = 20 * time.Second // 定缺持续时间
	DecideIgnoreDuration     = 1 * time.Second  // 定缺结束后的动画播放时间
	pongGangHuGuoExpiration  = 20 * time.Second // 碰、杠、胡、过持续时间
	playCardExpiration       = 20 * time.Second // 出牌行为持续时间
	SettlementDuration       = 10 * time.Second // 结算持续时间
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

type gangInfo struct {
	loserSeats    map[int32]int64 // 赔付的位置，赔付的分
	totalWinScore int64           // 本次杠总分
}
type (
	// 麻将-血战到底 参与游戏的玩家数据
	mahjongPlayer struct {
		*room.Player
		score         int64
		totalWinScore int64 // 单局的总输赢
		ready         bool
		readyExpireAt time.Time
		finalStatsMsg *outer.MahjongBTEFinialPlayerInfo

		ignoreColor ColorType            // 定缺花色
		exchange    *outer.Exchange3Info // 换三张信息

		handCards Cards // 手牌

		// 碰杠数据
		pong           map[int32]int64   // map[碰牌]ShortId
		lightGang      map[int32]int64   // map[杠牌]ShortId 明杠
		darkGang       map[int32]int64   // map[杠牌]ShortId 暗杠
		gangInfos      map[int]*gangInfo // map[peerIndex]杠信息
		gangTotalScore int64             // 杠总共赢的分数

		// 胡牌数据
		hu            HuType          // 胡牌
		huCard        Card            // 胡的那张牌
		huExtra       []ExtFanType    // 胡牌额外加番
		huGen         int32           // 胡牌有几根
		huPeerIndex   int             // 胡的那次peer下标
		winScore      map[int32]int64 // 胡牌赢的分 map[赔分的位置]赔的分
		huTotalScore  int64           // 胡牌赢的总分
		passHandHuFan int             // 过手胡限制，当前限制番数
	}

	action struct {
		seat    int
		acts    []outer.ActionType // 当前行动者能执行的行为
		hus     []outer.HuType     // 当前行动者能胡的牌
		gang    []int32            // 当前行动者能杠的牌
		newCard Card               // 当前行动者摸到的新牌
	}

	Mahjong struct {
		room                *room.Room
		fsm                 *room.FSM
		currentStateEnterAt time.Time                          // 当前状态的进入时间
		currentStateEndAt   time.Time                          // 当前状态的结束时间
		playerAutoReady     func(p *mahjongPlayer, ready bool) //

		dices          [2]int32 // 两颗骰子数
		masterIndex    int      // 庄家位置 0,1,2,3
		masterCard14   Card     // 庄家的第14张牌
		gameCount      int      // 游戏的连续局数
		huSeat         []int32  // 胡牌的位置，依次按顺序加入
		multiHuByIndex int      // 一炮多响点炮的人
		scoreZeroOver  bool     // 因为有玩家没分了，而触发的结束

		cards              Cards                  // 剩余牌组
		cardsInDesktop     Cards                  // 打出的牌
		cardsPlayOrder     []int32                // 出牌座位顺序
		mahjongPlayers     [maxNum]*mahjongPlayer // 参与游戏的玩家
		peerRecords        []peerRecords          // 需要关注的操作记录(出牌、摸牌、杠)按触发顺序添加
		actionMap          map[int]*action        // 代表某次操作触发了玩家的行为
		currentAction      []*action              // 当前行动者,从actionMap中筛选出来的，当前能行动的人
		currentActionEndAt time.Time              // 当前行动者结束时间
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
		GameCount:        int32(m.gameCount),
		Players:          m.playersToPB(shortId, false),
		Dices:            m.dices[:],
		MasterIndex:      int32(m.masterIndex),
		Ex3Info:          m.ex3Info(shortId),
		TotalCardsCount:  int32(m.cards.Len()),
		CardsPlayHistory: m.cardsInDesktop.ToSlice(),
		CardsPlayOrder:   m.cardsPlayOrder,
		ActionEndAt:      m.currentActionEndAt.UnixMilli(),
		HuSeats:          m.huSeat,
		HuFan:            HuTypePB(),
		HuExtraFan:       ExtFanTypePB(),
		ActionShortId:    0,
		ActionType:       nil,
		HuType:           nil,
		GangCards:        nil,
		NewCard:          0,
	}

	// 检查当前的行动者是不是出牌状态，如果是，需要广播行动者是谁
	for _, a := range m.currentAction {
		if a.isValidAction(outer.ActionType_ActionPlayCard) {
			info.ActionShortId = m.mahjongPlayers[a.seat].ShortId
			break
		}
	}

	// 检查玩家是否是行动者，如果不是，以下发送内容都不发
	var currentAction *action
	for _, a := range m.currentAction {
		if m.mahjongPlayers[a.seat].ShortId == shortId {
			currentAction = a
			break
		}
	}

	if currentAction == nil {
		return info
	}

	// 发行动数据
	p := m.mahjongPlayers[currentAction.seat]
	if p.ShortId == shortId {
		info.ActionType = currentAction.acts
		info.HuType = currentAction.hus
		info.GangCards = currentAction.gang
		info.NewCard = currentAction.newCard.Int32()
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

// 封顶番数
func (m *Mahjong) fanUpLimit() int32 {
	if m.gameParams().FanUpLimit < 1 || m.gameParams().FanUpLimit > 5 {
		return 6
	}
	return m.gameParams().FanUpLimit
}

func (m *Mahjong) playersToPB(shortId int64, settlement bool) (players []*outer.MahjongPlayerInfo) {
	for _, player := range m.mahjongPlayers {
		if player == nil {
			players = append(players, nil)
		} else {

			var huPeer peerRecords
			if player.huPeerIndex != -1 {
				huPeer = m.peerRecords[player.huPeerIndex]
			}

			allCards := player.allCardsToPB(m.gameParams(), shortId, settlement)
			players = append(players, &outer.MahjongPlayerInfo{
				ShortId:        player.ShortId,
				Ready:          player.ready,
				ReadyExpireAt:  player.readyExpireAt.UnixMilli(),
				Exchange3Ready: player.exchange != nil,
				DecideColor:    outer.ColorType(player.ignoreColor),
				AllCards:       allCards,
				HuType:         player.hu.PB(),
				HuExtraType:    ExtFanArrToPB(player.huExtra),
				HuCard:         huPeer.card.Int32(),
				HuGen:          player.huGen,
				Score:          m.immScore(shortId),
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
	if p.Gold <= 0 {
		return true
	}

	// 只有准备和结算时可以离开
	switch m.fsm.State() {
	case Settlement:
		if m.gameCount == int(m.gameParams().PlayCountLimit) {
			return true
		}

	case Ready:
		if m.gameCount == 1 {
			return true
		}
	}
	return false
}

func (m *Mahjong) CanReady(p *inner.PlayerInfo) bool {
	if m.fsm.State() == Ready {
		return true
	}
	return false
}

func (m *Mahjong) CanSetGold(p *inner.PlayerInfo) bool {
	if m.fsm.State() == Ready {
		return true
	}
	return false
}

func (m *Mahjong) PlayerEnter(roomPlayer *room.Player) {
	for i, player := range m.mahjongPlayers {
		if player == nil {
			player = m.newMahjongPlayer(roomPlayer)
			player.finalStatsMsg = &outer.MahjongBTEFinialPlayerInfo{}
			m.mahjongPlayers[i] = player
			if m.playerAutoReady != nil {
				m.playerAutoReady(player, false)
			}
			break
		}
	}
}

func (m *Mahjong) readyAfterTimeout(player *mahjongPlayer, expireAt time.Time) {
	m.room.AddTimer(player.RID, expireAt, func(dt time.Duration) {
		m.Log().Infow("the player was kicked out of the room due to a timeout in the ready period",
			"room", m.room.RoomId, "player", player.ShortId)
		m.room.PlayerLeave(player.ShortId, true)
	})
}

func (m *Mahjong) PlayerLeave(quitPlayer *room.Player) {
	for idx, player := range m.mahjongPlayers {
		if player != nil && player.ShortId == quitPlayer.ShortId {
			m.room.CancelTimer(quitPlayer.RID)
			m.mahjongPlayers[idx] = nil
			m.Log().Infow("player leave mahjong", "shortId", player.ShortId, "seat", idx, "gold", player.Gold)
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
		score:       p.Gold,
		Player:      p,
		lightGang:   map[int32]int64{},
		darkGang:    map[int32]int64{},
		pong:        map[int32]int64{},
		gangInfos:   map[int]*gangInfo{},
		huPeerIndex: -1,
	}
}

func (m *Mahjong) gameParams() *outer.MahjongParams {
	return m.room.GameParams.Mahjong
}

func (m *Mahjong) baseScore() int64 {
	base := m.gameParams().BaseScore
	if base == 0 {
		base = 1
	}

	baseScoreTimes := m.gameParams().BaseScoreTimes
	if baseScoreTimes == 0 {
		baseScoreTimes = 1.0
	}

	return int64(float32(base*1000) * baseScoreTimes)
}

func (m *Mahjong) clear() {
	// 重置玩家数据
	for i := 0; i < maxNum; i++ {
		gamer := m.mahjongPlayers[i]
		if gamer != nil {
			m.mahjongPlayers[i] = m.newMahjongPlayer(gamer.Player)
			m.mahjongPlayers[i].finalStatsMsg = gamer.finalStatsMsg
		}
	}

	m.cards = nil
	m.scoreZeroOver = false
	m.cardsInDesktop = make(Cards, 0, 8)
	m.cardsPlayOrder = make([]int32, 0, 8)
	m.currentAction = nil
	m.actionMap = make(map[int]*action)
	m.currentActionEndAt = time.Time{}
	m.peerRecords = nil
}

func (m *Mahjong) allSeats(ignoreSeat ...int) (result []int) {
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

func (m *Mahjong) immScore(shortId int64) int64 {
	roomPlayer := m.room.FindPlayer(shortId)
	if roomPlayer == nil {
		return 0
	}

	mahPlayer, _ := m.findMahjongPlayer(shortId)
	if mahPlayer == nil {
		return 0
	}

	totalScore := roomPlayer.Gold
	if m.gameParams().GangImmediatelyScore {
		totalScore += mahPlayer.gangTotalScore
	}

	if m.gameParams().HuImmediatelyScore {
		totalScore += mahPlayer.huTotalScore
	}
	return totalScore
}

func (m *Mahjong) peerRecordsLog() string {
	var peers []peerRecords
	if len(m.peerRecords) < 5 {
		peers = m.peerRecords
	} else {
		peers = m.peerRecords[len(m.peerRecords)-5:]
	}
	var log string
	log += "[ "
	for _, record := range peers {
		log += fmt.Sprintf("{typ:%v card:%v seat:%v qiangFn:%v},", record.typ, record.card.Int32(), record.seat, record.afterQiangPass != nil)
	}
	log += " ]"
	return log
}

func (m *mahjongPlayer) updateScore(val int64) {
	m.score += val
	m.totalWinScore += val // 单局总输赢
	m.finalStatsMsg.TotalScore += val
}

func (m *mahjongPlayer) allCardsToPB(params *outer.MahjongParams, shortId int64, settlement bool) *outer.CardsOfBTE {
	allCards := &outer.CardsOfBTE{}
	if m.ShortId == shortId || settlement || (params.HuImmediatelyScore && m.hu != HuInvalid) {
		allCards.Cards = m.handCards.ToSlice()
	} else {
		handLen := m.handCards.Len()
		allCards.Cards = make([]int32, handLen, handLen)
	}

	// 明杠暗杠，随时显示
	allCards.LightGang = m.lightGang
	for card, _ := range m.darkGang {
		allCards.DarkGang = append(allCards.DarkGang, card)
	}
	return allCards
}

// 检测该行为是否在当前可操作行为中
func (a *action) isValidAction(actionType outer.ActionType) bool {
	if a == nil {
		return false
	}

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
	return fmt.Sprintf("seat:%v actions:%v,hus:%v,gang:%v", a.seat, a.acts, a.hus, a.gang)
}
