package mahjong

import (
	"server/proto/outermsg/outer"
)

// 碰牌操作
func (s *StatePlaying) operatePong(p *mahjongPlayer, seatIndex int) outer.ERROR {
	if len(s.peerRecords) == 0 {
		s.Log().Errorw("operate pong failed peerRecords len = 0",
			"room", s.room.RoomId, "player", p.ShortId)
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}
	// 获得打出的那张牌
	peer := s.peerRecords[len(s.peerRecords)-1]
	if peer.typ != playCardType {
		s.Log().Errorw("operate pong failed peer is drawCard",
			"room", s.room.RoomId, "player", p.ShortId, "peerRecords", s.peerRecordsLog())
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	// 检查,桌面牌最后一张是否和碰的牌一致
	tail := len(s.cardsInDesktop) - 1
	desktopCard := s.cardsInDesktop[tail]
	if desktopCard != peer.card {
		s.Log().Errorw("operate pong logic error",
			"room", s.room.RoomId, "player", p.ShortId, "peerRecords", s.peerRecordsLog(), "desktopCard", desktopCard)
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	// 检查,不可能自己碰自己的牌
	if peer.seat == seatIndex {
		s.Log().Errorw("unexpected logic ",
			s.room.RoomId, "player", p.ShortId, "peer", peer, "seat", seatIndex)
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	var err error
	p.handCards, _, err = p.handCards.Pong(peer.card)
	if err != nil {
		s.Log().Errorw("unexpected logic pong failed",
			"room", s.room.RoomId, "seat", seatIndex, "player", p.ShortId, "peer", peer, "err", err)
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	p.pong[peer.card.Int32()] = s.mahjongPlayers[peer.seat].ShortId                                              // 加入自己的碰牌组
	s.actionMap[seatIndex] = &action{seat: seatIndex, acts: []outer.ActionType{outer.ActionType_ActionPlayCard}} // 碰后新增出牌行为
	s.cardsInDesktop = s.cardsInDesktop[:tail]
	p.resetPassHand()
	return outer.ERROR_OK
}
