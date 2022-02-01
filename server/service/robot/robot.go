package robot

import (
	"fmt"
	"math/rand"
	"server/common"
	"server/service/client"
	"time"

	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/tools"
)

const goc = 1

type Robot struct {
	actor.Base
	clients map[string]*client.Client
}

func (s *Robot) OnInit() {
	s.clients = make(map[string]*client.Client, 100)
	for i := 0; i < goc; i++ {
		randtime := rand.Int63n(10) * int64(time.Second)
		s.AddTimer(tools.UUID(), tools.NowTime()+randtime, func(dt int64) {
			acc := fmt.Sprintf("robot_%v", time.Now().Nanosecond())
			if _, ok := s.clients[acc]; ok {
				return
			}
			s.clients[acc] = &client.Client{ACC: acc}
			expect.Nil(s.System().Add(actor.New(common.Client, s.clients[acc], actor.SetLocalized())))
			s.AddTimer(tools.UUID(), tools.NowTime()+int64(5*time.Second), func(dt int64) {
				s.clients[acc].Exit()
			})
		})
	}
}
