package fasterrun

import (
	"reflect"
	"server/common"
	"sort"
	"time"

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

	// 分析牌型，检查牌型是否有效
	playCardsGroup := cards.AnalyzeCards(s.gameParams().AAAIsBombs)
	if playCardsGroup.Type == CardsTypeUnknown {
		s.Log().Warnw("play analyze failed invalid cards", "short", player.ShortId,
			"hand cards", player.handCards, "play cards", cards)
		return outer.ERROR_FASTERRUN_PLAY_CARDS_INVALID
	}

	// 如果需要跟牌，检查牌型是否符合跟牌牌型
	lastValidPlayCards := s.lastValidPlayCards()
	follow := lastValidPlayCards != nil && lastValidPlayCards.shortId != player.ShortId
	if follow {
		// 跟牌牌型不同
		if playCardsGroup.Type != lastValidPlayCards.cardsGroup.Type {
			return outer.ERROR_FASTERRUN_PLAY_CARDS_SHOULD_BE_FOLLOW
		}

		// 主牌必须比跟的牌大
		if !playCardsGroup.Bigger(lastValidPlayCards.cardsGroup) {
			return outer.ERROR_FASTERRUN_PLAY_CARDS_SHOULD_BE_BIGGER
		}

		// 副牌数量不匹配
		if len(playCardsGroup.SideCards) != len(lastValidPlayCards.cardsGroup.SideCards) {
			if s.gameParams().PlayTolerance {
				// TODO ...
			}

			if s.gameParams().FollowPlayTolerance {
				// TODO ...
			}
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

	var follow bool // 跟牌，还是牌权出牌

	// 找到最后一次有效出牌，以此决定本次出牌是否有牌权
	// 如果最后一个出牌人不是自己，就需要跟牌
	lastValidPlay := s.lastValidPlayCards()
	if lastValidPlay != nil && lastValidPlay.shortId != player.ShortId {
		follow = true
	}

	// 如果出的是炸弹，并且转了一圈都没人能大，就算炸弹赢分
	var bombWinScore *outer.BombsWinScore
	if lastPlayInfo != nil && lastValidPlay.shortId == player.ShortId && lastPlayInfo.cardsGroup.Type == Bombs {
		var totalWinScore int64
		loser := make(map[int32]int64)
		for seatIdx, runPlayer := range s.fasterRunPlayers {
			score := s.bombWinScore()
			// 不允许负分，能减多少减多少
			if !s.gameParams().AllowScoreSmallZero && runPlayer.score < score {
				score = common.Min(runPlayer.score, score)
			}
			runPlayer.updateScore(-score)
			loser[int32(seatIdx)] = score
			player.updateScore(score)
			totalWinScore += score
		}

		bombWinScore = &outer.BombsWinScore{
			WinScore: totalWinScore,
			Loser:    loser,
		}
	}

	// 本次需要出牌的玩家
	s.waitingPlayShortId = player.ShortId
	s.waitingPlayFollow = follow

	waitingExpiration := tools.Now().Add(WaitingPlayExpiration)
	ntf := &outer.FasterRunTurnNtf{
		WaitingEndAt:   waitingExpiration.UnixMilli(),
		FollowPlay:     follow,
		PlayingShortId: player.ShortId,
		PrevRecord:     lastPlayInfo.ToPB(),
		BombsWin:       bombWinScore,
	}
	s.room.Broadcast(ntf)

	s.actionTimer(waitingExpiration)

	s.Log().Infow("next player ", "seat", seat, "shortId", player.ShortId, "follow", follow, "hand cards", player.handCards, "prev play", lastPlayInfo)
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
		s.Log().Infow("player timeout", "shortId", s.waitingPlayShortId)
		player, _ := s.findFasterRunPlayer(s.waitingPlayShortId)
		var cards PokerCards
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

		if !s.gameParams().AllowScoreSmallZero && player.score <= 0 {
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
