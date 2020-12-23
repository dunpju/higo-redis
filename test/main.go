package main

import (
	"fmt"
	"github.com/dengpju/higo-redis/redis"
	"math/rand"
)

func main()  {
	redis.New(redis.PoolConfigure{Host:"192.168.8.99",Port:6379,Auth:"1qaz2wsx",
		Db:0,MaxConnections:10, MaxIdle:3, MaxIdleTime:60})
	v, _ := redis.Redis.Get("name")
	fmt.Println(v)
	if v == "" {
		redis.Redis.Setex("name", rand.Intn(1000), 200)
	}
	ttl, _ := redis.Redis.Ttl("name")
	fmt.Println(ttl)
	if ttl < 0 {
		ttl = 0
	}
	ttl = ttl + 500
	b, _ := redis.Redis.Expire("name", ttl)
	fmt.Println(b)
}
