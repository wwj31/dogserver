package test

import (
	"testing"

	"server/common/log"

	"github.com/stretchr/testify/assert"

	"server/proto/outermsg/outer"
	"server/service/client/client"
)

func TestInviteAlli(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj1"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdInviteAllianceReq, &outer.InviteAllianceReq{
		ShortId: 1517575,
	}).(*outer.InviteAllianceRsp)
	assert.True(t, ok)
	log.Infof("invite members rsp [%v]\n", rsp)
}

func TestSetPosition(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj1"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdSetMemberPositionReq, &outer.SetMemberPositionReq{
		ShortId:  1670667,
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

func TestSearchPlayer(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj1"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdSearchPlayerInfoReq, &outer.SearchPlayerInfoReq{
		ShortId: 1784645,
	}).(*outer.SearchPlayerInfoRsp)
	assert.True(t, ok)
	log.Infof("search player rsp [%v]\n", rsp)
}
