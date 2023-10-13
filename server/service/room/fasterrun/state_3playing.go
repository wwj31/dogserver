package fasterrun

import (
	"reflect"
	"sort"
	"time"

	"server/common"

	"github.com/wwj31/dogactor/tools"

	"server/proto/outermsg/outer"
)

type StatePlaying struct {
	*FasterRun
	actionTimerId string
}

func (s *StatePlaying) State() int {
	return Playing
}

func (s *StatePlaying) Enter() {
	s.Log().Infow("[FasterRun] enter state playing", "room", s.room.RoomId, "params", *s.gameParams())
	s.actionTimerId = ""
	s.currentStateEnterAt = time.Time{}

	s.room.Broadcast(&outer.FasterRunGameStartNtf{})

	s.nextPlayer(s.masterIndex, nil)
}

func (s *StatePlaying) Leave() {
	s.cancelActionTimer()
	s.Log().Infow("[FasterRun] leave state playing", "room", s.room.RoomId)
}

func (s *StatePlaying) Handle(shortId int64, v any) (result any) {
	player, _, err := s.getPlayerAndSeatE(shortId)
	if err != outer.ERROR_OK {
		return err
	}

	s.Log().Infow("playing handle msg", "shortId", shortId, reflect.TypeOf(v).String(), v)

	switch msg := v.(type) {
	case *outer.FasterRunPlayCardReq:
		if err = s.play(player, int32ArrToPokerCards(msg.PlayCards)); err != outer.ERROR_OK {
			return err
		}
		return &outer.FasterRunPlayCardRsp{HandCards: player.handCards.ToPB()}

	case *outer.FasterRunPassReq: // 过
		if !s.waitingPlayFollow {
			return outer.ERROR_FASTERRUN_PLAY_IS_YOUR_TURN
		}

		if bigger := player.handCards.FindBigger(s.lastValidPlayCards().cardsGroup); len(bigger) > 0 {
			return outer.ERROR_FASTERRUN_PLAY_EXIST_BIGGER_CANNOT_PASS
		}
		s.pass(player)
		return &outer.FasterRunPassRsp{}

	default:
		s.Log().Warnw("playing status has received an unknown message", "msg", reflect.TypeOf(msg).String())
	}
	return outer.ERROR_FASTERRUN_STATE_MSG_INVALID
}

// 过牌
func (s *StatePlaying) pass(player *fasterRunPlayer) {
	s.playRecords = append(s.playRecords, PlayCardsRecord{
		shortId: player.ShortId,
		follow:  true,
		playAt:  tools.Now(),
	})

	seat := s.SeatIndex(player.ShortId)
	s.Log().Infow("pass", "seat", seat, "shortId", player.ShortId)
	s.nextPlayer(s.nextSeatIndex(seat), &s.playRecords[len(s.playRecords)-1])
}

// 打牌
func (s *StatePlaying) play(player *fasterRunPlayer, cards PokerCards) outer.ERROR {
	if len(cards) == 0 {
		return outer.ERROR_FASTERRUN_PLAY_CARDS_LEN_EMPTY
	}

	// 检查是否轮到该玩家出牌
	if s.waitingPlayShortId != player.ShortId {
		s.Log().Warnw("play check s.waitingPlayShortId != player.ShortId", "waitingShortId", s.waitingPlayShortId,
			"shortId", player.ShortId)
		return outer.ERROR_FASTERRUN_PLAY_CARDS_LEN_EMPTY
	}

	// 检查要打的牌是否全部在手牌中
	if !player.handCards.CheckExist(cards...) {
		s.Log().Warnw("play check exist failed", "short", player.ShortId,
			"hand cards", player.handCards, "play cards", cards)
		return outer.ERROR_FASTERRUN_PLAY_CARDS_MISS
	}

	if len(s.playRecords) == 0 && !s.checkFirstSpade3(cards) {
		s.Log().Warnw("play checkFirstSpade3 failed", "short", player.ShortId,
			"hand cards", player.handCards, "play cards", cards)
		return outer.ERROR_FASTERRUN_PLAY_FIRST_SPADE3_LIMIT
	}

	// 如果需要跟牌，检查牌型是否符合跟牌牌型
	lastValidPlayCards := s.lastValidPlayCards()
	follow := lastValidPlayCards != nil && lastValidPlayCards.shortId != player.ShortId

	// 分析牌型，检查牌型是否有效
	playCardsGroup := cards.AnalyzeCards(s.gameParams().AAAIsBombs)

	// 主要用于跟牌时，判断是否需要严格校验sideCards
	// FollowPlayTolerance特殊规则开启情况加，可以不用判断sideCards数量
	var tolerance bool

	// 是否允许三张或三带一
	if playCardsGroup.Type == Trips && !s.gameParams().SpecialThreeCards {
		playCardsGroup.Type = CardsTypeUnknown
	} else if playCardsGroup.Type == TripsWithOne && !s.gameParams().SpecialThreeCardsWithOne {
		playCardsGroup.Type = CardsTypeUnknown
	}

	if playCardsGroup.Type == CardsTypeUnknown {
		// 要保证一手全部打完，所以出牌既是手牌
		if len(cards) != len(player.handCards) {
			return outer.ERROR_FASTERRUN_PLAY_CARDS_INVALID
		}

		if follow {
			// 跟牌出牌，并且能一手出完的特殊情况
			if s.gameParams().FollowPlayTolerance {
				playCardsGroup, tolerance = s.followPlayTolerance(player.handCards, lastValidPlayCards.cardsGroup)
			}
		} else {
			// 有牌权，并且能一手出完的特殊情况
			if s.gameParams().PlayTolerance {
				playCardsGroup = s.playTolerance(player.handCards)
			}
		}

		// 特殊牌型判断后，依然是个无效牌，只能返回
		if playCardsGroup.Type == CardsTypeUnknown {
			s.Log().Warnw("play analyze failed invalid cards", "short", player.ShortId,
				"hand cards", player.handCards, "play cards", cards)
			return outer.ERROR_FASTERRUN_PLAY_CARDS_INVALID
		}
	}

	// 跟牌模式的校验
	if follow {
		// 跟牌牌型不同
		if playCardsGroup.Type != Bombs && playCardsGroup.Type != lastValidPlayCards.cardsGroup.Type {
			return outer.ERROR_FASTERRUN_PLAY_CARDS_SHOULD_BE_FOLLOW
		}

		// 主牌必须比跟的牌大, 只有牌型相同才比较，包括炸弹
		if playCardsGroup.Type == lastValidPlayCards.cardsGroup.Type && !playCardsGroup.Bigger(lastValidPlayCards.cardsGroup) {
			return outer.ERROR_FASTERRUN_PLAY_CARDS_SHOULD_BE_BIGGER
		}

		// 副牌数量不匹配(容错的情况下不用校验)
		if !tolerance && len(playCardsGroup.SideCards) != len(lastValidPlayCards.cardsGroup.SideCards) {
			return outer.ERROR_FASTERRUN_PLAY_CARDS_SIDE_CARD_LEN_ERR
		}
	}

	player.handCards = player.handCards.Remove(cards...)
	s.playRecords = append(s.playRecords, PlayCardsRecord{
		shortId:    player.ShortId,
		follow:     follow,
		cardsGroup: playCardsGroup,
		playAt:     tools.Now(),
	})
	seat := s.SeatIndex(player.ShortId)

	s.Log().Infow("play cards", "seat", seat, "short", player.ShortId, "play", cards, "hand", player.handCards)
	s.nextPlayer(s.nextSeatIndex(seat), &s.playRecords[len(s.playRecords)-1])

	return outer.ERROR_OK
}

// 下一位打牌的人
func (s *StatePlaying) nextPlayer(seat int, lastPlayInfo *PlayCardsRecord) {
	if s.gameOver() {
		s.SwitchTo(Settlement)
		return
	}

	player := s.fasterRunPlayers[seat]

	// 找到最后一次有效出牌
	lastValidPlay := s.lastValidPlayCards()

	// 如果出的是炸弹，并且转了一圈都没人能大，就算炸弹赢分
	var bombWinScore *outer.BombsWinScore
	if lastPlayInfo != nil && lastValidPlay.shortId == player.ShortId && lastPlayInfo.cardsGroup.Type == Bombs {
		var totalWinScore int64
		losers := make(map[int32]int64)
		for seatIdx, loserPlayer := range s.fasterRunPlayers {
			if loserPlayer.ShortId == player.ShortId {
				continue
			}

			score := s.bombWinScore()
			// 不允许负分，能减多少减多少
			if !s.gameParams().AllowScoreSmallZero && loserPlayer.score < score {
				score = common.Min(loserPlayer.score, score)
			}
			loserPlayer.updateScore(-score)
			losers[int32(seatIdx)] = score
			player.updateScore(score)
			totalWinScore += score
		}

		bombWinScore = &outer.BombsWinScore{
			WinScore: totalWinScore,
			Loser:    losers,
		}
	}

	waitingExpiration := tools.Now().Add(WaitingPlayExpiration)
	// 如果最后一个出牌人不是自己，就需要跟牌
	var follow, bigger bool // 跟牌，还是牌权出牌
	if lastValidPlay != nil && lastValidPlay.shortId != player.ShortId {
		follow = true

		// 手牌没有更大的牌了，延迟2秒直接过
		biggerCardGroups := player.handCards.FindBigger(lastValidPlay.cardsGroup)
		bigger = len(biggerCardGroups) > 0
		if !bigger {
			waitingExpiration = tools.Now().Add(WaitingPassExpiration)
		}
	}

	// 本次需要出牌的玩家
	s.waitingPlayShortId = player.ShortId
	s.waitingPlayFollow = follow

	ntf := &outer.FasterRunTurnNtf{
		WaitingEndAt:   waitingExpiration.UnixMilli(),
		FollowPlay:     follow,
		Bigger:         bigger,
		PlayingShortId: player.ShortId,
		PrevRecord:     lastPlayInfo.ToPB(),
		BombsWin:       bombWinScore,
	}
	s.room.Broadcast(ntf)

	s.actionTimer(waitingExpiration)

	s.Log().Infow("next player ", "seat", seat, "shortId", player.ShortId, "follow", follow, "hand cards", player.handCards, "prev play", lastPlayInfo, "ntf", ntf.String())
	s.Log().Infof(" ")
}

// 首局先出黑桃三检测
func (s *StatePlaying) checkFirstSpade3(cards PokerCards) bool {
	if !s.gameParams().FirstSpades3 {
		return true
	}

	if s.gameParams().DecideMasterType != 1 && s.gameParams().DecideMasterType != 3 {
		return true
	}

	for _, card := range cards {
		if card == Spades_3 {
			return true
		}
	}
	return false
}

func (s *StatePlaying) getPlayerAndSeatE(shortId int64) (*fasterRunPlayer, int, outer.ERROR) {
	player, seatIndex := s.findFasterRunPlayer(shortId)
	if player == nil {
		return nil, -1, outer.ERROR_PLAYER_NOT_IN_ROOM
	}

	return player, seatIndex, outer.ERROR_OK
}

// 取消行动倒计时
func (s *StatePlaying) cancelActionTimer() {
	s.room.CancelTimer(s.actionTimerId)
}

// 行动倒计时
func (s *StatePlaying) actionTimer(expireAt time.Time) {
	s.cancelActionTimer()
	s.actionTimerId = s.room.AddTimer(tools.XUID(), expireAt, func(dt time.Duration) {
		player, _ := s.findFasterRunPlayer(s.waitingPlayShortId)
		var cards PokerCards

		defer func() {
			s.Log().Infow("player timeout", "shortId", s.waitingPlayShortId, "cards", cards)
		}()

		if s.waitingPlayFollow {
			latest := s.lastValidPlayCards()
			biggerCardGroups := player.handCards.FindBigger(latest.cardsGroup)
			if len(biggerCardGroups) == 0 {
				s.pass(player)
				return
			}

			sort.Slice(biggerCardGroups, func(i, j int) bool {
				return biggerCardGroups[i].Cards[0].Point() < biggerCardGroups[i].Cards[0].Point()
			})
			cards = append(biggerCardGroups[0].Cards, biggerCardGroups[0].SideCards...)
		} else {
			cards = player.handCards.SideCards(1)
			if len(cards) == 0 {
				cards = PokerCards{player.handCards[0]}
			}
		}

		err := s.play(player, cards)
		if err != outer.ERROR_OK {
			s.Log().Errorw("timeout play failed", "err", err)
		}
	})
}

func (s *StatePlaying) gameOver() bool {
	for _, p := range s.fasterRunPlayers {
		if len(p.handCards) == 0 {
			return true
		}
	}

	// 只要有一位玩家分<=警戒值就结束
	for _, player := range s.fasterRunPlayers {
		// NOTE: 玩家每把结算后，会更新playerInfo，所以每把的GoldLine是固定的
		if player.score <= player.GetGoldLine() {
			s.scoreZeroOver = true
			return true
		}
	}

	return false
}

func int32ArrToPokerCards(cards []int32) PokerCards {
	result := make(PokerCards, 0, len(cards))
	for _, card := range cards {
		result = append(result, PokerCard(card))
	}
	return result
}

// 三张/飞机可少带出完：有牌权的一方出牌时，手牌是三张/飞机但牌带不满，手牌能一次全部出完
func (s *StatePlaying) playTolerance(handCards PokerCards) (playCardsGroup CardsGroup) {
	// 检查手牌能否组成三张或飞机,
	// 以及分别判断sidecars长度是否小于需要的长度

	// 飞机
	biggerPlane := handCards.FindBigger(CardsGroup{
		Type:  Plane,
		Cards: PokerCards{0},
	})
	if len(biggerPlane) > 0 {
		n := len(biggerPlane[0].Cards) / 3
		needSideCardsNum := n * 2 // 飞机需要的带牌数量
		spareCards := handCards.Remove(biggerPlane[0].Cards...)
		if len(spareCards) <= needSideCardsNum {
			playCardsGroup.Type = Plane
			playCardsGroup.Cards = biggerPlane[0].Cards
			playCardsGroup.SideCards = spareCards
		}
		return
	}

	// 三张
	biggerTrips := handCards.FindBigger(CardsGroup{
		Type:  Trips,
		Cards: PokerCards{0},
	})
	if len(biggerTrips) > 0 {
		needSideCardsNum := 2 // 三带二需要的带牌数量
		spareCards := handCards.Remove(biggerPlane[0].Cards...)
		if len(spareCards) <= needSideCardsNum {
			playCardsGroup.Type = TripsWithTwo
			playCardsGroup.Cards = biggerPlane[0].Cards
			playCardsGroup.SideCards = spareCards
		}
	}
	return
}

func (s *StatePlaying) followPlayTolerance(handCards PokerCards, lastCards CardsGroup) (playCardsGroup CardsGroup, tolerance bool) {
	var (
		needCardType     PokerCardsType // 需要的牌型
		needSideCardsNum int            // 需要的带牌张数
	)

	switch lastCards.Type {
	case TripsWithOne:
		needCardType = Trips
		needSideCardsNum = 1

	case TripsWithTwo:
		needCardType = Trips
		needSideCardsNum = 2

	case PlaneWithTwo:
		needCardType = Plane
		needSideCardsNum = len(lastCards.SideCards)
	default:
		s.Log().Warnw("lastCards is not valid type", "last cards", lastCards)
		return
	}

	bigger := handCards.FindBigger(CardsGroup{
		Type:  needCardType,
		Cards: PokerCards{0},
	})

	if len(bigger) > 0 {
		spareCards := handCards.Remove(bigger[0].Cards...)
		if len(spareCards) <= needSideCardsNum {
			playCardsGroup.Type = lastCards.Type
			playCardsGroup.Cards = bigger[0].Cards
			playCardsGroup.SideCards = spareCards
		}
	}

	tolerance = true
	return
}
