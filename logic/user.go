package logic

import (
	"context"
	"gin-skeleton/dao"
	"gin-skeleton/svc"
	"gin-skeleton/types"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type UserLogic struct {
	Logic
	userDao *dao.UserDao
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logic:   *NewLogic(ctx, svcCtx),
		userDao: dao.NewUserDao(ctx, svcCtx),
	}
}

func (that *UserLogic) Say(req *types.UserIndexReq) (resp *types.UserIndexReply, err error) {

	return &types.UserIndexReply{
		Message: req.Name + "," + req.Message,
	}, nil
}

func (that *UserLogic) GetUserInfo(req *types.UserInfoReq) (resp *types.UserInfoReply, err error) {

	result := that.redis.Get("user:" + string(req.UserId))
	that.log.Debug("debug 关键字", "这是debug消息")
	that.log.Info("info 关键字", "这是info消息")
	that.log.Warn("warn 关键字", "这是warn消息")
	that.log.Error("error 关键字", req)
	if err := result.Err(); err != nil && err != redis.Nil {
		return nil, errors.WithStack(err)
	}
	if result.Val() == "" {
		user, err := that.userDao.GetUserInfo(req.UserId)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if user != nil {
			return &types.UserInfoReply{
				Message: "hello:" + user.Username,
			}, nil
		} else {
			return &types.UserInfoReply{
				Message: "default value",
			}, nil
		}
	} else {
		return &types.UserInfoReply{
			Message: result.Val(),
		}, nil
	}

}
