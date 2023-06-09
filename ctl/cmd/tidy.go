// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import "github.com/spf13/cobra"

var tidyCmd = &cobra.Command{
	Use:   "tidy",
	Short: `Tidy completes the script definition in the workspace`,
	Long:  `Tidy completes the script definition in the workspace`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	yockCmd.AddCommand(tidyCmd)
}
