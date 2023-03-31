package robot

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"server/service/client"

	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/tools"
)

const goc = 1

type Robot struct {
	actor.Base
	clients sync.Map
}

func (s *Robot) OnInit() {
	for i := 0; i < goc; i++ {
		acc := fmt.Sprintf("robot_%v", 1)
		s.stateLogin(acc)
		time.Sleep(time.Microsecond)
	}
}

var i int32

func (s *Robot) stateLogin(acc string) {
	// 随机randtime时间后，开启actor执行游戏
	randtime := time.Duration(rand.Int63n(1000)+100)*time.Millisecond + time.Duration(atomic.AddInt32(&i, 1))
	s.AddTimer(tools.XUID(), tools.Now().Add(randtime), func(dt time.Duration) {
		v, _ := s.clients.LoadOrStore(acc, &client.Client{ACC: acc})
		cli := v.(*client.Client)
		expect.Nil(s.System().NewActor(cli.ACC, cli, actor.SetLocalized()))

		fmt.Println("\nlogin", cli.ACC)
		// 随机2～8秒后，退出actor logout
		exitRandTime := time.Duration(rand.Intn(6) + 2)
		s.AddTimer(tools.XUID(), tools.Now().Add(exitRandTime*time.Second), func(dt time.Duration) {
			s.stateExit(acc)
		})
	})
}

func (s *Robot) stateExit(acc string) {
	v, _ := s.clients.Load(acc)
	cli := v.(*client.Client)
	cli.Exit()

	// 下一次进入登录状态的时间 0.5~1.5秒
	nextReloginTime := time.Duration(rand.Int63n(1000)+500) * time.Millisecond
	s.AddTimer(tools.XUID(), tools.Now().Add(nextReloginTime), func(dt time.Duration) {
		s.stateLogin(acc)
	})
}
