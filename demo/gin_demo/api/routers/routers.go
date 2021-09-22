package routers

import (
	"api/api"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) *gin.Engine {
	router.GET("/", api.IndexUsers)

	//路由组
	users := router.Group("api/v1/users")
	{
		users.GET("", api.GetAll)             //api/v1/users
		users.POST("/add", api.AddUsers)      //api/v1/users/add?name=wang&telephone=123456
		users.GET("/get/:id", api.GetOne)     //api/v1/users/get/13
		users.POST("/update", api.UpdateUser) //api/v1/users/update?id=13&telephone=654321
		users.POST("/del", api.DelUser)       //api/v1/users/del?id=13
	}

	return router
}
