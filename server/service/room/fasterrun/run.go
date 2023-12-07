package fasterrun

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/rdsop"

	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/logger"

	"server/proto/outermsg/outer"

	"github.com/wwj31/dogactor/tools"

	"server/proto/innermsg/inner"
	"server/service/room"
)

const (
	ReadyExpiration       = 20 * time.Second // 准备超时时间
	DealExpiration        = 3 * time.Second  // 发牌状态持续时间
	WaitingPlayExpiration = 15 * time.Second // 打牌等待持续时间
	WaitingPassExpiration = 3 * time.Second  // 不能大，等待短暂的几秒自动过
	SettlementDuration    = 10 * time.Second // 结算持续时间
)

func New(r *room.Room) *FasterRun {
	fasterRun := &FasterRun{
		room: r,
		fsm:  room.NewFSM(),
	}
	n := fasterRun.playerNumber()
	fasterRun.fasterRunPlayers = make([]*fasterRunPlayer, n, n)
	_ = fasterRun.fsm.Add(&StateReady{FasterRun: fasterRun})      // 准备中
	_ = fasterRun.fsm.Add(&StateDeal{FasterRun: fasterRun})       // 发牌中
	_ = fasterRun.fsm.Add(&StatePlaying{FasterRun: fasterRun})    // 游戏中
	_ = fasterRun.fsm.Add(&StateSettlement{FasterRun: fasterRun}) // 游戏结束结算界面

	fasterRun.SwitchTo(Ready)

	return fasterRun
}

type (
	// 跑得快 参与游戏的玩家数据
	fasterRunPlayer struct {
		*room.Player
		score          int64
		totalWinScore  int64 // 单局的总输赢
		ready          bool
		readyExpireAt  time.Time
		handCards      PokerCards
		BombsCount     int32
		doubleHearts10 bool // 红桃10 翻倍
		finalStatsMsg  *outer.FasterRunFinialPlayerInfo
	}

	FasterRun struct {
		room                *room.Room
		fsm                 *room.FSM
		currentStateEnterAt time.Time                            // 当前状态的进入时间
		currentStateEndAt   time.Time                            // 当前状态的结束时间
		playerAutoReady     func(p *fasterRunPlayer, ready bool) //
		scoreZeroOver       bool                                 // 因为有玩家没分了，而触发的结束

		masterIndex        int                // 庄家位置 0，1，2
		fasterRunPlayers   []*fasterRunPlayer // 参与游戏的玩家
		playRecords        []playCardsRecord  // 出牌历史
		gameCount          int                // 游戏的连续局数
		lastWinShortId     int64              // 最后一局的赢家
		waitingPlayShortId int64              // 当前等待的出牌人
		waitingPlayFollow  bool               // 当前等待的出牌人是否是跟牌
		spareCards         PokerCards         // 剩下没用的牌
	}

	playCardsRecord struct {
		shortId    int64      // 出牌人
		follow     bool       // true.跟牌出牌，false.有牌权出牌
		cardsGroup CardsGroup // 牌型
		playAt     time.Time  // 出牌时间
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

func (f *FasterRun) toRoomPlayers() (players []*room.Player) {
	for _, p := range f.fasterRunPlayers {
		players = append(players, p.Player)
	}
	return players
}

func (f *FasterRun) playerNumber() int {
	n := f.gameParams().PlayerNumber
	if n == 0 {
		return 2
	} else if n == 1 {
		return 3
	}
	log.Errorw("the player number is unexpected", "number", n)
	return 0
}

// 获得最后一次有效的出牌
func (f *FasterRun) lastValidPlayCards() *playCardsRecord {
	for i := len(f.playRecords) - 1; i >= 0; i-- {
		record := f.playRecords[i]
		if record.cardsGroup.Type == CardsTypeUnknown {
			continue
		}
		return &record
	}
	return nil
}

func (f *FasterRun) Data(shortId int64) proto.Message {
	var records []*outer.PlayCardsRecord
	for _, record := range f.playRecords {
		records = append(records, record.ToPB())
	}

	info := &outer.FasterRunGameInfo{
		State:        outer.FasterRunState(f.fsm.State()),
		StateEnterAt: f.currentStateEnterAt.UnixMilli(),
		StateEndAt:   f.currentStateEndAt.UnixMilli(),
		GameCount:    int32(f.gameCount),
		Players:      f.playersToPB(shortId, false),
		MasterIndex:  int32(f.masterIndex),
		History:      records,
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

func (f *FasterRun) CanSetGold(p *inner.PlayerInfo) bool {
	if f.fsm.State() == Ready {
		return true
	}
	return false
}

// RecordingPlayback 当前状态是否需要记录回播内容
func (f *FasterRun) CanRecordingPlayback() bool {
	if f.fsm.State() == Ready {
		return false
	}
	return true
}

func (f *FasterRun) PlayerOnline(shortId int64)  {}
func (f *FasterRun) PlayerOffline(shortId int64) {}

func (f *FasterRun) PlayerEnter(roomPlayer *room.Player) {
	for i, player := range f.fasterRunPlayers {
		if player == nil {
			player = f.newFasterRunPlayer(roomPlayer)
			player.finalStatsMsg = &outer.FasterRunFinialPlayerInfo{}
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
			f.Log().Infow("player leave faster run", "shortId", player.ShortId, "seat", idx, "gold", player.Gold)
			return
		}
	}
}

// Handle 游戏消息，全部交由当前状态处理
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

func (f *FasterRun) playersToPB(shortId int64, settlement bool) (players []*outer.FasterRunPlayerInfo) {
	for _, player := range f.fasterRunPlayers {
		if player == nil {
			players = append(players, nil)
		} else {
			var handCards []int32
			if player.ShortId == shortId || settlement {
				handCards = player.handCards.ToPB()
			} else {
				handCards = make([]int32, len(player.handCards))
			}

			players = append(players, &outer.FasterRunPlayerInfo{
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

// 开局扣门票
func (f *FasterRun) masterRebate() {
	if f.gameParams().MasterRebate <= 0 {
		return
	}

	rsp, err := f.room.RequestWait(actortype.AllianceName(f.room.AllianceId), &inner.AllianceInfoReq{})
	if err != nil {
		f.Log().Errorw("master rebate req failed", "err", err)
		return
	}
	infoRsp := rsp.(*inner.AllianceInfoRsp)
	if infoRsp.MasterShortId == 0 {
		f.Log().Errorw("masterRebate master shortId == 0")
		return
	}

	record := &outer.RebateDetailInfo{
		Type:      outer.GameType_FasterRun,
		BaseScore: f.gameParams().BaseScore,
		CreateAt:  tools.Now().UnixMilli(),
	}

	var pip redis.Pipeliner
	score := int64(f.gameParams().MasterRebate * common.Gold1000Times) // 门票
	for _, player := range f.fasterRunPlayers {
		player.score = common.Max(player.score-score, 0)
		record.Gold = score
		record.ShortId = player.ShortId
		rdsop.RecordRebateGold(common.JsonMarshal(record), infoRsp.MasterShortId, score, pip)
	}

	if _, err = pip.Exec(context.Background()); err != nil {
		f.Log().Errorw("master rebate redis failed", "err", err)
	}
}

func (f *FasterRun) gameParams() *outer.FasterRunParams {
	return f.room.GameParams.FasterRun
}

func (f *FasterRun) baseScore() int64 {
	base := f.gameParams().BaseScore
	if base <= 0 {
		base = 1
	}

	return int64(base * common.Gold1000Times)
}

func (f *FasterRun) bombWinScore() int64 {
	base := f.gameParams().BaseScore
	if base == 0 {
		base = 1
	}

	return int64(float32(base*common.Gold1000Times) * 5)
}

func (f *FasterRun) clear() {
	f.playRecords = nil
	f.scoreZeroOver = false
	f.waitingPlayShortId = 0
	f.waitingPlayFollow = false

	// 重置玩家数据
	for i := 0; i < f.playerNumber(); i++ {
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

	for seatIndex := 0; seatIndex < f.playerNumber(); seatIndex++ {
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

func (p *playCardsRecord) String() string {
	if p == nil {
		return "{nil}"
	}

	return fmt.Sprintf("{short:%v follow:%v cardsGroup:%v playAt:%v }", p.shortId, p.follow, p.cardsGroup, p.playAt)
}

func (p *playCardsRecord) ToPB() *outer.PlayCardsRecord {
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
