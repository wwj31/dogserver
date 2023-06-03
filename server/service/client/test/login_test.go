package test

import (
	"testing"
	"time"

	"server/service/client/client"
)

func TestLogin(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj1"}
	Init(cli)
	time.Sleep(1 * time.Second)
}

func TestLogin2(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj2", UpShortId: 1532126}
	Init(cli)
	time.Sleep(1 * time.Hour)
}

func TestLogin3(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj3", UpShortId: 1492924}
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
