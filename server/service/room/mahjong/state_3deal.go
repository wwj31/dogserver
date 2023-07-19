package mahjong

import (
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/proto/outermsg/outer"
)

// 发牌状态

type StateDeal struct {
	*Mahjong
}

func (s *StateDeal) State() int {
	return Deal
}

func (s *StateDeal) Enter() {
	s.cards = RandomCards(nil) // 总共108张
	s.cards = Cards{
		11, 11, 11, 12, 12, 12, 13, 13, 13, 14, 14, 14, 15,
		21, 22, 23, 24, 25, 26, 27, 21, 22, 23, 24, 25, 26,
		27, 28, 29, 31, 32, 33, 34, 28, 29, 31, 32, 33, 34,
		35, 36, 37, 38, 39, 21, 22, 23, 24, 25, 26, 27, 28,
		15, 16, 38, 14, 16, 16, 15, 17, 18, 18, 18, 19, 19,
		19, 31, 32, 16, 33, 37, 18, 35, 36, 17, 34, 35, 38,
		37, 38, 39, 39, 21, 17, 36, 17, 35, 23, 33, 31, 25,
		36, 13, 26, 28, 12, 37, 32, 27, 24, 39, 19, 22, 34,
		15, 29, 29, 11}
	log.Infow("[Mahjong] enter state deal", "room", s.room.RoomId, "cards", s.cards)

	var i int
	for seatIndex, player := range s.mahjongPlayers {
		player.handCards = append(Cards{}, s.cards[i:i+13]...).Sort()
		i += 13

		s.room.SendToPlayer(player.ShortId, &outer.MahjongBTEDealNtf{
			Cards: player.handCards.ToSlice(),
		})
		log.Infow("dealing", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId, "cards", player.handCards)
	}

	// 庄家多发一张
	master := s.mahjongPlayers[s.masterIndex]
	master.handCards = master.handCards.Insert(s.cards[52])

	// 剩下的算本局牌组
	s.cards = s.cards[53:]

	// 发牌动画后，进入下个状态
	s.currentStateEndAt = tools.Now().Add(DealShowDuration)
	s.room.AddTimer(tools.XUID(), s.currentStateEndAt, func(dt time.Duration) {
		var nextState State
		if s.room.GameParams.Mahjong.HuanSanZhang == 2 {
			nextState = DecideIgnore // 不换牌，直接定缺
		} else {
			nextState = Exchange3
		}

		s.SwitchTo(nextState)
	})

	log.Infow("deal finished cards", "room", s.room.RoomId, "spare cards", s.cards)
}

func (s *StateDeal) Leave() {
	log.Infow("[Mahjong] leave state deal", "room", s.room.RoomId)
}

func (s *StateDeal) Handle(shortId int64, v any) (result any) {
	return nil
}
