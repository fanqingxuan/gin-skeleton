package handler

import (
	"gin-skeleton/common/errorx"
	"gin-skeleton/common/response"
	"gin-skeleton/handler/user"
	"gin-skeleton/svc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine, svcCtx *svc.ServiceContext) {

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, response.NewCodeError(ctx, http.StatusNotFound, errorx.New("页面不存在")))
		return
	})
	r.GET("/user", user.IndexHandler(svcCtx))

	r.GET("/userinfo", user.InfoHandler(svcCtx))

	r.GET("/adduser", user.AddHandler(svcCtx))

}
