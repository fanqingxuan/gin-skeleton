package handler

import (
	"gin-skeleton/svc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine, svcCtx *svc.ServiceContext) {

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, svcCtx.Response.ErrorWithCode(http.StatusNotFound, "页面不存在"))
		return
	})
	r.POST("/user", UserIndexHandler(svcCtx))

	r.GET("/userinfo", UserInfoHandler(svcCtx))

}
