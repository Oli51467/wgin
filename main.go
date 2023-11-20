package main

import (
	"github.com/gin-gonic/gin"
	"wgin/bootstrap"
	"wgin/global"
)

func main() {
	// 初始化配置文件的配置
	bootstrap.InitializeConfig()
	// 初始化日志
	bootstrap.InitializeLogger()
	// 初始化验证器
	bootstrap.InitializeValidator()
	// 初始化数据库
	bootstrap.InitializeDB()
	bootstrap.InitializeRedis()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.Database != nil {
			db, _ := global.App.Database.DB()
			db.Close()
		}
	}()

	gin.SetMode(gin.ReleaseMode)
	bootstrap.RunServer()
	global.App.Logger.Info("success")
}
