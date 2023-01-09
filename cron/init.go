package cron

import (
	"context"
	"gin-skeleton/svc"
)

func RegisterCronJobs(svcCtx *svc.ServiceContext) {

	c := svc.NewCron(svcCtx)

	c.AddJob("@every 10s", NewRemoveExpiredCacheKey())

	c.AddFunc("@every 1s", func(ctx context.Context, svcCtx *svc.ServiceContext) {

	})

	c.AddFunc("@every 1s", SampleFunc)

	// c.AddJob("@every 1s", NewSampleJob(svcCtx))

	c.Start()
}
