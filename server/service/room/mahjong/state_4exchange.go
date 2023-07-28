package mahjong

import (
	"reflect"
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/proto/outermsg/outer"
)

// 换三张

type StateExchange3 struct {
	*Mahjong
	timerId string
}

func (s *StateExchange3) State() int {
	return Exchange3
}

func (s *StateExchange3) Enter() {
	s.currentStateEndAt = tools.Now().Add(Exchange3Expiration)
	s.room.Broadcast(&outer.MahjongBTEExchange3Ntf{EndAt: s.currentStateEndAt.UnixMilli()})
	s.timerId = s.room.AddTimer(tools.XUID(), s.currentStateEndAt, func(time.Duration) {
		s.stateEnd()
	})
	s.Log().Infow("[Mahjong] enter state  exchange3", "room", s.room.RoomId)
}

func (s *StateExchange3) Leave() {
	s.Log().Infow("[Mahjong] leave state exchange3", "room", s.room.RoomId)
	for seatIndex, player := range s.mahjongPlayers {
		s.Log().Infow("hand cards", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId, "cards", player.handCards)
	}
}

// 换三张结束
func (s *StateExchange3) stateEnd() {
	s.room.CancelTimer(s.timerId)
	for idx, player := range s.mahjongPlayers {
		s.checkAndInit(player)
		rival := s.mahjongPlayers[player.exchange.ToSeatIndex]
		s.checkAndInit(rival)
		rival.exchange.FromSeatIndex = int32(idx)
		rival.exchange.CardsFrom = player.exchange.CardsTo
		rival.handCards = rival.handCards.
			Remove(Card(rival.exchange.CardsTo[0])).
			Remove(Card(rival.exchange.CardsTo[1])).
			Remove(Card(rival.exchange.CardsTo[2])).
			Insert(Card(rival.exchange.CardsFrom[0])).
			Insert(Card(rival.exchange.CardsFrom[1])).
			Insert(Card(rival.exchange.CardsFrom[2]))

		s.Log().Infow("exchanging ", "room", s.room.RoomId, "seat", idx, "player", player.ShortId,
			"to seat", player.exchange.ToSeatIndex, "to player", rival.ShortId, "to cards", player.exchange.CardsTo,
		)
	}

	for _, player := range s.mahjongPlayers {
		s.room.SendToPlayer(player.ShortId, &outer.MahjongBTEExchange3EndNtf{
			Ex3Info: player.exchange,
			Cards:   player.handCards.ToSlice(),
		})
	}

	// 状态结束，给个换牌动画播放延迟，进入定缺
	s.room.AddTimer(tools.XUID(), tools.Now().Add(Exchange3ShowDuration), func(time.Duration) {
		s.SwitchTo(DecideIgnore)
	})
}

func (s *StateExchange3) Handle(shortId int64, v any) (result any) {
	switch msg := v.(type) {
	case *outer.MahjongBTEExchange3Req:
		if len(msg.Index) != 3 {
			return outer.ERROR_MAHJONG_EXCHANGE3_LEN_ERROR
		}

		// 三张不能有相同
		if msg.Index[0] == msg.Index[1] || msg.Index[1] == msg.Index[2] {
			return outer.ERROR_MAHJONG_EXCHANGE3_INDEX_EQUAL
		}

		// 同花色换三张
		if s.room.GameParams.Mahjong.HuanSanZhang == 0 {
			if Card(msg.Index[0]).Color() != Card(msg.Index[1]).Color() ||
				Card(msg.Index[1]).Color() != Card(msg.Index[2]).Color() {
				return outer.ERROR_MAHJONG_EXCHANGE3_COLOR_ERROR
			}
		}

		player, _ := s.findMahjongPlayer(shortId)
		if player == nil {
			return outer.ERROR_PLAYER_NOT_IN_ROOM
		}

		// 不能重复操作
		if player.exchange != nil {
			return outer.ERROR_MAHJONG_EXCHANGE3_OPERATED
		}

		var exchange3Cards Cards
		for _, idx := range msg.Index {
			if idx < 0 || idx >= int32(player.handCards.Len()) {
				return outer.ERROR_MAHJONG_EXCHANGE3_INDEX_ERROR
			}

			exchange3Cards = append(exchange3Cards, player.handCards[idx])
		}

		seatIndex := s.SeatIndex(player.ShortId)
		nextSeatIndex := s.nextSeatIndex(seatIndex)
		player.exchange = &outer.Exchange3Info{
			CardsFrom:     nil,
			FromSeatIndex: -1,
			CardsTo:       exchange3Cards.ToSlice(),
			ToSeatIndex:   int32(s.nextSeatIndex(s.SeatIndex(player.ShortId))),
		}

		s.Log().Infow("MahjongBTEExchange3Req ", "room", s.room.RoomId, "playerID", player.ShortId,
			"cards", exchange3Cards, "seatIndex", seatIndex, "nextSeatIndex", nextSeatIndex)

		// 广播玩家确认换三张
		s.room.Broadcast(&outer.MahjongBTEExchange3PlayerReadyNtf{ShortId: shortId})

		// 所有人都准备好了，结束换三张
		if s.isAllReady() {
			s.stateEnd()
		}
		return &outer.MahjongBTEExchange3Rsp{}
	default:
		s.Log().Warnw("exchange 3 status has received an unknown message", "msg", reflect.TypeOf(msg).String())
	}
	return outer.ERROR_MAHJONG_STATE_MSG_INVALID
}

func (s *StateExchange3) checkAndInit(player *mahjongPlayer) {
	if player.exchange == nil {
		player.exchange = &outer.Exchange3Info{
			CardsFrom:     nil,
			FromSeatIndex: -1,
			CardsTo:       player.handCards[:3].ToSlice(),
			ToSeatIndex:   int32(s.nextSeatIndex(s.SeatIndex(player.ShortId))),
		}
	}
}

func (s *StateExchange3) isAllReady() bool {
	for _, player := range s.mahjongPlayers {
		if player.exchange == nil {
			return false
		}
	}
	return true
}
