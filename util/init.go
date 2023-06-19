// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/ansurfen/cushion/utils"
	"go.uber.org/zap"
)

const yockConfTmpl = `lang: en_us
ycho:
  level: debug
  compress: false
  filename: yock.log
  fileMaxSize: 1024
  fileMaxBackups: 0
  stdout: true
  path: %s/log
`

func init() {
	// panic is recovered to unify the information output format,
	// before YCHO has been initialized.
	defer func() {
		crash := recover()
		switch crash.(type) {
		case string, error:
			_, file, line, _ := runtime.Caller(2)
			caller := filepath.Base(file) + ":" + strconv.Itoa(line)
			fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 -0700"), "\033[31mFATAL\033[0m", caller, crash)
			os.Exit(1)
		default:
		}
	}()

	// Initialize each path for the global workspace
	WorkSpace = filepath.ToSlash(path.Join(utils.GetEnv().Workdir(), ".yock"))
	PluginPath = path.Join(WorkSpace, "plugin")
	DriverPath = path.Join(WorkSpace, "driver")

	// determines YockPath according to YockBuild and exfPath.
	// Details to see /util/meta.go
	exfPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	if YockBuild == "dev" {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		exfPath = wd
	}
	YockPath = filepath.Join(exfPath, "..")

	// pull the configuration file from the global workspace
	conf = utils.NewEnv(utils.EnvOpt[yockConf]{
		Workdir:   ".yock",
		Subdirs:   []string{"log", "mount", "unmount", "tmp"},
		BlankConf: fmt.Sprintf(yockConfTmpl, WorkSpace),
	})

	err = initYcho(YchoOpt{
		Compress:    conf.Ycho.Compress,
		Path:        conf.Ycho.Path,
		FileName:    conf.Ycho.FileName,
		Level:       conf.Ycho.Level,
		FileMaxSize: conf.Ycho.FileMaxSize,
		Stdout:      conf.Ycho.Stdout,
	})
	if err != nil {
		panic(err)
	}
	Ycho = zap.L()

	// init yock watch
	yockCpu = newCPU()
	yockMem = newMem()
	yockDisk = newDisk()
	yockHost = newHost()
	yockNet = newNet()

	center = &SSHCenter{}
}
