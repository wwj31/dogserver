package main

import (
	"flag"
	"math/rand"
	"time"
)

var (
	tomlPath   = flag.String("toml", "../toml", "toml file path")
	appName    = flag.String("app", "all", "app type")
	appId      = flag.Int("id", 1, "app id")
	logLevel   = flag.Int("log", -1, "log level, debug log=-1")
	logPath    = flag.String("log_path", "", "path of log file")
	sysLogPath = flag.String("syslog_path", "", "path of sys log file")
	version    = flag.String("v", "0.0.1", "the version of server")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()

	startup()
}
