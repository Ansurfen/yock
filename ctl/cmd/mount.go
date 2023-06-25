// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"path/filepath"

	yockc "github.com/ansurfen/yock/cmd"
	"github.com/ansurfen/yock/util"
	"github.com/spf13/cobra"
)

type mountCmdParameter struct {
	recovery bool
	plain    bool
}

var (
	mountParameter mountCmdParameter
	mountCmd       = &cobra.Command{
		Use:   "mount [name] [file]",
		Short: `Mount mounts the specified file to make it globally available`,
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				util.Ycho.Fatal(util.ErrArgsTooLittle.Error())
			}
			name := args[0]
			file, err := filepath.Abs(args[1])
			if err != nil {
				util.Ycho.Fatal(err.Error())
			}

			mount_file := name
			mount_tmpl := ""
			unmount_file := args[1]

			switch util.CurPlatform.OS {
			case "windows":
				mount_file += ".bat"
				unmount_file += ".bat"
				if mountParameter.plain {
					mount_tmpl = wrapPlainBatch
				} else {
					mount_tmpl = wrapYockBatch
				}
			default:
				mount_file += ".sh"
				unmount_file += ".sh"
				if mountParameter.plain {
					mount_tmpl = wrapPlainBash
				} else {
					mount_tmpl = wrapYockBash
				}
			}

			mount_path := util.Pathf("@/mount")

			if err := util.SafeMkdirs(mount_path); err != nil {
				util.Ycho.Fatal(err.Error())
			}
			if mountParameter.recovery {
				if util.IsExist(filepath.Join(mount_path, mount_file)) {
					util.Ycho.Fatal(util.ErrFileExist.Error())
				}
				err := yockc.Mv(yockc.MvOpt{}, util.Pathf("@/unmount/")+unmount_file, filepath.Join(mount_path, mount_file))
				if err != nil {
					util.Ycho.Fatal(err.Error())
				}
			} else {
				if mountParameter.plain {
					mount_tmpl = fmt.Sprintf(mount_tmpl, file)
				} else {
					mount_tmpl = fmt.Sprintf(mount_tmpl, file, file)
				}
				if err := util.WriteFile(
					filepath.Join(mount_path, mount_file),
					[]byte(mount_tmpl)); err != nil {
					util.Ycho.Fatal(err.Error())
				}
			}
		},
	}
)

func init() {
	yockCmd.AddCommand(mountCmd)
	mountCmd.PersistentFlags().BoolVarP(&mountParameter.recovery, "recovery", "r", false, "restores the specified file from unmount to mount state")
	mountCmd.PersistentFlags().BoolVarP(&mountParameter.plain, "plain", "p", false, "only mount the file, do not add the yock run prefix")
}

const (
	wrapYockBatch = `@echo off
yock run %s -- %s %%*`
	wrapYockBash = `#/bin/bash
yock run %s -- %s %%*`
	wrapPlainBatch = `@echo off
%s %%*`
	wrapPlainBash = `#/bin/bash
%s %%*`
)
