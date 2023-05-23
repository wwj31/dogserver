package player

import (
	"server/proto/innermsg/inner"
	"server/service/game/iface"
	"server/service/game/logic/player/models"
	"server/service/game/logic/player/models/agent"
	"server/service/game/logic/player/models/alliance"
	"server/service/game/logic/player/models/mail"
	"server/service/game/logic/player/models/role"
)

const (
	modRole = iota
	modMail
	modAlliance
	modAgent

	allMod
)

func (s *Player) initModule() {
	s.models[modRole] = role.New(models.New(s))         // 角色
	s.models[modMail] = mail.New(models.New(s))         // 邮件
	s.models[modAlliance] = alliance.New(models.New(s)) // 联盟
	s.models[modAgent] = agent.New(models.New(s))       // 代理
}

func (s *Player) Account() *inner.Account  { return s.accountInfo }
func (s *Player) Gamer() iface.Gamer       { return s.gamer }
func (s *Player) Role() iface.Role         { return s.models[modRole].(iface.Role) }
func (s *Player) Mail() iface.Mailer       { return s.models[modMail].(iface.Mailer) }
func (s *Player) Alliance() iface.Alliance { return s.models[modAlliance].(iface.Alliance) }
func (s *Player) Agent() iface.Agent       { return s.models[modAgent].(iface.Agent) }
