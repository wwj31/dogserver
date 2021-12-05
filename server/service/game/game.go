package game

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/iniconfig"
	"github.com/wwj31/dogactor/log"
	"server/common"
	"server/db"
	"server/service/game/handler"
	"server/service/game/iface"
	"server/service/game/logic/player"
)

func New(conf iniconfig.Config) *Game {
	return &Game{config: conf}
}

type Game struct {
	actor.Base
	common.SendTools
	*db.DB

	sid int32 // 服务器Id

	config     iniconfig.Config
	playerMgr  iface.PlayerManager
	msgHandler common.MsgHandler
}

func (s *Game) OnInit() {
	s.DB = db.New(s.config.String("mysql"), s.config.String("database"))
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

func (s *Game) PlayerMgr() iface.PlayerManager {
	return s.playerMgr
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
