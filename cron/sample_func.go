package cron

import (
	"context"
	"fmt"
)

func SampleFunc(ctx context.Context) {
	fmt.Println(ctx.Value("traceId"))
}
