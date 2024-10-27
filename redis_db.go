package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
    // ip must be same as in the compose file, links section
    REDIS_DEFAULT_IP = "redis"
    REDIS_DEFAULT_PORT = 6379
    REDIS_DEFAULT_PASSWORD = ""
    REDIS_DEFAULT_DB = 0
)

var (

)

type RedisCache struct {
    client *redis.Client
    ctx context.Context
}

func RedisCacheNew() *RedisCache {
    obj := new(RedisCache)
    obj.ctx = context.Background()
    return obj
}

func (self *RedisCache) Connect(address, password string, db int) *RedisCache {
    self.client = redis.NewClient(&redis.Options{
        Addr: address,
        Password: password,
        DB: 0,
    })

    return self
}

func (self *RedisCache) Save(key string, value any, expiration time.Duration) error {
    err := self.client.Set(self.ctx, key, value, expiration).Err()
    return err
}

func (self *RedisCache) Get(key string) (string, error) {
    result, err := self.client.Get(self.ctx, key).Result()
    return result, err
}

func (self *RedisCache) Close() {
    self.client.Close()
}

func (self *RedisCache) Ping() (string, error) {
    return self.client.Ping(context.Background()).Result()
}
