package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (

)

type RedisDB struct {
    client *redis.Client
}

func RedisDBNew() *RedisDB {
    obj := new(RedisDB)
    return obj
}

func (self *RedisDB) Connect(address, password string, db int) *RedisDB {
    self.client = redis.NewClient(&redis.Options{
        Addr: address,
        Password: password,
        DB: 0,
    })

    return self
}

func (self *RedisDB) Close() {
    self.client.Close()
}

func (self *RedisDB) Ping() (string, error) {
    return self.client.Ping(context.Background()).Result()
}
