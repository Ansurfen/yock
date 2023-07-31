// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"fmt"
	"net/http"
	"os"

	yockc "github.com/ansurfen/yock/cmd"
	"github.com/ansurfen/yock/util"
	"github.com/gin-gonic/gin"
)

type YockRoutes struct{}

func (router *YockRoutes) InstallYockAPI(group *gin.RouterGroup) {
	yockRouter := group.Group("/yock")
	{
		yockRouter.GET("/version", yockVersion)
		yockRouter.POST("/eval", yockEval)
	}
}

// @Summary Get Yock Version
// @Description Get Yock Version
// @Tags Yock
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Router /yock/version [get]
func yockVersion(ctx *gin.Context) {
	v, err := util.Exec("yock", "version")
	if err != nil {
		ctx.String(http.StatusInternalServerError, "fail to get yock's version")
	}
	ctx.String(http.StatusOK, string(v))
}

// @Summary Eval Yock Script
// @Description Eval Yock Script
// @Tags Yock
// @Accept json
// @Produce json
// @Param script formData string true "Script"
// @Success 200 {string} Success
// @Router /yock/eval [post]
func yockEval(ctx *gin.Context) {
	script := ctx.PostForm("script")
	file, err := os.CreateTemp("", "*.lua")
	if err != nil {
		ctx.String(http.StatusInternalServerError, "server internal error")
	}
	_, err = file.Write([]byte(script))
	if err != nil {
		ctx.String(http.StatusInternalServerError, "server internal error")
	}
	file.Close()
	res, err := yockc.Exec(yockc.ExecOpt{Quiet: false}, fmt.Sprintf("yock run %s", file.Name()))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
			"msg": res,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"err": nil,
		"msg": res,
	})
}
