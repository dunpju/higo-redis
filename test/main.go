package main

import (
	"fmt"
	"github.com/dengpju/higo-redis/redis"
	"math/rand"
)

func main()  {
	redis.New(redis.PoolConfigure{Host:"192.168.42.131",Port:6379,Auth:"1qaz2wsx",
		Db:0,MaxConnections:10, MaxIdle:3, MaxIdleTime:60})
	v, _ := redis.Redis.Get("name")
	fmt.Println(v)
	if v == "" {
		_, _ = redis.Redis.Setex("name", rand.Intn(1000), 200)
	}
	ttl, _ := redis.Redis.Ttl("name")
	fmt.Println(ttl)
	if ttl < 0 {
		ttl = 0
	}
	ttl = ttl + 500
	b, _ := redis.Redis.Expire("name", ttl)
	fmt.Println(b)
	s := ""
	s += "local stockKey = KEYS[1]"
	s += "local num = tonumber(KEYS[2])"
	s += "local stockNum = redis.call('GET', stockKey)"
	s += "if not stockNum or tonumber(stockNum) < num  then"
	s += "  return 0"
	s += "end"
	s += "redis.call('decrBy', stockKey, num)"
	s += "return 1"
	r, _ := redis.Redis.Eval(s, 2, "stock", 1)
	fmt.Println(r)
}
