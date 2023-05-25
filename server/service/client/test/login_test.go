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
	cli := &client.Client{Addr: *Addr, DeviceID: "test2", UpShortId: 2865755}
	Init(cli)
	time.Sleep(1 * time.Second)
}
func TestLogin3(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "test3", UpShortId: 2865755}
	Init(cli)
	time.Sleep(1 * time.Second)
}
