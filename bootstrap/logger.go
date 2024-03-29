package bootstrap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
	"wgin/global"
	"wgin/util"
)

var (
	level   zapcore.Level // zap 日志等级
	options []zap.Option  // zap 配置项
)

func InitializeLogger() *zap.Logger {
	// 创建根目录
	createRootDir()

	// 设置日志等级
	setLogLevel()

	if global.App.Config.Logger.ShowLine {
		options = append(options, zap.AddCaller())
	}

	logger := zap.New(getZapCore(), options...)
	global.App.Logger = logger
	return logger
}

func createRootDir() {
	if ok, _ := util.PathExists(global.App.Config.Logger.RootDir); !ok {
		_ = os.Mkdir(global.App.Config.Logger.RootDir, os.ModePerm)
	}
}

func setLogLevel() {
	switch global.App.Config.Logger.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

// 扩展 Zap
func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(global.App.Config.Environment.Env + "." + l.String())
	}

	// 设置编码器
	if global.App.Config.Logger.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 初始化 zap
	var writes = []zapcore.WriteSyncer{getLogWriter(), zapcore.AddSync(os.Stdout)}
	return zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writes...), level)
}

// 使用 lumberjack 作为日志写入器
func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   global.App.Config.Logger.RootDir + "/" + global.App.Config.Logger.Filename,
		MaxSize:    global.App.Config.Logger.MaxSize,
		MaxBackups: global.App.Config.Logger.MaxBackups,
		MaxAge:     global.App.Config.Logger.MaxAge,
		Compress:   global.App.Config.Logger.Compress,
	}

	return zapcore.AddSync(file)
}
