package mahjong

import (
	"reflect"
	"time"

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
	s.Log().Infow("[Mahjong] enter state ready ", "room", s.room.RoomId)
	s.huSeat = nil
	s.mutilHuByIndex = -1
	readyExpireAt := time.Now().Add(ReadyExpiration)
	for _, player := range s.mahjongPlayers {
		if player != nil {
			player.readyExpireAt = readyExpireAt
			s.readyTimeout(player.RID, player.ShortId, player.readyExpireAt)
		}
	}

	s.room.Broadcast(&outer.MahjongBTEReadyNtf{ReadyExpireAt: readyExpireAt.UnixMilli()})
}

func (s *StateReady) Leave() {
	s.Log().Infow("[Mahjong] leave state ready", "room", s.room.RoomId)
}

func (s *StateReady) Handle(shortId int64, v any) (result any) {
	player, _ := s.findMahjongPlayer(shortId)
	if player == nil {
		s.Log().Warnw("player not in room", "roomId", s.room.RoomId, "shortId", shortId)
		return outer.ERROR_PLAYER_NOT_IN_ROOM
	}

	switch msg := v.(type) {
	case *outer.MahjongBTEReadyReq:
		s.room.Broadcast(&outer.MahjongBTEPlayerReadyNtf{ShortId: shortId, Ready: msg.Ready})

		player.ready = msg.Ready
		if msg.Ready {
			player.readyExpireAt = time.Time{}

			s.room.CancelTimer(player.RID)
			if s.checkAllReady() {
				if s.gameCount == 0 {
					s.SwitchTo(DecideMaster)
				} else {
					s.SwitchTo(Deal)
				}
			}
		} else {
			player.readyExpireAt = time.Now().Add(ReadyExpiration)
			s.readyTimeout(player.RID, player.ShortId, player.readyExpireAt)
		}
		return &outer.MahjongBTEReadyRsp{Ready: msg.Ready}
	default:
		s.Log().Warnw("ready state has received an unknown message", "msg", reflect.TypeOf(msg).String())
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
