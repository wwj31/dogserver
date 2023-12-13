package log

import (
	"github.com/wwj31/dogactor/logger"
)

var gLogger *logger.Logger

func Path() string {
	return gLogger.Option.LogPath
}
func Init(lv int, path, fileName string, dispay bool) {
	gLogger = logger.New(logger.Option{
		Level:          logger.Level(lv),
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
	gLogger.Color(logger.Green)
	gLogger.Debugf(msg, args...)
}

func Infof(msg string, args ...interface{}) {
	gLogger.Color(logger.Gray)
	gLogger.Infof(msg, args...)
}

func Warnf(msg string, args ...interface{}) {
	gLogger.Color(logger.Yellow)
	gLogger.Warnf(msg, args...)
	gLogger.CleanColor()
}

func Errorf(msg string, args ...interface{}) {
	gLogger.Color(logger.Red)
	gLogger.Errorf(msg, args...)
	gLogger.CleanColor()
}

func Debugw(msg string, args ...interface{}) {
	gLogger.Color(logger.Green)
	gLogger.Debugw(msg, args...)
}

func Infow(msg string, args ...interface{}) {
	gLogger.Color(logger.Gray)
	gLogger.Infow(msg, args...)
}

func Warnw(msg string, args ...interface{}) {
	gLogger.Color(logger.Yellow)
	gLogger.Warnw(msg, args...)
	gLogger.CleanColor()
}

func Errorw(msg string, args ...interface{}) {
	gLogger.Color(logger.Red)
	gLogger.Errorw(msg, args...)
	gLogger.CleanColor()
}
