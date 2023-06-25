// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	yockc "github.com/ansurfen/yock/cmd"
	"github.com/ansurfen/yock/util"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm [area] [pattern]",
	Short: `Rm removes files from the specified area`,
	Long:  `Rm removes files from the specified area`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			util.Ycho.Fatal(util.ErrArgsTooLittle.Error())
		}
		area := args[0]
		pattern := args[1]
		ext := ""
		switch util.CurPlatform.OS {
		case "windows":
			ext = ".bat"
		default:
			ext = ".sh"
		}
		switch area {
		case "mount":
			pattern = util.Pathf("@/mount/") + pattern + ext
		case "unmount":
			pattern = util.Pathf("@/unmount/") + pattern + ext
		default:
			util.Ycho.Fatal("invalid area")
		}
		if err := yockc.Rm(yockc.RmOpt{Safe: false}, []string{pattern}); err != nil {
			util.Ycho.Fatal(err.Error())
		}
	},
}

func init() {
	yockCmd.AddCommand(rmCmd)
}
