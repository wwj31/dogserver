package test

import (
	"server/common/log"
	"testing"

	"github.com/stretchr/testify/assert"

	"server/proto/outermsg/outer"
	"server/service/client/client"
)

func TestSetPosition(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj1"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdSetMemberPositionReq, &outer.SetMemberPositionReq{
		ShortId:  1632144,
		Position: outer.Position_Manager,
	}).(*outer.SetMemberPositionRsp)
	assert.True(t, ok)
	log.Infof("agent members rsp [%v]\n", rsp)
}
