package cmd

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/ansurfen/yock/scheduler"
	"github.com/ansurfen/yock/util"
	"github.com/spf13/cobra"
)

type runCmdParameter struct {
	file           string
	modes          []string
	protect        bool
	disableAnalyse bool
	debug          bool
	cooperate      bool
}

var (
	runParameter runCmdParameter
	runCmd       = &cobra.Command{
		Use:   "run [file] [modes...]",
		Short: `Run runs the yock script or module.`,
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 || filepath.Ext(args[0]) != ".lua" {
				util.YchoFatal("", util.ErrFileNotExist.Error())
			}
			for idx, arg := range args {
				if idx == 0 {
					runParameter.file = arg
					continue
				}
				runParameter.modes = append(runParameter.modes, arg)
			}
			
			opts := []scheduler.YockSchedulerOption{}
			if runParameter.cooperate {
				opts = append(opts, scheduler.OptionUpgradeSingalStream())
			}
			
			yock := scheduler.New(opts...)
			go yock.EventLoop()

			proto := yock.Compile(scheduler.CompileOpt{
				DisableAnalyse: runParameter.disableAnalyse,
			}, runParameter.file)
			if err := yock.DoCompliedFile(proto); err != nil {
				util.YchoFatal("", err.Error())
			}
			var jobs sync.WaitGroup
			for _, mode := range runParameter.modes {
				jobs.Add(1)
				if runParameter.debug {
					util.YchoInfo("", fmt.Sprintf("%s start to run", mode))
				}
				go func(mode string) {
					yock.RunJob(mode)
					jobs.Done()
				}(mode)
			}
			jobs.Wait()
		},
	}
)

func init() {
	yockCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().BoolVarP(&runParameter.protect, "protect", "p", false, "")
	runCmd.PersistentFlags().BoolVarP(&runParameter.disableAnalyse, "analyze", "a", false, "")
	runCmd.PersistentFlags().BoolVarP(&runParameter.debug, "debug", "d", false, "")
	runCmd.PersistentFlags().BoolVarP(&runParameter.cooperate, "cooperate", "c", false, "")
}
