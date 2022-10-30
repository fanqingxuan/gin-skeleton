package logic

import (
	"context"
	"gin-skeleton/svc"
)

type Logic struct {
	ctx   context.Context
	svc   *svc.ServiceContext
	redis *svc.AWRedis
}

func NewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Logic {
	return &Logic{
		ctx:   ctx,
		svc:   svcCtx,
		redis: svc.NewRedis(ctx, svcCtx.Redis),
	}
}
