package main

import (
	"fmt"
	"os"
	"os/signal"
	"server/common/redis"
	"server/db/dbmysql"
	"syscall"

	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/actor/cluster/mq"
	"github.com/wwj31/dogactor/actor/cluster/mq/nats"
	"github.com/wwj31/dogactor/expect"
	syslog "github.com/wwj31/dogactor/log"
	"github.com/wwj31/dogactor/logger"
	"github.com/wwj31/dogactor/tools"
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/toml"
	"server/config/confgo"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/client"
	"server/service/game"
	"server/service/game/logic/channel"
	"server/service/gateway"
	"server/service/login"
	"server/service/robot"
)

func startup() {
	tools.Try(func() {
		osSignal := make(chan os.Signal)
		signal.Notify(osSignal, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		syslog.SysLog.Level(logger.WarnLevel)

		// init toml file
		toml.Init(*tomlPath, *appName, *appId)

		// init log file
		logName := *appName + cast.ToString(appId)
		log.Init(*logLevel, *logPath, logName, cast.ToBool(toml.Get("dispaly")))

		// init redis
		if err := redis.Builder().ClusterMode().Addr(toml.Get("redisaddr", "localhost:9001")).
			Connect(); err != nil {
			panic(err)
		}

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
		<-system.Stopped
	})
	logger.Close()

	fmt.Println("stop")
}

func run(appType string, appId int32) *actor.System {
	// startup the system of actor
	system, _ := actor.NewSystem(
		//fullmesh.WithRemote(toml.Get("etcdaddr"), toml.Get("etcdprefix")),
		mq.WithRemote(toml.Get("natsurl"), nats.New()),
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
	dbIns := dbmysql.New(toml.Get("mysql"), toml.Get("database"), system)
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
