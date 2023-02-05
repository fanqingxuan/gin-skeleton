package user

import (
	"gin-skeleton/common/errorx"
	"gin-skeleton/common/responsex"
	"gin-skeleton/logic/user"
	"gin-skeleton/svc"
	"gin-skeleton/svc/logx"
	"gin-skeleton/types"

	"github.com/gin-gonic/gin"
)

func IndexHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req types.UserIndexReq
		if err := ctx.ShouldBind(&req); err != nil {
			logx.WithContext(ctx).Error("Handler ShouldBind Parse")
			responsex.New(ctx, errorx.NewCodeError(1, err.Error()))
			return
		}

		logic := user.NewIndexLogic(ctx, svcCtx)
		resp, err := logic.Handle(&req)
		if err != nil {
			responsex.New(ctx, err)
		} else {
			responsex.New(ctx, resp)
		}
	}
}
