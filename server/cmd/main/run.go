package main

import (
	"fmt"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/actor/cluster"
	"github.com/wwj31/dogactor/actor/cmd"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/iniconfig"
	"server/common"
	"server/db"
	"server/service/client"
	"server/service/gateway"
	"server/service/login"
)

func run(appType string, appId int32) *actor.System {
	// 服务通用配置
	conf, conferr := iniconfig.NewAppConf(appType, appId)
	if conferr != nil {
		fmt.Println("NewAppConf error", conferr)
		return nil
	}
	etcdAddr, etcdPrefix := iniconfig.BaseString("etcd_addr"), iniconfig.BaseString("etcd_prefix")
	// 启动actor服务
	system, _ := actor.NewSystem(actor.WithCMD(cmd.New()), cluster.WithRemote(etcdAddr, etcdPrefix), actor.Addr(conf.String("actor_addr")))

	switch appType {
	case common.Client:
		expect.Nil(system.Regist(actor.New(common.Client, &client.Client{}, actor.SetLocalized())))
	case common.GateWay_Actor:
		newGateway(appId, conf, system)
	case common.Login_Actor:
		newLogin(conf, system)
	case "All":
		newGateway(appId, conf, system)
		newLogin(conf, system)
	}
	return system
}

func newLogin(conf iniconfig.Config, system *actor.System) {
	dbIns := db.New(conf.String("mysql"), conf.String("database"))
	loginActor := login.New(dbIns, conf)
	expect.Nil(system.Regist(actor.New(common.Login_Actor, loginActor)))
}

func newGateway(appId int32, conf iniconfig.Config, system *actor.System) {
	loginActor := gateway.New(conf)
	expect.Nil(system.Regist(actor.New(common.GatewayName(appId), loginActor)))
}
