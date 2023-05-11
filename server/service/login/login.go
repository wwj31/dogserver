package login

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"math/rand"

	"server/common"
	"server/common/log"
	"server/common/rds"
	"server/common/rdskey"
	"server/proto/outermsg/outer"
)

type Login struct {
	actor.Base
	shortIDs []int32
}

func New() *Login {
	return &Login{}
}

func (s *Login) OnInit() {
	log.Infow("login OnInit")

	result := rds.Ins.Get(context.Background(), rdskey.ShortIDKey())
	if result.Err() == redis.Nil {
		s.randShortID()
		b, _ := json.Marshal(s.shortIDs)
		rds.Ins.Set(context.Background(), rdskey.ShortIDKey(), string(b), 0)
	} else {
		var shortID []int32
		_ = json.Unmarshal([]byte(result.Val()), &shortID)
		s.shortIDs = shortID
	}
}

func (s *Login) OnStop() bool {
	log.Debugw("login stop", "id", s.ID())
	return true
}

func (s *Login) OnHandle(m actor.Message) {
	payload := m.Payload()
	v, _, gSession, err := common.UnwrappedGateMsg(payload)

	expect.Nil(err)
	switch msg := v.(type) {
	case *outer.LoginReq:
		err = s.LoginReq(m.GetSourceId(), gSession, msg)
	default:
		err = fmt.Errorf("undefined localmsg type %v", msg)
	}

	if err != nil {
		log.Errorw("handle outer error", "err", err)
	}
}

func (s *Login) LoginReq(sourceId string, gSession common.GSession, msg *outer.LoginReq) error {
	log.Debugf(msg.String())
	s.Login(gSession, msg)
	return nil
}

func (s *Login) randShortID() {
	var pool []int32
	for i := int32(140150); i < 999999; i++ {
		pool = append(pool, i)
	}
	// 使用 Fisher-Yates 算法洗牌
	for i := len(pool) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		pool[i], pool[j] = pool[j], pool[i]
	}
	s.shortIDs = pool
}
