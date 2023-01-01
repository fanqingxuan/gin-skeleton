package logic

import (
	"gin-skeleton/svc"
)

type Logic struct {
	Redis  *svc.AWRedis
	Log    *svc.Log
	XCache *svc.XCache
}

func NewLogic(svcCtx *svc.ServiceContext) *Logic {
	return &Logic{
		Redis:  svcCtx.Redis,
		Log:    svcCtx.Log,
		XCache: svcCtx.XCache,
	}
}
