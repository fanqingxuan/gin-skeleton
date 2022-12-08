package user

import (
	"fmt"
	"gin-skeleton/logic/user"
	"gin-skeleton/svc"
	"gin-skeleton/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InfoHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		svc := svcCtx.WithContext(ctx)
		var req types.UserInfoReq
		if err := ctx.ShouldBind(&req); err != nil {
			svc.Log.Error("Parse Error", fmt.Sprintf("%+v", err))
			ctx.JSON(http.StatusOK, svc.Response.Error(err.Error()))
			return
		}

		userLogic := user.NewInfoLogic(svc)
		resp, err := userLogic.GetUserInfo(&req)
		if err != nil {
			svc.Log.Error("err", fmt.Sprintf("%+v", err))
			ctx.JSON(http.StatusOK, svc.Response.Error(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, svc.Response.Success(resp))
		}
	}
}
