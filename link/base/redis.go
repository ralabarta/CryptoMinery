package base

import (
	"fmt"
	"github.com/IEatLemons/BaseGin/config"
	"github.com/go-redis/redis"
)

var Redis *redis.Client

// 初始化连接
func InitRedis(cfg *config.Redis) (err error) {
	RedisCli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.Db,       // use default DB
	})

	_, err = RedisCli.Ping().Result()
	if err != nil {
		return err
	}
	Redis = RedisCli
	return nil
}
