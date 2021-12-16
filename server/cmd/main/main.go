package main

import (
	"flag"
	"github.com/wwj31/dogactor/tools"
)

func main() {
	flag.String("toml", "../toml", "toml file path")
	flag.String("app", "all", "app type")
	flag.Int("id", 0, "app id")
	flag.Int("log", 0, "log level, if debug log=0")
	flag.Parse()

	tools.Try(func() {
		startup()
	}, nil)
}
