package alliance

import (
	"context"

	"server/common"

	gogo "github.com/gogo/protobuf/proto"
	"github.com/spf13/cast"

	"server/common/actortype"
	"server/common/log"
	"server/common/rds"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/rdsop"
	"server/service/alliance"
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
		upShortId := rdsop.AgentUp(s.Player.Role().ShortId())
		if upShortId != 0 {
			upPlayerInfo := rdsop.PlayerInfo(upShortId)
			if upPlayerInfo.AllianceId != 0 {
				result, err := s.Player.RequestWait(actortype.AllianceName(upPlayerInfo.AllianceId), &inner.AddMemberReq{
					Player: s.Player.PlayerInfo(),
					Ntf:    false, // 自己请求加入联盟，不需要额外通知
				})
				if yes, code := common.IsErr(result, err); yes {
					log.Warnf("player request join alliance failed ", "rid", s.Player.RID(),
						"upShortId", upPlayerInfo,
						"alliance", upPlayerInfo.AllianceId,
						"code", code.String())
					return
				}

				_ = result.(*inner.AddMemberRsp)
				s.data.AllianceId = upPlayerInfo.AllianceId
				s.data.Position = 1
			}
		} else {
			// 如果上级没有联盟，再检测离线期间是否被设为盟主
			joinAllianceKey := rdsop.JoinAllianceKey(s.Player.Role().ShortId())
			allianceId, _ := rds.Ins.Get(context.Background(), joinAllianceKey).Result()
			if allianceId != "" {
				s.data.AllianceId = cast.ToInt32(allianceId)
				s.data.Position = alliance.Master.Int32()
				rds.Ins.Del(context.Background(), joinAllianceKey)
			}
		}
	} else {
		if rdsop.IsAllianceDeleted(s.AllianceId()) {
			s.data.Position = 0
			s.data.AllianceId = 0
			return
		}
	}

	if s.data.AllianceId != 0 {
		result, err := s.Player.RequestWait(actortype.AllianceName(s.AllianceId()), &inner.MemberInfoOnLoginReq{
			GateSession: s.Player.GateSession().String(),
			RID:         s.Player.RID(),
		})

		if err != nil {
			log.Warnf("alliance login send failed", "rid", s.Player.RID(), "err", err)
			return
		}

		memberInfoRsp, ok := result.(*inner.MemberInfoOnLoginRsp)
		if !ok {
			return
		}

		s.data.Position = memberInfoRsp.Position
		s.data.AllianceId = memberInfoRsp.AllianceId
	}

	enterGameRsp.RoleInfo.AllianceId = s.data.AllianceId
	enterGameRsp.RoleInfo.Position = outer.Position(s.data.Position)
}

func (s *Alliance) OnLogout() {
	if s.data.AllianceId != 0 {
		err := s.Player.Send(actortype.AllianceName(s.AllianceId()), &inner.MemberInfoOnLogoutReq{
			GateSession: s.Player.GateSession().String(),
			RID:         s.Player.RID(),
		})

		if err != nil {
			log.Warnf("alliance logout send failed", "rid", s.Player.RID(), "err", err)
		}
	}
}
