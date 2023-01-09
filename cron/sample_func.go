package cron

import (
	"context"
	"fmt"
	"gin-skeleton/svc"
)

func SampleFunc(ctx context.Context, svcCtx *svc.ServiceContext) {
	fmt.Println(ctx.Value("traceId"))
}
