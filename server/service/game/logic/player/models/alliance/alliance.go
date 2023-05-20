package alliance

import (
	gogo "github.com/gogo/protobuf/proto"
	"server/common/actortype"
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player/models"
)

type Alliance struct {
	models.Model
	data inner.AllianceInfo
}

func New(base models.Model) *Alliance {
	mod := &Alliance{Model: base}
	mod.data.RID = base.Player.RID()
	return mod
}
func (s *Alliance) AllianceId() int32 {
	return s.data.AllianceId
}
func (s *Alliance) SetAllianceId(id int32) {
	s.data.AllianceId = id
	s.Player.UpdateInfoToRedis()
}

func (s *Alliance) Data() gogo.Message {
	return &s.data
}

func (s *Alliance) OnLogin(first bool, enterGameRsp *outer.EnterGameRsp) {
	if first {
	}

	if s.data.AllianceId != 0 {
		err := s.Player.Send(actortype.AllianceName(s.AllianceId()), &inner.OnlineNtf{
			GateSession: s.Player.GateSession().String(),
			RID:         s.Player.RID(),
		})

		if err != nil {
			log.Warnf("alliance login send failed", "rid", s.Player.RID(), "err", err)
		}
	}
}

func (s *Alliance) OnLogout() {
	if s.data.AllianceId != 0 {
		err := s.Player.Send(actortype.AllianceName(s.AllianceId()), &inner.OfflineNtf{
			GateSession: s.Player.GateSession().String(),
			RID:         s.Player.RID(),
		})

		if err != nil {
			log.Warnf("alliance logout send failed", "rid", s.Player.RID(), "err", err)
		}
	}
}
