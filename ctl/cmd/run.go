package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/ansurfen/yock/scheduler"
	"github.com/spf13/cobra"
)

type runCmdParameter struct {
	file    string
	modes   []string
	protect bool
}

var (
	runParameter runCmdParameter
	runCmd       = &cobra.Command{
		Use:   "run [file] [modes...]",
		Short: ``,
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 || filepath.Ext(args[0]) != ".lua" {
				fmt.Println("file not found")
				os.Exit(1)
			}
			for idx, arg := range args {
				if idx == 0 {
					runParameter.file = arg
					continue
				}
				runParameter.modes = append(runParameter.modes, arg)
			}
			yock := scheduler.New()
			proto := yock.Compile(runParameter.file)
			if err := yock.DoCompliedFile(proto); err != nil {
				panic(err)
			}
			var jobs sync.WaitGroup
			for _, mode := range runParameter.modes {
				jobs.Add(1)
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
}
