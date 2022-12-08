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

		svc := svcCtx.WithContext(ctx)
		var req types.UserIndexReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusOK, svc.Response.Error(err.Error()))
			return
		}

		indexLogic := user.NewIndexLogic(svcCtx)
		resp, err := indexLogic.Say(&req)
		if err != nil {
			ctx.JSON(http.StatusOK, svc.Response.Error(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, svc.Response.Success(resp))
		}
	}
}
