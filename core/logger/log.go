package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-demo/config"
	"gin-demo/util"
	"runtime"

	"go.uber.org/zap"
)

var appLogger *zap.Logger
var RequestLogger, PanicLogger *zap.Logger

func InitLog() {
	appLogger = GetLogger("app", config.Log.Level, true)
	RequestLogger = GetLogger("request", "info", false)
	PanicLogger = GetLogger("panic", "info", false)
}

func SyncLog() {
	defer func() {
		appLogger.Sync()
		RequestLogger.Sync()
		PanicLogger.Sync()
	}()
}

func getData(ctx context.Context, keywords string, data interface{}) string {
	b, _ := json.Marshal(data)
	message := string(b)
	requestId := ctx.Value("requestId").(string)
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		panic("can not get file path and line")
	}

	fileAndLine := fmt.Sprintf("%s:%d", file, line)
	filename := util.TrimmedPath(fileAndLine)

	switch val := data.(type) {

	case CustomLog:
		filename = util.TrimmedPath(fmt.Sprintf("%s:%d", val.File, val.Line))
		message = val.Message

	}
	return fmt.Sprintf("%s\t%s\t%s\t%s", requestId, filename, keywords, message)
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(ctx context.Context, keywords string, value interface{}) {

	appLogger.Debug(getData(ctx, keywords, value))
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(ctx context.Context, keywords string, value interface{}) {
	appLogger.Info(getData(ctx, keywords, value))
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(ctx context.Context, keywords string, value interface{}) {
	appLogger.Warn(getData(ctx, keywords, value))
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(ctx context.Context, keywords string, value interface{}) {
	appLogger.Error(getData(ctx, keywords, value))
}
