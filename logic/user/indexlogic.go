package user

import (
	"gin-skeleton/logic"
	"gin-skeleton/svc"
	"gin-skeleton/types"
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
	val, found := that.LocalStorage.Get(req.Name)
	var message string
	if found {
		message = val.(string) + " get value from cache"
	} else {
		message = req.Message
		that.LocalStorage.Set(req.Name, req.Message)
	}

	return &types.UserIndexReply{
		Message: req.Name + " say:" + message,
	}, nil
}
