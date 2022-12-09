package cron

import (
	"gin-skeleton/svc"
)

type RemoveExpiredCacheKey struct {
}

func NewRemoveExpiredCacheKey() *RemoveExpiredCacheKey {
	return &RemoveExpiredCacheKey{}
}

func (that *RemoveExpiredCacheKey) Run(svcCtx *svc.ServiceContext) {
	svcCtx.LocalStorage.DeleteExpired()
}
