package client

import (
	"fmt"
	"math/rand"
	"server/common"
	"server/proto/outermsg/outer"
	"time"

	"github.com/spf13/cast"
)

func (s *Client) InitCmd() {
	s.RegistCmd("login", s.login, "login <uid>")
	s.RegistCmd("rlogin", s.randLogin, "")

	s.RegistCmd("enter", s.enter, "enter game")

	s.RegistCmd("item", s.item, "")

	s.RegistCmd("listmail", s.listMail, "")
	s.RegistCmd("readmail", s.readMail, "")
	s.RegistCmd("recvmail", s.recvMail, "")
	s.RegistCmd("delmail", s.delMail, "")
}

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
	s.SendToServer(outer.MSG_LOGIN_REQ.Int32(), logReq)
}

func (s *Client) randLogin(arg ...string) {
	for {
		logReq := &outer.LoginReq{
			PlatformUUID: fmt.Sprintf("%v", time.Now().UnixNano()),
			PlatformName: "test",
			OS:           "test",
		}
		s.SendToServer(outer.MSG_LOGIN_REQ.Int32(), logReq)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func (s *Client) enter(arg ...string) {
	msg := &outer.EnterGameReq{
		UID:       s.UID,
		RID:       s.RID,
		NewPlayer: s.NewPlayer,
		Checksum:  common.LoginMD5(s.UID, s.RID, s.NewPlayer),
	}
	s.SendToServer(outer.MSG_ENTER_GAME_REQ.Int32(), msg)
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
	s.SendToServer(outer.MSG_USE_ITEM_REQ.Int32(), msg)
}

func (s *Client) listMail(arg ...string) {
	if len(arg) != 1 {
		return
	}
	count := cast.ToInt32(arg[0])
	msg := &outer.MailListReq{
		Count: count,
	}
	s.SendToServer(outer.MSG_MAIL_LIST_REQ.Int32(), msg)
}

func (s *Client) readMail(arg ...string) {
	if len(arg) != 1 {
		return
	}
	mailId := cast.ToUint64(arg[0])
	msg := &outer.ReadMailReq{
		Uuid: mailId,
	}
	s.SendToServer(outer.MSG_READ_MAIL_REQ.Int32(), msg)
}

func (s *Client) recvMail(arg ...string) {
	if len(arg) != 1 {
		return
	}
	mailId := cast.ToUint64(arg[0])
	msg := &outer.ReceiveMailItemReq{
		Uuid: mailId,
	}
	s.SendToServer(outer.MSG_RECEIVE_MAIL_ITEM_REQ.Int32(), msg)
}

func (s *Client) delMail(arg ...string) {
	if len(arg) != 1 {
		return
	}
	mailId := cast.ToUint64(arg[0])
	msg := &outer.DeleteMailReq{
		Uuids: []uint64{mailId},
	}
	s.SendToServer(outer.MSG_DELETE_MAIL_REQ.Int32(), msg)
}
