package controller

import (
	"server/common/log"
	"server/proto/message"
	"server/service/game/iface"
)

var _ = regist(MsgName(&message.MailListReq{}), func(player iface.Player, v interface{}) {
	msg := v.(*message.MailListReq)

	mails := player.Mail().Mails(msg.Count, 10)
	player.Send2Client(&message.MailListResp{Mails: mails})
	log.Debugw("player request mails ", "player", player.ID(), "mails", mails)
})

var _ = regist(MsgName(&message.ReadMailReq{}), func(player iface.Player, v interface{}) {
	msg := v.(*message.ReadMailReq)

	player.Mail().Read(msg.Uuid)
	player.Send2Client(&message.ReadMailResp{})
	log.Debugw("player read mail ", "player", player.ID(), "mailId", msg.Uuid)
})

var _ = regist(MsgName(&message.ReceiveMailItemReq{}), func(player iface.Player, v interface{}) {
	msg := v.(*message.ReceiveMailItemReq)

	player.Mail().ReceiveItem(msg.Uuid)
	player.Send2Client(&message.ReadMailResp{})
	log.Debugw("player recv mail item", "player", player.ID(), "mailId", msg.Uuid)
})

var _ = regist(MsgName(&message.DeleteMailReq{}), func(player iface.Player, v interface{}) {
	msg := v.(*message.DeleteMailReq)

	player.Mail().Delete(msg.Uuids...)
	log.Debugw("player del mail", "player", player.ID(), "mailId", msg.Uuids)
})
