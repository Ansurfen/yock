// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/util"
	"github.com/spf13/cobra"
)

type envCmdParameter struct {
	safe   bool
	key    string
	value  string
	path   string
	expand bool
	local  bool
}

var (
	envParameter envCmdParameter
	envCmd       = &cobra.Command{
		Use:   "env [key] [value]",
		Short: `Env is used to CRUD environment variables`,
		Long:  `Env is used to CRUD environment variables`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				util.Ycho.Fatal("use yock env [key] [value]")
			} else {
				envParameter.key = args[0]
				envParameter.value = args[1]
			}
			env := utils.NewEnvVar()
			if len(envParameter.path) > 0 {
				switch utils.CurPlatform.OS {
				case "windows":
					switch v := envParameter.path; v {
					case "sys":
						env.SetPath(envParameter.path)
					case "user":
						env.SetPath(envParameter.path)
					default:
						util.Ycho.Fatal(util.ErrInvalidPath.Error())
					}
				default:
					env.SetPath(envParameter.path)
				}
			}
			if envParameter.safe {
				env.SafeSet(envParameter.key, envParameter.value)
			} else {
				env.Set(envParameter.key, envParameter.value)
			}
		},
	}
)

func init() {
	yockCmd.AddCommand(envCmd)
	envCmd.PersistentFlags().BoolVarP(&envParameter.safe, "safe", "s", true, "")
	envCmd.PersistentFlags().StringVarP(&envParameter.path, "path", "p", "", "")
	envCmd.PersistentFlags().BoolVarP(&envParameter.expand, "expand", "e", false, "")
	envCmd.PersistentFlags().BoolVarP(&envParameter.local, "local", "l", false, "")
}
