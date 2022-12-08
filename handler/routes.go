package handler

import (
	"gin-skeleton/handler/user"
	"gin-skeleton/svc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine, svcCtx *svc.ServiceContext) {

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, svcCtx.WithContext(ctx).Response.ErrorWithCode(http.StatusNotFound, "页面不存在"))
		return
	})
	r.GET("/user", user.IndexHandler(svcCtx))

	r.GET("/userinfo", user.InfoHandler(svcCtx))

}
