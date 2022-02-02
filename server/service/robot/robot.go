package robot

import (
	"fmt"
	"math/rand"
	"server/service/client"
	"sync"
	"time"

	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/tools"
)

const goc = 100

type Robot struct {
	actor.Base
	clients sync.Map
}

func (s *Robot) OnInit() {
	for i := 0; i < goc; i++ {
		randtime := rand.Int63n(10000)*int64(time.Millisecond) + int64(i)
		s.AddTimer(tools.UUID(), tools.NowTime()+randtime, func(dt int64) {
			acc := fmt.Sprintf("robot_%v", time.Now().Nanosecond())
			v, _ := s.clients.LoadOrStore(acc, &client.Client{ACC: acc})
			cli := v.(*client.Client)
			expect.Nil(s.System().Add(actor.New(cli.ACC, cli, actor.SetLocalized())))
			s.AddTimer(tools.UUID(), tools.NowTime()+int64(5*time.Second), func(dt int64) {
				cli.Exit()
				nextReloginTime := rand.Int63n(10000) * int64(time.Millisecond)
				s.AddTimer(tools.UUID(), tools.NowTime()+nextReloginTime, func(dt int64) {
					expect.Nil(s.System().Add(actor.New(cli.ACC, cli, actor.SetLocalized())))
				})
			})
		})
	}
}
