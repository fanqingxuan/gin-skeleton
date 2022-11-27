package svc

import (
	"context"
	"fmt"
	"reflect"
	"time"

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

		logger := that.logger.WithContext(c)
		jobName := reflect.TypeOf(cmd).String()
		start := time.Now()
		logger.Info("before", fmt.Sprintf("spec:%s, jobName:%s, start...", spec, jobName))
		cmd.Run(c)
		elapsed := time.Now().Sub(start).Seconds()
		logger.Info("after", fmt.Sprintf("spec:%s, jobName:%s, finish,cost:%fs", spec, jobName, elapsed))
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
}

func (that *CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	that.logger.Error(msg, fmt.Sprintf("%+v", errors.WithStack(err)))
}
