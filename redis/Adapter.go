package redis

import (
	"github.com/dengpju/higo-throw/throw"
	"github.com/gomodule/redigo/redis"
)

var Redis RedisAdapter

type RedisAdapter struct {
	conn redis.Conn
}

func NewRedisAdapter() *RedisAdapter {
	return &RedisAdapter{}
}

func (this *RedisAdapter) Connection() *RedisAdapter {
	this.conn = RedisPool.Get()
	return this
}

func (this *RedisAdapter) Conn() redis.Conn {
	this.conn = RedisPool.Get()
	return this.conn
}

func (this *RedisAdapter) Set(key string, v interface{}, args ...*Parameter) string {
	return NewStringResult(redis.String(this.Executor("set", Parameters(args).Handle(key, v)...))).String()
}

func (this *RedisAdapter) Setnx(key string, v interface{}) bool {
	return NewBoolResult(redis.Bool(this.Executor("setnx", key, v))).Bool()
}

func (this *RedisAdapter) Setex(key string, v interface{}, expire int) string {
	return NewStringResult(redis.String(this.Executor("setex", key, expire, v))).Unwrap().String()
}

func (this *RedisAdapter) Get(key string) string {
	return this.get(key).Unwrap().String()
}

func (this *RedisAdapter) get(key string) *StringResult {
	return NewStringResult(redis.String(this.Executor("get", key)))
}

func (this *RedisAdapter) GetDefault(key string, v string) string {
	return NewStringResult(redis.String(this.Executor("get", key))).Default(v).String()
}

func (this *RedisAdapter) Mget(key string, keys ...*Parameter) []string {
	return NewStringsResult(redis.Strings(this.Executor("mget", Parameters(keys).Handle(key)...))).Unwrap().Strings()
}

func (this *RedisAdapter) MgetIterable(key string, keys ...*Parameter) *Iterator {
	return NewSliceResult(redis.Strings(this.Executor("mget", Parameters(keys).Handle(key)...))).Unwrap().Iterable()
}

func (this *RedisAdapter) args(keys ...string) (args []interface{}) {
	for _, k := range keys {
		args = append(args, k)
	}
	return
}

func (this *RedisAdapter) GetByte(key string) ([]byte, error) {
	v, err := redis.Bytes(this.Executor("get", key))
	return v, err
}

func (this *RedisAdapter) Expire(key string, seconds int) (bool, error) {
	b, err := redis.Bool(this.Executor("expire", key, seconds))
	return b, err
}

func (this *RedisAdapter) Ttl(key string) int {
	return NewIntResult(redis.Int(this.Executor("ttl", key))).Unwrap().Int()
}

func (this *RedisAdapter) Eval(script string, args ...*Parameter) (interface{}, error) {
	return this.Executor("eval", Parameters(args).Handle(script)...)
}

func (this *RedisAdapter) Incr(key string) int {
	return NewIntResult(redis.Int(this.Executor("incr", key))).Unwrap().Int()
}

func (this *RedisAdapter) IncrBy(key string, score int) int {
	return NewIntResult(redis.Int(this.Executor("incrby", key, score))).Unwrap().Int()
}

func (this *RedisAdapter) Decr(key string) int {
	return NewIntResult(redis.Int(this.Executor("decr", key))).Unwrap().Int()
}

func (this *RedisAdapter) DecrBy(key string, score int) int {
	return NewIntResult(redis.Int(this.Executor("decrby", key, score))).Unwrap().Int()
}

func (this *RedisAdapter) Lpush(key string, v interface{}) int {
	return NewIntResult(redis.Int(this.Executor("lpush", key, v))).Unwrap().Int()
}

func (this *RedisAdapter) Llen(key string) int {
	return NewIntResult(redis.Int(this.Executor("llen", key))).Unwrap().Int()
}

func (this *RedisAdapter) Lrange(key string, start int, end int) []string {
	return NewStringsResult(redis.Strings(this.Executor("lrange", key, start, end))).Unwrap().Strings()
}

func (this *RedisAdapter) Lpop(key string) string {
	return NewStringResult(redis.String(this.Executor("lpop", key))).Unwrap().String()
}

func (this *RedisAdapter) Rpush(key string, v interface{}) int {
	return NewIntResult(redis.Int(this.Executor("rpush", key, v))).Unwrap().Int()
}

func (this *RedisAdapter) Rpop(key string) string {
	return NewStringResult(redis.String(this.Executor("rpop", key))).Unwrap().String()
}

func (this *RedisAdapter) Del(key string) int {
	return NewIntResult(redis.Int(this.Executor("del", key))).Unwrap().Int()
}

func (this *RedisAdapter) Zset() {

}

func (this *RedisAdapter) Hash() {

}

func (this *RedisAdapter) Close() {
	defer this.conn.Close()
}

func (this *RedisAdapter) Executor(commandName string, args ...interface{}) (interface{}, error) {
	defer this.Close()
	reply, err := this.Conn().Do(commandName, args...)
	if err != nil {
		this.Close()
		throw.Throw(err, 0)
	}
	return reply, err
}
