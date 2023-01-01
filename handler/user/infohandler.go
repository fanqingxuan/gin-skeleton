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
		var req types.UserInfoReq
		if err := ctx.ShouldBind(&req); err != nil {
			svcCtx.Log.WithContext(ctx).Error("Parse Error", fmt.Sprintf("%+v", err))
			ctx.JSON(http.StatusOK, svcCtx.Response.WithContext(ctx).Error(err.Error()))
			return
		}

		userLogic := user.NewInfoLogic(svcCtx)
		resp, err := userLogic.GetUserInfo(&req)
		if err != nil {
			svcCtx.Log.WithContext(ctx).Error("err", fmt.Sprintf("%+v", err))
			ctx.JSON(http.StatusOK, svcCtx.Response.WithContext(ctx).Error(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, svcCtx.Response.WithContext(ctx).Success(resp))
		}
	}
}
