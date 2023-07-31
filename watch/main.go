// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"path/filepath"

	"github.com/ansurfen/yock/ctl/conf"
	yocke "github.com/ansurfen/yock/env"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/watch/log"
	"github.com/ansurfen/yock/watch/metrics"
	"github.com/ansurfen/yock/watch/server"
)

// @BasePath /
// @title YockWatch
// @version 1.0
// @description Document: ansurfen.github.io/YockNav/

func init() {
	metrics.DefaultMetricsWatch = metrics.New()
	log.DefaultLoggerWatch = log.New()
}

var port = flag.Int("p", 0, "")

func main() {
	flag.Parse()
	if *port == 0 {
		panic("invalid port")
	}
	env := yocke.InitEnv(&yocke.EnvOpt[conf.YockConf]{
		Workdir: ".yock",
		Subdirs: []string{"log", "mnt", "unmnt"},
		Conf:    conf.YockConf{},
	})
	util.WorkSpace = filepath.ToSlash(filepath.Join(env.User().HomeDir, ".yock"))
	server.New().UseRouter().Run(*port)
}
