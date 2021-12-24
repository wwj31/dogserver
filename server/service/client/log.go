package client

import (
	"github.com/wwj31/dogactor/l"
)

var logger = l.New(l.Option{
	Level:          l.DebugLevel,
	LogPath:        "./",
	FileName:       "client.log",
	FileMaxAge:     3,
	FileMaxSize:    100,
	FileMaxBackups: 1,
	DisplayConsole: true,
})
