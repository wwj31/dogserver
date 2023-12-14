package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/logger"
	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/client/client"
)

// -addr ws://localhost:7001/
var (
	addr      = flag.String("addr", "localhost:7001/", "链接地址")
	reconnect = flag.Int64("recon", -1, "重连时间(毫秒)")
	device    = flag.String("device", "test1", "设备号")
	help      = flag.String("help", " ", "帮助")
)

func main() {
	flag.Usage = func() {
		flag.PrintDefaults()
		fmt.Println("例子: ./cli -addr=ws://1.14.17.15:7001/ -recon=3000 -device=test1")
	}
	flag.Parse()
	if *help == "" {
		flag.Usage()
		return
	}

	rand.Seed(time.Now().UnixNano())

	osQuitSignal := make(chan os.Signal)
	signal.Notify(osQuitSignal, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	log.Init(-1, "", "", true)
	system, _ := actor.NewSystem(
		actor.Name("fake_client"),
		actor.ProtoIndex(newProtoIndex()),
		actor.LogLevel(logger.InfoLevel),
	)

	_ = system.NewActor("client", &client.Client{
		Addr:      *addr,
		Reconnect: *reconnect,
		DeviceID:  *device,
	}, actor.SetLocalized())

	// safe quit
	<-osQuitSignal
	system.Stop()
	<-system.Stopped
}

func newProtoIndex() *tools.ProtoIndex {
	return tools.NewProtoIndex(func(name string) (v interface{}, ok bool) {
		v, ok = inner.Spawner(name)
		if !ok {
			v, ok = outer.Spawner(name, true)
		}
		return
	}, tools.EnumIdx{
		PackageName: "outer",
		Prefix:      "Id",
		Enum2Name:   outer.Msg_name,
		Name2Enum:   outer.Msg_value,
	})
}
