package test

import (
	"server/common/log"
	"testing"

	"github.com/stretchr/testify/assert"

	"server/proto/outermsg/outer"
	"server/service/client/client"
)

func TestCreateRoom(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj1"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdCreateRoomReq, &outer.CreateRoomReq{
		GameType: outer.GameType_Mahjong,
	}).(*outer.CreateRoomRsp)
	assert.True(t, ok)
	log.Infof("create room rsp [%v]\n", rsp)
}

func TestDisbandRoom(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj1"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdDisbandRoomReq, &outer.DisbandRoomReq{Id: 100003}).(*outer.DisbandRoomRsp)
	assert.True(t, ok)
	log.Infof("disband room rsp [%v]\n", rsp)
}

func TestRoomList(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj1"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdRoomListReq, &outer.RoomListReq{}).(*outer.RoomListRsp)
	assert.True(t, ok)
	log.Infof("room list rsp [%v]\n", rsp)
}
