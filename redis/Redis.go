package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

var RedisPool *redis.Pool

var redisOnce sync.Once

func New(configure PoolConfigure) *redis.Pool {
	redisOnce.Do(func() {
		RedisPool = &redis.Pool {
			MaxActive:   configure.MaxConnections,
			MaxIdle:     configure.MaxIdle,
			IdleTimeout: time.Duration(configure.MaxIdleTime) * time.Second,
			Dial: func() (conn redis.Conn, e error) {
				return redis.Dial("tcp",
					fmt.Sprintf("%s:%d", configure.Host, configure.Port),
					redis.DialDatabase(configure.Db),
					redis.DialPassword(configure.Auth),
				)
			},
		}
		Redis = RedisAdapter{}
	})
	return RedisPool
}