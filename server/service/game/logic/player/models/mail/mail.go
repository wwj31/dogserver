package mail

import (
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/container/rank"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/tools"
	"server/common/log"
	"server/db/table"
	"server/proto/message"
	"server/service/game/iface"
	"server/service/game/logic/player/models"
)

type Mail struct {
	models.Model

	zSet     *rank.Rank
	mailInfo message.MailInfo
}

func New(rid uint64, base models.Model) *Mail {
	mail := &Mail{
		Model:    base,
		zSet:     rank.New(),
		mailInfo: message.MailInfo{Mails: make(map[uint64]*message.Mail)},
	}

	if !base.Player.IsNewRole() {
		tMail := table.Mail{UUId: rid}
		err := base.Player.Gamer().Load(&tMail)
		expect.Nil(err)

		err = proto.Unmarshal(tMail.Bytes, &mail.mailInfo)
		expect.Nil(err)

		for _, m := range mail.mailInfo.Mails {
			mail.zSet.Add(cast.ToString(m.Uuid), m.CreateAt)
		}
	}

	return mail
}

func (s *Mail) Add(mail *message.Mail) {
	s.mailInfo.Mails[mail.Uuid] = mail
	s.zSet.Add(cast.ToString(mail.Uuid), mail.CreateAt)
}

func (s *Mail) NewBuilder() iface.MailBuilder {
	return &Builder{
		mail: &message.Mail{
			Uuid:     s.Player.Gamer().GenUuid(),
			CreateAt: tools.NowTime(),
			Status:   0,
		},
		mailer: s,
	}
}

func (s *Mail) Mails(count, limit int32) []*message.Mail {
	var (
		arr   []int
		mails []*message.Mail
	)
	for c := count; c < count+limit; c++ {
		arr = append(arr, int(c))
	}
	keys := s.zSet.Get(arr...)

	for _, key := range keys {
		mails = append(mails, s.mailInfo.Mails[cast.ToUint64(key)])
	}
	return mails
}

func (s *Mail) Read(uuid uint64) {
	mail, ok := s.mailInfo.Mails[uuid]
	if !ok {
		log.Warnw("read mail can not find mail", "uuid", uuid, "roleId", s.Player.Role().RoleId())
		return
	}
	mail.Status = 1
}

func (s *Mail) ReceiveItem(uuid uint64) {
	mail, ok := s.mailInfo.Mails[uuid]
	if !ok {
		log.Warnw("ReceiveItem can not find mail", "uuid", uuid, "roleId", s.Player.Role().RoleId())
		return
	}

	s.Player.Item().Add(mail.Items, true)
	mail.Status = 2
}

func (s *Mail) Delete(uuid uint64) {
	delete(s.mailInfo.Mails, uuid)
}
