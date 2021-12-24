package main

import (
	"flag"
)

func main() {
	flag.String("toml", "../toml", "toml file path")
	flag.String("app", "all", "app type")
	flag.Int("id", 0, "app id")
	flag.Int("log", -1, "log level, if debug log=-1")
	flag.Parse()

	startup()
}
