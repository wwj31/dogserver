package rds

import redisv9 "github.com/redis/go-redis/v9"

var (
	unlockScript = redisv9.NewScript(`
	if redis.call("GET",KEYS[1]) == ARGV[1] then
		return redis.call("DEL",KEYS[1])
	else
		return 0
	end
`)
)
