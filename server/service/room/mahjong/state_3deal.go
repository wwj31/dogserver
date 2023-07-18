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
	s.cards = RandomCards() // 总共108张
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

	s.cards = s.cards[52:]

	// 发牌动画后，进入下个状态
	s.room.AddTimer(tools.XUID(), tools.Now().Add(DealShowDuration), func(dt time.Duration) {
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
