// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"path/filepath"

	"github.com/ansurfen/cushion/utils"
	yockc "github.com/ansurfen/yock/cmd"
	"github.com/ansurfen/yock/util"
	"github.com/spf13/cobra"
)

var unmountCmd = &cobra.Command{
	Use:   "unmount [name]",
	Short: `Unmount unmounts specifies the file from mount`,
	Long: `Unmount unmounts specifies the file from mount,
which enables file fail to access in global.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			util.Ycho.Fatal(util.ErrArgsTooLittle.Error())
		}
		if err := utils.SafeMkdirs(util.Pathf("@/unmount")); err != nil {
			util.Ycho.Fatal(err.Error())
		}
		file := args[0]
		exf := ""
		switch utils.CurPlatform.OS {
		case "windows":
			exf = ".bat"
		default:
			exf = ".sh"
		}
		if err := yockc.Mv(yockc.MvOpt{}, filepath.Join(util.Pathf("@/mount"), file+exf),
			filepath.Join(util.Pathf("@/unmount"), file+exf)); err != nil {
			util.Ycho.Fatal(err.Error())
		}
	},
}

func init() {
	yockCmd.AddCommand(unmountCmd)
}
