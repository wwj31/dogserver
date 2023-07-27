package mahjong

import (
	"time"

	"server/common"
	"server/common/actortype"
	"server/proto/innermsg/inner"
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
	s.Log().Infow("[Mahjong] enter state settlement", "room", s.room.RoomId, "master", s.masterIndex)
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

	modifyRspCount := make(map[string]struct{}) // 必须等待所有玩家金币修改成功后，才能发送结算
	// 结算分数为最终金币
	for seat, player := range s.mahjongPlayers {
		rid := player.RID
		finalScore := common.Max(0, player.score)
		s.room.Request(actortype.PlayerId(rid), &inner.ModifyGoldReq{Gold: finalScore}).Handle(func(resp any, err error) {
			modifyRspCount[rid] = struct{}{}
			if err != nil {
				s.Log().Errorw("modify gold failed kick-out player",
					"room", s.room.RoomId, "player", player.ShortId, "seat", seat, "err", err)
			} else {
				modifyRsp := resp.(*inner.ModifyGoldRsp)
				player.PlayerInfo = modifyRsp.Info
			}
			s.Log().Infow("modify gold success", "room", s.room.RoomId, "player", player.ShortId, "seat", seat, "player info", *player.PlayerInfo)
			if len(modifyRspCount) == maxNum {
				s.settlementBroadcast(settlementMsg)
			}
		})
	}
}
func (s *StateSettlement) settlementBroadcast(ntf *outer.MahjongBTESettlementNtf) {
	// 组装结算消息
	allPlayerInfo := s.playersToPB(0, true)

	// 根据每个人的杠牌，先分析出每个人扣的杠分
	darkGangMap := map[int32][]int32{}
	lightGangMap := map[int32][]int32{}
	for seat, player := range s.mahjongPlayers {
		var huPeer peerRecords
		if player.huPeerIndex != -1 {
			huPeer = s.peerRecords[player.huPeerIndex]
		}

		ntf.PlayerData = append(ntf.PlayerData, &outer.MahjongBTESettlementPlayerData{
			Player:               allPlayerInfo[seat],
			DianPaoSeatIndex:     int32(huPeer.seat),
			ByDarkGangSeatIndex:  nil,
			ByLightGangSeatIndex: nil,
			TotalFan:             0, // TODO
			TotalScore:           0, // TODO
		})

		// 本局该玩家所有的杠牌,以及每次杠成功后赔钱的位置
		for peerIndex, loserSeats := range player.gangScore {
			for _, loserSeat := range loserSeats {
				if s.peerRecords[peerIndex].typ == GangType4 {
					darkGangMap[loserSeat] = append(darkGangMap[loserSeat], int32(seat))
				} else {
					lightGangMap[loserSeat] = append(lightGangMap[loserSeat], int32(seat))
				}
			}
		}
	}

	for seat, playerData := range ntf.PlayerData {
		playerData.ByDarkGangSeatIndex = darkGangMap[int32(seat)]
		playerData.ByLightGangSeatIndex = lightGangMap[int32(seat)]
	}

	s.room.Broadcast(ntf)

	s.clear()           // 分算完清理数据
	s.nextMasterIndex() // 计算下一局庄家

	// 结算给个短暂的时间
	s.currentStateEndAt = tools.Now().Add(SettlementDuration)
	s.room.AddTimer(tools.XUID(), s.currentStateEndAt, func(dt time.Duration) {
		s.SwitchTo(Ready)
	})

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
