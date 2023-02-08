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

func AddHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req types.UserAddReq
		if err := ctx.ShouldBind(&req); err != nil {
			logx.WithContext(ctx).Errorf("Handler ShouldBind Parse %v", err)
			responsex.Result(ctx, errorx.NewCodeError(1, err.Error()))

			return
		}

		logic := user.NewAddLogic(ctx, svcCtx)
		resp, err := logic.Handle(&req)
		if err != nil {
			responsex.Result(ctx, err)
		} else {
			responsex.Result(ctx, resp)
		}
	}
}
