package logx

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ServeLevel zapcore.Level = zapcore.DebugLevel - 3
)

type fileLogger struct {
	zapLogger *zap.Logger
	ctx       context.Context
	skip      int
}

const (
	infoFileName  = "info.log"
	errorFileName = "error.log"
	serveFileName = "access.log"
)

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

func (fl *fileLogger) Serve(message ...interface{}) {
	fl.print(ServeLevel, message...)
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
	zapLogger := fl.zapLogger
	if fl.ctx != nil {
		switch traceId := fl.ctx.Value("traceId"); traceId.(type) {
		case string:
			s, _ := traceId.(string)
			zapLogger = fl.zapLogger.With(zap.String("traceId", s))
		}
	}
	for _, message := range messageSlice {
		switch message.(type) {
		case string:
			msg = message.(string)

		case error:
			msg = message.(error).Error()
		default:
			s, err := json.Marshal(message)
			if err != nil {
				panic("write log error:" + err.Error())
			}
			msg = string(s)
		}

		zapLogger.WithOptions(zap.AddCallerSkip(fl.skip)).Log(level, msg)
	}

}

func getWriter(filename string) io.Writer {
	hook := NewFileWriter(
		fmt.Sprintf("%s/%s", "./logs", filename),
		//"2006-01-02T15-04-05.000"
		"20060102",
	)

	return hook
}

// logpath 日志文件路径
// loglevel 日志级别
func initLogger(loglevel string) *zap.Logger {

	// 设置日志级别
	// debug 可以打印出 info debug warn
	// info  级别可以打印 warn info
	// warn  只能打印 warn
	// debug->info->warn->error
	var level zapcore.Level
	level, err := zapcore.ParseLevel(loglevel)
	if err != nil {
		level = zap.InfoLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		CallerKey:     "caller",
		MessageKey:    "msg",
		FunctionKey:   zapcore.OmitKey,
		StacktraceKey: "",
		LineEnding:    zapcore.DefaultLineEnding,

		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeLevel: func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			if l != ServeLevel {
				enc.AppendString(l.CapitalString())
			} else {
				enc.AppendString(zapcore.InfoLevel.CapitalString())
			}
		},
	}

	// 根据日志级别拆分日期
	infoWriter := getWriter(infoFileName)
	warnWriter := getWriter(errorFileName)
	serveWriter := getWriter(serveFileName)

	eableInfoLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= level && l < zap.WarnLevel && l != ServeLevel
	})
	enableWarnLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= level && l >= zap.WarnLevel && l != ServeLevel
	})
	enableserveLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == ServeLevel
	})
	cores := []zapcore.Core{
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(infoWriter), eableInfoLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(warnWriter), enableWarnLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(serveWriter), enableserveLevel),
	}

	core := zapcore.NewTee(
		cores...,
	)
	caller := zap.AddCaller()
	// 构造日志
	l := zap.New(core, caller)
	defer l.Sync()
	return l
}
