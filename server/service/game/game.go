package game

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/iniconfig"
	"github.com/wwj31/dogactor/log"

	"server/common"
	"server/proto/inner_message"
	"server/service/game/iface"
	"server/service/game/player"
)

type Game struct {
	actor.Base
	config     iniconfig.Config
	sid        int32 // 服务器Id
	playerMgr  iface.Manager
	msgHandler *common.MsgHandler
}

func (s *Game) OnInit() {
	s.playerMgr = player.NewMgr(s)
	s.msgHandler = common.NewMsgHandler()
}

func (s *Game) SID() int32 {
	return s.sid
}

func (s *Game) RegistMsg(msg proto.Message, handle common.Handle) {
	s.msgHandler.Reg(msg, handle)
}

func (s *Game) SendToPlayer(gSession common.GSession, pb proto.Message) {
	actorId, _ := gSession.SplitGateSession()
	err := s.Send(actorId, inner_message.NewGateWrapperByPb(pb, gSession))
	if err != nil {
		log.Error(err.Error())
	}
}

func (s *Game) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	actMsg, gSession, err := inner_message.UnwrapperGateMsg(msg)
	expect.Nil(err)

	pbMsg, ok := actMsg.(proto.Message)
	expect.True(ok)

	rsp := s.msgHandler.Handle(sourceId, gSession, pbMsg)
	if rsp != nil {
		if gSession.Valid() {
			s.SendToPlayer(gSession, rsp)
		} else {
			s.Send(sourceId, rsp)
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
