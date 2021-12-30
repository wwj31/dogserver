package main

import (
	"fmt"
	"os"
	"os/signal"
	"server/common"
	"server/common/log"
	"server/common/toml"
	"server/db"
	"server/service/client"
	"server/service/game"
	"server/service/gateway"
	"server/service/login"
	"syscall"

	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/l"
	"github.com/wwj31/dogactor/tools"

	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/actor/cluster"
	"github.com/wwj31/dogactor/actor/cmd"
	"github.com/wwj31/dogactor/expect"
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
		//err = config_go.Load(iniconfig.BaseString("configjson"))
		//expect.Nil(err)
		//common.RefactorConfig()

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
	)

	switch appType {
	case common.Client:
		expect.Nil(system.Add(actor.New(common.Client, &client.Client{}, actor.SetLocalized())))
	case common.GateWay_Actor:
		newGateway(appId, system)
	case common.Login_Actor:
		newLogin(system)
	case common.Game_Actor:
		newGame(appId, system)
	case "all":
		newGateway(appId, system)
		newGame(appId, system)
		newLogin(system)
	}
	return system
}

func newLogin(system *actor.System) {
	dbIns := db.New(toml.Get("mysql"), toml.Get("database"))
	loginActor := login.New(dbIns)
	expect.Nil(system.Add(actor.New(common.Login_Actor, loginActor)))
}

func newGateway(appId int32, system *actor.System) {
	loginActor := gateway.New()
	expect.Nil(system.Add(actor.New(common.GatewayName(appId), loginActor)))
}

func newGame(appId int32, system *actor.System) {
	dbIns := db.New(toml.Get("mysql"), toml.Get("database"))
	gameActor := game.New(dbIns)
	expect.Nil(system.Add(actor.New(common.GameName(appId), gameActor)))
}
