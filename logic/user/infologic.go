package user

import (
	"context"
	"fmt"
	"gin-skeleton/model"
	"gin-skeleton/svc"
	"gin-skeleton/svc/logx"
	"gin-skeleton/types"
	"time"
)

type InfoLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	UserModel model.UserModel
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		svcCtx:    svcCtx,
		ctx:       ctx,
		Logger:    logx.WithContext(ctx),
		UserModel: model.NewUserModel(ctx, svcCtx.Mysql),
	}
}

func (that *InfoLogic) Handle(req *types.UserInfoReq) (resp *types.UserInfoReply, err error) {
	that.svcCtx.Redis.Expire("test", time.Nanosecond/1000)
	that.Logger.Debug("这是debug消息")
	that.Logger.Info("这是info消息")
	that.Logger.Warn("这是warn消息")
	that.Logger.Error(req)

	user, err := that.UserModel.List(100)
	if err != nil {
		return
	}
	fmt.Println(user)

	if user == nil {
		resp = &types.UserInfoReply{
			Message: "数据不存在",
		}
		return
	}
	resp = &types.UserInfoReply{
		Message: "hello",
	}
	return

}
