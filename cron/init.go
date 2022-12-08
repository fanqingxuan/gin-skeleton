package cron

import (
	"context"
	"gin-skeleton/dao"
	"gin-skeleton/svc"
)

func RegisterCronJobs(sctx *svc.ServiceContext) {

	svcCtx := sctx.WithLog(svc.NewLog("cron/", sctx.Config.Level, svc.BusinessLogType))
	recoverLog := svc.NewLog("cron_panic/", sctx.Config.Level, svc.PanicLogType)
	c := svc.NewCron(svcCtx.Log, recoverLog)

	c.AddJob("@every 10s", NewRemoveExpiredCacheKey(sctx))

	c.AddFunc("@every 1s", func(ctx context.Context) {
		userDao := dao.NewUserDao(svcCtx.WithContext(ctx).DB)
		userDao.GetUserInfo(2)
	})

	c.AddFunc("@every 1s", SampleFunc)

	// c.AddJob("@every 1s", NewSampleJob(svcCtx))

	c.Start()
}
