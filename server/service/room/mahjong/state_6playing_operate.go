package mahjong

import (
	"github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/tools"
	"server/common/log"
	"server/proto/outermsg/outer"
	"time"
)

// 碰杠胡过
func (s *StatePlaying) operate(player *mahjongPlayer, seatIndex int, op outer.ActionType, card Card) (ok bool, err outer.ERROR) {
	var (
		act   *action
		exist bool
	)

	if op == outer.ActionType_ActionPlayCard {
		// 此函数不受理打牌
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	// 检查是否是行动者,以及行为是否有效
	if act, exist = s.actionMap[seatIndex]; !exist {
		return false, outer.ERROR_MAHJONG_ACTION_PLAYER_NOT_MATCH
	} else if !act.isValidAction(op) {
		return false, outer.ERROR_MAHJONG_ACTION_PLAYER_NOT_OPERA
	}

	delete(s.actionMap, seatIndex)

	ntf := &outer.MahjongBTEOperaNtf{
		OpShortId: player.ShortId,
		OpType:    op,
	}

	var (
		playCardAfterPongNtf proto.Message
		nextDrawShortIndex   = s.nextSeatIndex(s.peerCards[len(s.peerCards)-1].seat) //提前计算下家摸牌的座位
		qiangGang            bool
	)

	peer := s.peerCards[len(s.peerCards)-1]
	switch op {
	case outer.ActionType_ActionPass:
		ok = true
	case outer.ActionType_ActionPong:
		ok, err, playCardAfterPongNtf = s.operatePong(player, seatIndex)
		ntf.Card = peer.card.Int32() // 碰的牌

	case outer.ActionType_ActionGang:
		ok, qiangGang, err = s.operateGang(player, seatIndex, card, ntf)
		ntf.Card = card.Int32() // 杠的牌

	case outer.ActionType_ActionHu:
		ok, err = s.operateHu(player, seatIndex, ntf)
	default:
		log.Errorw("unknown action op",
			"roomId", s.room.RoomId, "player", player.ShortId, "op", op)
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	if !ok {
		return
	}

	// 除了过以外的操作都需要广播通知
	if op != outer.ActionType_ActionPass {
		s.room.Broadcast(ntf)
	}

	// 所有操作者都执行完了，进入下一次摸牌，或者结束本局
	switch op {
	case outer.ActionType_ActionPong:
		s.room.Broadcast(playCardAfterPongNtf)
		// 操作碰完成，需要出牌，不用检查结算
		return
	case outer.ActionType_ActionGang:
		// 杠操作完成，需要自身摸一张牌(如果还能摸的情况下)
		nextDrawShortIndex = seatIndex

		// 抢杠胡，检查是否有人能抢
		if qiangGang {
			var (
				qiang            []int64
				qiangActionEndAt time.Time
			)
			for seat, other := range s.mahjongPlayers {
				if other.ShortId == player.ShortId {
					continue
				}

				if hu := other.handCards.Insert(card).IsHu(other.lightGang, other.darkGang, other.pong); hu != HuInvalid {
					if qiangActionEndAt.IsZero() {
						qiangActionEndAt = tools.Now().Add(pongGangHuGuoExpire)
					}
					newAction := action{}
					newAction.currentActions = append(newAction.currentActions, outer.ActionType_ActionHu, outer.ActionType_ActionPass)
					newAction.currentHus = append(newAction.currentHus, hu.PB())
					s.actionMap[seat] = &newAction // 抢杠胡操作
					qiang = append(qiang, other.ShortId)

					s.room.SendToPlayer(other.ShortId, &outer.MahjongBTETurnNtf{
						TotalCards:    int32(s.cards.Len()),
						ActionShortId: other.ShortId,
						ActionEndAt:   qiangActionEndAt.UnixMilli(),
						ActionType:    newAction.currentActions,
						HuType:        newAction.currentHus,
						NewCard:       -1, // 客户端自己取桌面牌最后一张
					})
				}
			}

			// 有人能抢胡
			if len(qiang) > 0 {
				s.actionTimer(qiangActionEndAt)
				s.room.Broadcast(&outer.MahjongBTETurnNtf{
					TotalCards:  int32(s.cards.Len()),
					ActionEndAt: qiangActionEndAt.UnixMilli(),
				}, qiang...)
				return true, outer.ERROR_OK
			}
		}
	}

	// 操作杠、胡、过、完成，需要检查结算
	if len(s.actionMap) == 0 {
		if s.cards.Len() > 0 {
			s.drawCard(nextDrawShortIndex)
		} else {
			s.SwitchTo(Settlement)
		}
	}
	return true, outer.ERROR_OK
}

// 碰牌操作
func (s *StatePlaying) operatePong(p *mahjongPlayer, seatIndex int) (bool, outer.ERROR, proto.Message) {
	if len(s.peerCards) == 0 {
		log.Errorw("operate pong failed peerCards len = 0",
			"roomId", s.room.RoomId, "player", p.ShortId)
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID, nil
	}
	// 获得打出的那张牌
	peer := s.peerCards[len(s.peerCards)-1]
	if peer.typ != playCardType {
		log.Errorw("operate pong failed peer is drawCard",
			"roomId", s.room.RoomId, "player", p.ShortId, "peerCards", s.peerCards)
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID, nil
	}

	// 检查,桌面牌最后一张是否和碰的牌一致
	tail := len(s.cardsInDesktop) - 1
	desktopCard := s.cardsInDesktop[tail]
	if desktopCard != peer.card {
		log.Errorw("operate pong logic error",
			"roomId", s.room.RoomId, "player", p.ShortId, "peerCards", s.peerCards, "desktopCard", desktopCard)
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID, nil
	}

	// 检查,不可能自己碰自己的牌
	if peer.seat == seatIndex {
		log.Errorw("unexpected logic ",
			s.room.RoomId, "player", p.ShortId, "peer", peer, "seatIndex", seatIndex)
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID, nil
	}

	s.cardsInDesktop = s.cardsInDesktop[:tail]                      // 删除桌面牌
	p.pong[peer.card.Int32()] = s.mahjongPlayers[peer.seat].ShortId // 加入自己的碰牌组

	// 碰成功后，需要删除其余执行者(不能一家碰了，一家还能胡)
	s.actionMap = make(map[int]*action)
	s.actionMap[seatIndex] = &action{currentActions: []outer.ActionType{outer.ActionType_ActionPlayCard}} // 碰后出牌行为

	actionExpireAt := tools.Now().Add(playCardExpire)
	s.actionTimer(actionExpireAt) // 碰后出牌操作时间
	playCardNtf := &outer.MahjongBTETurnNtf{
		TotalCards:    int32(s.cards.Len()),
		ActionShortId: p.ShortId,
		ActionEndAt:   actionExpireAt.UnixMilli(),
		ActionType:    s.actionMap[seatIndex].currentActions,
	}

	return true, outer.ERROR_OK, playCardNtf
}

// 杠牌操作
func (s *StatePlaying) operateGang(p *mahjongPlayer, seatIndex int, card Card, ntf *outer.MahjongBTEOperaNtf) (ok, qiang bool, err outer.ERROR) {
	if len(s.peerCards) == 0 {
		log.Errorw("operate gang failed peerCards len = 0",
			"roomId", s.room.RoomId, "player", p.ShortId)
		return false, false, outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	// 获得最后一次操作的牌
	peer := s.peerCards[len(s.peerCards)-1]
	switch peer.typ {
	case drawCardType: // 摸牌
		if _, ok := p.pong[card.Int32()]; ok {
			delete(p.pong, card.Int32())
		} else {
			if p.handCards.CanGangTo(card) {
				log.Errorw("operate gang failed player cannot Gang",
					"roomId", s.room.RoomId, "player", p.ShortId)
				return false, false, outer.ERROR_MSG_REQ_PARAM_INVALID
			}
			// 暗杠（下雨）
			p.handCards, _, _ = p.handCards.Gang(card)
			ntf.GangType = 2
		}
		p.darkGang[card.Int32()] = p.ShortId

	case playCardType: // 打牌
		if _, ok := p.pong[card.Int32()]; ok {
			// 下杠,可以抢杠胡
			delete(p.pong, card.Int32())
			ntf.GangType = 1 // 明杠（刮风）
			s.AppendPeerCard(lightGangType, card, seatIndex)
			qiang = true
		} else {
			if p.handCards.CanGangTo(card) {
				log.Errorw("operate gang failed player cannot Gang",
					"roomId", s.room.RoomId, "player", p.ShortId)
				return false, false, outer.ERROR_MSG_REQ_PARAM_INVALID
			}
			// 明杠（刮风）
			p.handCards, _, _ = p.handCards.Gang(card)
			ntf.GangType = 2
		}
		p.lightGang[card.Int32()] = s.mahjongPlayers[peer.seat].ShortId
	}

	// 杠成功后，需要删除其余执行者(不能一家杠了，一家还能胡)
	s.actionMap = make(map[int]*action)
	return true, qiang, outer.ERROR_OK
}

// 胡牌操作
func (s *StatePlaying) operateHu(p *mahjongPlayer, seatIndex int, ntf *outer.MahjongBTEOperaNtf) (bool, outer.ERROR) {
	// TODO

	// 胡成功后，删除Gang和Pong(可以一炮多响,但是胡了，就不能碰、杠)
	for seat, act := range s.actionMap {
		if act.isValidAction(outer.ActionType_ActionPong) {
			act.remove(outer.ActionType_ActionPong)
		}
		if act.isValidAction(outer.ActionType_ActionGang) {
			act.remove(outer.ActionType_ActionGang)
			act.currentGang = []int32{}
		}
		if len(act.currentActions) == 0 {
			delete(s.actionMap, seat)
		}
	}
	return true, outer.ERROR_OK
}
