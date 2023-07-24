package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"server/common/log"

	"server/proto/outermsg/outer"
	"server/service/client/client"
)

func TestCreateRoom(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "test1"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdCreateRoomReq, &outer.CreateRoomReq{
		GameType:   outer.GameType_Mahjong,
		GameParams: &outer.GameParams{Mahjong: &outer.MahjongParams{}},
	}).(*outer.CreateRoomRsp)
	assert.True(t, ok)
	log.Infof("create room rsp [%v]\n", rsp)
}

func TestDisbandRoom(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "test1"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdDisbandRoomReq, &outer.DisbandRoomReq{RoomId: 2}).(*outer.DisbandRoomRsp)
	assert.True(t, ok)
	log.Infof("disband room rsp [%v]\n", rsp)
}

func TestRoomList(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "test1"}
	Init(cli)
	_, ok := cli.Req(outer.Msg_IdRoomListReq, &outer.RoomListReq{}).(*outer.RoomListRsp)
	assert.True(t, ok)
}

func TestJoinRoom(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "test2"}
	Init(cli)
	_, ok := cli.Req(outer.Msg_IdJoinRoomReq, &outer.JoinRoomReq{RoomId: 1}).(*outer.JoinRoomRsp)
	assert.True(t, ok)
}

func TestLeaveRoom(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "test1"}
	Init(cli)
	_, ok := cli.Req(outer.Msg_IdLeaveRoomReq, &outer.LeaveRoomReq{}).(*outer.LeaveRoomRsp)
	assert.True(t, ok)
}
