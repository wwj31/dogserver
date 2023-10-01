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

func TestSetScoreForDownReq(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj1"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdSetScoreForDownReq, &outer.SetScoreForDownReq{
		ShortId: 1147959,
		Gold:    -100,
	}).(*outer.SetScoreForDownRsp)
	assert.True(t, ok)
	log.Infof("agent members rsp [%v]\n", rsp)
}

func TestSetAgentDownRebateReq(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "test1"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdSetAgentDownRebateReq, &outer.SetAgentDownRebateReq{
		ShortId: 1612475,
		Rebate:  10,
	}).(*outer.SetAgentDownRebateRsp)
	assert.True(t, ok)
	log.Infof("agent SetAgentDownRebate rsp [%v]\n", rsp)
}
