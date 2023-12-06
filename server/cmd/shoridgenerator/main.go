package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/spf13/cast"

	"server/common/rds"
)

var (
	key     = flag.String("key", "shortid", "the key of the short id")
	uri     = flag.String("uri", "redis://:@localhost:6379/0", "the addr of the redis")
	cluster = flag.Bool("cluster", false, "is redis cluster")
)

func main() {
	flag.Parse()

	if err := rds.Connect(*uri, *cluster); err != nil {
		fmt.Println("redis connect failed", "err", err)
		return
	}

	var (
		genStart int64
		genEnd   int64
	)

	progress := int64(1000000)
	data, err := os.ReadFile(".id")
	if err == nil {
		genStart = cast.ToInt64(string(data))
	} else {
		genStart = progress
	}

	genEnd = genStart + progress

	fmt.Println("gen id from ", genStart, " to ", genEnd)
	set := randShortID(genStart, genEnd)
	fmt.Println("upload to redis key:", *key)
	defer func() {
		err := os.WriteFile(".id", []byte(cast.ToString(genEnd)), 0644)
		if err != nil {
			fmt.Println("write .id file failed err:", err)
		}
	}()

	proc := 100000
	var (
		begin = 0
		end   = proc
	)

	for {
		if begin >= len(set) {
			break
		}
		if end > len(set) {
			end = len(set)
		}

		fmt.Println("number: ", begin, " - ", end, " upload success!")
		cmd := rds.Ins.SAdd(context.Background(), *key, set[begin:end]...)
		if cmd.Err() != nil {
			fmt.Println("upload failed err:", cmd.Err())
			return
		}
		begin = end + 1
		end += proc
	}
	fmt.Println("Done!")

}

func randShortID(begin, end int64) []interface{} {
	var pool []interface{}
	for i := begin; i < end; i++ {
		pool = append(pool, i)
	}

	// 使用 Fisher-Yates 算法洗牌
	for i := len(pool) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		pool[i], pool[j] = pool[j], pool[i]
	}
	return pool
}
