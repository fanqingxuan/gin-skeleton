package user

import (
	"gin-skeleton/dao"
	"gin-skeleton/logic"
	"gin-skeleton/svc"
	"gin-skeleton/types"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type InfoLogic struct {
	*logic.Logic
	userDao *dao.UserDao
}

func NewInfoLogic(svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logic:   logic.NewLogic(svcCtx),
		userDao: dao.NewUserDao(svcCtx.DB),
	}
}

func (that *InfoLogic) GetUserInfo(req *types.UserInfoReq) (resp *types.UserInfoReply, err error) {
	that.Redis.Expire("test", time.Nanosecond/1000)
	result := that.Redis.Get("user:" + string(req.UserId))
	that.Log.Debug("debug 关键字", "这是debug消息")
	that.Log.Info("info 关键字", "这是info消息")
	that.Log.Warn("warn 关键字", "这是warn消息")
	that.Log.Error("error 关键字", req)
	if err := result.Err(); err != nil && !errors.Is(err, redis.Nil) {
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
