// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fslib

import (
	"io/fs"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadFs(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("fs")
	lib.SetField(map[string]any{
		// functions
		"Sub":                fs.Sub,
		"WalkDir":            fs.WalkDir,
		"ValidPath":          fs.ValidPath,
		"Glob":               fs.Glob,
		"ReadDir":            fs.ReadDir,
		"FileInfoToDirEntry": fs.FileInfoToDirEntry,
		"ReadFile":           fs.ReadFile,
		"Stat":               fs.Stat,
		// constants
		"ModeDir":        fs.ModeDir,
		"ModeAppend":     fs.ModeAppend,
		"ModeExclusive":  fs.ModeExclusive,
		"ModeTemporary":  fs.ModeTemporary,
		"ModeSymlink":    fs.ModeSymlink,
		"ModeDevice":     fs.ModeDevice,
		"ModeNamedPipe":  fs.ModeNamedPipe,
		"ModeSocket":     fs.ModeSocket,
		"ModeSetuid":     fs.ModeSetuid,
		"ModeSetgid":     fs.ModeSetgid,
		"ModeCharDevice": fs.ModeCharDevice,
		"ModeSticky":     fs.ModeSticky,
		"ModeIrregular":  fs.ModeIrregular,
		"ModeType":       fs.ModeType,
		"ModePerm":       fs.ModePerm,
		// variable
		"ErrInvalid":    fs.ErrInvalid,
		"ErrPermission": fs.ErrPermission,
		"ErrExist":      fs.ErrExist,
		"ErrNotExist":   fs.ErrNotExist,
		"ErrClosed":     fs.ErrClosed,
		"SkipDir":       fs.SkipDir,
		"SkipAll":       fs.SkipAll,
	})
}
