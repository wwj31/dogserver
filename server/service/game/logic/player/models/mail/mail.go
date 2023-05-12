package mail

import (
	gogo "github.com/gogo/protobuf/proto"
	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/container/rank"
	"github.com/wwj31/dogactor/tools"
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
	"server/service/game/logic/player/models"
)

type Mail struct {
	models.Model
	data inner.MailInfo

	zSet rank.Rank
}

func New(base models.Model) *Mail {
	mail := &Mail{
		Model: base,
		zSet:  *rank.New(),
	}
	mail.data.RID = base.Player.RID()

	return mail
}

func (s *Mail) OnLoaded() {
	for _, m := range s.data.Mails {
		s.zSet.Add(m.GetUUID(), m.CreateAt)
	}
}

func (s *Mail) Data() gogo.Message {
	//_, err := common.GZip(common.ProtoMarshal(&s.data))
	//if err != nil {
	//	log.Errorw("mail zip failed", "err", err)
	//	return nil
	//}
	return &s.data
}

func (s *Mail) OnLogin(first bool, enterGameRsp *outer.EnterGameRsp) {
	if first {
		s.data.Mails = make(map[string]*inner.Mail, 4)
		s.NewBuilder().
			SetMailTitle("welcome to dog game!").
			SetContent("best wish for you !").
			SetItems(map[int64]int64{10001: 1, 10002: 10}).
			Build()
	}
}

func (s *Mail) Add(mail *inner.Mail) {
	s.data.Mails[mail.GetUUID()] = mail
	s.zSet.Add(cast.ToString(mail.GetUUID()), mail.CreateAt)
	s.Player.Send2Client(&outer.AddMailNotify{Uuid: mail.GetUUID()})
	log.Debugw("add actormail ", "player", s.Player.Role().RoleId(), "actormail", mail.Title, "items", mail.Items)
}

func (s *Mail) NewBuilder() iface.MailBuilder {
	return &Builder{
		mail: &inner.Mail{
			UUID:     tools.XUID(),
			CreateAt: tools.Now().Unix(),
			Status:   0,
		},
		mailer: s,
	}
}

func (s *Mail) Mails(count, limit int32) []*inner.Mail {
	var mails []*inner.Mail
	keys := s.zSet.Get(int(count+1), int(count+limit))

	for _, k := range keys {
		mailId := k.Key
		mail, ok := s.data.Mails[mailId]
		if !ok {
			log.Warnw("can not found actor mail key:%v", mailId)
			continue
		}
		mails = append(mails, mail)
	}
	return mails
}

func (s *Mail) Read(uuid string) {
	mail, ok := s.data.Mails[uuid]
	if !ok {
		log.Warnw("Read can not find mail", "uuid", uuid, "roleId", s.Player.Role().RoleId())
		return
	}
	mail.Status = 1
}

func (s *Mail) ReceiveItem(uuid string) {
	mail, ok := s.data.Mails[uuid]
	if !ok {
		log.Warnw("ReceiveItem can not find mail", "uuid", uuid, "roleId", s.Player.Role().RoleId())
		return
	}
	if mail.Status == 2 {
		s.Player.Send2Client(&outer.FailRsp{Error: outer.ERROR_FAILED})
		return
	}

	//s.Player.Item().Add(mail.Items, true)
	mail.Status = 2
}

func (s *Mail) Delete(uuids ...string) {
	for _, uuid := range uuids {
		s.zSet.Del(cast.ToString(uuid))
		delete(s.data.Mails, uuid)
	}
}
