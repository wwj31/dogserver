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
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj2", UpShortId: 1797839}
	Init(cli)
	time.Sleep(1 * time.Second)
}
func TestLogin3(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "wwj3", UpShortId: 1797839}
	Init(cli)
	time.Sleep(1 * time.Second)
}
