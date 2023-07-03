package mahjong

import (
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/proto/outermsg/outer"
	"server/service/room"
)

// 换三张

type StateExchange3 struct {
	*Mahjong
}

func (s *StateExchange3) State() int {
	return Exchange3
}

func (s *StateExchange3) Enter(fsm *room.FSM) {
	s.room.Broadcast(&outer.MahjongBTEExchange3Ntf{})
	s.room.AddTimer(tools.XUID(), tools.Now().Add(15*time.Second), func(dt time.Duration) {
		s.stateEnd()
	})
	log.Infow("[Mahjong] leave state  exchange3", "room", s.room.RoomId)
}

func (s *StateExchange3) Leave(fsm *room.FSM) {
	log.Infow("[Mahjong] leave state exchange3", "room", s.room.RoomId)
}

func (s *StateExchange3) stateEnd() {
	for _, player := range s.mahjongPlayers {
		s.room.SendToPlayer(player.ShortId, &outer.MahjongBTEExchange3EndNtf{
			ShortId: 0,
			Ex3Info: nil,
			Cards:   nil,
		})
	}

	// 状态结束，给个换牌动画播放延迟，进入定缺
	s.room.AddTimer(tools.XUID(), tools.Now().Add(3*time.Second), func(dt time.Duration) {
		s.SwitchTo(DecideIgnore)
	})
}

func (s *StateExchange3) Handle(fsm *room.FSM, v any, shortId int64) (result any) {
	return nil
}
