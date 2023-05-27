// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/ansurfen/yock/util"
	"github.com/spf13/cobra"
)

type versionCmdParameter struct {
	logo bool
}

const logoYock = "                                       ,-.  \n" +
	"        ,---,                      ,--/ /|  \n" +
	"       /_ ./|   ,---.            ,--. :/ |  \n" +
	" ,---, |  ' :  '   ,'\\           :  : ' /   \n" +
	"/___/ \\.  : | /   /   |   ,---.  |  '  /    \n" +
	" .  \\  \\ ,' '.   ; ,. :  /     \\ '  |  :    \n" +
	"  \\  ;  `  ,''   | |: : /    / ' |  |   \\   \n" +
	"   \\  \\    ' '   | .; :.    ' /  '  : |. \\  \n" +
	"    '  \\   | |   :    |'   ; :__ |  | ' \\ \\ \n" +
	"     \\  ;  ;  \\   \\  / '   | '.'|'  : |--'  \n" +
	"      :  \\  \\  `----'  |   :    :;  |,'     \n" +
	"       \\  ' ;           \\   \\  / '--'       \n" +
	"        `--`             `----'             "

var (
	versionParameter versionCmdParameter
	versionCmd       = &cobra.Command{
		Use:   "version",
		Short: `Present the version information of yock`,
		Long:  `Present the version information of yock`,
		Run: func(cmd *cobra.Command, args []string) {
			if versionParameter.logo {
				fmt.Println(logoYock)
			}
			fmt.Printf("yock version: %s", util.YockVersion)
		},
	}
)

func init() {
	yockCmd.AddCommand(versionCmd)
	versionCmd.PersistentFlags().BoolVarP(&versionParameter.logo, "logo", "l", false, "display the logo")
}
