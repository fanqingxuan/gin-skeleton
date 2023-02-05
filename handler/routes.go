package handler

import (
	"gin-skeleton/common/errorx"
	"gin-skeleton/common/responsex"
	"gin-skeleton/handler/user"
	"gin-skeleton/svc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine, svcCtx *svc.ServiceContext) {

	r.NoMethod(func(ctx *gin.Context) {
		responsex.New(ctx, errorx.NewCodeError(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed)))
		return
	})
	r.NoRoute(func(ctx *gin.Context) {
		responsex.New(ctx, errorx.NewCodeError(http.StatusNotFound, http.StatusText(http.StatusNotFound)))
		return
	})
	r.GET("/user", user.IndexHandler(svcCtx))

	r.GET("/userinfo", user.InfoHandler(svcCtx))

	r.GET("/adduser", user.AddHandler(svcCtx))

}
