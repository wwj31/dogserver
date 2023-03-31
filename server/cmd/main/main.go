package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	tomlPath = flag.String("toml", "../toml", "toml file path")
	appName  = flag.String("app", "all", "app type")
	appId    = flag.Int("id", 1, "app id")
	logLevel = flag.Int("log", -1, "log level, if debug log=-1")
	logPath  = flag.String("logpath", "./", "path of log file")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Usage = func() { fmt.Println("flag param error") }
	flag.Parse()

	startup()
}
