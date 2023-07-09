package mahjong

import (
	"server/proto/outermsg/outer"
)

// 碰杠胡过
func (s *StatePlaying) operate(player *mahjongPlayer, seatIndex int, op outer.ActionType) (bool, outer.ERROR) {
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

	// 所有操作者都执行完了，进入下一次摸牌，或者结束本局
	defer func() {
		if len(s.actionMap) == 0 {
			if s.cards.Len() > 0 {
				s.drawCard(s.nextSeatIndex(s.latestPlayIndex))
			} else {
				s.SwitchTo(Settlement)
			}
		}
	}()

	switch op {
	case outer.ActionType_ActionPass:
		return s.operatePass(player, seatIndex)
	case outer.ActionType_ActionPong:
		return s.operatePong(player, seatIndex)
	case outer.ActionType_ActionGang:
		return s.operateGang(player, seatIndex)
	case outer.ActionType_ActionHu:
		return s.operateHu(player, seatIndex)
	}
}

// 过牌操作
func (s *StatePlaying) operatePass(p *mahjongPlayer, seatIndex int) (bool, outer.ERROR) {
	delete(s.actionMap, seatIndex)
	return true, outer.ERROR_OK
}

// 碰牌操作
func (s *StatePlaying) operatePong(p *mahjongPlayer, seatIndex int) (bool, outer.ERROR) {

}

// 杠牌操作
func (s *StatePlaying) operateGang(p *mahjongPlayer, seatIndex int) (bool, outer.ERROR) {

}

// 胡牌操作
func (s *StatePlaying) operateHu(p *mahjongPlayer, seatIndex int) (bool, outer.ERROR) {

}
