package game

import (
	"server/common"
	"server/proto/inner_message/inner"
	"server/proto/message"
	"server/service/game/iface"
	"server/service/game/logic/player"
	"server/service/game/logic/player/localmsg"

	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"server/common/log"
)

func New(s iface.SaveLoader) *Game {
	return &Game{
		SaveLoader: s,
	}
}

type Game struct {
	actor.Base
	common.SendTools
	iface.SaveLoader

	sid int32 // 服务器Id

	playerMgr iface.PlayerManager
}

func (s *Game) OnInit() {
	s.SendTools = common.NewSendTools(s)
	s.playerMgr = player.NewMgr(s)

	log.Debugf("game OnInit")
}

// 区服id
func (s *Game) OnStop() bool {
	log.Infof("stop game")
	return true
}

// 区服id
func (s *Game) SID() int32 {
	return s.sid
}

func (s *Game) PlayerMgr() iface.PlayerManager {
	return s.playerMgr
}

func (s *Game) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	actMsg, gSession, err := common.UnwrapperGateMsg(msg)
	expect.Nil(err)

	switch pbMsg := actMsg.(type) {
	case *message.EnterGameReq:
		s.EnterGameReq(gSession, pbMsg)
	case *inner.GT2GSessionClosed:
		s.Logout(pbMsg)
	}
}

// 玩家请求进入游戏
func (s *Game) EnterGameReq(gSession common.GSession, msg *message.EnterGameReq) {
	log.Debugw("EnterGameReq", "msg", msg)

	// 重复登录
	if _, ok := s.PlayerMgr().PlayerBySession(gSession); ok {
		log.Warnw("player repeated enter game", "gSession", gSession, "localmsg", msg.RID)
		return
	}

	// todo .. 解密
	playerId := common.PlayerId(msg.RID)
	// 重连删除旧session,否则创建Player Actor
	if oldSession, ok := s.PlayerMgr().GSessionByPlayer(playerId); ok {
		s.PlayerMgr().DelGSession(oldSession)
	} else {
		playerActor := actor.New(playerId, player.New(msg.RID), actor.SetMailBoxSize(200))
		err := s.System().Add(playerActor)
		if err != nil {
			log.Errorw("regist player actor error", "rid", msg.RID, "err", err)
			return
		}
	}

	err := s.Send(playerId, localmsg.Login{GSession: gSession})
	if err != nil {
		log.Errorw("login send error", "rid", msg.RID, "err", err, "playerId", playerId)
		return
	}

	s.PlayerMgr().SetPlayer(gSession, playerId)
}

// 玩家离线
func (s *Game) Logout(msg *inner.GT2GSessionClosed) proto.Message {
	gSession := common.GSession(msg.GateSession)
	playerId, ok := s.PlayerMgr().PlayerBySession(gSession)
	if ok {
		s.Send(playerId, msg)
	}

	s.PlayerMgr().DelGSession(gSession)
	return nil
}
