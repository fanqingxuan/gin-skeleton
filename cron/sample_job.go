package cron

import (
	"context"
	"fmt"
	"gin-skeleton/svc"
)

type SampleJob struct {
}

func NewSampleJob() *SampleJob {
	return &SampleJob{}
}

func (that *SampleJob) Run(ctx context.Context, svcCtx *svc.ServiceContext) {
	fmt.Println(ctx.Value("traceId"))
}
