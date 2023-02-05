package handler

import (
	"gin-skeleton/common/responsex"
	"gin-skeleton/handler/user"
	"gin-skeleton/svc"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine, svcCtx *svc.ServiceContext) {

	r.NoMethod(func(ctx *gin.Context) {
		responsex.NoMethod(ctx)
		return
	})
	r.NoRoute(func(ctx *gin.Context) {
		responsex.NoRoute(ctx)
		return
	})

	r.GET("/user", user.IndexHandler(svcCtx))

	r.GET("/userinfo", user.InfoHandler(svcCtx))

	r.GET("/adduser", user.AddHandler(svcCtx))

}
