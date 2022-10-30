package dao

import (
	"fmt"
	"gin-skeleton/svc"

	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB
}

func NewDao(svcCtx *svc.ServiceContext) *Dao {
	fmt.Println(svcCtx.DB)
	return &Dao{
		db: svcCtx.DB,
	}
}
