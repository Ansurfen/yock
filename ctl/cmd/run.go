// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"path/filepath"
	"sync"

	yockpack "github.com/ansurfen/yock/pack"
	yockr "github.com/ansurfen/yock/runtime"
	yocks "github.com/ansurfen/yock/scheduler"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
	"github.com/spf13/cobra"
)

type runCmdParameter struct {
	file          string
	modes         []string
	protect       bool
	enableAnalyse bool
	debug         bool
	cooperate     bool
}

var (
	runParameter runCmdParameter
	runCmd       = &cobra.Command{
		Use:   "run [file] [modes...]",
		Short: `Run runs the yock script or module`,
		Long:  `Run runs the yock script or module`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 || filepath.Ext(args[0]) != ".lua" {
				ycho.Fatal(util.ErrFileNotExist)
			}
			for idx, arg := range args {
				if idx == 0 {
					runParameter.file = arg
					continue
				}
				if len(arg) > 0 && arg[0] == '-' && arg[1] == '-' {
					break
				}
				runParameter.modes = append(runParameter.modes, arg)
			}

			opts := []yocks.YockSchedulerOption{
				yocks.OptionLibPath(util.Pathf("~/lib/yock")),
			}
			if runParameter.cooperate {
				opts = append(opts, yocks.OptionUpgradeSingalStream())
			}
			if runParameter.enableAnalyse {
				opts = append(opts, yocks.OptionEnableYockDriverMode())
			}

			yocks := yocks.Default(opts...)
			go yocks.EventLoop()
			go ycho.Eventloop()

			yockp := yockpack.New()
			fn := yockp.Compile(yockpack.CompileOpt{
				DisableAnalyse: runParameter.enableAnalyse,
				VM:             yocks.YockRuntime,
			}, runParameter.file)

			if err := yockr.LuaDoFunc(yocks.State().LState(), fn); err != nil {
				ycho.Fatal(err)
			}

			var tasks sync.WaitGroup
			for _, mode := range runParameter.modes {
				tasks.Add(1)
				if runParameter.debug {
					ycho.Infof("%s start to run", mode)
				}
				go func(mode string) {
					yocks.LaunchTask(mode)
					tasks.Done()
				}(mode)
			}
			tasks.Wait()
		},
	}
)

func init() {
	yockCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().BoolVarP(&runParameter.protect, "protect", "p", false, "")
	runCmd.PersistentFlags().BoolVarP(&runParameter.enableAnalyse, "analyze", "a", false, "enable dependency analyse mode")
	runCmd.PersistentFlags().BoolVarP(&runParameter.debug, "debug", "d", false, "print the information of launch")
	runCmd.PersistentFlags().BoolVarP(&runParameter.cooperate, "cooperate", "c", false, "enable daemon to meet distributed system")
}
