package agent

import (
	gogo "github.com/gogo/protobuf/proto"

	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player/models"
)

type agent struct {
	models.Model
	data inner.AllianceInfo
}

func New(base models.Model) *agent {
	mod := &agent{Model: base}
	mod.data.RID = base.Player.RID()
	return mod
}

func (s *agent) Data() gogo.Message {
	return &s.data
}

func (s *agent) OnLogin(first bool, enterGameRsp *outer.EnterGameRsp) {
	if first {
	}

}

func (s *agent) OnLogout() {

}
