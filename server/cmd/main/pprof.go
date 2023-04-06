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

// 安装graphviz后，通过以下命令打开浏览器可以查看分析图和火焰图 6061端口可以随意定
// go tool pprof -http localhost:6061 ./pprof.samples.cpu.001.pb.gz

/*

cd "$(brew --repo)"
git remote set-url origin https://mirrors.ustc.edu.cn/brew.git
cd "$(brew --repo)/Library/Taps/homebrew/homebrew-core"
git remote set-url origin https://mirrors.ustc.edu.cn/homebrew-core.git

brew update
brew install graphviz
*/
