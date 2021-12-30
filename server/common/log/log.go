package log

import (
	"github.com/wwj31/dogactor/l"
)

var gLogger *l.Logger

func Init(lv int, path, fileName string, dispay bool) {
	gLogger = l.New(l.Option{
		Level:          l.Level(lv),
		LogPath:        path,
		FileName:       fileName,
		FileMaxAge:     5,
		FileMaxSize:    512,
		FileMaxBackups: 10,
		DisplayConsole: dispay,
		Skip:           2,
	})
}

func Debugf(msg string, args ...interface{}) {
	gLogger.Color(l.Green)
	gLogger.Debugf(msg, args...)
}

func Infof(msg string, args ...interface{}) {
	gLogger.Color(l.Gray)
	gLogger.Infof(msg, args...)
}

func Warnf(msg string, args ...interface{}) {
	gLogger.Color(l.Yellow)
	gLogger.Warnf(msg, args...)
	gLogger.CleanColor()
}

func Errorf(msg string, args ...interface{}) {
	gLogger.Color(l.Red)
	gLogger.Errorf(msg, args...)
	gLogger.CleanColor()
}

func Debugw(msg string, args ...interface{}) {
	gLogger.Color(l.Green)
	gLogger.Debugw(msg, args...)
}

func Infow(msg string, args ...interface{}) {
	gLogger.Color(l.Gray)
	gLogger.Infow(msg, args...)
}

func Warnw(msg string, args ...interface{}) {
	gLogger.Color(l.Yellow)
	gLogger.Warnw(msg, args...)
	gLogger.CleanColor()
}

func Errorw(msg string, args ...interface{}) {
	gLogger.Color(l.Red)
	gLogger.Errorw(msg, args...)
	gLogger.CleanColor()
}
