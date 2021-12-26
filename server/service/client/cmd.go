package client

import (
	"fmt"
	"math/rand"
	"server/proto/message"
	"time"
)

func (s *Client) InitCmd() {
	s.RegistCmd("login", s.login, "login <uid>")
	s.RegistCmd("rlogin", s.randLogin, "")

	s.RegistCmd("enter", s.enter, "enter game")
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
	enterReq := &message.EnterGameReq{
		UID: s.UID,
		RID: s.RID,
	}
	s.SendToServer(message.MSG_ENTER_GAME_REQ.Int32(), enterReq)
}
