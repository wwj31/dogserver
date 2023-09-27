package test

import (
	"context"
	"testing"

	"server/common/log"
	"server/common/rds"
)

func init() {
	if err := rds.NewBuilder().
		Addr("localhost:6379").
		//ClusterMode().
		Connect(); err != nil {
		log.Errorw("redis connect failed", "err", err)
		return
	}
}

func TestSADD(t *testing.T) {
	rds.Ins.SAdd(context.Background(), "fuck", 123, "fufufu", 6565, 76878, 4342)
}
