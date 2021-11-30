package common

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/expect"
	"reflect"
)

type Handle func(sourceId string, gSession GSession, msg interface{}) proto.Message

type MsgHandler struct {
	handleFunc map[string]Handle
}

func NewMsgHandler() MsgHandler {
	return MsgHandler{
		handleFunc: make(map[string]Handle),
	}
}

func (s *MsgHandler) Reg(msg proto.Message, h Handle) {
	name := reflect.TypeOf(msg).String()
	_, exist := s.handleFunc[name]
	expect.True(!exist)

	s.handleFunc[name] = h
}

func (s *MsgHandler) Handle(sourceId string, gSession GSession, msg proto.Message) proto.Message {
	name := reflect.TypeOf(msg).String()
	handle, ok := s.handleFunc[name]
	expect.True(ok)
	return handle(sourceId, gSession, msg)
}
