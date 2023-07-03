// MIT License

// Copyright (c) 2023 Ansurfen

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"fmt"
	"github.com/ansurfen/yock/ctl/cmd"
	"github.com/ansurfen/yock/ctl/conf"
	yocke "github.com/ansurfen/yock/env"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

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

	env := yocke.InitEnv(&yocke.EnvOpt[conf.YockConf]{
		Workdir:  ".yock",
		Subdirs:  []string{"log", "mnt", "unmnt", "tmp"},
		Conf:     conf.YockConf{},
		ConfTmpl: fmt.Sprintf(conf.YockConfTmpl, util.WorkSpace),
	})

	// Initialize each path for the global workspace
	util.WorkSpace = filepath.ToSlash(path.Join(env.User().HomeDir, ".yock"))
	util.PluginPath = path.Join(util.WorkSpace, "plugin")
	util.DriverPath = path.Join(util.WorkSpace, "driver")

	// determines YockPath according to YockBuild and exfPath.
	// Details to see /util/meta.go
	exfPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	if util.YockBuild == "dev" {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		exfPath = wd
	}
	util.YockPath = filepath.Join(exfPath, "..")

	conf := env.Conf()
	zlog, err := ycho.NewZLog(ycho.YchoOpt{
		Compress:    conf.Ycho.Compress,
		Path:        conf.Ycho.Path,
		FileName:    conf.Ycho.FileName,
		Level:       conf.Ycho.Level,
		FileMaxSize: conf.Ycho.FileMaxSize,
		Stdout:      conf.Ycho.Stdout,
	})
	ycho.SetYcho(zlog)
	if err != nil {
		panic(err)
	}
}

func main() {
	cmd.Execute()
}
