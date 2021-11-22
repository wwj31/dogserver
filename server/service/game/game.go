package game

import (
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/iniconfig"
	"server/service/game/iface"
	"server/service/game/player"
)

type Game struct {
	actor.Base
	config    iniconfig.Config
	sid       int32 // 服务器Id
	playerMgr iface.PlayerMgr
}

func (s *Game) OnInit() {
	s.playerMgr = player.NewMgr(s)
}

func (s *Game) SID() int32 {
	return s.sid
}
