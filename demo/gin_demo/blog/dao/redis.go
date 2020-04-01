package dao

import (
	"blog/config"
	"fmt"

	"github.com/go-redis/redis"
)

const (
	KeyArticleCount = "blog:article:read:count:%s" // 24小时文章阅读数key eq:blog:article:count:20200315
)

var (
	Client *redis.Client
)

// 初始化连接
func InitRedis(cfg *config.RedisConfig) (err error) {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})

	_, err = Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
