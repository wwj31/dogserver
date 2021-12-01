package player

import "server/service/game/iface"

type playerBuildProcess struct {
	player *Player
}

func NewBuildProcess() playerBuildProcess {
	return playerBuildProcess{player: &Player{}}
}

func (s playerBuildProcess) Player() *Player {
	return s.player
}

func (s playerBuildProcess) SetGame(g iface.Gamer) playerBuildProcess {
	s.player.game = g
	return s
}

func (s playerBuildProcess) SetRole(r iface.Role) playerBuildProcess {
	s.player.Role = r
	return s
}
