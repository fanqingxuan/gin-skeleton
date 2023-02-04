package user

import (
	"context"
	"gin-skeleton/model"
	"gin-skeleton/svc"
	"gin-skeleton/svc/logx"
	"gin-skeleton/types"
)

type AddLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	UserModel model.UserModel
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		svcCtx:    svcCtx,
		ctx:       ctx,
		Logger:    logx.WithContext(ctx),
		UserModel: model.NewUserModel(ctx, svcCtx.Mysql),
	}
}

func (that *AddLogic) Handle(req *types.UserAddReq) (resp *types.UserAddReply, err error) {

	userid, err := that.UserModel.Insert(&model.User{
		Username: req.Username,
		Age:      req.Age,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserAddReply{
		UserId: userid,
	}, nil

}
