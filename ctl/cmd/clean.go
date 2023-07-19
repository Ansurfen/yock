// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	yockc "github.com/ansurfen/yock/cmd"
	"github.com/ansurfen/yock/util"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean [pattern]",
	Short: `Clean deletes Yock's cache`,
	Long: `Clean deletes Yock's cache. You can append the pattern parameter
to delete the specified file by matching the regular expression.`,
	Run: func(cmd *cobra.Command, args []string) {
		pattern := ".*"
		if len(args) > 0 {
			pattern = args[0]
		}
		yockc.Rm(yockc.RmOpt{Pattern: pattern, Safe: false}, util.Pathf("@/tmp"))
	},
}

func init() {
	yockCmd.AddCommand(cleanCmd)
}
