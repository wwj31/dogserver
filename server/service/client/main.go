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

var addr = flag.String("addr", "ws://localhost:7001/", "addr")

func main() {
	flag.Usage = func() { fmt.Println("flag param error") }
	flag.Parse()

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
		Addr:     *addr,
		DeviceID: "wwj3",
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
