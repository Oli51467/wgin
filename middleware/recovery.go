package middleware

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"wgin/common/response"
	"wgin/global"
)

// CustomRecovery 将错误日志写入文件中
func CustomRecovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(
		&lumberjack.Logger{
			Filename:   global.App.Config.Logger.RootDir + "/" + global.App.Config.Logger.Filename,
			MaxSize:    global.App.Config.Logger.MaxSize,
			MaxBackups: global.App.Config.Logger.MaxBackups,
			MaxAge:     global.App.Config.Logger.MaxAge,
			Compress:   global.App.Config.Logger.Compress,
		},
		response.ServerError)
}
