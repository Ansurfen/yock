// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/util"
	"github.com/spf13/cobra"
)

type confCmdParameter struct {
	restore bool
	lang    string
}

var (
	confParameter confCmdParameter
	confCmd       = &cobra.Command{
		Use:   "conf",
		Short: `Conf is used to modify yock's configuration`,
		Long:  `Conf is used to modify yock's configuration`,
		Run: func(cmd *cobra.Command, args []string) {
			if confParameter.restore {
				if err := util.Conf().Restore(); err != nil {
					util.Ycho.Fatal(err.Error())
				}
				return
			}
			if len(confParameter.lang) > 0 {
				utils.GetEnv().Commit("lang", confParameter.lang)
			}
			utils.GetEnv().Write()
		},
	}
)

func init() {
	yockCmd.AddCommand(confCmd)
	confCmd.PersistentFlags().StringVarP(&confParameter.lang, "lang", "l", "", "")
	confCmd.PersistentFlags().BoolVarP(&confParameter.restore, "restore", "r", false, "")
}
