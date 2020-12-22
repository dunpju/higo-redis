package redis

import (
	"fmt"
	"github.com/dengpju/higo-throw/throw"
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

var RedisPool *redis.Pool

var redisOnce sync.Once

func InitRedisPool() *redis.Pool {
	redisOnce.Do(func() {
		_redis := Config("REDIS")
		confDefault := _redis.Configure("DEFAULT")
		pool := confDefault.Configure("POOL")
		RedisPool = &redis.Pool {
			MaxActive:   pool.IntValue("MAX_CONNECTIONS"),
			MaxIdle:     pool.IntValue("MAX_IDLE"),
			IdleTimeout: time.Duration(pool.IntValue("MAX_IDLE_TIME")) * time.Second,
			Dial: func() (conn redis.Conn, e error) {
				return redis.Dial("tcp",
					fmt.Sprintf("%s:%s", confDefault.StrValue("HOST"), confDefault.StrValue("PORT")),
					redis.DialDatabase(confDefault.IntValue("DB")),
					redis.DialPassword(confDefault.StrValue("AUTH")),
				)
			},
		}
		Redis = RedisAdapter{}
	})
	return RedisPool
}