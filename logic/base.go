package logic

import (
	"context"
	"gin-skeleton/svc"
)

type Logic struct {
	Ctx   context.Context
	Svc   *svc.ServiceContext
	Redis *svc.AWRedis
	Log   *svc.Log
}

func NewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Logic {
	return &Logic{
		Ctx:   ctx,
		Svc:   svcCtx,
		Redis: svc.NewRedis(ctx, svcCtx.Redis),
		Log:   svcCtx.Log.WithContext(ctx),
	}
}
