package cron

import (
	"context"
	"gin-skeleton/svc"
)

type RemoveExpiredCacheKey struct {
}

func NewRemoveExpiredCacheKey() *RemoveExpiredCacheKey {
	return &RemoveExpiredCacheKey{}
}

func (that *RemoveExpiredCacheKey) Run(ctx context.Context, svcCtx *svc.ServiceContext) {
	svcCtx.CacheX.DeleteExpired()
}
