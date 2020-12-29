package main

import (
	"fmt"
	"github.com/dengpju/higo-redis/redis"
	"math/rand"
)

func main()  {
	redis.New(redis.PoolConfigure{Host:"192.168.8.99",Port:6379,Auth:"1qaz2wsx",
		Db:0,MaxConnections:10, MaxIdle:3, MaxIdleTime:60})
	v1 := redis.Redis.MgetIterable("name", "name1")
	for v1.HasNext() {
		//fmt.Printf("%s\n",v1.Next())
		fmt.Println(v1.Next())
	}

	v := redis.Redis.GetDefault("name", "")
	fmt.Println(v)
	if v == "" {
		_ = redis.Redis.Setex("name", rand.Intn(1000), 200)
	}
	ttl := redis.Redis.Ttl("name")
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
	r, _ := redis.Redis.Eval(s, 2, "stock1", 1)
	fmt.Println(r)
	r1 := redis.Redis.Incr("abc")
	fmt.Println(r1)
	r2 := redis.Redis.IncrBy("abc", 5)
	fmt.Println(r2)
	r1 = redis.Redis.Decr("abc")
	fmt.Println(r1)
	r2 = redis.Redis.DecrBy("abc", 3)
	fmt.Println(r2)
	l1 := redis.Redis.Lrange("ll", 0, -1)
	fmt.Println(l1)
	l2 := redis.Redis.Lpush("ll", "ddd")
	fmt.Println(l2)
	l3 := redis.Redis.Rpop("ll")
	fmt.Println(l3)
}
