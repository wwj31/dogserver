package fasterrun

import (
	"reflect"
	"time"

	"server/proto/outermsg/outer"
)

// 准备状态

type StateReady struct {
	*FasterRun
}

func (s *StateReady) State() int {
	return Ready
}

func (s *StateReady) Enter() {
	s.Log().Infow("[FasterRun] enter state ready ", "room", s.room.RoomId)
	s.gameCount++
	s.playerAutoReady = s.ready

	var (
		autoReady      bool
		resetPlayCount bool
	)

	if s.gameCount > int(s.gameParams().PlayCountLimit) {
		resetPlayCount = true
	}

	// 先把金币<=0的玩家踢出去
	for seat := 0; seat < s.playerNumber(); seat++ {
		player := s.fasterRunPlayers[seat]
		if player == nil {
			continue
		}

		if player.Gold <= 0 {
			// 新一轮游戏，或者不允许负分，都需要踢出玩家
			if resetPlayCount || !s.gameParams().AllowScoreSmallZero {
				s.Log().Infow("kick player with ready case gold <= 0", "shortId", player.ShortId, "gold", player.Gold)
				s.room.PlayerLeave(player.ShortId, true)
			}
		}

		// 大结算后，需要踢出托管的玩家
		if player.trusteeship && resetPlayCount {
			s.Log().Infow("kick player with ready case trusteeship", "shortId", player.ShortId)
			s.room.PlayerLeave(player.ShortId, true)
		}
	}

	// 判断游戏是否需要重置
	if resetPlayCount {
		s.Log().Infow("reset game count", "current", s.gameCount, "param", s.gameParams().PlayCountLimit, "reset", resetPlayCount)
		s.gameCount = 1
	} else {
		autoReady = true // 不需要重置，就自动准备
	}

	// 设置每个玩家的准备状态
	for seat := 0; seat < s.playerNumber(); seat++ {
		player := s.fasterRunPlayers[seat]
		if player == nil {
			continue
		}
		s.Log().Infow("player ready", "player", player.ShortId, "ready", autoReady, "gold", player.Gold)
		s.ready(player, autoReady)
	}

	// 只有重置后的第一局才需要自动准备
	if s.gameCount == 1 {
		readyExpireAt := time.Now().Add(ReadyExpiration)
		s.room.Broadcast(&outer.FasterRunReadyNtf{ReadyExpireAt: readyExpireAt.UnixMilli()})
	}
}

func (s *StateReady) Leave() {
	s.masterRebate()
	s.Log().Infow("[FasterRun] leave state ready", "room", s.room.RoomId)
}

func (s *StateReady) Handle(shortId int64, v any) (result any) {
	player, _ := s.findFasterRunPlayer(shortId)
	if player == nil {
		s.Log().Warnw("player not in room", "roomId", s.room.RoomId, "shortId", shortId)
		return outer.ERROR_PLAYER_NOT_IN_ROOM
	}

	switch msg := v.(type) {
	case *outer.FasterRunReadyReq:
		s.ready(player, msg.Ready)
		return &outer.FasterRunReadyRsp{Ready: msg.Ready}
	default:
		s.Log().Warnw("ready state has received an unknown message", "msg", reflect.TypeOf(msg).String())
	}
	return outer.ERROR_FASTERRUN_STATE_MSG_INVALID
}

func (s *StateReady) checkAllReady() bool {
	for _, player := range s.fasterRunPlayers {
		if player == nil || !player.ready {
			return false
		}
	}
	return true
}

// 玩家准备操作，选择false到期自动准备，选择true检查是否开局
func (s *StateReady) ready(player *fasterRunPlayer, r bool) {
	s.room.Broadcast(&outer.FasterRunPlayerReadyNtf{ShortId: player.ShortId, Ready: r})

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
