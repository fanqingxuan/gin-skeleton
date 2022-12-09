package cron

import (
	"fmt"
	"gin-skeleton/svc"
)

func SampleFunc(svcCtx *svc.ServiceContext) {
	fmt.Println(svcCtx.Ctx.Value("traceId"))
}
