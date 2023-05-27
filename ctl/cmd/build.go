// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/ansurfen/yock/util"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build [script]",
	Short: `Build packages yock and specified scripts together into executable file`,
	Long: `Build packages yock and specified scripts together into executable file,
but a Go compiler is required`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			util.Ycho.Fatal(util.ErrArgsTooLittle.Error())
		}
	},
}

func init() {
	yockCmd.AddCommand(buildCmd)
}
