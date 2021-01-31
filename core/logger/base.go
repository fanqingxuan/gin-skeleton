package logger

import (
	"io"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type CustomLog struct {
	File    string
	Line    int
	Message string
}

func getLogLevel(level string) zapcore.Level {

	m := map[string]zapcore.Level{
		"info":  zapcore.InfoLevel,
		"debug": zapcore.DebugLevel,
		"error": zapcore.ErrorLevel,
		"warn":  zapcore.WarnLevel,
	}

	l, ok := m[level]
	if ok {
		return l
	} else {
		return zapcore.InfoLevel
	}
}

func GetLogger(path string, level string, bapp bool) *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "file",
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "keywords",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		}, //zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	logLevel := getLogLevel(level)

	// 实现两个判断日志等级的interface (其实 zapcore.*Level 自身就是 interface)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl >= logLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel && lvl >= logLevel
	})

	// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现
	infoWriter := getWriter("./logs/" + path + "/info")
	warnWriter := getWriter("./logs/" + path + "/error")
	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
	)
	return zap.New(core) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名
	hook, err := rotatelogs.New(
		strings.Replace(filename, ".log", "", -1) + "-%Y%m%d.log", // 日志文件名格式
	)

	if err != nil {
		panic(err)
	}
	return hook
}
