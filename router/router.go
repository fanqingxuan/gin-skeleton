package router

import (
	"gin-demo/controller"

	"github.com/gin-gonic/gin"
)

//LoadRoutes load the route
func LoadRoutes(router *gin.Engine) {

	router.NoRoute(controller.NOTFOUND)
	router.GET("/", controller.Index)
	router.GET("/create", controller.Create)
	router.GET("/test", controller.TestRedis)
	router.GET("/testlog", controller.TestLog)
	router.GET("/find", controller.FindUser)
}
