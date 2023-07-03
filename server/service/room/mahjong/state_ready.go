package mahjong

import (
	"server/common/log"
	"server/proto/outermsg/outer"
	"server/service/room"
)

// 准备状态

type StateReady struct {
	*Mahjong
}

func (s *StateReady) State() int {
	return Ready
}

func (s *StateReady) Enter(fsm *room.FSM) {
	log.Infow("[Mahjong] leave state  ready ", "room", s.room.RoomId)
	s.room.Broadcast(&outer.MahjongBTEReadyNtf{})
}

func (s *StateReady) Leave(fsm *room.FSM) {
	log.Infow("[Mahjong] leave state ready", "room", s.room.RoomId)
}

func (s *StateReady) Handle(fsm *room.FSM, v any, shortId int64) (result any) {
	switch msg := v.(type) {
	case *outer.ReadyReq:
		if msg.Ready && s.checkAllReady() {
			// 所有人准备了，进入定庄
			if s.gameCount == 0 {
				s.SwitchTo(DecideMaster)
			} else {
				s.SwitchTo(Deal)
			}
		}
		return &outer.ReadyRsp{Ready: msg.Ready}
	}
	return nil
}

func (s *StateReady) checkAllReady() bool {
	for _, player := range s.mahjongPlayers {
		if player == nil || !player.Ready {
			return false
		}
	}
	return true
}
