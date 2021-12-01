package handler

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/log"
	"server/common"
	"server/proto/message"
	"server/service/game/player"
	"server/service/game/player/role"
)

// 玩家请求进入游戏
func (s *Controller) EnterGameReq(sourceId string, gSession common.GSession, pbMsg interface{}) proto.Message {
	msg := pbMsg.(*message.EnterGameReq)
	log.KV("msg", msg).Debug("EnterGameReq")

	_player, exist := s.PlayerMgr().PlayerByRID(msg.RID)
	if !exist {
		_player = player.
			NewBuildProcess().
			SetGame(s).
			SetRole(role.New(msg.RID, s.Gamer)).
			Player()
	}

	return &message.EnterGameRsp{
		UID:  _player.Table().UUId,
		RID:  _player.Table().RoleId,
		Name: _player.Table().Name,
	}
}
