package mahjong

import (
	"math"

	"server/common"
	"server/proto/outermsg/outer"
)

// 胡牌操作
func (s *StatePlaying) operateHu(p *mahjongPlayer, seatIndex int, ntf *outer.MahjongBTEOperaNtf) outer.ERROR {
	// 获得最后一次操作的牌
	lastPeerIndex := len(s.peerRecords) - 1
	peer := s.peerRecords[lastPeerIndex]
	var hu HuType
	switch peer.typ {
	case drawCardType: // 自摸
		hu = p.handCards.IsHu(p.lightGang, p.darkGang, p.pong, peer.card, s.gameParams())

	case playCardType, GangType1, GangType3: // 点炮,抢杠
		hu = p.handCards.Insert(peer.card).IsHu(p.lightGang, p.darkGang, p.pong, peer.card, s.gameParams())
		// 胡成功，删除最后打的那张牌
	}

	if hu == HuInvalid {
		s.Log().Errorw("operate hu invalid",
			"room", s.room.RoomId, "seat", seatIndex, "player", p.ShortId, "hand", p.handCards)
		return outer.ERROR_MAHJONG_HU_INVALID
	}

	s.huSeat = append(s.huSeat, int32(seatIndex))
	p.hu = hu
	p.huPeerIndex = len(s.peerRecords) - 1
	if peer.typ == GangType1 || peer.typ == GangType3 {
		peer.afterQiangPass = nil // 抢杠成功，杠的人，杠失败

		s.room.Broadcast(&outer.MahjongBTEGangResultNtf{
			OpShortId:        s.mahjongPlayers[peer.seat].ShortId,
			QiangGangShortId: p.ShortId,
			Card:             peer.card.Int32(),
			LoseScores:       nil,
			CurrentScores:    nil,
		})
		s.Log().Infow("qiang gang success", "room", s.room.RoomId, "seat", seatIndex, "gang seat", peer.seat)
	}

	// 计算额外加番
	p.huExtra = s.huExtra(seatIndex)
	p.huGen = int32(s.huGen(seatIndex))
	p.winScore = make(map[int32]int64)
	p.finalStatsMsg.TotalHu++
	p.passHandHuFan = 0

	// 胡成功后，删除Gang和Pong(可以一炮多响,但是有人胡了就不能再碰、杠)
	for seat, act := range s.actionMap {
		// 如果行为中有胡，保留碰杠
		if act.isValidAction(outer.ActionType_ActionHu) {
			continue
		}

		act.remove(outer.ActionType_ActionGang)
		act.remove(outer.ActionType_ActionPong)
		act.gang = []int32{}

		// 只剩下[过]操作，删掉
		if len(act.acts) == 1 && act.isValidAction(outer.ActionType_ActionPass) {
			act.remove(outer.ActionType_ActionPass)
		}

		// 如果删空了，直接删掉整个行为
		if len(act.acts) == 0 {
			delete(s.actionMap, seat)
		}
	}

	s.canHus[seatIndex] = true

	if s.husWasAllDo() {
		s.huSettlement(ntf)
	}
	return outer.ERROR_OK
}

// 是否所有能胡的人都胡了
func (s *StatePlaying) huIndex(seat int) int32 {
	for idx, huSeat := range s.huSeat {
		if seat == int(huSeat) {
			return int32(idx + 1)
		}
	}
	return -1
}

// 是否所有能胡的人都胡了,至少有1人胡，才为true
func (s *StatePlaying) husWasAllDo() bool {
	if len(s.canHus) == 0 {
		return false
	}

	for _, isHu := range s.canHus {
		if !isHu {
			return false
		}
	}
	return true
}

// 是否所有能胡的人都过了
func (s *StatePlaying) husWasAllPass() bool {
	return len(s.canHus) == 0
}

// 是否已经有人胡过了
func (s *StatePlaying) husWasDo() bool {
	for _, isHu := range s.canHus {
		if isHu {
			return true
		}
	}
	return false
}

// 是否还有人能胡，但是没胡
func (s *StatePlaying) needWaiting4Hu() bool {
	for _, isHu := range s.canHus {
		if !isHu {
			return true
		}
	}
	return false
}

// 胡牌小结算，等能胡的所有人都胡了，才执行结算操作
func (s *StatePlaying) huSettlement(ntf *outer.MahjongBTEOperaNtf) {
	huResultNtf := &outer.MahjongBTEHuResultNtf{
		LoseScores: make(map[int32]int64),
		Card:       s.peerRecords[len(s.peerRecords)-1].card.Int32(),
	}

	if ntf != nil {
		ntf.HuResult = huResultNtf
	}

	loseScores := map[int32]int64{}
	defer func() {

		for _, p := range s.mahjongPlayers {
			huResultNtf.CurrentScores = append(huResultNtf.CurrentScores, s.immScore(p.ShortId))
		}

		if s.gameParams().HuImmediatelyScore {
			huResultNtf.LoseScores = loseScores
		}

		if ntf == nil {
			s.room.Broadcast(huResultNtf)
		}
		s.Log().Infow("hu settlement ", "MahjongBTEHuResultNtf", huResultNtf.String())
		s.currentAction = nil
	}()

	// 呼叫转移,只在1对1的时候才生效,一炮多响，不会触发呼叫转移
	if len(s.canHus) == 1 {
		// 先找到胡的那个人
		var (
			p      *mahjongPlayer
			huSeat int
		)
		for huSeat, _ = range s.canHus {
			p = s.mahjongPlayers[huSeat]
		}

		// 判断杠上炮的情况
		if p.huExtra == GangShangPao {
			gangPeerIndex := len(s.peerRecords) - 3
			peerRecord := s.peerRecords[gangPeerIndex]          // 杠的那次记录
			rivalGang := s.mahjongPlayers[peerRecord.seat]      // 杠的人
			rivalGangInfo := rivalGang.gangInfos[gangPeerIndex] // 杠信息
			totalGangScore := rivalGangInfo.totalWinScore       // 本次转移的总分

			rivalGang.updateScore(-totalGangScore)
			rivalGang.gangTotalScore -= totalGangScore // 呼叫转移

			s.Log().Infow("lose score update by gangShangPao", "room", s.room.RoomId,
				"shortId", rivalGang.ShortId, "current score", rivalGang.score, "sub score", totalGangScore)

			p.updateScore(totalGangScore)
			p.gangTotalScore += totalGangScore // 退杠

			s.Log().Infow("win score update by gangShangPao", "room", s.room.RoomId,
				"shortId", p.ShortId, "current score", p.score, "sub score", totalGangScore)

			// 如果需要实时结算，就把结算分放入通知
			if s.gameParams().HuImmediatelyScore {
				huResultNtf.ShiftGangScore = totalGangScore
				huResultNtf.ShiftGangScoreSeat = int32(peerRecord.seat)
			}

			s.Log().Infow("gangShangPao,exchange gang score",
				"room", s.room.RoomId, "hu", p.hu, "hu shortId", p.ShortId, "gang shortId", rivalGang.ShortId,
				"score", totalGangScore, "rival gang loser seats", rivalGangInfo.loserSeats)
		}

		/////////////////////////////////////  以下为自摸胡的情况 /////////////////////////////////////
		huPeer := s.peerRecords[p.huPeerIndex]
		if huPeer.typ == drawCardType {
			winScore := s.huScore(p, true)

			// 其余没胡的都要赔钱
			for seat, other := range s.mahjongPlayers {
				if other.hu != HuInvalid || other.ShortId == p.ShortId {
					continue
				}

				s.AWinB(huSeat, seat, winScore)
				loseScores[int32(seat)] = winScore // 统计输家和输的分
			}

			p.winScore = loseScores

			// 组装通知消息数据
			huResultNtf.ZiMo = true
			huResultNtf.Winner = append(huResultNtf.Winner, &outer.MahjongBTEHuInfo{
				Seat:        int32(huSeat),
				HuType:      p.hu.PB(),
				HuExtraType: p.huExtra.PB(),
				HuOrder:     s.huIndex(huSeat),
			})

			s.Log().Infow("hu win score zimo",
				"room", s.room.RoomId, "hu", p.hu, "extra", p.huExtra, "shortId", p.ShortId, "winner score", p.winScore)
			return
		}
	}

	/////////////////////////////////////  以下为点炮胡的情况 /////////////////////////////////////
	lastPeerIndex := len(s.peerRecords) - 1
	peer := s.peerRecords[lastPeerIndex]
	loserSeat := peer.seat
	loser := s.mahjongPlayers[loserSeat]
	// 一炮多响，记录点炮的人
	if len(s.canHus) > 1 && s.multiHuByIndex == -1 {
		s.multiHuByIndex = loserSeat
	}

	var (
		winnerScore   = map[int]int64{}
		totalWinScore int64
	)

	// 先统计胡牌的所有人，总共要赢多少分
	for huSeat := range s.canHus {
		huPlayer := s.mahjongPlayers[huSeat]
		winScore := s.huScore(huPlayer, false)
		winnerScore[huSeat] = winScore
		totalWinScore += winScore
	}

	// 结算每个胡牌玩家
	for huSeat, winScore := range winnerScore {
		// 不允许负分，并且玩家身上的钱不够赔付总额，就把玩家身上的总分，按赔付比例分别赔付给每个胡牌人
		// 如果允许负分,或者玩家身上的分足够赔付每个胡牌的人,就直接扣
		if !s.gameParams().AllowScoreSmallZero && loser.score < totalWinScore {
			ratio := float64(winScore) / float64(totalWinScore)
			winScore = int64(float64(loser.score) * ratio)
		}

		s.AWinB(huSeat, loserSeat, winScore)
		loseScores[int32(loserSeat)] += winScore

		winner := s.mahjongPlayers[huSeat]
		winner.winScore = loseScores
		huResultNtf.Winner = append(huResultNtf.Winner, &outer.MahjongBTEHuInfo{
			Seat:        int32(huSeat),
			HuType:      winner.hu.PB(),
			HuExtraType: winner.huExtra.PB(),
			HuOrder:     s.huIndex(huSeat),
		})
		s.Log().Infow("hu win score dianPao", "room", s.room.RoomId,
			"hu", winner.hu, "extra", winner.huExtra, "shortId", winner.ShortId, "winner score", winner.winScore)
	}

	s.cardsInDesktop = s.cardsInDesktop[:len(s.cardsInDesktop)-1] // 胡成功，删除最后一张牌
}

func (s *StatePlaying) AWinB(winnerSeat, loserSeat int, score int64) {
	winner := s.mahjongPlayers[winnerSeat]
	loser := s.mahjongPlayers[loserSeat]
	winner.updateScore(score)
	loser.updateScore(-score)
	s.Log().Infow("a win b", "room", s.room.RoomId,
		"a", winner.ShortId, "a score", winner.score, "b", loser.ShortId, "b score", loser.score, "score", score)
}

// 胡牌计算得分
func (s *StatePlaying) huScore(p *mahjongPlayer, ziMo bool) int64 {
	fan := huFan[p.hu] + extraFan[p.huExtra] + int(p.huGen)
	baseScore := s.baseScore()
	if ziMo {
		if s.gameParams().ZiMoJia == 0 { // 自摸加番
			fan += 1
		} else if s.gameParams().ZiMoJia == 1 { // 自摸加底
			baseScore *= 2
		}
	}

	fan = common.Min(int(s.fanUpLimit()), fan)
	ratio := math.Pow(float64(2), float64(fan))
	winScore := s.baseScore() * int64(ratio)
	return winScore
}

// 分析是否有额外加番
func (s *StatePlaying) huExtra(seatIndex int) ExtFanType {
	var extraFans []ExtFanType

	// 根据番数大到小，优先计算大番型
	if len(s.peerRecords) == 1 && s.gameParams().TianDiHu {
		return TianHu
	}

	if len(s.peerRecords) == 2 && s.gameParams().TianDiHu {
		return Dihu
	}

	lastPeerCard := s.peerRecords[len(s.peerRecords)-1]

	// 没牌了，执行[扫底胡]和[海底炮]检测
	if s.cards.Len() == 0 {
		switch lastPeerCard.typ {
		case drawCardType:
			extraFans = append(extraFans, ShaoDiHu) // 最后一张牌，摸起来胡了，扫底胡
		case playCardType:
			extraFans = append(extraFans, HaiDiPao) // 最后一张牌，摸起来后出牌点炮了，海底炮
		}
	}

	p := s.mahjongPlayers[seatIndex]
	if p.handCards.Len() == 2 {
		extraFans = append(extraFans, JinGouGou) // 只剩2张牌做将，金钩胡
	}

	// 抢杠
	if lastPeerCard.typ == GangType1 {
		extraFans = append(extraFans, QiangGangHu) // 抢杠胡
	}

	// 如果上次是杠，这次一定是摸牌，判断是否杠上花
	if len(s.peerRecords) >= 2 {
		beforeLastPeerCard := s.peerRecords[len(s.peerRecords)-2]
		if beforeLastPeerCard.typ >= GangType1 {
			extraFans = append(extraFans, GangShangHua) // 刚上花
		}
	}

	// 如果上上次是杠，这次一定是出牌，判断是否杠上炮
	if len(s.peerRecords) >= 3 {
		beforeBeforeLastPeerCard := s.peerRecords[len(s.peerRecords)-3]
		if beforeBeforeLastPeerCard.typ >= GangType1 {
			extraFans = append(extraFans, GangShangPao) // 杠上炮
		}
	}

	return 0
}

// 算根
func (s *StatePlaying) huGen(seatIndex int) (count int) {
	p := s.mahjongPlayers[seatIndex]
	count = len(p.lightGang) + len(p.darkGang)
	for pongCard := range p.pong {
		p.handCards.Range(func(card Card) bool {
			if pongCard == card.Int32() {
				count++
				return true
			}
			return false
		})
	}

	cardsStat := p.handCards.ConvertStruct()
	for _, num := range cardsStat {
		if num == 4 {
			count++
		}
	}
	return
}
