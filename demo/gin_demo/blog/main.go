package main

import (
	"blog/config"
	"blog/dao"
	"blog/logger"
	"blog/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	// 载入配置文件
	if err := config.InitFromIni("conf/conf.ini"); err != nil {
		fmt.Printf("config.Init failed, err:%v\n", err)
		return
	}

	// init Logger
	if err := logger.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	// init MySQL
	if err := dao.InitMySQL(config.Conf.MySQLConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}

	// init redis
	if err := dao.InitRedis(config.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
}

func main() {
	logger.Info("start project...")
	router := gin.New()

	r := routers.SetupRouter(router)
	r.Run(fmt.Sprintf(":%d", config.Conf.ServerConfig.Port))
}
