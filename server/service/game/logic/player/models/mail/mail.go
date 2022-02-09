package mail

import (
	"server/common"
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
	"server/service/game/logic/player/models"

	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/container/rank"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/tools"
)

type Mail struct {
	models.Model

	zSet     rank.Rank
	mailInfo inner.MailInfo
}

func New(base models.Model) *Mail {
	mail := &Mail{
		Model: base,
		zSet:  *rank.New(),
	}

	if !base.Player.IsNewRole() {
		err := proto.Unmarshal(base.Player.PlayerData().MailBytes, &mail.mailInfo)
		expect.Nil(err)

		for _, m := range mail.mailInfo.Mails {
			mail.zSet.Add(cast.ToString(m.Uuid), m.CreateAt)
		}
	} else {
		mail.mailInfo.Mails = make(map[uint64]*inner.Mail, 4)
		mail.NewBuilder().
			SetMailTitle("welcome to dog game!").
			SetSender(0).
			SetContent("best wish for you !").
			SetItems(map[int64]int64{10001: 1, 10002: 10}).
			Build()
	}

	return mail
}

func (s *Mail) OnSave() {
	data, err := common.GZip(common.ProtoMarshal(&s.mailInfo))
	if err != nil {
		log.Errorw("mail zip failed", "err", err)
		return
	}
	s.Player.PlayerData().MailBytes = data
}

func (s *Mail) Add(mail *inner.Mail) {
	s.mailInfo.Mails[mail.Uuid] = mail
	s.zSet.Add(cast.ToString(mail.Uuid), mail.CreateAt)
	s.Player.Send2Client(&outer.AddMailNotify{Uuid: mail.Uuid})
	log.Debugw("add actormail ", "player", s.Player.Role().RoleId(), "actormail", mail.Title, "items", mail.Items)
}

func (s *Mail) NewBuilder() iface.MailBuilder {
	return &Builder{
		mail: &inner.Mail{
			Uuid:     s.Player.Gamer().GenUuid(),
			CreateAt: tools.NowTime(),
			Status:   0,
		},
		mailer: s,
	}
}

func (s *Mail) Mails(count, limit int32) []*inner.Mail {
	var mails []*inner.Mail
	keys := s.zSet.Get(int(count+1), int(count+limit))

	for _, k := range keys {
		mailId := cast.ToUint64(k.Key)
		mail, ok := s.mailInfo.Mails[mailId]
		if !ok {
			log.Warnw("can not found actor mail key:%v", mailId)
			continue
		}
		mails = append(mails, mail)
	}
	return mails
}

func (s *Mail) Read(uuid uint64) {
	mail, ok := s.mailInfo.Mails[uuid]
	if !ok {
		log.Warnw("Read can not find mail", "uuid", uuid, "roleId", s.Player.Role().RoleId())
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
		s.Player.Send2Client(&outer.Fail{Error: outer.ERROR_MAIL_REPEAT_RECV_ITEM})
		return
	}

	s.Player.Item().Add(mail.Items, true)
	mail.Status = 2
}

func (s *Mail) Delete(uuids ...uint64) {
	for _, uuid := range uuids {
		s.zSet.Del(cast.ToString(uuid))
		delete(s.mailInfo.Mails, uuid)
	}
}
