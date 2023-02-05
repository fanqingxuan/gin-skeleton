package user

import (
	"context"
	"gin-skeleton/common/errorx"
	"gin-skeleton/svc"
	"gin-skeleton/svc/logx"
	"gin-skeleton/types"
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
		that.svcCtx.Redis.Get(req.Name).Val()
		that.svcCtx.CacheX.Set(req.Name, req.Message)
	}

	return &types.UserIndexReply{
		Message: req.Name + " say:" + message,
	}, nil
}
