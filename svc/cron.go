package svc

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	uuid "github.com/satori/go.uuid"
)

type Cron struct {
	cron   *cron.Cron
	logger *Log
}

type JobFunc func(context.Context)

func (f JobFunc) Run(ctx context.Context) { f(ctx) }

type Job interface {
	Run(context.Context)
}

func NewCron(logger, recoverlogger *Log) *Cron {
	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.Recover(NewCronLogger(recoverlogger.WithContext(context.Background()))),
		),
	)
	return &Cron{
		cron:   c,
		logger: logger,
	}
}

func (that *Cron) AddFunc(spec string, cmd JobFunc) (cron.EntryID, error) {
	return that.AddJob(spec, cmd)
}

func (that *Cron) AddJob(spec string, cmd Job) (cron.EntryID, error) {
	return that.cron.AddFunc(spec, func() {
		c := context.WithValue(context.Background(), "traceId", uuid.NewV4())
		cmd.Run(c)
	})
}

func (that *Cron) Start() {
	that.cron.Start()
}

type CronLogger struct {
	logger *Log
}

func NewCronLogger(logger *Log) *CronLogger {
	return &CronLogger{
		logger: logger,
	}
}

func (that *CronLogger) Info(msg string, keysAndValues ...interface{}) {
	fmt.Println(msg)
}

func (that *CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	that.logger.Error(msg, fmt.Sprintf("%+v", errors.WithStack(err)))
}

func formatString(numKeysAndValues int) string {
	var sb strings.Builder
	sb.WriteString("%s")
	if numKeysAndValues > 0 {
		sb.WriteString(", ")
	}
	for i := 0; i < numKeysAndValues/2; i++ {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("%v=%v")
	}
	return sb.String()
}
