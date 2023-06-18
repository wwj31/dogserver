package test

import (
	"server/common/log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

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
	rsp, ok := cli.Req(outer.Msg_IdDisbandRoomReq, &outer.DisbandRoomReq{RoomId: 1}).(*outer.DisbandRoomRsp)
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
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj2"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdJoinRoomReq, &outer.JoinRoomReq{RoomId: 100001}).(*outer.JoinRoomRsp)
	assert.True(t, ok)
	log.Infof("join room rsp [%v]\n", rsp)
	time.Sleep(30 * time.Second)
}

func TestLeaveRoom(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj2"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdLeaveRoomReq, &outer.LeaveRoomReq{}).(*outer.LeaveRoomRsp)
	assert.True(t, ok)
	log.Infof("leave room rsp [%v]\n", rsp)
}

func TestReadyRoom(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj1"}
	Init(cli)
	rsp, ok := cli.Req(outer.Msg_IdReadyReq, &outer.ReadyReq{}).(*outer.ReadyRsp)
	assert.True(t, ok)
	log.Infof("leave room rsp [%v]\n", rsp)
}
