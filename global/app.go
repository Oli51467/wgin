package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"wgin/config"
)

// Application 用来存放一些项目启动时的变量，便于调用
type Application struct {
	ViperConfig *viper.Viper
	Config      config.Configuration
	Logger      *zap.Logger
	Database    *gorm.DB
	Redis       *redis.Client
}

var App = new(Application)
