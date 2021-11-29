package handler

import (
	"server/proto/message"
	"server/service/game/iface"
)

type Controller struct {
	iface.Gamer
}

func Init(g iface.Gamer) {
	handler := &Controller{
		Gamer: g,
	}
	g.RegistMsg((*message.EnterGameReq)(nil), handler.EnterGameReq)
}
