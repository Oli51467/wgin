package global

import (
	"github.com/spf13/viper"
	"wgin/config"
)

// Application 用来存放一些项目启动时的变量，便于调用
type Application struct {
	ViperConfig *viper.Viper
	Config      config.Configuration
}

var App = new(Application)
