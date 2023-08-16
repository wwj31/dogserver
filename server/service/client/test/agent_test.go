package test

import (
	"testing"

	"server/common/log"

	"github.com/stretchr/testify/assert"

	"server/proto/outermsg/outer"
	"server/service/client/client"
)

func TestMembersReq(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "test1"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdAgentMembersReq, &outer.AgentMembersReq{}).(*outer.AgentMembersRsp)
	assert.True(t, ok)
	log.Infof("agent members rsp [%v]\n", rsp)

	//cli = &client.Client{Addr: *Addr, DeviceID: "wwj2"}
	//Init(cli)
	//rsp, ok = cli.Req(outer.Msg_IdAgentMembersReq, &outer.AgentMembersReq{}).(*outer.AgentMembersRsp)
	//assert.True(t, ok)
	//log.Infof("agent members rsp [%v]\n", rsp)
	//
	//cli = &client.Client{Addr: *Addr, DeviceID: "wwj3"}
	//Init(cli)
	//rsp, ok = cli.Req(outer.Msg_IdAgentMembersReq, &outer.AgentMembersReq{}).(*outer.AgentMembersRsp)
	//assert.True(t, ok)
	//log.Infof("agent members rsp [%v]\n", rsp)
	//
	//cli = &client.Client{Addr: *Addr, DeviceID: "wwj4"}
	//Init(cli)
	//rsp, ok = cli.Req(outer.Msg_IdAgentMembersReq, &outer.AgentMembersReq{}).(*outer.AgentMembersRsp)
	//assert.True(t, ok)
	//log.Infof("agent members rsp [%v]\n", rsp)
	//
	//cli = &client.Client{Addr: *Addr, DeviceID: "wwj5"}
	//Init(cli)
	//rsp, ok = cli.Req(outer.Msg_IdAgentMembersReq, &outer.AgentMembersReq{}).(*outer.AgentMembersRsp)
	//assert.True(t, ok)
	//log.Infof("agent members rsp [%v]\n", rsp)
}
