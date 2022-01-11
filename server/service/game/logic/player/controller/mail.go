package controller

import (
	"server/common/log"
	"server/proto/message"
	"server/service/game/iface"
)

// 前端请求邮件列表，msg.Count 前端已有的邮件数量
var _ = regist(MsgName(&message.MailListReq{}), func(player iface.Player, v interface{}) {
	msg := v.(*message.MailListReq)

	mails := player.Mail().Mails(msg.Count, 10)
	player.Send2Client(&message.MailListResp{Mails: mails})
	log.Debugw("player request mails ", "player", player.ID(), "mails", mails)
})

// 已读操作
var _ = regist(MsgName(&message.ReadMailReq{}), func(player iface.Player, v interface{}) {
	msg := v.(*message.ReadMailReq)

	player.Mail().Read(msg.Uuid)
	player.Send2Client(&message.ReadMailResp{})
	log.Debugw("player read mail ", "player", player.ID(), "mailId", msg.Uuid)
})

// 领取邮件道具
var _ = regist(MsgName(&message.ReceiveMailItemReq{}), func(player iface.Player, v interface{}) {
	msg := v.(*message.ReceiveMailItemReq)

	player.Mail().ReceiveItem(msg.Uuid)
	player.Send2Client(&message.ReadMailResp{})
	log.Debugw("player recv mail item", "player", player.ID(), "mailId", msg.Uuid)
})

// 删除邮件
var _ = regist(MsgName(&message.DeleteMailReq{}), func(player iface.Player, v interface{}) {
	msg := v.(*message.DeleteMailReq)

	player.Mail().Delete(msg.Uuids...)
	log.Debugw("player del mail", "player", player.ID(), "mailId", msg.Uuids)
})
