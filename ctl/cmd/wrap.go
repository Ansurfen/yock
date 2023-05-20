package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/util"
	"github.com/spf13/cobra"
)

var wrapCmd = &cobra.Command{
	Use:   "wrap [name] [file]",
	Short: ``,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 || filepath.Ext(args[1]) != ".lua" {
			util.YchoFatal("", util.ErrArgsTooLittle.Error())
		}
		if err := utils.SafeMkdirs(util.Pathf("@/mount")); err != nil {
			util.YchoFatal("", err.Error())
		}
		switch utils.CurPlatform.OS {
		case "windows":
			name := args[0]
			file, err := filepath.Abs(args[1])
			if err != nil {
				util.YchoFatal("", err.Error())
			}
			if err := utils.WriteFile(filepath.Join(util.Pathf("@/mount"), name+".bat"), []byte(fmt.Sprintf(WrapBatch, file, file))); err != nil {
				util.YchoFatal("", err.Error())
			}
		default: // PosixOS
			name := args[0]
			file, err := filepath.Abs(args[1])
			if err != nil {
				util.YchoFatal("", err.Error())
			}
			if err := utils.WriteFile(filepath.Join(util.Pathf("@/mount"), name+".sh"), []byte(fmt.Sprintf(WrapBatch, file, file))); err != nil {
				util.YchoFatal("", err.Error())
			}
		}
	},
}

func init() {
	yockCmd.AddCommand(wrapCmd)
}

const (
	WrapBatch = `@echo off
yock run %s -- %s %%*`
	WrapBash = `#/bin/bash
yock run %s -- %s %%*`
)
