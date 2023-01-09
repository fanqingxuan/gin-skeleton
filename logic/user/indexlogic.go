package user

import (
	"context"
	"gin-skeleton/common/errorx"
	"gin-skeleton/logic"
	"gin-skeleton/svc"
	"gin-skeleton/types"
)

type IndexLogic struct {
	logic.Logic
	Log *svc.Log
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Log:   svcCtx.Log.WithContext(ctx),
		Logic: *logic.NewLogic(svcCtx),
	}
}

func (that *IndexLogic) Handle(req *types.UserIndexReq) (resp *types.UserIndexReply, err error) {
	val, found := that.CacheX.Get(req.Name)
	var message string
	found = false
	if found {
		message = " get value from cache:" + val.(string)
	} else {
		if req.Message == "" {
			return nil, errorx.New("消息不能为空")
		}
		that.Redis.Get(req.Name).Val()
		that.CacheX.Set(req.Name, req.Message)
	}

	return &types.UserIndexReply{
		Message: req.Name + " say:" + message,
	}, nil
}
