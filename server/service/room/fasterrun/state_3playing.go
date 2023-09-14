package fasterrun

import (
	"reflect"
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

	s.nextPlayer(s.masterIndex)
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
		// TODO ...
		return &outer.FasterRunPassRsp{}

	default:
		s.Log().Warnw("playing status has received an unknown message", "msg", reflect.TypeOf(msg).String())
	}
	return outer.ERROR_FASTERRUN_STATE_MSG_INVALID
}

// 打牌
func (s *StatePlaying) play(player *fasterRunPlayer, cards PokerCards) outer.ERROR {
	if len(cards) == 0 {
		return outer.ERROR_FASTERRUN_PLAY_CARDS_LEN_EMPTY
	}

	// 检查是否轮到该玩家出牌
	if s.waitingPlayShortId != player.ShortId {
		return outer.ERROR_FASTERRUN_PLAY_CARDS_LEN_EMPTY
	}

	// 检查要打的牌是否全部在手牌中
	if !player.handCards.CheckExist(cards...) {
		s.Log().Warnw("play check exist failed", "short", player.ShortId,
			"hand cards", player.handCards, "play cards", cards)
		return outer.ERROR_FASTERRUN_PLAY_CARDS_MISS
	}

	// 分析牌型，检查牌型是否有效
	playCardsGroup := cards.AnalyzeCards(s.gameParams().AAAIsBombs)
	if playCardsGroup.Type == CardsTypeUnknown {
		return outer.ERROR_FASTERRUN_PLAY_CARDS_INVALID
	}

	// 如果需要跟牌，检查牌型是否符合跟牌牌型
	lastValidPlayCards := s.lastValidPlayCards()
	if lastValidPlayCards != nil && lastValidPlayCards.shortId != player.ShortId {
		// 跟牌牌型不同
		if playCardsGroup.Type != lastValidPlayCards.records.Type {
			return outer.ERROR_FASTERRUN_PLAY_CARDS_SHOULD_BE_FOLLOW
		}

		// 主牌必须比跟的牌大
		if !playCardsGroup.Bigger(lastValidPlayCards.records) {
			return outer.ERROR_FASTERRUN_PLAY_CARDS_SHOULD_BE_BIGGER
		}

		// 副牌数量不匹配
		if len(playCardsGroup.SideCards) != len(lastValidPlayCards.records.SideCards) {
			return outer.ERROR_FASTERRUN_PLAY_CARDS_SIDE_CARD_LEN_ERR
		}

	}

}

// 下一位打牌的人
func (s *StatePlaying) nextPlayer(seat int) {
	player := s.fasterRunPlayers[seat]

	var follow bool // 跟牌，还是牌权出牌

	// 找到最后一次有效出牌，以此决定本次出牌是否有牌权
	// 如果最后一个出牌人不是自己，就需要跟牌
	lastValidPlay := s.lastValidPlayCards()
	if lastValidPlay != nil && lastValidPlay.shortId != player.ShortId {
		follow = true
	}

	s.waitingPlayShortId = player.ShortId
	s.waitingPlayFollow = follow

	waitingExpiration := tools.Now().Add(WaitingPlayExpiration)
	ntf := &outer.FasterRunPlayCardNtf{
		WaitingEndAt:   waitingExpiration.UnixMilli(),
		FollowPlay:     follow,
		PlayingShortId: player.ShortId,
	}
	s.room.Broadcast(ntf)

	s.actionTimer(waitingExpiration)
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
