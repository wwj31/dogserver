package main

import (
	"flag"
	"fmt"
	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/l"
	"github.com/wwj31/dogactor/tools"
	"math/rand"
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

	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/actor/cluster"
	"github.com/wwj31/dogactor/actor/cmd"
	"github.com/wwj31/dogactor/expect"
)

func startup() {
	tools.Try(func() {
		rand.Seed(tools.Now().UnixNano()) //设置随机数种子
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		rand.Seed(tools.Now().UnixNano())

		// 获取程序启动参数
		appType, tomlPath, appId, logLv, err := outputFlags()
		if err != nil {
			fmt.Println("initConf error", err)
			return
		}

		// 初始化toml配置
		toml.Init(tomlPath, appType, appId)

		// 初始化日志
		log.Init(logLv, "./",
			appType+cast.ToString(appId),
			cast.ToBool(toml.Get("dispaly")),
		)

		// 加载配置表
		//err = config_go.Load(iniconfig.BaseString("configjson"))
		//expect.Nil(err)
		//common.RefactorConfig()

		system := run(appType, appId)
		<-c
		system.Stop()
		<-system.CStop
	})
	l.Close()

	fmt.Println("stop")
}

func outputFlags() (appType, tomlPath string, appId, logLv int32, err error) {
	appType = flag.Lookup("app").Value.String()
	appId, err = cast.ToInt32E(flag.Lookup("id").Value.String())
	if err != nil {
		return
	}

	tomlPath = flag.Lookup("toml").Value.String()

	logLv, err = cast.ToInt32E(flag.Lookup("log").Value.String())
	if err != nil {
		return
	}

	return
}

func run(appType string, appId int32) *actor.System {
	// 启动actor服务
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
	case "All":
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
