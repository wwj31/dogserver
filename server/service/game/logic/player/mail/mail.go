package mail

import (
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/container/rank"
	"github.com/wwj31/dogactor/expect"
	"server/db/table"
	"server/proto/message"
	"server/service/game/logic/model"
)

type Mail struct {
	model.Model

	zSet     *rank.Rank
	mailInfo message.MailInfo
}

func New(rid uint64, base model.Model) *Mail {
	mail := &Mail{
		Model:    base,
		zSet:     rank.New(),
		mailInfo: message.MailInfo{Mails: make(map[uint64]*message.Mail)},
	}

	if !base.Player.IsNewRole() {
		tMail := table.Mail{UUId: rid}
		err := base.Player.Load(&tMail)
		expect.Nil(err)

		err = proto.Unmarshal(tMail.Bytes, &mail.mailInfo)
		expect.Nil(err)

		for _, m := range mail.mailInfo.Mails {
			mail.zSet.Add(cast.ToString(m.Uuid), rank.Score(m.CreateAt))
		}
	}

	return mail
}

func (s *Mail) Add() {

}
