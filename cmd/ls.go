// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

type LsFileInfo struct {
	Perm     string
	Size     int64
	ModTime  string
	Filename string
}

type LsOpt struct {
	Dir     string
	recurse bool
	Callack func(path string, info fs.FileInfo) bool
}

func (opt *LsOpt) SetRecurse() {
	opt.recurse = true
}

func Ls(opt LsOpt) ([]LsFileInfo, error) {
	fileinfo := []LsFileInfo{}
	if opt.recurse {
		err := filepath.Walk(opt.Dir, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !opt.Callack(path, info) {
				return errors.New("error")
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	} else {
		files, err := os.ReadDir(opt.Dir)
		if err != nil {
			return fileinfo, err
		}
		for _, file := range files {
			info, err := file.Info()
			if err != nil {
				return fileinfo, err
			}
			perm := info.Mode().Perm().String()
			size := info.Size()
			modeTime := info.ModTime().Format("Jan _2 15:04")
			filename := file.Name()
			fileinfo = append(fileinfo, LsFileInfo{
				Perm:     perm,
				Size:     size,
				ModTime:  modeTime,
				Filename: filename,
			})
		}
	}
	return fileinfo, nil
}
