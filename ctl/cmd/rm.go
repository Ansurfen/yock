// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	yockc "github.com/ansurfen/yock/cmd"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm [area] [pattern]",
	Short: `Rm removes files from the specified area`,
	Long:  `Rm removes files from the specified area`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			ycho.Fatal(util.ErrArgsTooLittle)
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
			ycho.Fatalf("invalid area")
		}
		if err := yockc.Rm(yockc.RmOpt{Safe: false}, []string{pattern}); err != nil {
			ycho.Fatal(err)
		}
	},
}

func init() {
	yockCmd.AddCommand(rmCmd)
}
