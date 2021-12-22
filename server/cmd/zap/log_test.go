package main

import (
	"github.com/wwj31/dogactor/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"testing"
	"time"
)

func BenchmarkZap(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	core, _ := zap.NewProduction(zap.ErrorOutput(zapcore.AddSync(io.Discard)))
	sug := core.Sugar()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sug.Infow("BBBBBBB", "fewfewfew", "kkk", "fjeoiwfj", 3, "fefew", time.Second)
	}
}

func BenchmarkLog(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	log.Init(log.TAG_DEBUG_I, nil, "./_log", "demo", 1)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		log.KVs(log.Fields{"fewfewfew": "kkk", "fjeoiwfj": 3, "fefew": time.Second}).Info("BBBBBBB")
	}
}
