package login

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/tools"
	"math/rand"
	"sync"
	"time"

	"server/common"
	"server/common/log"
	"server/common/rds"
	"server/common/rdskey"
	"server/proto/outermsg/outer"
)

type Login struct {
	actor.Base
	shortIDs      []int32
	lastSaveCount int
	lock          sync.Mutex
}

func New() *Login {
	return &Login{}
}

func (s *Login) OnInit() {
	log.Infow("login OnInit")

	result := rds.Ins.Get(context.Background(), rdskey.ShortIDKey())
	if result.Err() == redis.Nil {
		s.randShortID()
		s.saveShortID()
	} else {
		var shortID []int32
		_ = json.Unmarshal([]byte(result.Val()), &shortID)
		s.shortIDs = shortID
		s.lastSaveCount = len(s.shortIDs)
	}

	// 定期检查，数量发送变化就存
	s.AddTimer(tools.XUID(), time.Now().Add(time.Minute), func(dt time.Duration) {
		if s.lastSaveCount != len(s.shortIDs) {
			s.saveShortID()
		}
	})
}

func (s *Login) saveShortID() {
	b, _ := json.Marshal(s.shortIDs)
	rds.Ins.Set(context.Background(), rdskey.ShortIDKey(), string(b), 0)
	s.lastSaveCount = len(s.shortIDs)
}
func (s *Login) GetShortID() int32 {
	s.lock.Lock()
	defer s.lock.Unlock()
	id := s.shortIDs[len(s.shortIDs)-1]
	s.shortIDs = s.shortIDs[:len(s.shortIDs)-1]
	return id
}

func (s *Login) OnStop() bool {
	s.saveShortID()
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
