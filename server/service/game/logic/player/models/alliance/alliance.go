package alliance

import (
	gogo "github.com/gogo/protobuf/proto"

	"server/common/actortype"
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/rdsop"
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

func (s *Alliance) Data() gogo.Message {
	return &s.data
}

func (s *Alliance) OnLogin(first bool, enterGameRsp *outer.EnterGameRsp) {
	if first {
	}

	// 如果没联盟，检测是否需要加入联盟
	if s.data.AllianceId == 0 {
		// 检测上级是否有联盟，有就加入联盟
		upShortId := rdsop.AgentUp(s.Player.ShortId())
		if upShortId != 0 {
			upPlayerInfo := rdsop.PlayerInfo(upShortId)
			if upPlayerInfo.AllianceId != 0 {
				s.Player.Send(actortype.AllianceName(upPlayerInfo.AllianceId), &inner.SetMemberReq{
					Players: []*inner.PlayerInfo{s.Player.PlayerInfo()},
				})
			}
		}
	} else {
		_, err := s.Player.RequestWait(actortype.AllianceName(s.AllianceId()), &inner.OnlineNtf{
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
