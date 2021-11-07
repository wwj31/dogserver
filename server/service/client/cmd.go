package client

import (
	"github.com/wwj31/dogactor/expect"
	"server/proto/message"
)

func (s *Client) InitCmd() {
	s.System().RegistCmd(s.GetID(), "login", s.login)
}

func (s *Client) login(arg ...string) {
	expect.True(len(arg) == 1)
	logReq := &message.LoginReq{
		PlatformUUID: arg[0],
		PlatformName: "test",
		OS:           "test",
	}
	s.SendToServer(message.MSG_LOGIN_REQ.Int32(), logReq)
}
