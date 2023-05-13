package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/wwj31/dogactor/actor/cluster/fullmesh"

	"server/mgo"

	"server/common"
	"server/common/rds"
	"server/config/conf"

	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/logger"
	"github.com/wwj31/dogactor/tools"

	"server/common/actortype"
	"server/common/log"
	"server/common/mongodb"
	"server/common/toml"
	_ "server/controller"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game"
	"server/service/gateway"
	"server/service/login"
)

func startup() {
	osQuitSignal := make(chan os.Signal)
	signal.Notify(osQuitSignal, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// init toml
	toml.Init(*tomlPath, *appName, *appId)

	// init log
	logName := *appName + cast.ToString(appId)
	log.Init(*logLevel, *logPath, logName, cast.ToBool(toml.Get("display")))

	// load config of excels
	if path, ok := toml.GetB("config_json"); ok {
		err := conf.Load(path)
		if err != nil {
			log.Errorw("toml get config json failed", "err", err)
			return
		}
		common.RefactorConfig()
	}

	// init mongo
	if err := mongodb.Builder().Addr(toml.Get("mongo_addr")).
		Database(toml.Get("database")).
		//EnableSharding().
		Connect(); err != nil {
		log.Errorw("mongo connect failed", "err", err)
		return
	}

	// init redis
	if err := rds.NewBuilder().
		Addr(toml.GetArray("redis_addr", "localhost:6379")...).
		Password(toml.Get("redis_password", "")).
		//ClusterMode().
		Connect(); err != nil {
		log.Errorw("redis connect failed", "err", err)
		return
	}

	//monitor(*logPath, logName+"mo")
	//pprof("6060")

	// startup
	system := run(*appName, int32(*appId))

	// safe quit
	<-osQuitSignal
	system.Stop()
	<-system.Stopped

	mgo.Stop()
	logger.Close()

	fmt.Println("stop")
}

func run(appType string, appId int32) *actor.System {
	system, _ := actor.NewSystem(
		fullmesh.WithRemote(toml.Get("etcd_addr"), toml.Get("etcd_prefix")),
		actor.Name(appType+cast.ToString(appId)),
		//mq.WithRemote(toml.Get("nats_url"), nats.New()),
		actor.ProtoIndex(newProtoIndex()),
		actor.LogLevel(logger.InfoLevel),
	)

	switch appType {
	case actortype.GatewayActor:
		newGateway(appId, system)
	case actortype.LoginActor:
		newLogin(system)
	case actortype.GameActor:
		newGame(appId, system)
	case "allinone":
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
		Prefix:      "Id",
		Enum2Name:   outer.Msg_name,
		Name2Enum:   outer.Msg_value,
	})
}

func newLogin(system *actor.System) {
	loginActor := login.New()
	_ = system.NewActor(actortype.LoginActor, loginActor, actor.SetMailBoxSize(2000))
}

func newGateway(appId int32, system *actor.System) {
	gateActor := gateway.New()
	_ = system.NewActor(actortype.GatewayName(appId), gateActor, actor.SetMailBoxSize(2000))
}

func newGame(appId int32, system *actor.System) {
	gameActor := game.New(appId)
	_ = system.NewActor(actortype.GameName(appId), gameActor, actor.SetMailBoxSize(1000))
}
