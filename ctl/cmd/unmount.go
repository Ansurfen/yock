// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"path/filepath"

	yockc "github.com/ansurfen/yock/cmd"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
	"github.com/spf13/cobra"
)

var unmountCmd = &cobra.Command{
	Use:   "unmount [name]",
	Short: `Unmount unmounts specifies the file from mount`,
	Long: `Unmount unmounts specifies the file from mount,
which enables file fail to access in global.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			ycho.Fatal(util.ErrArgsTooLittle)
		}
		if err := util.SafeMkdirs(util.Pathf("@/unmnt")); err != nil {
			ycho.Fatal(err)
		}
		file := args[0]
		exf := ""
		switch util.CurPlatform.OS {
		case "windows":
			exf = ".bat"
		default:
			exf = ".sh"
		}
		if err := yockc.Mv(yockc.MvOpt{}, filepath.Join(util.Pathf("@/mnt"), file+exf),
			filepath.Join(util.Pathf("@/unmnt"), file+exf)); err != nil {
			ycho.Fatal(err)
		}
	},
}

func init() {
	yockCmd.AddCommand(unmountCmd)
}
