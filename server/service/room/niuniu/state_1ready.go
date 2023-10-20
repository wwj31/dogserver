package niuniu

import (
	"reflect"
	"time"

	"server/proto/outermsg/outer"
)

// 准备状态

type StateReady struct {
	*NiuNiu
}

func (s *StateReady) State() int {
	return Ready
}

func (s *StateReady) Enter() {
	s.Log().Infow("[NiuNiu] enter state ready ", "room", s.room.RoomId)

	readyExpireAt := time.Now().Add(ReadyExpiration)
	s.room.Broadcast(&outer.NiuNiuReadyNtf{ReadyExpireAt: readyExpireAt.UnixMilli()})
}

func (s *StateReady) Leave() {
	s.Log().Infow("[NiuNiu] leave state ready", "room", s.room.RoomId)
}

func (s *StateReady) Handle(shortId int64, v any) (result any) {
	player, _ := s.findNiuNiuPlayer(shortId)
	if player == nil {
		s.Log().Warnw("player not in room", "roomId", s.room.RoomId, "shortId", shortId)
		return outer.ERROR_PLAYER_NOT_IN_ROOM
	}

	switch msg := v.(type) {
	case *outer.NiuNiuReadyReq:
		s.ready(player, msg.Ready)
		return &outer.NiuNiuReadyRsp{Ready: msg.Ready}
	default:
		s.Log().Warnw("ready state has received an unknown message", "msg", reflect.TypeOf(msg).String())
	}
	return outer.ERROR_NIUNIU_STATE_MSG_INVALID
}

func (s *StateReady) checkAllReady() bool {
	for _, player := range s.niuniuPlayers {
		if player == nil || !player.ready {
			return false
		}
	}
	return true
}

// 玩家准备操作，选择false到期自动准备，选择true检查是否开局
func (s *StateReady) ready(player *niuniuPlayer, r bool) {
	s.room.Broadcast(&outer.NiuNiuPlayerReadyNtf{ShortId: player.ShortId, Ready: r})

	s.Log().Infow("the player request ready ",
		"room", s.room.RoomId, "player", player.ShortId, "ready", r, "gold", player.Gold)

	player.ready = r

	if r {
		player.readyExpireAt = time.Time{}
		s.room.CancelTimer(player.RID)
		if s.checkAllReady() {
			s.SwitchTo(Deal)
		}
	} else {
		player.readyExpireAt = time.Now().Add(ReadyExpiration)
		s.readyAfterTimeout(player, player.readyExpireAt)
	}
}
