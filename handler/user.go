package handler

import (
	"fmt"
	"gin-skeleton/logic"
	"gin-skeleton/svc"
	"gin-skeleton/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserIndexHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req types.UserIndexReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusOK, svcCtx.Response.Error(err.Error()))
			return
		}

		userLogic := logic.NewUserLogic(ctx, svcCtx)
		resp, err := userLogic.Say(&req)
		if err != nil {
			ctx.JSON(http.StatusOK, svcCtx.Response.Error(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, svcCtx.Response.Success(resp))
		}
	}
}

func UserInfoHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserInfoReq
		if err := ctx.ShouldBind(&req); err != nil {
			svcCtx.Log.WithContext(ctx).Error("Parse Error", fmt.Sprintf("%+v", err))
			ctx.JSON(http.StatusOK, svcCtx.Response.WithContext(ctx).Error(err.Error()))
			return
		}

		userLogic := logic.NewUserLogic(ctx, svcCtx)
		resp, err := userLogic.GetUserInfo(&req)
		if err != nil {
			svcCtx.Log.WithContext(ctx).Error("err", fmt.Sprintf("%+v", err))
			ctx.JSON(http.StatusOK, svcCtx.Response.WithContext(ctx).Error(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, svcCtx.Response.WithContext(ctx).Success(resp))
		}
	}
}
