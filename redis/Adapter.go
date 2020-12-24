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

func (this *RedisAdapter) Set(key string, v interface{}) (bool, error) {
	b, err := redis.Bool(this.Executor("set", key, v))
	return b, err
}

func (this *RedisAdapter) Get(key string) (string, error) {
	v, err := redis.String(this.Executor("get", key))
	return v, err
}

func (this *RedisAdapter) GetByte(key string) ([]byte, error) {
	v, err := redis.Bytes(this.Executor("get", key))
	return v, err
}

func (this *RedisAdapter) Setex(key string, data interface{}, expire int) (bool, error) {
	b, err := redis.Bool(this.Executor("setex", key, expire, data))
	return b, err
}

func (this *RedisAdapter) Expire(key string, seconds int) (bool, error) {
	b, err := redis.Bool(this.Executor("expire", key, seconds))
	return b, err
}

func (this *RedisAdapter) Ttl(key string) (int, error) {
	ttl, err := redis.Int(this.Executor("ttl", key))
	return ttl, err
}

func (this *RedisAdapter) Eval(script string, numkeys int, args ...interface{}) (interface{}, error) {
	params := make([]interface{}, 0)
	params = append(params, script)
	params = append(params, numkeys)
	params = append(params, args...)
	return this.Executor("eval", params...)
}

func (this *RedisAdapter) Incr(key string) (int, error) {
	r, err := redis.Int(this.Executor("incr", key))
	return r, err
}

func (this *RedisAdapter) IncrBy(key string, score int) (int, error) {
	r, err := redis.Int(this.Executor("incrby", key, score))
	return r, err
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
