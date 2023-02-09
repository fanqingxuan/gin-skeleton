package handler

import (
	"gin-skeleton/common/responsex"
	"gin-skeleton/handler/user"
	"gin-skeleton/svc"
	"net/http"

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

	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, "<b>Hello World</b>")
	})
	r.GET("/user", user.IndexHandler(svcCtx))

	r.GET("/userinfo", user.InfoHandler(svcCtx))

	r.GET("/adduser", user.AddHandler(svcCtx))

}
