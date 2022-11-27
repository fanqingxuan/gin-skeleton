package cron

import (
	"context"
	"gin-skeleton/svc"
)

type RemoveExpiredCacheKey struct {
	svcCtx *svc.ServiceContext
}

func NewRemoveExpiredCacheKey(svcCtx *svc.ServiceContext) *RemoveExpiredCacheKey {
	return &RemoveExpiredCacheKey{
		svcCtx: svcCtx,
	}
}

func (that *RemoveExpiredCacheKey) Run(ctx context.Context) {
	that.svcCtx.LocalStorage.DeleteExpired()
}
