package common

import (
	"github.com/wwj31/dogactor/log"
	"time"
)

func InitLog(level int32, logdir, appType string, appId int32) {
	log.Init(level, reportLog, logdir, appType, appId)
}

func reportLog(t time.Time, level string, file string, line int, msg string) {
	if log.Levels[level] < log.TAG_WARN_I {
		return
	}
}
