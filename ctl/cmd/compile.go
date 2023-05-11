package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/parser"
	"github.com/spf13/cobra"
)

type compileCmdParameter struct {
	file  string
	modes []string
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
			compileParameter.file = args[0]
			include := utils.OpenConfFromPath("test/include.yaml")
			if err := include.ReadInConfig(); err != nil {
				panic(err)
			}

			files := include.GetStringSlice("file")
			methods := include.GetStringSlice("method")
			anlyzer := parser.NewLuaDependencyAnalyzer()
			// import stdlib
			out, err := utils.ReadStraemFromFile("../parser/stdlib.json")
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
				anlyzer.Load(file)
			}
			fmt.Println(anlyzer.Tidy(compileParameter.file))
		},
	}
)

func init() {
	yockCmd.AddCommand(compileCmd)
	compileCmd.PersistentFlags().StringSliceVarP(&compileParameter.modes, "modes", "m", nil, "")
}
