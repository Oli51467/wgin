package bootstrap

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"wgin/global"
	"wgin/lib/logger"
)

func InitializeConfig() *viper.Viper {
	// 配置文件路径
	config := "config.yaml"
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}
	// 初始化 viper
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(config)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}
	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed: ", in.Name)
		// 重新加载配置 将json解析成global.Environment.config
		if err := v.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}
	loggerConfig := global.App.Config.Logger
	logger.Setup(&logger.Settings{
		Path:       loggerConfig.Path,
		Name:       loggerConfig.Name,
		Ext:        loggerConfig.Ext,
		TimeFormat: loggerConfig.TimeFormat,
	})
	return v
}
