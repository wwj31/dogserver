package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"server/common/log"
	"server/proto/outermsg/outer"
	"server/service/client/client"
)

func TestMembersReq(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "Client5"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdAgentMembersReq, &outer.AgentMembersReq{}).(*outer.AgentMembersRsp)
	assert.True(t, ok)

	log.Infof("agent members rsp [%v]", rsp)
}
