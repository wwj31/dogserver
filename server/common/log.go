package common

import (
	"github.com/wwj31/dogactor/log"
	"time"
)

func ReportLog(t time.Time, level string, file string, line int, msg string) {
	if log.Levels[level] < log.TAG_WARN_I {
		return
	}
}
