package main

import (
	"github.com/gin-gonic/gin"
	"wgin/bootstrap"
	"wgin/global"
)

var logger = global.App.Logger

func main() {
	// 初始化配置文件的配置
	bootstrap.InitializeConfig()
	logger = bootstrap.InitializeLogger()
	logger.Info("success")

	// 初始化数据库
	global.App.Database = bootstrap.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.Database != nil {
			db, _ := global.App.Database.DB()
			db.Close()
		}
	}()

	gin.SetMode(gin.ReleaseMode)
	bootstrap.RunServer()
}
