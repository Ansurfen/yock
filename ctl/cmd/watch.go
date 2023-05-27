// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/ansurfen/yock/util"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch [hardware]",
	Short: `Watch monitors hardware status and information`,
	Long: `Watch monitors hardware status and information.
You can use yock watch [hardware] command to view details.
At present yock only supports cpu, mem, disk, net and host as hardware parameter`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			util.Ycho.Fatal(util.ErrArgsTooLittle.Error())
		}
		hardware := args[0]
		switch hardware {
		case "cpu":
			fmt.Printf("LogicalCore: %d\nPhysicalCore: %d\n", util.CPU().LogicalCore, util.CPU().PhysicalCore)
		case "mem":
			stat, err := util.Mem().VirtualMemory()
			if err != nil {
				util.Ycho.Fatal(err.Error())
			}
			fmt.Printf("Total: %d\nAvailable: %d\nUsed: %d (%.1f%%)\nFree: %d", stat.Total, stat.Available, stat.Used, stat.UsedPercent, stat.Free)
		case "disk":
		case "net":
		case "host":
		default:
			util.Ycho.Fatal(util.ErrNoSupportHardward.Error())
		}
	},
}

func init() {
	yockCmd.AddCommand(watchCmd)
}
