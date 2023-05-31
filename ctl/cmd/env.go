// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/util"
	"github.com/spf13/cobra"
)

type envCmdParameter struct {
	safe   bool
	path   string
	expand bool
	local  bool

	env utils.EnvVar
}

func getChar() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter any keyboard to exit")
	reader.ReadRune()
}

var (
	envParameter envCmdParameter
	envCmd       = &cobra.Command{
		Use:   "env [action]",
		Short: `Env is used to CRUD environment variables`,
		Long: `Env is used to CRUD environment variables.
Examples:
	yock env set [key] [value]
	yock env export [file]
	yock env print
	yock env unset [key]
	yock env search`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				util.Ycho.Fatal(util.ErrArgsTooLittle.Error())
			}
			envParameter.env = utils.NewEnvVar()
			if len(envParameter.path) > 0 {
				switch utils.CurPlatform.OS {
				case "windows":
					switch v := envParameter.path; v {
					case "sys":
						envParameter.env.SetPath(envParameter.path)
					case "user":
						envParameter.env.SetPath(envParameter.path)
					default:
						util.Ycho.Fatal(util.ErrInvalidPath.Error())
					}
				default:
					envParameter.env.SetPath(envParameter.path)
				}
			}
			action := args[0]
			switch action {
			case "set":
				if len(args) < 3 {
					util.Ycho.Fatal("Usage env set [key] [value]")
				}
				key := args[1]
				value := args[2]
				// global
				if envParameter.safe {
					if envParameter.expand {
						envParameter.env.SafeSet(key, []string{value})
					} else {
						envParameter.env.SafeSet(key, value)
					}
				} else {
					if envParameter.expand {
						envParameter.env.Set(key, []string{value})
					} else {
						envParameter.env.Set(key, value)
					}
				}
			case "unset":
				if len(args) < 2 {
					util.Ycho.Fatal("Usage env unset [key]")
				}
				if err := envParameter.env.Unset(args[1]); err != nil {
					util.Ycho.Fatal(err.Error())
				}
			// export current enviroment string into specify file
			case "export":
				if len(args) < 2 {
					util.Ycho.Fatal("Usage env export [file]")
				}
				file := args[1]
				if len(filepath.Ext(file)) == 0 {
					file += ".ini"
				}
				if err := envParameter.env.Export(file); err != nil {
					util.Ycho.Fatal(err.Error())
				}
			case "search":
			case "print":
				envParameter.env.Print()
				getChar()
			default:
				util.Ycho.Fatal(util.ErrGeneral.Error())
			}
		},
	}
)

func init() {
	yockCmd.AddCommand(envCmd)
	envCmd.PersistentFlags().BoolVarP(&envParameter.safe, "safe", "s", true, "set enviroment variable when key isn't exist")
	envCmd.PersistentFlags().StringVarP(&envParameter.path, "path", "p", "", " set operate target, windows: sys or user, this is required in windows.")
	envCmd.PersistentFlags().BoolVarP(&envParameter.expand, "expand", "e", false, "like slice to append value, not replace compare with string")
	envCmd.PersistentFlags().BoolVarP(&envParameter.local, "local", "l", false, "set local enviroment variable")
}
