package logx

import (
	"context"
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type fileLogger struct {
	zapLogger *zap.Logger
	ctx       context.Context
	skip      int
}

// 断言fileLogger实现了Logger接口
var _ Logger = &fileLogger{}

func (fl *fileLogger) WithContext(ctx context.Context) Logger {
	return &fileLogger{
		zapLogger: fl.zapLogger,
		skip:      fl.skip,
		ctx:       ctx,
	}
}
func (fl *fileLogger) WithCallerSkip(skip int) Logger {
	return &fileLogger{
		zapLogger: fl.zapLogger,
		ctx:       fl.ctx,
		skip:      skip,
	}
}

func (fl *fileLogger) Debug(message ...interface{}) {
	fl.print(zap.DebugLevel, message...)
}

func (fl *fileLogger) Info(message ...interface{}) {
	fl.print(zap.InfoLevel, message...)
}

func (fl *fileLogger) Warn(message ...interface{}) {
	fl.print(zap.WarnLevel, message...)
}

func (fl *fileLogger) Error(message ...interface{}) {
	fl.print(zap.ErrorLevel, message...)
}

func (fl *fileLogger) Debugf(format string, message ...interface{}) {
	fl.WithCallerSkip(fl.skip + 1).Debug(formatter(format, message...))
}

func (fl *fileLogger) Infof(format string, message ...interface{}) {
	fl.WithCallerSkip(fl.skip + 1).Info(formatter(format, message...))
}

func (fl *fileLogger) Warnf(format string, message ...interface{}) {
	fl.WithCallerSkip(fl.skip + 1).Warn(formatter(format, message...))
}

func (fl *fileLogger) Errorf(format string, message ...interface{}) {

	fl.WithCallerSkip(fl.skip + 1).Error(formatter(format, message...))
}

func formatter(format string, message ...interface{}) string {
	return fmt.Sprintf(format, message...)
}

func (fl *fileLogger) print(level zapcore.Level, messageSlice ...interface{}) {
	var msg string
	for _, message := range messageSlice {
		switch message.(type) {
		case string:
			msg = message.(string)
		case error:
			msg = message.(error).Error()
		default:
			var json = jsoniter.ConfigCompatibleWithStandardLibrary
			s, err := json.Marshal(message)
			if err != nil {
				panic("write log error:" + err.Error())
			}
			msg = string(s)
		}

		zapLogger := fl.zapLogger
		if fl.ctx != nil {
			switch traceId := fl.ctx.Value("traceId"); traceId.(type) {
			case string:
				s, _ := traceId.(string)
				zapLogger = fl.zapLogger.With(zap.String("traceId", s))
			}
		}
		zapLogger.WithOptions(zap.AddCallerSkip(fl.skip)).Log(level, msg)
	}

}

// logpath 日志文件路径
// loglevel 日志级别
func initLogger(loglevel string) *zap.Logger {
	// 日志分割
	hook, err := rotatelogs.New(
		"./logs/%Y%m%d.log",
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(10)*time.Second),
	)
	if err != nil {
		panic(err)
	}
	write := zapcore.AddSync(hook)
	// 设置日志级别
	// debug 可以打印出 info debug warn
	// info  级别可以打印 warn info
	// warn  只能打印 warn
	// debug->info->warn->error
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		CallerKey:     "caller",
		MessageKey:    "msg",
		FunctionKey:   zapcore.OmitKey,
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,

		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		write,
		level,
	)
	caller := zap.AddCaller()
	// 构造日志
	l := zap.New(core, caller)
	defer l.Sync()
	return l
}
