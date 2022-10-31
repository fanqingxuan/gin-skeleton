package svc

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogType int8

const (
	BusinessLogType LogType = iota - 1
	RequestLogType
	PanicLogType
)

type Log struct {
	log     *zap.Logger
	ctx     context.Context
	logtype LogType
}

func NewLog(logpath string, loglevel string, logtype LogType) *Log {
	return &Log{
		log:     initLogger("./logs/"+logpath, loglevel, logtype),
		logtype: logtype,
	}
}

func (that *Log) WithContext(ctx context.Context) *Log {
	that.ctx = ctx
	return &Log{
		ctx:     ctx,
		log:     that.log,
		logtype: that.logtype,
	}
}

func (that *Log) Debug(keywords string, message interface{}) {
	that.Printf(zap.DebugLevel, keywords, message)
}

func (that *Log) Info(keywords string, message interface{}) {
	that.Printf(zap.InfoLevel, keywords, message)
}

func (that *Log) Warn(keywords string, message interface{}) {
	that.Printf(zap.WarnLevel, keywords, message)
}

func (that *Log) Error(keywords string, message interface{}) {
	that.Printf(zap.ErrorLevel, keywords, message)
}

func (that *Log) Printf(level zapcore.Level, keywords string, message interface{}) {
	var msg string
	switch message.(type) {
	case string:
		msg = message.(string)
	default:
		s, err := json.Marshal(message)
		if err != nil {
			panic("write log error:" + err.Error())
		}
		msg = string(s)
	}

	traceId := that.ctx.Value("traceId")

	var s string
	if that.logtype == BusinessLogType {
		s = fmt.Sprintf("%s\t%s\t%s", traceId, keywords, msg)
	} else {
		s = fmt.Sprintf("%s\t%s", traceId, msg)
	}

	that.log.Log(level, s)
}

// logpath 日志文件路径
// loglevel 日志级别
func initLogger(logpath string, loglevel string, logtype LogType) *zap.Logger {
	// 日志分割
	hook, err := rotatelogs.New(
		strings.Trim(logpath, "/")+"/%F.log",
		rotatelogs.WithMaxAge(30*24*time.Hour),
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
		NameKey:       "logger",
		CallerKey:     "file",
		MessageKey:    "keywords",
		FunctionKey:   zapcore.OmitKey,
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,

		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	if logtype == BusinessLogType {
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		// zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&write)), // 打印到控制台和文件
		write,
		level,
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	skip_caller := zap.AddCallerSkip(2)

	// 构造日志
	l := zap.New(core, caller, skip_caller)
	defer l.Sync()
	return l
}
