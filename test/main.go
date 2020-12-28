package main

import (
	"fmt"
	"github.com/dengpju/higo-redis/redis"
	"math/rand"
)

func main()  {
	redis.New(redis.PoolConfigure{Host:"192.168.8.99",Port:6379,Auth:"1qaz2wsx",
		Db:0,MaxConnections:10, MaxIdle:3, MaxIdleTime:60})
	v1 := redis.Redis.MgetIterable("name2", "name1")
	for v1.HasNext() {
		//fmt.Printf("%s\n",v1.Next())
		fmt.Println(v1.Next())
	}
	return
	v := redis.Redis.Get("name")
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
	s := `local stockKey = KEYS[1]
local num = tonumber(KEYS[2])
local stockNum = redis.call('GET', stockKey)
if not stockNum or tonumber(stockNum) < num  then
  return 0
end
redis.call('decrBy', stockKey, num)
return 1
`
	r, _ := redis.Redis.Eval(s, 2, "stock", 1)
	fmt.Println(r)
	r1, _ := redis.Redis.Incr("abc")
	fmt.Println(r1)
	r2, _ := redis.Redis.IncrBy("abc", 5)
	fmt.Println(r2)
	r1, _ = redis.Redis.Decr("abc")
	fmt.Println(r1)
	r2, _ = redis.Redis.DecrBy("abc", 3)
	fmt.Println(r2)
}
