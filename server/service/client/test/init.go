package test

import (
	"flag"

	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/client/client"
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
	Addr = flag.String("Addr", "ws://localhost:7001/", "Addr")
	//Addr = flag.String("Addr", "ws://1.14.17.15:7001/", "Addr")
)

func Init(cli *client.Client) {
	cli.Reconnect = -1
	log.Init(-1, "", "", true)
	system, _ := actor.NewSystem(
		actor.Name("fake_client"),
		actor.ProtoIndex(newProtoIndex()),
		actor.LogFileName("", ""),
	)

	_ = system.NewActor("client", cli, actor.SetLocalized())
}
