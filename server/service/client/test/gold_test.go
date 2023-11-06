package test

import (
	"github.com/stretchr/testify/assert"
	"server/common/log"
	"server/proto/outermsg/outer"
	"server/service/client/client"
	"testing"
)

func TestGoldRecord(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "test1", Test: true}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdGoldRecordsReq, &outer.GoldRecordsReq{StartIndex: 0, EndIndex: 19}).(*outer.GoldRecordsRsp)
	assert.True(t, ok)
	log.Infof("DisbandRoomRsp [%v]\n", rsp)
}