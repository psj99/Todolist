package cache

import (
	"Todolist/config"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// func NewReidsClient(ctx context.Context) *redis.Client {
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     fmt.Sprintf("%s:%s", redisCfg.RedisHost, redisCfg.RedisPort),
// 		Password: redisCfg.RedisPassword,
// 		DB:       redisCfg.RedisDbName,
// 	})

// 	_, err := client.Ping(ctx).Result()
// 	if err != nil {
// 		panic(err)
// 	}

// 	return client
// }

// 初始化redis链接
func InitRedis() {
	var redisCfg = config.Cfg.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisCfg.RedisHost, redisCfg.RedisPort),
		Password: redisCfg.RedisPassword,
		DB:       redisCfg.RedisDbName,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}
