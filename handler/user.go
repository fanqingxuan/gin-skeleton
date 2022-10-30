package handler

import (
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
			ctx.JSON(http.StatusOK, svcCtx.Response.Error(err.Error()))
			return
		}

		userLogic := logic.NewUserLogic(ctx, svcCtx)
		resp, err := userLogic.GetUserInfo(&req)
		if err != nil {
			ctx.JSON(http.StatusOK, svcCtx.Response.Error(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, svcCtx.Response.Success(resp))
		}
	}
}
