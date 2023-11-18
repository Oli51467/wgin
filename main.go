package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wgin/bootstrap"
	"wgin/global"
	"wgin/lib/logger"
)

func main() {
	// 初始化配置文件的配置
	bootstrap.InitializeConfig()
	logger.Info("success")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 启动服务器
	err := r.Run(":" + global.App.Config.Environment.Port)
	if err != nil {
		return
	}
}
