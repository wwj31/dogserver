package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/spf13/cast"
	"math/rand"
	"os"
	"server/common/rds"
)

var (
	key  = flag.String("key", "shortId", "the key of the short id")
	pwd  = flag.String("pwd", "ewqmh388", "the password of the redis")
	addr = flag.String("addr", "localhost:6379", "the addr of the redis")
)

func main() {
	flag.Usage = func() { fmt.Println("flag param error") }
	flag.Parse()

	if err := rds.NewBuilder().
		Addr(*addr).
		Password(*pwd).
		Connect(); err != nil {
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
