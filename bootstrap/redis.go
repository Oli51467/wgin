package bootstrap

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"wgin/global"
)

// InitializeRedis 初始化redis配置
func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.App.Config.Redis.Host + ":" + global.App.Config.Redis.Port,
		Password: global.App.Config.Redis.Password, // no password set
		DB:       global.App.Config.Redis.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.App.Logger.Error("Redis connection failed, err:", zap.Any("err", err))
		return nil
	}
	global.App.Redis = client
	return client
}
