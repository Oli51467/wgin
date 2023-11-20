package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wgin/controller"
	"wgin/middleware"
	"wgin/service"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.POST("/user/register", controller.Register)
	router.POST("/user/login", controller.Login)

	// 对token进行鉴权
	authRouter := router.Group("").Use(middleware.JWTAuth(service.AppGuardName))
	{
		authRouter.POST("/user/info", controller.Info)
		authRouter.POST("/user/logout", controller.Logout)
	}
}
