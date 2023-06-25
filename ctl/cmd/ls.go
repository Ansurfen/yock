// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"os"

	"github.com/ansurfen/yock/util"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls [key]",
	Short: `Ls prints specified environment information based on the provided key`,
	Long: `ls prints the specified environment information according to the provided key, 
such as 'ls mount/unmount' to display mount/unmount information`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			util.Ycho.Fatal(util.ErrArgsTooLittle.Error())
		}
		switch v := args[0]; v {
		case "mount", "unmount":
			files, err := os.ReadDir(util.Pathf("@/" + v))
			if err != nil {
				util.Ycho.Fatal(err.Error())
			}
			rows := [][]string{}
			for _, file := range files {
				if !file.IsDir() {
					info, err := file.Info()
					if err != nil {
						util.Ycho.Fatal(err.Error())
					}
					rows = append(rows, []string{info.Mode().Perm().String(), file.Name(), info.ModTime().Format("Jan _2 15:04")})
				}
			}
			util.Prinf(util.PrintfOpt{MaxLen: 30}, []string{"Perm", "Filename", "ModTime"}, rows)
		default:
		}
	},
}

func init() {
	yockCmd.AddCommand(lsCmd)
}
