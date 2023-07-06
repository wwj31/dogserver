package mahjong

import (
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/proto/outermsg/outer"
)

// 换三张

type StateExchange3 struct {
	*Mahjong
}

func (s *StateExchange3) State() int {
	return Exchange3
}

func (s *StateExchange3) Enter() {
	s.room.Broadcast(&outer.MahjongBTEExchange3Ntf{})
	s.room.AddTimer(tools.XUID(), tools.Now().Add(15*time.Second), func(time.Duration) {
		s.stateEnd()
	})
	log.Infow("[Mahjong] leave state  exchange3", "room", s.room.RoomId)
}

func (s *StateExchange3) Leave() {
	log.Infow("[Mahjong] leave state exchange3", "room", s.room.RoomId)
}

// 换三张结束
func (s *StateExchange3) stateEnd() {
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

		log.Infow("exchanging ", "roomId", s.room.RoomId, "seatIndex", idx,
			"player", player.ShortId, " exchange to seatIndex", player.exchange.ToSeatIndex,
			"rival", rival.ShortId, "rival cards from", rival.exchange.CardsFrom,
			"rival cards to seatIndex", rival.exchange.ToSeatIndex,
			"rival cards to", rival.exchange.CardsTo,
		)
	}

	for _, player := range s.mahjongPlayers {
		s.room.SendToPlayer(player.ShortId, &outer.MahjongBTEExchange3EndNtf{
			Ex3Info: player.exchange,
			Cards:   player.handCards.ToSlice(),
		})
	}

	// 状态结束，给个换牌动画播放延迟，进入定缺
	s.room.AddTimer(tools.XUID(), tools.Now().Add(3*time.Second), func(time.Duration) {
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
		if msg.Index[0] == msg.Index[1] ||
			msg.Index[1] == msg.Index[2] ||
			msg.Index[0] == msg.Index[2] {
			return outer.ERROR_MAHJONG_EXCHANGE3_INDEX_EQUAL
		}

		player := s.findMahjongPlayer(shortId)
		if player == nil {
			return outer.ERROR_PLAYER_NOT_IN_ROOM
		}

		// 不能重复操作
		if player.exchange != nil {
			return outer.ERROR_MAHJONG_EXCHANGE3_OPERATED
		}

		var exchange3Cards Cards
		for _, idx := range msg.Index {
			if idx < 0 || idx > 12 {
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
			ToSeatIndex:   s.nextSeatIndex(s.SeatIndex(player.ShortId)),
		}

		log.Infow("MahjongBTEExchange3Req ", "roomId", s.room.RoomId, "playerID", player.ShortId,
			"cards", exchange3Cards,
			"seatIndex", seatIndex,
			"nextSeatIndex", nextSeatIndex,
		)

		// 所有人都准备好了，结束换三张
		if s.isAllReady() {
			s.stateEnd()
		}
		return &outer.MahjongBTEExchange3Rsp{}
	}
	return nil
}

func (s *StateExchange3) checkAndInit(player *mahjongPlayer) {
	if player.exchange == nil {
		player.exchange = &outer.Exchange3Info{
			CardsFrom:     nil,
			FromSeatIndex: -1,
			CardsTo:       player.handCards[:3].ToSlice(),
			ToSeatIndex:   s.nextSeatIndex(s.SeatIndex(player.ShortId)),
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
