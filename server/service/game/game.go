package game

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/iniconfig"
	"github.com/wwj31/dogactor/log"
	"server/common"
	"server/service/game/handler"
	"server/service/game/iface"
	"server/service/game/player"
)

func New(conf iniconfig.Config) *Game {
	return &Game{config: conf}
}

type Game struct {
	actor.Base
	common.SendTools
	config     iniconfig.Config
	sid        int32 // 服务器Id
	playerMgr  iface.Manager
	msgHandler common.MsgHandler
}

func (s *Game) OnInit() {
	s.SendTools = common.NewSendTools(s)
	s.playerMgr = player.NewMgr(s)
	s.msgHandler = common.NewMsgHandler()

	// handler模块初始化
	handler.Init(s)
	log.Debug("game OnInit")
}

// 区服id
func (s *Game) SID() int32 {
	return s.sid
}

// 注册消息
func (s *Game) RegistMsg(msg proto.Message, handle common.Handle) {
	s.msgHandler.Reg(msg, handle)
}

func (s *Game) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	actMsg, gSession, err := common.UnwrapperGateMsg(msg)
	expect.Nil(err)

	pbMsg, ok := actMsg.(proto.Message)
	expect.True(ok)

	rsp := s.msgHandler.Handle(sourceId, gSession, pbMsg)
	if rsp != nil {
		var err error
		if gSession.Valid() {
			err = s.Send2Client(gSession, rsp)
		} else {
			err = s.Send(sourceId, rsp)
		}
		if err != nil {
			log.Error(err.Error())
		}
	}
}

func (s *Game) OnHandleRequest(sourceId, targetId, requestId string, msg interface{}) (respErr error) {
	pbMsg, ok := msg.(proto.Message)
	expect.True(ok)

	rsp := s.msgHandler.Handle(sourceId, "", pbMsg)
	if rsp != nil {
		return s.Response(requestId, msg)
	}
	return nil
}
