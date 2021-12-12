package main

import (
	"server/common"
	"server/common/toml"
	"server/db"
	"server/service/client"
	"server/service/game"
	"server/service/gateway"
	"server/service/login"

	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/actor/cluster"
	"github.com/wwj31/dogactor/actor/cmd"
	"github.com/wwj31/dogactor/expect"
)

func run(appType string, appId int32) *actor.System {
	// 启动actor服务
	system, _ := actor.NewSystem(
		actor.WithCMD(cmd.New()),
		cluster.WithRemote(toml.Get("etcdaddr"), toml.Get("etcdprefix")),

		actor.Addr(toml.Get("actoraddr")),
	)

	switch appType {
	case common.Client:
		expect.Nil(system.Regist(actor.New(common.Client, &client.Client{}, actor.SetLocalized())))
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
	expect.Nil(system.Regist(actor.New(common.Login_Actor, loginActor)))
}

func newGateway(appId int32, system *actor.System) {
	loginActor := gateway.New()
	expect.Nil(system.Regist(actor.New(common.GatewayName(appId), loginActor)))
}

func newGame(appId int32, system *actor.System) {
	dbIns := db.New(toml.Get("mysql"), toml.Get("database"))
	gameActor := game.New(dbIns)
	expect.Nil(system.Regist(actor.New(common.GameName(appId), gameActor)))
}
