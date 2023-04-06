package client

import (
	"fmt"
	"math/rand"
	"time"

	"server/common"
	"server/proto/outermsg/outer"

	"github.com/spf13/cast"
)

//func (s *Client) InitCmd() {
//	s.RegistryCmd("login", s.login, "login <uid>")
//	s.RegistryCmd("rlogin", s.randLogin, "")
//
//	s.RegistryCmd("enter", s.enter, "enter game")
//
//	s.RegistryCmd("item", s.item, "")
//
//	s.RegistryCmd("listmail", s.listMail, "")
//	s.RegistryCmd("readmail", s.readMail, "")
//	s.RegistryCmd("recvmail", s.recvMail, "")
//	s.RegistryCmd("delmail", s.delMail, "")
//
//	s.RegistryCmd("chat", s.chat, "")
//}

func (s *Client) login(arg ...string) {
	if len(arg) != 1 {
		return
	}
	logReq := &outer.LoginReq{
		PlatformUUID: arg[0],
		PlatformName: "test",
		OS:           "test",
	}
	token := common.LoginToken(logReq)
	logReq.Token = token
	s.SendToServer(outer.Msg_IdLoginReq.Int32(), logReq)
}

func (s *Client) randLogin(arg ...string) {
	for {
		logReq := &outer.LoginReq{
			PlatformUUID: fmt.Sprintf("%v", time.Now().UnixNano()),
			PlatformName: "test",
			OS:           "test",
		}
		s.SendToServer(outer.Msg_IdLoginReq.Int32(), logReq)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func (s *Client) enter(arg ...string) {
	msg := &outer.EnterGameReq{
		UID:       s.UID,
		RID:       s.RID,
		NewPlayer: s.NewPlayer,
		Checksum:  common.EnterGameToken(s.UID, s.RID, s.NewPlayer),
	}
	s.SendToServer(outer.Msg_IdEnterGameReq.Int32(), msg)
}

func (s *Client) item(arg ...string) {
	if len(arg) != 2 {
		return
	}
	itemId := cast.ToInt64(arg[0])
	itemcount := cast.ToInt64(arg[1])
	msg := &outer.UseItemReq{
		Items: map[int64]int64{itemId: itemcount},
	}
	s.SendToServer(outer.Msg_IdUseItemReq.Int32(), msg)
}

func (s *Client) listMail(arg ...string) {
	if len(arg) != 1 {
		return
	}
	count := cast.ToInt32(arg[0])
	msg := &outer.MailListReq{
		Count: count,
	}
	s.SendToServer(outer.Msg_IdMailListReq.Int32(), msg)
}

func (s *Client) readMail(arg ...string) {
	if len(arg) != 1 {
		return
	}
	mailId := arg[0]
	msg := &outer.ReadMailReq{
		Uuid: mailId,
	}
	s.SendToServer(outer.Msg_IdReadMailReq.Int32(), msg)
}

func (s *Client) recvMail(arg ...string) {
	if len(arg) != 1 {
		return
	}
	mailId := arg[0]
	msg := &outer.ReceiveMailItemReq{
		Uuid: mailId,
	}
	s.SendToServer(outer.Msg_IdReceiveMailItemReq.Int32(), msg)
}

func (s *Client) delMail(arg ...string) {
	if len(arg) != 1 {
		return
	}
	mailId := arg[0]
	msg := &outer.DeleteMailReq{
		Uuids: []string{mailId},
	}
	s.SendToServer(outer.Msg_IdDeleteMailReq.Int32(), msg)
}

func (s *Client) chat(arg ...string) {
	if len(arg) != 1 {
		return
	}
	content := arg[0]
	msg := &outer.ChatReq{
		Content:     content,
		ChannelType: 1,
	}
	s.SendToServer(outer.Msg_IdChatReq.Int32(), msg)
}
