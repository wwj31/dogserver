package player

import (
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

func (p *Player) initModule() {
	p.models[modRole] = role.New(models.New(p))         // 角色
	p.models[modMail] = mail.New(models.New(p))         // 邮件
	p.models[modAlliance] = alliance.New(models.New(p)) // 联盟
	p.models[modAgent] = agent.New(models.New(p))       // 代理
}

func (p *Player) Gamer() iface.Gamer       { return p.gamer }
func (p *Player) Role() iface.Role         { return p.models[modRole].(iface.Role) }
func (p *Player) Mail() iface.Mailer       { return p.models[modMail].(iface.Mailer) }
func (p *Player) Alliance() iface.Alliance { return p.models[modAlliance].(iface.Alliance) }
func (p *Player) Agent() iface.Agent       { return p.models[modAgent].(iface.Agent) }
