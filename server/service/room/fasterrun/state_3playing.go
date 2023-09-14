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
	case *outer.FasterRunPlayCardReq: // 打牌
		// TODO ...
		return &outer.FasterRunPlayCardRsp{HandCards: player.handCards.ToPB()}

	default:
		s.Log().Warnw("playing status has received an unknown message", "msg", reflect.TypeOf(msg).String())
	}
	return outer.ERROR_FASTERRUN_STATE_MSG_INVALID
}

func (s *StatePlaying) nextPlayer(seat int) {
	player := s.fasterRunPlayers[seat]

	var (
		latestPlay *PlayCardsHistory
		follow     bool // 跟牌，还是牌权出牌
	)
	// 反着找最后一个出牌人
	for i := len(s.playRecords) - 1; i >= 0; i-- {
		record := s.playRecords[i]
		if record.records.Type == CardsTypeUnknown {
			continue
		}
		latestPlay = &record
	}

	// 如果最后一个出牌人不是自己，就需要跟牌
	if latestPlay != nil && latestPlay.shortId != player.ShortId {
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
