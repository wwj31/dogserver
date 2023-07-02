package mahjong

import (
	"server/common/log"
	"server/proto/outermsg/outer"
	"server/service/room"
)

// 定缺状态

type StateDecideIgnore struct {
	*Mahjong
}

func (s *StateDecideIgnore) State() int {
	return DecideIgnore
}

func (s *StateDecideIgnore) Enter(fsm *room.FSM) {

	log.Infow("Mahjong enter exchange3", "room", s.room.RoomId)
}

func (s *StateDecideIgnore) Leave(fsm *room.FSM) {
	log.Infow("Mahjong leave exchange3", "room", s.room.RoomId)
}

func (s *StateDecideIgnore) Handle(fsm *room.FSM, v any, shortId int64) (result any) {
	switch msg := v.(type) {
	case *outer.MahjongBTEDecideIgnoreReq:
		if msg.Color == 0 {
			return &outer.FailRsp{Error: outer.ERROR_MSG_REQ_PARAM_INVALID}
		}
		player := s.findMahjongPlayer(shortId)
		if player == nil {
			return &outer.FailRsp{Error: outer.ERROR_ROOM_PLAYER_NOT_IN_GAME}
		}

		player.ignoreColor = ColorType(msg.Color)
		return &outer.MahjongBTEDecideIgnoreRsp{}
	}
	return nil
}
