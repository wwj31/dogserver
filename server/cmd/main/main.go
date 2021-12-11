package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"server/common/toml"
	"syscall"

	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/log"
	"github.com/wwj31/dogactor/tools"

	"server/common"
)

func main() {
	tools.Try(func() {
		exit := make(chan os.Signal)
		signal.Notify(exit, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
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
		common.InitLog(logLv, toml.Get("logpath"), appType, appId)
		// 加载配置表
		//err = config_go.Load(iniconfig.BaseString("configjson"))
		//expect.Nil(err)
		//common.RefactorConfig()

		system := run(appType, appId)
		<-exit

		system.Stop()
		<-system.CStop
	}, nil)

	log.Stop()
	fmt.Println("stop")
}

func outputFlags() (appType, tomlPath string, appId, logLv int32, err error) {
	rand.Seed(tools.Now().UnixNano()) //设置随机数种子

	flag.String("toml", "../toml", "toml file path")
	flag.String("app", "all", "app type")
	flag.Int("id", 0, "app id")
	flag.Int("log", 0, "log level, if debug log=0")

	flag.Parse()

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
