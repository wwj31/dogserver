package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"server/common/log"
	"server/proto/outermsg/outer"
)

func TestAgentUpAndDown(t *testing.T) {
	rsp, ok := Cli.Req(outer.Msg_IdAgentMembersReq, &outer.AgentMembersReq{}).(*outer.AgentMembersRsp)
	assert.True(t, ok)

	log.Infof("agent members rsp [%v]", rsp)
}
