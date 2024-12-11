package logger

import (
	"github.com/qsk5yrs/testing/common/enum"
	"github.com/qsk5yrs/testing/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var _logger *zap.Logger

func init() {
	// 编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 创建Json格式编码器
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	// 创建文件写入同步器
	fileWriteSyncer := getFileLogWriter()

	var cores []zapcore.Core

	switch config.App.Env {
	case enum.ModeTest, enum.ModeProd:
		// 测试环境和生产环境日志输出到文件中
		cores = append(cores, zapcore.NewCore(encoder, fileWriteSyncer, zapcore.InfoLevel))
	case enum.ModeDev:
		// 开发环境同时向控制台和文件输出日志，设置日志级别为Debug
		cores = append(cores,
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
			zapcore.NewCore(encoder, fileWriteSyncer, zapcore.DebugLevel),
		)
	}
	// 组合日志核心
	core := zapcore.NewTee(cores...)
	// 创建日志记录器
	_logger = zap.New(core)
}

func getFileLogWriter() (writeSyncer zapcore.WriteSyncer) {
	// 使用lumberJack实现logger rotate
	lumberJackLogger := &lumberjack.Logger{
		Filename:  config.App.Log.FilePath,
		MaxSize:   config.App.Log.FileMaxSize,      // 文件最大100M
		MaxAge:    config.App.Log.BackUpFileMaxAge, // 归档文件存放60天
		Compress:  false,
		LocalTime: true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
