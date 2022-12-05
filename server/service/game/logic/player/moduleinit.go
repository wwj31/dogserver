package player

import (
	"server/service/game/iface"
	"server/service/game/logic/player/models"
	"server/service/game/logic/player/models/chat"
	"server/service/game/logic/player/models/item"
	"server/service/game/logic/player/models/mail"
	"server/service/game/logic/player/models/role"
)

const (
	modRole = iota
	modItem
	modMail
	modChat

	allmod
)

func (s *Player) initModule() {
	s.models[modRole] = role.New(models.New(s)) // 角色
	s.models[modItem] = item.New(models.New(s)) // 道具
	s.models[modMail] = mail.New(models.New(s)) // 邮件
	s.models[modChat] = chat.New(models.New(s)) // 聊天
}

func (s *Player) Gamer() iface.Gamer { return s.gamer }
func (s *Player) Role() iface.Role   { return s.models[modRole].(iface.Role) }
func (s *Player) Item() iface.Item   { return s.models[modItem].(iface.Item) }
func (s *Player) Mail() iface.Mailer { return s.models[modMail].(iface.Mailer) }
func (s *Player) Chat() iface.Chat   { return s.models[modChat].(iface.Chat) }
