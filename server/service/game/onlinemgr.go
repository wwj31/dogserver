package game

import (
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/service/game/iface"

	"github.com/gogo/protobuf/proto"
)

type onlineMgr struct {
	game             iface.Gamer
	playerByGSession map[common.GSession]actortype.ActorId
	gSessionByPlayer map[actortype.ActorId]common.GSession
}

func newMgr(g iface.Gamer) onlineMgr {
	return onlineMgr{
		game:             g,
		playerByGSession: make(map[common.GSession]actortype.ActorId, 100),
		gSessionByPlayer: make(map[actortype.ActorId]common.GSession, 100),
	}
}

func (s *onlineMgr) AssociateSession(id actortype.ActorId, gSession common.GSession) {
	s.playerByGSession[gSession] = id
	s.gSessionByPlayer[id] = gSession
}

func (s *onlineMgr) PlayerBySession(gateSession common.GSession) (actortype.ActorId, bool) {
	p, ok := s.playerByGSession[gateSession]
	return p, ok
}
func (s *onlineMgr) GSessionByPlayer(id actortype.ActorId) (common.GSession, bool) {
	p, ok := s.gSessionByPlayer[id]
	return p, ok
}
func (s *onlineMgr) DelGSession(gateSession common.GSession) {
	id, ok := s.playerByGSession[gateSession]
	if ok {
		delete(s.gSessionByPlayer, id)
		delete(s.playerByGSession, gateSession)
	}
}

func (s *onlineMgr) RangeOnline(f func(common.GSession, actortype.ActorId)) {
	for id, s := range s.gSessionByPlayer {
		f(s, id)
	}
}

func (s *onlineMgr) BroadcastToC(msg proto.Message) {
	sender := common.NewSendTools(s.game)
	s.RangeOnline(func(gs common.GSession, id actortype.ActorId) {
		if err := sender.Send2Client(gs, msg); err != nil {
			log.Errorw("broadcast localmsg error", "err", err, "gSession", gs, "Player", id)
		}
	})
}

func (s *onlineMgr) BroadcastToActor(msg proto.Message) {
	s.RangeOnline(func(_ common.GSession, id actortype.ActorId) {
		if err := s.game.Send(id, msg); err != nil {
			log.Errorw("broadcast localmsg error", "err", err, "Player", id)
		}
	})
}
