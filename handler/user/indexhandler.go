package user

import (
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
			ctx.JSON(http.StatusOK, svcCtx.Response.WithContext(ctx).Error(err.Error()))
			return
		}

		indexLogic := user.NewIndexLogic(svcCtx)
		resp, err := indexLogic.Say(&req)
		if err != nil {
			ctx.JSON(http.StatusOK, svcCtx.Response.WithContext(ctx).Error(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, svcCtx.Response.WithContext(ctx).Success(resp))
		}
	}
}
