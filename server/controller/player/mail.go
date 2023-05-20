package player

import (
	"server/common/log"
	"server/common/router"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"
)

// 前端请求邮件列表，msg.Count 前端已有的邮件数量
var _ = router.Reg(func(player *player.Player, msg *outer.MailListReq) any {
	var msgMails []*outer.Mail
	mails := player.Mail().Mails(msg.Count, 10)

	for _, mail := range mails {
		msgMails = append(msgMails, &outer.Mail{
			Uuid:         mail.UUID,
			CreateAt:     mail.CreateAt,
			SenderRoleId: mail.SenderRID,
			Name:         mail.Name,
			Title:        mail.Title,
			Content:      mail.Content,
			Items:        mail.Items,
			Status:       mail.Status,
		})
	}

	return &outer.MailListRsp{Mails: msgMails}
})

// 已读操作
var _ = router.Reg(func(player *player.Player, msg *outer.ReadMailReq) any {
	player.Mail().Read(msg.Uuid)
	log.Debugw("player read actormail ", "player", player.ID(), "mailId", msg.Uuid)
	return &outer.ReadMailRsp{Uuid: msg.Uuid}
})

// 领取邮件道具
var _ = router.Reg(func(player *player.Player, msg *outer.ReceiveMailItemReq) any {
	player.Mail().ReceiveItem(msg.Uuid)
	log.Debugw("player recv actormail item", "player", player.ID(), "mailId", msg.Uuid)
	return &outer.ReceiveMailItemRsp{Uuid: msg.Uuid}
})

// 删除邮件
var _ = router.Reg(func(player *player.Player, msg *outer.DeleteMailReq) any {
	player.Mail().Delete(msg.Uuids...)
	log.Debugw("player del actormail", "player", player.ID(), "mailId", msg.Uuids)
	return nil
})
