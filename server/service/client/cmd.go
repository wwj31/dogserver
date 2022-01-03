package client

import (
	"fmt"
	"github.com/spf13/cast"
	"math/rand"
	"server/common"
	"server/proto/message"
	"time"
)

func (s *Client) InitCmd() {
	s.RegistCmd("login", s.login, "login <uid>")
	s.RegistCmd("rlogin", s.randLogin, "")

	s.RegistCmd("enter", s.enter, "enter game")
	s.RegistCmd("item", s.item, "")
}

func (s *Client) login(arg ...string) {
	if len(arg) != 1 {
		return
	}
	logReq := &message.LoginReq{
		PlatformUUID: arg[0],
		PlatformName: "test",
		OS:           "test",
	}
	checksum := common.LoginChecksum(logReq)
	logReq.Checksum = checksum
	s.SendToServer(message.MSG_LOGIN_REQ.Int32(), logReq)
}

func (s *Client) randLogin(arg ...string) {
	for {
		logReq := &message.LoginReq{
			PlatformUUID: fmt.Sprintf("%v", time.Now().UnixNano()),
			PlatformName: "test",
			OS:           "test",
		}
		s.SendToServer(message.MSG_LOGIN_REQ.Int32(), logReq)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func (s *Client) enter(arg ...string) {
	msg := &message.EnterGameReq{
		UID: s.UID,
		RID: s.RID,
	}
	s.SendToServer(message.MSG_ENTER_GAME_REQ.Int32(), msg)
}

func (s *Client) item(arg ...string) {
	if len(arg) != 2 {
		return
	}
	itemId := cast.ToInt64(arg[0])
	itemcount := cast.ToInt64(arg[1])
	msg := &message.UseItemReq{
		Items: map[int64]int64{itemId: itemcount},
	}
	s.SendToServer(message.MSG_USE_ITEM_REQ.Int32(), msg)
}
