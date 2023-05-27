// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import "github.com/spf13/cobra"

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: `Uninstall yock`,
	Long:  `Uninstall yock`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	yockCmd.AddCommand(uninstallCmd)
}
