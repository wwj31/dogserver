package player

import (
	"server/db/table"
	"server/service/game/iface"
	"server/service/game/logic/player/models"
	"server/service/game/logic/player/models/chat"
	"server/service/game/logic/player/models/item"
	"server/service/game/logic/player/models/mail"
	"server/service/game/logic/player/models/role"
)

func (s *Player) initModule(data *table.Player) {
	s.models[modRole] = role.New(models.New(s), data.RoleBytes) // 角色
	s.models[modItem] = item.New(models.New(s), data.ItemBytes) // 道具
	s.models[modMail] = mail.New(models.New(s), data.MailBytes) // 邮件
	s.models[modChat] = chat.New(models.New(s))                 // 聊天
}

func (s *Player) Gamer() iface.Gamer { return s.gamer }
func (s *Player) Role() iface.Role   { return s.models[modRole].(iface.Role) }
func (s *Player) Item() iface.Item   { return s.models[modItem].(iface.Item) }
func (s *Player) Mail() iface.Mailer { return s.models[modMail].(iface.Mailer) }
func (s *Player) Chat() iface.Chat   { return s.models[modChat].(iface.Chat) }
