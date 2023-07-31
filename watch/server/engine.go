package server

import (
	"fmt"

	"github.com/ansurfen/yock/watch/docs"
	"github.com/ansurfen/yock/watch/routes"
	"github.com/gin-gonic/gin"
)

type YockwEngine struct {
	srv *gin.Engine
}

func New() *YockwEngine {
	gin.SetMode(gin.ReleaseMode)

	return &YockwEngine{
		srv: gin.Default(),
	}
}

func (engine *YockwEngine) UseRouter() *YockwEngine {
	docs.SwaggerInfo.BasePath = "/"

	baseRouter := engine.srv.Group("")

	routes.YockwRouter.InstallInternalAPI(baseRouter)
	routes.YockwRouter.InstallMetricsAPI(baseRouter)
	routes.YockwRouter.InstallLoggerAPI(baseRouter)
	routes.YockwRouter.InstallYockAPI(baseRouter)
	routes.YockwRouter.InstallYockdAPI(baseRouter)
	routes.YockwRouter.InstallYockiAPI(baseRouter)

	return engine
}

func (engine *YockwEngine) Run(port int) {
	engine.srv.Run(fmt.Sprintf(":%d", port))
}
