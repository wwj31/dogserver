package mail

import (
	"server/proto/innermsg/inner"
	"server/service/game/iface"
)

type Builder struct {
	mailer iface.Mailer
	mail   *inner.Mail
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
	if s.mail.Items == nil {
		s.mail.Items = make(map[int64]int64, len(items))
	}

	for itemId, c := range items {
		if count, ok := s.mail.Items[itemId]; ok {
			s.mail.Items[itemId] = count + c
		} else {
			s.mail.Items[itemId] = c
		}
	}
	return s
}
func (s *Builder) SetSender(RoleId string) iface.MailBuilder {
	s.mail.SenderRID = RoleId
	return s
}

func (s *Builder) Build() {
	s.mailer.Add(s.mail)
}
