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
	svcCtx *ServiceContext
}

type JobFunc func(context.Context, *ServiceContext)

func (f JobFunc) Run(ctx context.Context, svcCtx *ServiceContext) { f(ctx, svcCtx) }

type Job interface {
	Run(context.Context, *ServiceContext)
}

func NewCron(svcCtx *ServiceContext) *Cron {
	ctx := svcCtx.WithLog(NewLog("cron/", svcCtx.Config.Level, BusinessLogType))
	recoverLog := NewLog("cron_panic/", svcCtx.Config.Level, PanicLogType)
	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.Recover(NewCronLogger(recoverLog.WithContext(context.Background()))),
		),
	)
	return &Cron{
		cron:   c,
		svcCtx: ctx,
	}
}

func (that *Cron) AddFunc(spec string, cmd JobFunc) (cron.EntryID, error) {
	return that.AddJob(spec, cmd)
}

func (that *Cron) AddJob(spec string, cmd Job) (cron.EntryID, error) {
	return that.cron.AddFunc(spec, func() {
		ctx := context.WithValue(context.Background(), "traceId", uuid.NewV4())

		jobName := reflect.TypeOf(cmd).String()
		start := time.Now()
		that.svcCtx.Log.WithContext(ctx).Info("before", fmt.Sprintf("spec:%s, jobName:%s, start...", spec, jobName))
		cmd.Run(ctx, that.svcCtx)
		elapsed := time.Now().Sub(start).Seconds()
		that.svcCtx.Log.WithContext(ctx).Info("after", fmt.Sprintf("spec:%s, jobName:%s, finish,cost:%fs", spec, jobName, elapsed))
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
