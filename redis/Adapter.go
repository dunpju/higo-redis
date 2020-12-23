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

func (this *RedisAdapter) Set(key string, v interface{}) bool {
	this.Connection()
	defer this.conn.Close()
	_, err := this.conn.Do("set", key, v)
	if err != nil {
		this.conn.Close()
		throw.Throw(err, 0)
	}
	return true
}

func (this *RedisAdapter) Get(key string) (string, error) {
	this.Connection()
	defer this.conn.Close()
	v, err := redis.String(this.conn.Do("get", key))
	return v, err
}

func (this *RedisAdapter) GetByte(key string) ([]byte, error) {
	this.Connection()
	defer this.conn.Close()
	v, err := redis.Bytes(this.conn.Do("get", key))
	return v, err
}

func (this *RedisAdapter) Setex(key string, data interface{}, expire int) bool {
	this.Connection()
	defer this.conn.Close()
	_, err := this.conn.Do("setex", key, expire, data)
	if err != nil {
		this.conn.Close()
		throw.Throw(err, 0)
	}
	return true
}

func (this *RedisAdapter) Expire(key string, seconds int) (bool, error) {
	this.Connection()
	defer this.conn.Close()
	b, err := redis.Bool(this.conn.Do("expire", key, seconds))
	return b, err
}

func (this *RedisAdapter) Ttl(key string) (int, error) {
	this.Connection()
	defer this.conn.Close()
	ttl, err := redis.Int(this.conn.Do("ttl", key))
	return ttl, err
}
