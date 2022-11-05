package main

import (
	"fmt"
	"gin-skeleton/config"
	"gin-skeleton/handler"
	"gin-skeleton/middleware"
	"gin-skeleton/svc"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	time.LoadLocation("")
	var c config.Config
	svc.MustLoad("./.env.yaml", &c)
	if c.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.New()
	ctx := svc.NewServiceContext(c)
	middleware.RegisterGlobalMiddlewares(server)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at localhost%s...\n", c.App.Port)
	server.Run(":8000")
}
