package cron

import (
	"fmt"
	"gin-skeleton/svc"
)

type SampleJob struct {
}

func NewSampleJob() *SampleJob {
	return &SampleJob{}
}

func (that *SampleJob) Run(svcCtx *svc.ServiceContext) {
	fmt.Println(svcCtx.Ctx.Value("traceId"))
}
