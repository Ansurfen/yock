// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
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
			ycho.Fatal(util.ErrArgsTooLittle)
		}
		hardware := args[0]
		switch hardware {
		case "cpu":
			info, err := util.CPU().Info()
			if err != nil {
				ycho.Fatal(err)
			}
			fmt.Printf("LogicalCore: %d\nPhysicalCore: %d\nModelName: %s\nMHZ: %f",
				util.CPU().LogicalCore, util.CPU().PhysicalCore, info[0].ModelName, info[0].Mhz)
		case "mem":
			stat, err := util.Mem().VirtualMemory()
			if err != nil {
				ycho.Fatal(err)
			}
			fmt.Printf("Total: %d\nAvailable: %d\nUsed: %d (%.1f%%)\nFree: %d", stat.Total, stat.Available, stat.Used, stat.UsedPercent, stat.Free)
		case "disk":
			parts, err := util.Disk().Partitions(true)
			if err != nil {
				ycho.Fatal(err)
			}
			for _, part := range parts {
				stat, err := util.Disk().Usage(part.Device)
				if err != nil {
					ycho.Fatal(err)
				}
				fmt.Printf("%s\n  Type: %s\n  Opts: %s\n  Total: %d\n  Free: %d\n  Used: %d (%.2f%%)\n",
					part.Device, part.Fstype, strings.Join(part.Opts, " "), stat.Total, stat.Free, stat.Used, stat.UsedPercent)
			}
		case "net":
			interfaces, err := util.Net().Interfaces()
			if err != nil {
				ycho.Fatal(err)
			}
			for _, inter := range interfaces {
				fmt.Printf("%s\n  MTU: %s\n  Type: %s\n  MAC: %s\n  IPv6: %s\n  IPv4: %s\n",
					inter.Name, strconv.Itoa(inter.MTU), strings.Join(inter.Flags, " "), inter.HardwareAddr, inter.Addrs[0].Addr, inter.Addrs[1].Addr)
			}
		case "host":
			os, fam, ver, err := util.Host().PlatformInformation()
			if err != nil {
				ycho.Fatal(err)
			}
			fmt.Printf("OS: %s\nFamily: %s\nVersion: %s", os, fam, ver)
		default:
			ycho.Fatal(util.ErrNoSupportHardward)
		}
	},
}

func init() {
	yockCmd.AddCommand(watchCmd)
}
