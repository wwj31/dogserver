package main

import (
	"flag"
	"fmt"
)

var (
	tomlPath = flag.String("toml", "../toml", "toml file path")
	appName  = flag.String("app", "all", "app type")
	appId    = flag.Int("id", 0, "app id")
	logLevel = flag.Int("log", -1, "log level, if debug log=-1")
	logPath  = flag.String("logpath", "./", "path of log file")
)

func main() {
	flag.Usage = func() { fmt.Println("flag param error") }
	flag.Parse()

	// start the world
	startup()
}
