// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	parser "github.com/ansurfen/yock/pack"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
	"github.com/spf13/cobra"
)

type compileCmdParameter struct {
	file      string
	modes     []string
	decompose bool
	output    string
}

var (
	compileParameter compileCmdParameter
	compileCmd       = &cobra.Command{
		Use:   "compile [file]",
		Short: `Compile preprocess yock script to meet different demand`,
		Long:  `Compile preprocess yock script to meet different demand`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				ycho.Fatal(util.ErrFileNotExist)
			}
			for i := 0; i < len(args); i++ {
				if i == 0 {
					compileParameter.file = args[i]
					continue
				}
				compileParameter.modes = append(compileParameter.modes, args[i])
			}
			if compileParameter.decompose {
				yockpack := parser.YockPack[parser.NilFrame]{}
				yockpack.Decompose(parser.DecomposeOpt{
					Modes: compileParameter.modes,
					File:  compileParameter.output,
					Tpl:   util.Pathf("@/sdk/yock/decompose.tpl"),
				}, yockpack.ParseFile(compileParameter.file))
			} else {
				include, err := util.OpenConf(util.Pathf("@/include.yaml"))
				if err != nil {
					ycho.Fatal(err)
				}
				files := include.GetStringSlice("file")
				methods := include.GetStringSlice("method")
				anlyzer := parser.NewLuaDependencyAnalyzer()
				// import stdlib
				out, err := util.ReadStraemFromFile(util.Pathf("@/sdk/yock/deps/stdlib.json"))
				if err != nil {
					ycho.Fatal(err)
				}
				if err = json.Unmarshal(out, anlyzer); err != nil {
					ycho.Fatal(err)
				}
				for _, method := range methods {
					if !strings.HasSuffix(method, "()") {
						method = method + "()"
					}
					anlyzer.Preload(method, parser.LuaMethod{Pkg: "g"})
				}
				for _, file := range files {
					anlyzer.Load(util.Pathf(file))
				}
				fmt.Println(anlyzer.Completion(compileParameter.file))
			}
		},
	}
)

func init() {
	yockCmd.AddCommand(compileCmd)
	compileCmd.PersistentFlags().StringSliceVarP(&compileParameter.modes, "modes", "m", nil, "modes are used to divide source files")
	compileCmd.PersistentFlags().BoolVarP(&compileParameter.decompose, "decomposition", "d", false, "")
	compileCmd.PersistentFlags().StringVarP(&compileParameter.output, "output", "o", "", "")
}
