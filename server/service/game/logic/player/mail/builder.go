package mail

import (
	"server/proto/message"
	"server/service/game/iface"
)

type Builder struct {
	mailer iface.Mailer
	mail   *message.Mail
}

func (s *Builder) SetMailTitle(title string) iface.MailBuilder {
	s.mail.Title = title
	return s
}
func (s *Builder) SetContent(content string) iface.MailBuilder {
	s.mail.Content = content
	return s
}
func (s *Builder) SetItems(items map[int64]int64) iface.MailBuilder {
	s.mail.Items = items
	return s
}
func (s *Builder) SetSender(RoleId uint64) iface.MailBuilder {
	s.mail.SenderRoleId = RoleId
	return s
}

func (s *Builder) Build() {
	s.mailer.Add(s.mail)
}