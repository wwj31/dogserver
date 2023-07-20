package mahjong

import (
	"reflect"
	"time"

	"server/common/log"
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
	player, _ := s.findMahjongPlayer(shortId)
	if player == nil {
		log.Warnw("player not in room", "roomId", s.room.RoomId, "shortId", shortId)
		return outer.ERROR_PLAYER_NOT_IN_ROOM
	}

	switch msg := v.(type) {
	case *outer.MahjongBTEReadyReq:
		s.room.Broadcast(&outer.MahjongBTEPlayerReadyNtf{ShortId: shortId, Ready: msg.Ready})

		player.ready = msg.Ready
		if msg.Ready {
			s.room.CancelTimer(player.RID)
			if s.checkAllReady() {
				if s.gameCount == 0 {
					s.SwitchTo(DecideMaster)
				} else {
					s.SwitchTo(Deal)
				}
			}
		} else {
			s.room.AddTimer(player.RID, time.Now().Add(ReadyExpiration), func(dt time.Duration) {
				s.room.PlayerLeave(shortId, true)
			})
		}
		return &outer.MahjongBTEReadyRsp{Ready: msg.Ready}
	default:
		log.Warnw("ready state has received an unknown message", "msg", reflect.TypeOf(msg).String())
	}
	return outer.ERROR_MAHJONG_STATE_MSG_INVALID
}

func (s *StateReady) checkAllReady() bool {
	for _, player := range s.mahjongPlayers {
		if player == nil || !player.ready {
			return false
		}
	}
	return true
}
