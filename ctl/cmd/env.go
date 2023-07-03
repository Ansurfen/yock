// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	yocke "github.com/ansurfen/yock/env"
	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
	"github.com/spf13/cobra"
)

type envCmdParameter struct {
	safe   bool
	path   string
	expand bool
	local  bool

	env yocki.EnvVar
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
				ycho.Fatal(util.ErrArgsTooLittle)
			}
			envParameter.env = yocke.NewEnvVar()
			if len(envParameter.path) > 0 {
				switch util.CurPlatform.OS {
				case "windows":
					switch v := envParameter.path; v {
					case "sys":
						envParameter.env.SetPath(envParameter.path)
					case "user":
						envParameter.env.SetPath(envParameter.path)
					default:
						ycho.Fatal(util.ErrInvalidPath)
					}
				default:
					envParameter.env.SetPath(envParameter.path)
				}
			}
			action := args[0]
			switch action {
			case "set":
				if len(args) < 3 {
					ycho.Fatalf("Usage env set [key] [value]")
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
					ycho.Fatalf("Usage env unset [key]")
				}
				if err := envParameter.env.Unset(args[1]); err != nil {
					ycho.Fatal(err)
				}
			// export current enviroment string into specify file
			case "export":
				if len(args) < 2 {
					ycho.Fatalf("Usage env export [file]")
				}
				file := args[1]
				if len(filepath.Ext(file)) == 0 {
					file += ".ini"
				}
				if err := envParameter.env.Export(file); err != nil {
					ycho.Fatal(err)
				}
			case "search":
			case "print":
				envParameter.env.Print()
				getChar()
			default:
				ycho.Fatal(util.ErrGeneral)
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
