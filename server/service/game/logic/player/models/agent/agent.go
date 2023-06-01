package agent

import (
	gogo "github.com/gogo/protobuf/proto"

	"server/proto/outermsg/outer"
	"server/service/game/logic/player/models"
)

type Agent struct {
	models.Model
}

func New(base models.Model) *Agent {
	mod := &Agent{Model: base}
	return mod
}

func (a *Agent) Data() gogo.Message {
	return nil
}

func (a *Agent) OnLogin(first bool, enterGameRsp *outer.EnterGameRsp) {
	if first {
	}

}

func (a *Agent) OnLogout() {

}
