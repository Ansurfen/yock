// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/ansurfen/yock/watch/metrics"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type InternalRoutes struct{}

func (router *InternalRoutes) InstallInternalAPI(group *gin.RouterGroup) {
	internalRouter := group.Group("/internal")
	{
		internalRouter.GET("/swagger/*any", Swagger())
		internalRouter.GET("/metrics", Metrics())
	}
}

// @Summary Get API Document
// @Description Get API Document
// @Tags System
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Router /internal/swagger/index.html [get]
func Swagger() gin.HandlerFunc {
	return ginSwagger.WrapHandler(swaggerfiles.Handler)
}

// @Summary Get Metrics
// @Description Get Metrics
// @Tags System
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Router /internal/metrics [get]
func Metrics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		metrics.DefaultMetricsWatch.Document().ServeHTTP(ctx.Writer, ctx.Request)
	}
}
