package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/parser"
	"github.com/ansurfen/yock/scheduler"
	"github.com/spf13/cobra"
)

type compileCmdParameter struct {
	file          string
	modes         []string
	decomposition bool
	output        string
}

var (
	compileParameter compileCmdParameter
	compileCmd       = &cobra.Command{
		Use:   "compile [file]",
		Short: ``,
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("file not found")
				os.Exit(1)
			}
			for i := 0; i < len(args); i++ {
				if i == 0 {
					compileParameter.file = args[i]
					continue
				}
				compileParameter.modes = append(compileParameter.modes, args[i])
			}
			if compileParameter.decomposition {
				parser.Decomposition(parser.DecompositionOpt{
					Modes: compileParameter.modes,
					File:  compileParameter.output,
					Tpl:   scheduler.Pathf("@/sdk/yock/decomposition.tpl"),
				}, parser.ParserASTFromFile(compileParameter.file))
			} else {
				include := utils.OpenConfFromPath(scheduler.Pathf("@/include.yaml"))
				if err := include.ReadInConfig(); err != nil {
					panic(err)
				}
				files := include.GetStringSlice("file")
				methods := include.GetStringSlice("method")
				anlyzer := parser.NewLuaDependencyAnalyzer()
				// import stdlib
				out, err := utils.ReadStraemFromFile(scheduler.Pathf("@/sdk/yock/deps/stdlib.json"))
				if err != nil {
					panic(err)
				}
				if err = json.Unmarshal(out, anlyzer); err != nil {
					panic(err)
				}
				for _, method := range methods {
					if !strings.HasSuffix(method, "()") {
						method = method + "()"
					}
					anlyzer.Preload(method, parser.LuaMethod{Pkg: "g"})
				}
				for _, file := range files {
					anlyzer.Load(scheduler.Pathf(file))
				}
				fmt.Println(anlyzer.Tidy(compileParameter.file))
			}
		},
	}
)

func init() {
	yockCmd.AddCommand(compileCmd)
	compileCmd.PersistentFlags().StringSliceVarP(&compileParameter.modes, "modes", "m", nil, "")
	compileCmd.PersistentFlags().BoolVarP(&compileParameter.decomposition, "decomposition", "d", false, "")
	compileCmd.PersistentFlags().StringVarP(&compileParameter.output, "output", "o", "", "")
}
