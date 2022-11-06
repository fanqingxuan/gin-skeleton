package cron

import (
	"context"
	"fmt"
	"gin-skeleton/dao"
	"gin-skeleton/svc"
)

func RegisterCronJobs(sctx *svc.ServiceContext) {

	svcCtx := sctx.WithLog(svc.NewLog("cron/", sctx.Config.Level, svc.BusinessLogType))

	recoverLog := svc.NewLog("cron_panic/", sctx.Config.Level, svc.PanicLogType)
	log := svc.NewLog("cron_request/", sctx.Config.Level, svc.PanicLogType)
	c := svc.NewCron(log, recoverLog)
	c.AddFunc("@every 1s", func(ctx context.Context) {
		userDao := dao.NewUserDao(ctx, svcCtx)
		userDao.GetUserInfo(3)
	})
	c.AddFunc("@hourly", func(ctx context.Context) {
		userDao := dao.NewUserDao(ctx, svcCtx)
		userDao.GetUserInfo(2)
	})

	c.AddFunc("@every 1s", func(ctx context.Context) {
		fmt.Println(ctx.Value("traceId"))
	})

	c.AddJob("@every 1s", NewSampleJob(svcCtx))

	c.Start()
}
