package test

import (
	"flag"

	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/client/c"
)

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

var (
	addr = flag.String("addr", "ws://localhost:7001/", "addr")
	Cli  *c.Client
)

func init() {
	log.Init(-1, "", "", true)
	system, _ := actor.NewSystem(
		actor.Name("fake_client"),
		actor.ProtoIndex(newProtoIndex()),
		actor.LogFileName("", ""),
	)

	Cli = &c.Client{
		Addr:     *addr,
		DeviceID: "Client5",
	}
	_ = system.NewActor("client", Cli, actor.SetLocalized())
}
