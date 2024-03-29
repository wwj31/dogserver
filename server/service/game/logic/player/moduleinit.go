package player

import (
	"server/proto/innermsg/inner"
	"server/service/game/iface"
	"server/service/game/logic/player/models"
	"server/service/game/logic/player/models/mail"
	"server/service/game/logic/player/models/role"
)

const (
	modRole = iota
	modMail

	allmod
)

func (s *Player) initModule() {
	s.models[modRole] = role.New(models.New(s)) // 角色
	s.models[modMail] = mail.New(models.New(s)) // 邮件
}

func (s *Player) Account() *inner.Account { return s.accountInfo }
func (s *Player) Gamer() iface.Gamer      { return s.gamer }
func (s *Player) Role() iface.Role        { return s.models[modRole].(iface.Role) }
func (s *Player) Mail() iface.Mailer      { return s.models[modMail].(iface.Mailer) }
