package test

import (
	"context"
	"log"
	"testing"

	"server/common/rds"
)

func init() {
	if err := rds.Connect("redis://localhost:6379/3", false); err != nil {
		log.Fatalf("redids connect failed:%v", err)
	}
}

func TestSADD(t *testing.T) {
	rds.Ins.SAdd(context.Background(), "fuck", 123, "fufufu", 6565, 76878, 4342)
}
