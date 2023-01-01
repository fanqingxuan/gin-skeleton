package main

import (
	"fmt"
	"gin-skeleton/config"
	"gin-skeleton/handler"
	"gin-skeleton/middleware"
	"gin-skeleton/svc"

	"github.com/gin-gonic/gin"
)

func main() {
	var c config.Config
	svc.MustLoad("./.env.yaml", &c)
	if c.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.New()
	svcCtx := svc.NewServiceContext(c)
	middleware.RegisterGlobalMiddlewares(server)
	handler.RegisterHandlers(server, svcCtx)
	// cron.RegisterCronJobs(svcCtx)

	fmt.Printf("Starting server at localhost%s...\n", c.App.Port)

	server.Run(":8000")
}
