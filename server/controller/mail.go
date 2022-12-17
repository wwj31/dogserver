package controller

import (
	"server/common/log"
	"server/common/router"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"
)

// 前端请求邮件列表，msg.Count 前端已有的邮件数量
var _ = router.Reg(func(player *player.Player, msg *outer.MailListReq) {
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
var _ = router.Reg(func(player *player.Player, msg *outer.ReadMailReq) {
	player.Mail().Read(msg.Uuid)
	player.Send2Client(&outer.ReadMailResp{Uuid: msg.Uuid})
	log.Debugw("player read actormail ", "player", player.ID(), "mailId", msg.Uuid)
})

// 领取邮件道具
var _ = router.Reg(func(player *player.Player, msg *outer.ReceiveMailItemReq) {
	player.Mail().ReceiveItem(msg.Uuid)
	player.Send2Client(&outer.ReceiveMailItemResp{Uuid: msg.Uuid})
	log.Debugw("player recv actormail item", "player", player.ID(), "mailId", msg.Uuid)
})

// 删除邮件
var _ = router.Reg(func(player *player.Player, msg *outer.DeleteMailReq) {
	player.Mail().Delete(msg.Uuids...)
	log.Debugw("player del actormail", "player", player.ID(), "mailId", msg.Uuids)
})
