package user

import (
	"gin-skeleton/common/errorx"
	"gin-skeleton/common/response"
	"gin-skeleton/logic/user"
	"gin-skeleton/svc"
	"gin-skeleton/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req types.UserIndexReq
		if err := ctx.ShouldBind(&req); err != nil {
			svcCtx.Log.WithContext(ctx).Error("Handler ShouldBind Parse", err)
			ctx.JSON(http.StatusOK, response.NewDefaultError(ctx, errorx.New(err.Error())))
			return
		}

		logic := user.NewIndexLogic(ctx, svcCtx)
		resp, err := logic.Say(&req)
		if err != nil {
			ctx.JSON(http.StatusOK, response.NewDefaultError(ctx, err))
		} else {
			ctx.JSON(http.StatusOK, response.NewDefault(ctx, resp))
		}
	}
}
