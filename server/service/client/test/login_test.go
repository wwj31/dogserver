package test

import (
	"testing"
	"time"

	"server/service/client/client"
)

func TestLogin(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "test1"}
	Init(cli)
	time.Sleep(1 * time.Second)
}

func TestLogin2(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "test2"}
	Init(cli)
	time.Sleep(1 * time.Second)
}

func TestLogin3(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "test2"}
	Init(cli)
	time.Sleep(1 * time.Second)
}

func TestLogin4(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj4", UpShortId: 1036478}
	Init(cli)
	time.Sleep(1 * time.Second)
}

func TestLogin5(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj5", UpShortId: 1476742}
	Init(cli)
	time.Sleep(1 * time.Second)
}
