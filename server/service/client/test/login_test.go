package test

import (
	"testing"
	"time"

	"server/service/client/client"
)

func TestLogin(t *testing.T) {
	cli := &client.Client{Addr: *Addr, DeviceID: "Client7", UpShortId: 1476844}
	Init(cli)
	time.Sleep(1 * time.Second)
}
