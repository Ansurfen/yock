// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
	"os"
)

type LsFileInfo struct {
	Perm     string
	Size     int64
	ModTime  string
	Filename string
}

type LsOpt struct {
	Dir string
	Str bool
}

func Ls(opt LsOpt) ([]LsFileInfo, string, error) {
	fileinfo := []LsFileInfo{}
	fileinfoStr := ""
	files, err := os.ReadDir(opt.Dir)
	if err != nil {
		return fileinfo, fileinfoStr, err
	}
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			return fileinfo, fileinfoStr, err
		}
		perm := info.Mode().Perm().String()
		size := info.Size()
		modeTime := info.ModTime().Format("Jan _2 15:04")
		filename := file.Name()
		if opt.Str {
			fileinfoStr += fmt.Sprintf("%s\t%d\t%s\t%s\n", perm, size, modeTime, filename)
		} else {
			fileinfo = append(fileinfo, LsFileInfo{
				Perm:     perm,
				Size:     size,
				ModTime:  modeTime,
				Filename: filename,
			})
		}
	}
	return fileinfo, fileinfoStr, nil
}
