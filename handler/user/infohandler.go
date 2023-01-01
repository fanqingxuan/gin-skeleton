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

func InfoHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserInfoReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusOK, response.NewDefaultError(ctx, errorx.New(err.Error())))
			return
		}

		userLogic := user.NewInfoLogic(ctx, svcCtx)
		resp, err := userLogic.GetUserInfo(&req)
		if err != nil {
			ctx.JSON(http.StatusOK, response.NewDefaultError(ctx, err))
		} else {
			ctx.JSON(http.StatusOK, response.NewDefault(ctx, resp))
		}
	}
}
