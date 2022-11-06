package cron

import (
	"context"
	"fmt"
	"gin-skeleton/svc"
)

type SampleJob struct {
	svcCtx *svc.ServiceContext
}

func NewSampleJob(svcCtx *svc.ServiceContext) *SampleJob {
	return &SampleJob{
		svcCtx: svcCtx,
	}
}

func (that *SampleJob) Run(ctx context.Context) {
	fmt.Println(ctx.Value("traceId"))
}
