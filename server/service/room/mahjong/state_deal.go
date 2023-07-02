package mahjong

import (
	"server/common/log"
	"server/proto/outermsg/outer"
	"server/service/room"
)

// 发牌状态

type StateDeal struct {
	*Mahjong
}

func (s *StateDeal) State() int {
	return Deal
}

func (s *StateDeal) Enter(fsm *room.FSM) {
	log.Infow("Mahjong enter deal", "room", s.room.RoomId)
	s.cards = RandomCards() // 总共108张
	var i int
	for _, player := range s.mahjongPlayers {
		player.handCards = append(Cards{}, s.cards[i:i+13]...).Sort()
		i += 13

		s.room.SendToPlayer(player.ShortId, &outer.MahjongBTEDealNtf{
			Cards: player.handCards.ToSlice(),
		})
		log.Infow("dealing", "room", s.room.RoomId, "player", player.ShortId, "cards", player.handCards)
	}

	s.cards = s.cards[52:]
	log.Infow("deal finished cards", "room", s.room.RoomId, "spare cards", s.cards)

}

func (s *StateDeal) Leave(fsm *room.FSM) {
	log.Infow("Mahjong leave deal", "room", s.room.RoomId)
}

func (s *StateDeal) Handle(fsm *room.FSM, v any, shortId int64) (result any) {
	switch msg := v.(type) {
	case *outer.JoinRoomReq:
		_ = msg
	}
	return nil
}
