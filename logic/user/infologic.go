package user

import (
	"context"
	"fmt"
	"gin-skeleton/common/errorx"
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
	that.Logger.Info("这是info消息", "这是info日志22")
	that.Logger.Warn("这是warn消息")
	that.Logger.Error(req)

	that.Logger.Debugf("debg测试 %s %d", "姓名", 43)
	that.Logger.Infof("info测试 %s %s", "姓名", 43)
	that.Logger.Warnf("warn测试 %s %s", "姓名", 43)
	that.Logger.Errorf("error测试 %s %s", "姓名", 43)
	user, err := that.UserModel.List(100)
	that.Logger.Info(user)
	if err != nil {
		return
	}
	fmt.Println(user)

	if user == nil {

		err = errorx.NewDefaultError("数据不存在")
		return
	}
	resp = &types.UserInfoReply{
		Message: "hello",
	}
	return

}
