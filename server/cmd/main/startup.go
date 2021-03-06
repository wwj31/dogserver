package main

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/actor/cluster"
	"github.com/wwj31/dogactor/actor/cmd"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/l"
	"github.com/wwj31/dogactor/tools"
	"os"
	"os/signal"
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/toml"
	"server/config/confgo"
	"server/db"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/client"
	"server/service/game"
	"server/service/game/logic/channel"
	"server/service/gateway"
	"server/service/login"
	"server/service/robot"
	"syscall"
)

func startup() {
	tools.Try(func() {
		osSignal := make(chan os.Signal)
		signal.Notify(osSignal, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		// init toml file
		toml.Init(*tomlPath, *appName, *appId)

		// init log file
		logName := *appName + cast.ToString(appId)
		log.Init(*logLevel, *logPath, logName, cast.ToBool(toml.Get("dispaly")))

		// load config of excels
		if path, ok := toml.GetB("configjson"); ok {
			err := confgo.Load(path)
			if err != nil {
				panic(err)
			}
			common.RefactorConfig()
		}

		monitor(*logPath, logName+"mo")

		system := run(*appName, int32(*appId))
		<-osSignal
		system.Stop()
		<-system.CStop
	})
	l.Close()

	fmt.Println("stop")
}

func run(appType string, appId int32) *actor.System {
	// startup the system of actor
	system, _ := actor.NewSystem(
		actor.WithCMD(cmd.New()),
		cluster.WithRemote(toml.Get("etcdaddr"), toml.Get("etcdprefix")),
		actor.Addr(toml.Get("actoraddr")),
		actor.ProtoIndex(newProtoIndex()),
	)

	switch appType {
	case actortype.Client:
		expect.Nil(system.Add(actor.New(actortype.Client, &client.Client{}, actor.SetLocalized())))
	case actortype.Robot:
		system.Add(actor.New(actortype.Robot, &robot.Robot{}, actor.SetLocalized()))
	case actortype.GateWay_Actor:
		newGateway(appId, system)
	case actortype.Login_Actor:
		newLogin(system)
	case actortype.Game_Actor:
		newGame(appId, system)
	case "all":
		newGateway(appId, system)
		newGame(appId, system)
		newLogin(system)
	}
	return system
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
		Enum2Name:   outer.MSG_name,
	})
}

func newLogin(system *actor.System) {
	dbIns := db.New(toml.Get("mysql"), toml.Get("database"), system)
	loginActor := login.New(dbIns)
	expect.Nil(system.Add(actor.New(actortype.Login_Actor, loginActor, actor.SetMailBoxSize(1000))))
}

func newGateway(appId int32, system *actor.System) {
	loginActor := gateway.New()
	expect.Nil(system.Add(actor.New(actortype.GatewayName(appId), loginActor, actor.SetMailBoxSize(2000))))
}

func newGame(appId int32, system *actor.System) {
	gameActor := game.New(uint16(appId))
	ch := channel.New()
	expect.Nil(system.Add(actor.New(actortype.GameName(appId), gameActor, actor.SetMailBoxSize(4000))))
	expect.Nil(system.Add(actor.New(actortype.ChatName(uint16(appId)), ch, actor.SetMailBoxSize(1000))))
}
