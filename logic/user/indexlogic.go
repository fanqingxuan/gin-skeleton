package user

import (
	"errors"
	"fmt"
	"gin-skeleton/common/errorx"
	"gin-skeleton/logic"
	"gin-skeleton/svc"
	"gin-skeleton/types"

	"github.com/go-redis/redis/v8"
)

type IndexLogic struct {
	logic.Logic
}

func NewIndexLogic(svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logic: *logic.NewLogic(svcCtx),
	}
}

func (that *IndexLogic) Say(req *types.UserIndexReq) (resp *types.UserIndexReply, err error) {
	val, found := that.XCache.Get(req.Name)
	var message string
	found = false
	if found {
		message = " get value from cache:" + val.(string)
	} else {
		if req.Message == "" {
			return nil, errorx.New("消息不能为空")
		}
		result := that.Redis.Get(req.Name)
		if err = result.Err(); err != nil && err != redis.Nil {
			that.Log.Error("获取cache错误key:"+req.Name, err)
			return
		}
		fmt.Println(errors.Is(err, redis.Nil))

		message = result.Val()
		that.XCache.Set(req.Name, req.Message)
	}

	return &types.UserIndexReply{
		Message: req.Name + " say:" + message,
	}, nil
}
