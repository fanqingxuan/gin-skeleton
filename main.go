package main

import (
	"flag"
	"fmt"
	"gin-skeleton/config"
	"gin-skeleton/handler"
	"gin-skeleton/middleware"
	"gin-skeleton/svc"
	"gin-skeleton/svc/logx"

	"github.com/gin-gonic/gin"
)

var configFile = flag.String("f", "./.env.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	svc.MustLoad(*configFile, &c)
	if c.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.New()
	server.HandleMethodNotAllowed = true
	svcCtx := svc.NewServiceContext(c)
	middleware.RegisterGlobalMiddlewares(server)
	handler.RegisterHandlers(server, svcCtx)
	fmt.Printf("Starting server at localhost%s...\n", c.App.Port)

	if err := server.Run(c.App.Port); err != nil {
		fmt.Printf("Start server error,err=%v", err)
		logx.Errorf("Start server error:%v", err)
	}

}
