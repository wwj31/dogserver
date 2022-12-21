package main

import (
	"net/http"
	_ "net/http/pprof"
	"server/common/log"
)

const addr = "localhost:"

func pprof(port string) {
	ip := addr + port
	log.Infow("pprof start", "ip", ip)
	go func() {
		if err := http.ListenAndServe(ip, nil); err != nil {
			log.Errorw("pprof failed", "err", err, "addr", ip)
		}

	}()
}

// go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
// go tool pprof http://localhost:6060/debug/pprof/heap?seconds=30
// go tool pprof http://localhost:6060/debug/pprof/goroutine?seconds=30
