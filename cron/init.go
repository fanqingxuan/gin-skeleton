package cron

import (
	"gin-skeleton/dao"
	"gin-skeleton/svc"
)

func RegisterCronJobs(svcCtx *svc.ServiceContext) {

	c := svc.NewCron(svcCtx)

	c.AddJob("@every 10s", NewRemoveExpiredCacheKey())

	c.AddFunc("@every 1s", func(svcCtx *svc.ServiceContext) {
		userDao := dao.NewUserDao(svcCtx.DB)
		userDao.GetUserInfo(2)
	})

	c.AddFunc("@every 1s", SampleFunc)

	// c.AddJob("@every 1s", NewSampleJob(svcCtx))

	c.Start()
}
