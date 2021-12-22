package main

import (
	"bufio"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/wwj31/dogactor/tools"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"os"
	"path"
	"time"
)

type Syncer interface {
	Sync() error
}
type Sync struct {
	*bufio.Writer
}

func (s *Sync) Sync() error {
	s.Writer.Flush()
	return nil
}

func main() {
	baseLogPath := path.Join("./", fmt.Sprintf("%v%v", "test", 1))
	os.MkdirAll(baseLogPath, os.ModePerm)
	filename := fmt.Sprintf("%v%v", "test", 1)
	outfile, err := rotatelogs.New(
		path.Join(baseLogPath, filename+".%Y%m%d%H%M.log"),
		rotatelogs.WithLinkName(path.Join(baseLogPath, filename)), // 生成软链，指向最新日志文件
		//rotatelogs.WithMaxAge(5*time.Second),     // 文件最大保存时间
		rotatelogs.WithRotationTime(2*time.Hour),  // 日志切割时间间隔
		rotatelogs.WithRotationSize(1024*1024*50), // 日志切割大小 50M一个文件
	)
	if err != nil {
		fmt.Println(err)
	}

	log.SetOutput(io.MultiWriter(os.Stdout, outfile))
	buff := bufio.NewWriter(log.Writer())

	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), &Sync{Writer: buff}, zap.DebugLevel)
	lo := zap.New(core)
	sug := lo.Sugar()

	go func() {
		for {
			sug.Error("failed to fetch URL")
		}
	}()

	for {
		tools.PrintMemUsage()
		time.Sleep(1 * time.Second)
	}

}
