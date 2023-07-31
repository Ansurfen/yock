// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ansurfen/yock/watch/log"
	"github.com/gin-gonic/gin"
)

type LoggerRoutes struct{}

func (router *LoggerRoutes) InstallLoggerAPI(group *gin.RouterGroup) {
	loggerRouter := group.Group("/logger")
	{
		loggerRouter.GET("/parse", loggerParse)
		loggerRouter.GET("/find", loggerFind)
	}
}

// @Summary Get API Document
// @Description Get API Document
// @Tags Logger
// @Accept json
// @Produce json
// @Param path query string true "Path of logger"
// @Success 200 {string} Success
// @Router /logger/parse [get]
func loggerParse(ctx *gin.Context) {
	path, ok := ctx.GetQuery("path")
	if !ok {
		ctx.String(http.StatusBadRequest, "invalid path")
	}
	err := log.DefaultLoggerWatch.Parse(path)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "server internal error")
	}
	ctx.String(http.StatusOK, "success to parse")
}

// @Summary Get API Document
// @Description Get API Document
// @Tags Logger
// @Accept json
// @Produce json
// @Param file query string false "File"
// @Param time query string false "Time"
// @Param level query string false "Level"
// @Param caller query string false "Caller"
// @Param msg query string false "Message"
// @Success 200 {string} Success
// @Router /logger/find [get]
func loggerFind(ctx *gin.Context) {
	file, ok := ctx.GetQuery("file")
	if !ok {
		file = "*"
	}
	time, ok := ctx.GetQuery("time")
	if !ok {
		time = "*"
	}
	level, ok := ctx.GetQuery("level")
	if !ok {
		level = "*"
	}
	caller, ok := ctx.GetQuery("caller")
	if !ok {
		caller = "*"
	}
	msg, ok := ctx.GetQuery("msg")
	if !ok {
		msg = "*"
	}
	// TODO
	// ctx.GetQuery("limit")
	// ctx.GetQuery("page")
	entries := log.DefaultLoggerWatch.Find(
		file, time, level, caller, msg)
	data, err := json.Marshal(entries)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "server internal error")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"err":  nil,
		"msg":  "find successfully",
		"data": data,
	})
}
