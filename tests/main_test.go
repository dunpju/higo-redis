package tests

import (
	"fmt"
	"github.com/dunpju/higo-redis/redis"
	"math/rand"
	"testing"
)

func Test(t *testing.T) {
	redis.New(
		redis.NewPoolConfigure(
			redis.PoolHost("192.168.8.99"),
			redis.PoolPort(6379),
			redis.PoolAuth("1qaz2wsx"),
			redis.PoolDb(2),
			redis.PoolMaxConnections(10),
			redis.PoolMaxIdle(3),
			redis.PoolMaxIdleTime(60),
			redis.PoolMaxConnLifetime(10),
			redis.PoolWait(true),
		))
	fmt.Println(111)
	redis.Redis.Set("ttt", 11)
	fmt.Println(redis.Redis.ZrevrangeByScore("salary", redis.WithSubInf(), redis.WithAddInf()))
	return
	/**
	rand.Seed(time.Now().Unix())
	fmt.Println(redis.Redis.Lpush("testlist", rand.Intn(1000) + 1))
	fmt.Println(redis.Redis.Brpoplpush("testlist", "dtestlist", redis.WithTimeout(time.Duration(60))))
	fmt.Println("Brpoplpush finish")

	*/

	/**
	//不用每个请求都实例化缓存操作
	syncNewsCache := redis.NewsCache()
	defer redis.ReleaseNewsCache(syncNewsCache)
	syncNewsCache.DbGetter = func() interface{} {
		log.Println("sync get from db")
		return "sync news by id=1235"
	}
	fmt.Println(syncNewsCache.GetCache("news1235"))

	newsCache := redis.NewSimpleCache(
		redis.NewStringOperation(),
		redis.WithExpire(15),
		redis.SERILIZER_JSON, redis.NewCrossPolicy())
	newsCache.DbGetter = func() interface{} {
		log.Println("get from db")
		return "news by id=123"
	}
	fmt.Println(newsCache.GetCacheForObject("news123"))

	*/

	redis.Redis.Set("set_name", "ggg", redis.WithExpire(60))
	for {
		fmt.Println(1)
	}

	v1 := redis.Redis.MgetIterable("name", redis.WithKey("name1"))
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
	r, _ := redis.Redis.Eval(s, redis.WithEvalNumkeys(2), redis.WithEvalKey("stock1"), redis.WithEvalKey(1))
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
