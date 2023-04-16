package handler

import (
	"fmt"
	"gin-skeleton/common/responsex"
	"gin-skeleton/handler/user"
	"gin-skeleton/svc"

	"gin-skeleton/common/utils"

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
		m := map[string]string{
			"message": "Hello World",
		}
		a, _ := utils.Encode(m)

		var mm map[string]string
		utils.Decode(a, &mm)
		fmt.Println(mm)
		c.String(200, string(a))
	})
	r.GET("/user", user.IndexHandler(svcCtx))

	r.GET("/userinfo", user.InfoHandler(svcCtx))

	r.GET("/adduser", user.AddHandler(svcCtx))

}
