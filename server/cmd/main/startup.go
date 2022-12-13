package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/actor/cluster/mq"
	"github.com/wwj31/dogactor/actor/cluster/mq/nats"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/logger"
	"github.com/wwj31/dogactor/tools"
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/mongodb"
	"server/common/redis"
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

		// init toml
		toml.Init(*tomlPath, *appName, *appId)

		// init log
		logName := *appName + cast.ToString(appId)
		log.Init(*logLevel, *logPath, logName, cast.ToBool(toml.Get("dispaly")))

		// init mongo
		if err := mongodb.Builder().Addr(toml.Get("mongoaddr")).
			Database(toml.Get("database")).Connect(); err != nil {
			panic(err)
		}

		// init redis
		if err := redis.Builder().
			Addr(toml.GetArray("redisaddr", "localhost:6379")...).
			ClusterMode().Connect(); err != nil {
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
		//actor.Addr(toml.Get("actoraddr")),
		mq.WithRemote(toml.Get("natsurl"), nats.New()),
		actor.ProtoIndex(newProtoIndex()),
		actor.LogLevel(logger.InfoLevel),
	)

	switch appType {
	case actortype.Client:
		expect.Nil(system.Add(actor.New(actortype.Client, &client.Client{ACC: "Client"}, actor.SetLocalized())))
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
	loginActor := login.New()
	expect.Nil(system.Add(actor.New(actortype.Login_Actor, loginActor, actor.SetMailBoxSize(1000))))
}

func newGateway(appId int32, system *actor.System) {
	loginActor := gateway.New()
	expect.Nil(system.Add(actor.New(actortype.GatewayName(appId), loginActor, actor.SetMailBoxSize(2000))))
}

func newGame(appId int32, system *actor.System) {
	gameActor := game.New(appId)
	ch := channel.New()
	expect.Nil(system.Add(actor.New(actortype.GameName(appId), gameActor, actor.SetMailBoxSize(4000))))
	expect.Nil(system.Add(actor.New(actortype.ChatName(appId), ch, actor.SetMailBoxSize(1000))))
}
