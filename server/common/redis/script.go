package redis

import redisv8 "github.com/go-redis/redis/v8"

var (
	unlockScript = redisv8.NewScript(`
	if redis.call("GET",KEYS[1]) == ARGV[1] then
		return redis.call("DEL",KEYS[1])
	else
		return 0
	end
`)
)
