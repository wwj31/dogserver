package client

import (
	log "github.com/wwj31/dogactor/logger"
)

var logger = log.New(log.Option{
	Level:          log.DebugLevel,
	LogPath:        "./",
	FileName:       "client.log",
	FileMaxAge:     3,
	FileMaxSize:    100,
	FileMaxBackups: 1,
	DisplayConsole: true,
	Skip:           2,
})
