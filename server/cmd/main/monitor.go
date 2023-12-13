package main

import (
	"time"

	"github.com/wwj31/dogactor/logger"

	"server/common"
)

var monitorLog *logger.Logger

func monitor(path, fileName string) {
	monitorLog = logger.New(logger.Option{
		Level:          logger.Level(0),
		LogPath:        path,
		FileName:       fileName,
		FileMaxAge:     5,
		FileMaxSize:    512,
		FileMaxBackups: 10,
		DisplayConsole: false,
		Skip:           2,
	})
	go func() {
		tick := time.Tick(10 * time.Second)
		for range tick {
			str := common.PrintMemUsage()
			monitorLog.Infof("monitor :%v", str)
		}
	}()
}
