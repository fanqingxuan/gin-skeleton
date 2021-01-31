package redis

import (
	"context"
	"gin-demo/core/logger"
	log "gin-demo/core/logger"
	"runtime"
	"strings"

	"github.com/go-redis/redis/v8"
)

type LogHook struct {
	redis.Hook
}

func (h *LogHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	return ctx, nil
}
func (h *LogHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	err := cmd.Err()
	if err != nil && err != redis.Nil {
		error_message := err.Error()
		_, file, line, _ := runtime.Caller(4)
		logData := log.CustomLog{
			File:    file,
			Line:    line,
			Message: cmd.String(),
		}
		if strings.Contains(strings.ToLower(error_message), "No connection") || strings.Contains(strings.ToLower(error_message), "connection reset by peer") || strings.Contains(strings.ToLower(error_message), "target machine actively refused") {
			panic(logData)
		} else {
			logger.Warn(ctx, "redis", logData)
		}

	}
	return nil
}

func (h *LogHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	return ctx, nil
}
func (h *LogHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	return nil
}
