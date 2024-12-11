package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
	"runtime"
)

type logger struct {
	ctx     context.Context
	traceId string
	spanId  string
	pSpanId string
	_logger *zap.Logger
}

func (l *logger) Debug(msg string, kv ...interface{}) {
	l.log(zapcore.DebugLevel, msg, kv...)
}

func (l *logger) Info(msg string, kv ...interface{}) {
	l.log(zapcore.InfoLevel, msg, kv...)
}

func (l *logger) Warn(msg string, kv ...interface{}) {
	l.log(zapcore.WarnLevel, msg, kv...)
}

func (l *logger) Error(msg string, kv ...interface{}) {
	l.log(zapcore.ErrorLevel, msg, kv...)
}

// New logger构造函数
func New(ctx context.Context) *logger {
	var traceId, spandId, pSpandId string
	if ctx.Value("traceid") != nil {
		traceId = ctx.Value("traceid").(string)
	}
	if ctx.Value("spanid") != nil {
		spandId = ctx.Value("spanid").(string)
	}
	if ctx.Value("pspanid") != nil {
		pSpandId = ctx.Value("pspanid").(string)
	}
	return &logger{
		traceId: traceId,
		spanId:  spandId,
		pSpanId: pSpandId,
		_logger: _logger,
	}
}

// 封装通用写日志方法
func (l *logger) log(lvl zapcore.Level, msg string, kv ...interface{}) {
	// 保证要打印的日志信息成对出现
	if len(kv)%2 != 0 {
		kv = append(kv, "undefined")
	}
	// 日志信息中新增追踪参数
	kv = append(kv, "traceid", l.traceId, "spanid", l.spanId, "pspanid", l.pSpanId)
	// 日中信息中新增调用者信息，以便定位程序位置
	funcName, file, line := l.getLoggerCallerInfo()
	kv = append(kv, "func", funcName, "file", file, "line", line)
	// 初始化日志字段,创建一个初始长度为0、容量为kv长度一半的切片
	fields := make([]zap.Field, 0, len(kv)/2)
	// 遍历kv中的键值对
	for i := 0; i < len(kv); i += 2 {
		// 将键值格式化为字符串
		k := fmt.Sprintf("%v", kv[i])
		// 追加日志字段
		fields = append(fields, zap.Any(k, kv[i+1]))
	}
	// 调用check方法，判断日志级别写日志
	ce := l._logger.Check(lvl, msg)
	ce.Write(fields...)
}

// getLoggerCallerInfo 日志调用者信息 -- 方法名, 文件名, 行号
func (l *logger) getLoggerCallerInfo() (funcName, file string, line int) {

	pc, file, line, ok := runtime.Caller(3) // 回溯拿调用日志方法的业务函数的信息
	if !ok {
		return
	}
	file = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	return
}
