// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"os"
	"path/filepath"

	yockc "github.com/ansurfen/yock/cmd"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [sdk]",
	Short: ``,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			ycho.Fatal(util.ErrArgsTooLittle)
		}
		path := filepath.Join(util.Pathf("~/sdk"), args[0])
		wd, err := os.Getwd()
		if err != nil {
			ycho.Fatal(err)
		}
		err = yockc.Cp(yockc.CpOpt{Recurse: true}, path, wd)
		if err != nil {
			ycho.Fatal(err)
		}
	},
}

func init() {
	yockCmd.AddCommand(initCmd)
}
