// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import "github.com/spf13/cobra"

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: `Daemon is used to manage and view the status of yock's daemon`,
	Long:  `Daemon is used to manage and view the status of yock's daemon`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	yockCmd.AddCommand(daemonCmd)
}
