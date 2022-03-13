package models

import (
	"server/db/table"
	"server/service/game/iface"
)

type Model struct {
	Player iface.Player
}

func New(player iface.Player) Model {
	model := Model{
		Player: player,
	}
	return model
}

// OnSave 功能模块回存时触发回调，将对应功能数据放入data中
func (s *Model) OnSave(data *table.Player) {}

// OnLogin 玩家登录触发
func (s *Model) OnLogin() {}

// OnLogout 玩家离线触发
func (s *Model) OnLogout() {}
