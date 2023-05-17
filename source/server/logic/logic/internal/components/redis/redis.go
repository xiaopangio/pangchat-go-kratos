package redis

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"logic/internal/conf"
	"time"
)

var Nil redis.Error = redis.Nil

type Redis struct {
	client *redis.Client
	helper *log.Helper
}

func NewRedisClient(cf *conf.Bootstrap, helper *log.Helper) *Redis {
	redisCf := cf.Data
	client := redis.NewClient(&redis.Options{
		Addr:         redisCf.Redis.Addr,
		Password:     redisCf.Redis.Password,
		DB:           int(redisCf.Redis.Db),
		PoolSize:     int(redisCf.Redis.PoolSize),
		MinIdleConns: int(redisCf.Redis.MinIdleConns),
		MaxRetries:   int(redisCf.Redis.MaxRetries),
		ReadTimeout:  redisCf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: redisCf.Redis.WriteTimeout.AsDuration(),
	})
	ping := client.Ping(context.Background())
	helper.Infof("redis connection: %v", ping.Val())
	if ping.Err() != nil {
		helper.Error("cannot establish redis connection: %v", ping.Err())
	}
	return &Redis{
		client: client,
		helper: helper,
	}
}
func (r *Redis) Set(key string, value any, expiration time.Duration) error {
	return r.client.Set(context.Background(), key, value, expiration).Err()
}
func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(context.Background(), key).Result()
}
func (r *Redis) Del(key string) error {
	return r.client.Del(context.Background(), key).Err()
}
func (r *Redis) Exists(key string) (int64, error) {
	return r.client.Exists(context.Background(), key).Result()
}
