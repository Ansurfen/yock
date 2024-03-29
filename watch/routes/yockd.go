// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import "github.com/gin-gonic/gin"

type YockdRoutes struct{}

func (router *YockdRoutes) InstallYockdAPI(group *gin.RouterGroup) {
	yockdRouter := group.Group("/yockd")
	{
		yockdRouter.GET("/ping")
	}
}
