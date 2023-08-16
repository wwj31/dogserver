package mahjong

import (
	"reflect"
	"server/common/log"
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
	s.multiHuByIndex = -1
	s.gameCount++
	s.playerAutoReady = s.ready

	var (
		ready          bool
		resetGameCount bool
	)
	if s.gameCount <= int(s.gameParams().PlayCountLimit) {
		ready = true
	}

	if s.gameCount > int(s.gameParams().PlayCountLimit) || resetGameCount {
		s.gameCount = 1
	}

	readyExpireAt := time.Now().Add(ReadyExpiration)
	for seat := 0; seat < maxNum; seat++ {
		player := s.mahjongPlayers[seat]
		if player == nil {
			continue
		}

		if player.Gold <= 0 {
			s.room.PlayerLeave(player.ShortId, true)
			resetGameCount = true
			continue
		}

		log.Infow("player ready", "player", player.ShortId, "ready", ready)
		s.ready(player, ready)
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
		s.ready(player, msg.Ready)
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

// 玩家准备操作，选择false到期自动准备，选择true检查是否开局
func (s *StateReady) ready(player *mahjongPlayer, r bool) {
	s.room.Broadcast(&outer.MahjongBTEPlayerReadyNtf{ShortId: player.ShortId, Ready: r})

	s.Log().Infow("the player request ready ",
		"room", s.room.RoomId, "player", player.ShortId, "ready", r, "gold", player.Gold)

	player.ready = r

	if r {
		player.readyExpireAt = time.Time{}
		s.room.CancelTimer(player.RID)
		if s.checkAllReady() {
			if s.gameCount == 1 {
				s.SwitchTo(DecideMaster)
			} else {
				s.SwitchTo(Deal)
			}
		}
	} else {
		player.readyExpireAt = time.Now().Add(ReadyExpiration)
		s.room.AddTimer(player.RID, player.readyExpireAt, func(dt time.Duration) {
			s.ready(player, true)
		})
	}
}
