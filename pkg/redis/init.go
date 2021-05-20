package redis

import (
	"opencodes/pkg/config"

	"github.com/go-redis/redis"
)

var client *redis.Client

func InitRedis() {
	redisCfg := config.RedisConfig
	client = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       0,
	})
}
