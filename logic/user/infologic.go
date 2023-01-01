package user

import (
	"gin-skeleton/logic"
	"gin-skeleton/model"
	"gin-skeleton/svc"
	"gin-skeleton/types"
	"time"
)

type InfoLogic struct {
	*logic.Logic
	UserModel *model.UserModel
}

func NewInfoLogic(svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logic:     logic.NewLogic(svcCtx),
		UserModel: model.NewUserModel(svcCtx.DB),
	}
}

func (that *InfoLogic) GetUserInfo(req *types.UserInfoReq) (resp *types.UserInfoReply, err error) {
	that.Redis.Expire("test", time.Nanosecond/1000)
	that.Log.Debug("debug 关键字", "这是debug消息")
	that.Log.Info("info 关键字", "这是info消息")
	that.Log.Warn("warn 关键字", "这是warn消息")
	that.Log.Error("error 关键字", req)

	user, err := that.UserModel.FindOne(req.UserId)
	if err != nil {
		return
	}

	if user == nil {
		resp = &types.UserInfoReply{
			Message: "no data",
		}
		return
	}
	resp = &types.UserInfoReply{
		Message: user.Username,
	}
	return

}
