package robot

import (
	"fmt"
	"math/rand"
	"server/service/client"
	"sync"
	"sync/atomic"
	"time"

	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/tools"
)

const goc = 1000

type Robot struct {
	actor.Base
	clients sync.Map
}

func (s *Robot) OnInit() {
	for i := 0; i < goc; i++ {
		acc := fmt.Sprintf("robot_%v", time.Now().Nanosecond())
		s.stateLogin(acc)
		time.Sleep(time.Microsecond)
	}
}

var i int32

func (s *Robot) stateLogin(acc string) {
	// 随机randtime时间后，开启actor执行游戏
	randtime := (rand.Int63n(100000)+100)*int64(time.Millisecond) + int64(atomic.AddInt32(&i, 1))
	s.AddTimer(tools.UUID(), tools.NowTime()+randtime, func(dt int64) {
		v, _ := s.clients.LoadOrStore(acc, &client.Client{ACC: acc})
		cli := v.(*client.Client)
		expect.Nil(s.System().Add(actor.New(cli.ACC, cli, actor.SetLocalized())))

		// 随机2～8秒后，退出actor logout
		exitRandTime := rand.Intn(6) + 2
		s.AddTimer(tools.UUID(), tools.NowTime()+int64(exitRandTime*int(time.Second)), func(dt int64) {
			s.stateExit(acc)
		})
	})

}

func (s *Robot) stateExit(acc string) {
	v, _ := s.clients.Load(acc)
	cli := v.(*client.Client)
	cli.Exit()

	// 下一次进入登录状态的时间 0.5~1.5秒
	nextReloginTime := (rand.Int63n(10000) + 500) * int64(time.Millisecond)
	s.AddTimer(tools.UUID(), tools.NowTime()+nextReloginTime, func(dt int64) {
		s.stateLogin(acc)
	})
}
