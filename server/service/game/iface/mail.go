package iface

import (
	"server/proto/outermsg/outer"
)

type Mailer interface {
	Modeler

	Add(mail *outer.Mail)
	NewBuilder() MailBuilder

	Mails(count, limit int32) []*outer.Mail

	Read(uuid uint64)
	ReceiveItem(uuid uint64)
	Delete(uuid ...uint64)
}

type MailBuilder interface {
	SetMailTitle(title string) MailBuilder
	SetContent(content string) MailBuilder
	SetItems(items map[int64]int64) MailBuilder
	SetSender(sender uint64) MailBuilder
	Build()
}
