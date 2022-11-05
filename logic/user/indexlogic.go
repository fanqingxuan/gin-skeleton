package user

import (
	"context"
	"gin-skeleton/logic"
	"gin-skeleton/svc"
	"gin-skeleton/types"
)

type IndexLogic struct {
	logic.Logic
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logic: *logic.NewLogic(ctx, svcCtx),
	}
}

func (that *IndexLogic) Say(req *types.UserIndexReq) (resp *types.UserIndexReply, err error) {

	return &types.UserIndexReply{
		Message: req.Name + "," + req.Message,
	}, nil
}
