package middlewares

import (
	"blog/config"
	"blog/logger"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"

	//"github.com/gin-contrib/sessions/cookie"  	// session具体存储的介质
	//"github.com/gin-contrib/sessions/redis" 		// session具体存储的介质
	//"github.com/gin-contrib/sessions/memcached"   // session具体存储的介质
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Session() gin.HandlerFunc {
	address := fmt.Sprintf("%s:%d", config.Conf.RedisConfig.Host, config.Conf.RedisConfig.Port)
	store, err := redis.NewStore(10, "tcp", address, "", []byte("secret"))
	if err != nil {
		panic(err)
	}

	logger.Info("SessionRedis", zap.String("address", address))

	return sessions.Sessions("mySession", store)
}
