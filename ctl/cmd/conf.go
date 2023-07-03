// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/ansurfen/yock/ctl/conf"
	yocke "github.com/ansurfen/yock/env"
	"github.com/ansurfen/yock/ycho"
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
				if err := conf.Instance().Restore(); err != nil {
					ycho.Fatal(err)
				}
				return
			}
			if len(confParameter.lang) > 0 {
				// yocke.GetEnv[conf.YockConf]().Commit("lang", confParameter.lang)
			}
			yocke.GetEnv[conf.YockConf]().Save()
		},
	}
)

func init() {
	yockCmd.AddCommand(confCmd)
	confCmd.PersistentFlags().StringVarP(&confParameter.lang, "lang", "l", "", "")
	confCmd.PersistentFlags().BoolVarP(&confParameter.restore, "restore", "r", false, "")
}
