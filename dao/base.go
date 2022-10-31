package dao

import (
	"context"
	"gin-skeleton/svc"

	"gorm.io/gorm"
)

type Dao struct {
	db  *gorm.DB
	ctx context.Context
}

func NewDao(ctx context.Context, svcCtx *svc.ServiceContext) *Dao {
	return &Dao{
		db:  svcCtx.DB.WithContext(ctx),
		ctx: ctx,
	}
}
