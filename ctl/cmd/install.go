// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	yockc "github.com/ansurfen/yock/cmd"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: `Install updates yock`,
	Long:  `Install updates yock`,
	Run: func(cmd *cobra.Command, args []string) {

		yockc.Curl(yockc.CurlOpt{}, []string{})
	},
}

func init() {
	yockCmd.AddCommand(installCmd)
}
