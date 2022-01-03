package mail

import (
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/container/rank"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/tools"
	"server/common"
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

	if base.Player.IsNewRole() {
		mail.NewBuilder().
			SetMailTitle("welcome to dog game!").
			SetSender(0).
			SetContent("best wish for you !").
			SetItems(map[int64]int64{10001: 1, 10002: 10}).
			Build()
	} else {
		tMail := table.Mail{RoleId: rid}
		err := base.Player.Gamer().Load(&tMail)
		expect.Nil(err)
		tMail.Bytes, _ = common.UnGZip(tMail.Bytes)

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
	s.Player.Send2Client(&message.AddMailNotify{Uuid: mail.Uuid})
	log.Debugw("add mail ", "player", s.Player.Role().RoleId(), "mail", mail.Title, "items", mail.Items)
	s.save()
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
	var mails []*message.Mail
	keys := s.zSet.Get(int(count+1), int(count+limit))

	for _, key := range keys {
		mailId := cast.ToUint64(key.Key)
		mail, ok := s.mailInfo.Mails[mailId]
		if !ok {
			log.Warnw("can not found mail key:%v", mailId)
			continue
		}
		mails = append(mails, mail)
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
	if mail.Status == 2 {
		s.Player.Send2Client(&message.Fail{Error: message.ERROR_MAIL_REPEAT_RECV_ITEM})
		return
	}

	s.Player.Item().Add(mail.Items, true)
	mail.Status = 2
	s.save()
}

func (s *Mail) Delete(uuids ...uint64) {
	for _, uuid := range uuids {
		s.zSet.Del(cast.ToString(uuid))
		delete(s.mailInfo.Mails, uuid)
	}
	s.save()
}

func (s *Mail) save() {
	s.SetTable(&table.Mail{
		RoleId: s.Player.Role().RoleId(),
		Bytes:  s.marshal(&s.mailInfo),
	})
}

func (s *Mail) marshal(msg proto.Message) []byte {
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Errorw("proto marshal error", "err", err)
	}
	compress, _ := common.GZip(data)
	return compress
}
