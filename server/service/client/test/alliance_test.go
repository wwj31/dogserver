package test

import (
	"server/common/log"
	"testing"

	"github.com/stretchr/testify/assert"

	"server/proto/outermsg/outer"
	"server/service/client/client"
)

func TestInviteAlli(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj1"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdInviteAllianceReq, &outer.InviteAllianceReq{
		ShortId: 1184395,
	}).(*outer.InviteAllianceRsp)
	assert.True(t, ok)
	log.Infof("invite members rsp [%v]\n", rsp)
}

func TestSetPosition(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj1"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdSetMemberPositionReq, &outer.SetMemberPositionReq{
		ShortId:  1631196,
		Position: outer.Position_Manager,
	}).(*outer.SetMemberPositionRsp)
	assert.True(t, ok)
	log.Infof("agent members rsp [%v]\n", rsp)
}

func TestKickOut(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj1"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdKickOutMemberReq, &outer.KickOutMemberReq{ShortId: 1036478}).(*outer.KickOutMemberRsp)
	assert.True(t, ok)
	log.Infof("kick out menber [%v]\n", rsp)
}

func TestDisband(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj1"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdDisbandAllianceReq, &outer.DisbandAllianceReq{}).(*outer.DisbandAllianceRsp)
	assert.True(t, ok)
	log.Infof("kick out menber [%v]\n", rsp)
}
