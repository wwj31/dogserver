package mahjong

import (
	"time"

	"server/proto/outermsg/outer"

	"github.com/wwj31/dogactor/tools"
)

// 结算状态

type StateSettlement struct {
	*Mahjong
}

func (s *StateSettlement) State() int {
	return Settlement
}

func (s *StateSettlement) Enter() {
	s.gameCount++
	notHu := s.isNoHu()
	settlementMsg := &outer.MahjongBTESettlementNtf{
		NotHu:            notHu,
		GameCount:        int32(s.gameCount),
		GameSettlementAt: tools.Now().UnixMilli(),
		HuSeatIndex:      s.huSeat,
	}

	if notHu {
		// 流局，筛选出有叫、没叫、花猪的玩家
		var (
			hasTingSeat, hasNotTingSeat []int
			pig                         = make(map[int]struct{})
		)

		for seat, player := range s.mahjongPlayers {
			tingCards, err := player.handCards.ting(player.ignoreColor,
				player.lightGang,
				player.darkGang,
				player.pong,
				s.gameParams(),
			)

			if err != nil {
				s.Log().Errorw("player ting error",
					"room", s.room.RoomId, "player", player.ShortId, "seat", seat, "hand", player.handCards, "err", err)
				continue
			}

			if len(tingCards) > 0 {
				hasTingSeat = append(hasTingSeat, seat)
			} else {
				hasNotTingSeat = append(hasNotTingSeat, seat)
				if player.handCards.HasColorCard(player.ignoreColor) {
					pig[seat] = struct{}{} // 花猪
				}
			}
		}
	} else {

	}

	allPlayerInfo := s.playersToPB(0, true)
	for seat, player := range s.mahjongPlayers {
		var peer peerCard
		if player.huPeerIndex != -1 {
			peer = s.peerCards[player.huPeerIndex]
		}

		settlementMsg.PlayerData = append(settlementMsg.PlayerData, &outer.MahjongBTESettlementPlayerData{
			Player:               allPlayerInfo[seat],
			DianPaoSeatIndex:     int32(peer.seat),
			ByDarkGangSeatIndex:  nil, // TODO
			ByLightGangSeatIndex: nil, // TODO
			TotalFan:             0,   // TODO
			TotalScore:           0,   // TODO
		})
	}
	s.room.Broadcast(settlementMsg)

	s.clear()           // 分算完清理数据
	s.nextMasterIndex() // 计算下一局庄家

	// 结算给个短暂的时间
	s.currentStateEndAt = tools.Now().Add(SettlementDuration)
	s.room.AddTimer(tools.XUID(), s.currentStateEndAt, func(dt time.Duration) {
		s.SwitchTo(Ready)
	})

	s.Log().Infow("[Mahjong] enter state settlement",
		"room", s.room.RoomId, "master", s.masterIndex)
}

func (s *StateSettlement) Leave() {
	s.Log().Infow("[Mahjong] leave state settlement", "room", s.room.RoomId)
}

func (s *StateSettlement) Handle(shortId int64, v any) (result any) {
	return outer.ERROR_MAHJONG_STATE_MSG_INVALID
}

func (s *StateSettlement) nextMasterIndex() {
	if s.mutilHuByIndex != -1 {
		s.masterIndex = s.mutilHuByIndex
	} else if len(s.huSeat) > 0 {
		s.masterIndex = int(s.huSeat[0])
	} else {
		s.masterIndex = s.nextSeatIndexWithoutHu(s.masterIndex)
	}
}

// 检查是否流局
func (s *StateSettlement) isNoHu() bool {
	for _, player := range s.mahjongPlayers {
		if player.hu != HuInvalid {
			return false
		}
	}
	return true
}
