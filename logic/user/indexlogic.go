package user

import (
	"context"
	"fmt"
	"gin-skeleton/common/errorx"
	"gin-skeleton/svc"
	"gin-skeleton/svc/logx"
	"gin-skeleton/types"
	"time"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		svcCtx: svcCtx,
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
	}
}

func (that *IndexLogic) Handle(req *types.UserIndexReq) (resp *types.UserIndexReply, err error) {
	val, found := that.svcCtx.CacheX.Get(req.Name)
	var message string
	found = false
	if found {
		message = " get value from cache:" + val.(string)
	} else {
		if req.Message == "" {
			return nil, errorx.NewDefaultError("消息不能为空")
		}
		result, err := that.svcCtx.Redis.SetCtx(that.ctx, "name", "测试命令行344", time.Hour)
		fmt.Println(result, err)
		that.svcCtx.CacheX.Set(req.Name, req.Message)
	}

	return &types.UserIndexReply{
		Message: req.Name + " say:" + message,
	}, nil
}
