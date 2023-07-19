package mahjong

import (
	"reflect"

	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
)

// 准备状态

type StateReady struct {
	*Mahjong
}

func (s *StateReady) State() int {
	return Ready
}

func (s *StateReady) Enter() {
	log.Infow("[Mahjong] enter state ready ", "room", s.room.RoomId)
	s.huSeat = nil
	s.mutilHuByIndex = -1
	s.room.Broadcast(&outer.MahjongBTEReadyNtf{})
}

func (s *StateReady) Leave() {
	log.Infow("[Mahjong] leave state ready", "room", s.room.RoomId)
}

func (s *StateReady) Handle(shortId int64, v any) (result any) {
	switch msg := v.(type) {
	case *inner.ReadyReq:
		if msg.Ready && s.checkAllReady() {
			// 所有人准备了，进入定庄
			if s.gameCount == 0 {
				s.SwitchTo(DecideMaster)
			} else {
				s.SwitchTo(Deal)
			}
		}
		return &inner.ReadyRsp{}
	default:
		log.Warnw("the current status has received an unknown message", "msg", reflect.TypeOf(msg).String())
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
