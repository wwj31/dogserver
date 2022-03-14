package main

import (
	"github.com/wwj31/dogactor/l"
	"server/common"
	"time"
)

var monitorLog *l.Logger

func monitor(path, fileName string) {
	monitorLog = l.New(l.Option{
		Level:          l.Level(0),
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
