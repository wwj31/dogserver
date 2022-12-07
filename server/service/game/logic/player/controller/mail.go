package controller

import (
	"server/common/log"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
)

// 前端请求邮件列表，msg.Count 前端已有的邮件数量
var _ = registry(&outer.MailListReq{}, func(player iface.Player, v interface{}) {
	msg := v.(*outer.MailListReq)

	var msgMails []*outer.Mail
	mails := player.Mail().Mails(msg.Count, 10)

	for _, mail := range mails {
		msgMails = append(msgMails, &outer.Mail{
			Uuid:         mail.UUID,
			CreateAt:     mail.CreateAt,
			SenderRoleId: mail.SenderRoleId,
			Name:         mail.Name,
			Title:        mail.Title,
			Content:      mail.Content,
			Items:        mail.Items,
			Status:       mail.Status,
		})
	}
	player.Send2Client(&outer.MailListResp{Mails: msgMails})
	log.Debugw("player request mails ", "player", player.ID(), "mails", mails)
})

// 已读操作
var _ = registry(&outer.ReadMailReq{}, func(player iface.Player, v interface{}) {
	msg := v.(*outer.ReadMailReq)

	player.Mail().Read(msg.Uuid)
	player.Send2Client(&outer.ReadMailResp{Uuid: msg.Uuid})
	log.Debugw("player read actormail ", "player", player.ID(), "mailId", msg.Uuid)
})

// 领取邮件道具
var _ = registry(&outer.ReceiveMailItemReq{}, func(player iface.Player, v interface{}) {
	msg := v.(*outer.ReceiveMailItemReq)

	player.Mail().ReceiveItem(msg.Uuid)
	player.Send2Client(&outer.ReceiveMailItemResp{Uuid: msg.Uuid})
	log.Debugw("player recv actormail item", "player", player.ID(), "mailId", msg.Uuid)
})

// 删除邮件
var _ = registry(&outer.DeleteMailReq{}, func(player iface.Player, v interface{}) {
	msg := v.(*outer.DeleteMailReq)

	player.Mail().Delete(msg.Uuids...)
	log.Debugw("player del actormail", "player", player.ID(), "mailId", msg.Uuids)
})
