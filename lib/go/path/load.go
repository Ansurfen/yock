// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package path

import (
	"path/filepath"

	yocki "github.com/ansurfen/yock/interface"
	filepathlib "github.com/ansurfen/yock/lib/go/path/filepath"
	"github.com/ansurfen/yock/util"
)

func LoadPath(yocks yocki.YockScheduler) {
	filepathlib.LoadFilepath(yocks)
	lib := yocks.CreateLib("path")
	lib.SetField(map[string]any{
		"Separator": filepath.Separator,
	})
	lib.SetYFunction(map[string]yocki.YGFunction{
		"filename": pathFilename,
		"join":     pathJoin,
		"dir":      pathDir,
		"base":     pathBase,
		"clean":    pathClean,
		"ext":      pathExt,
		"abs":      pathAbs,
	})
}

// @param path string
//
// @return string
func pathFilename(s yocki.YockState) int {
	s.PushString(util.Filename(s.CheckString(1)))
	return 1
}

// @param elem ...string
//
// @return string
func pathJoin(s yocki.YockState) int {
	elem := []string{}
	for i := 1; i <= s.Argc(); i++ {
		elem = append(elem, s.CheckString(i))
	}
	s.PushString(filepath.Join(elem...))
	return 1
}

// @param path string
//
// @return string
func pathDir(s yocki.YockState) int {
	s.PushString(filepath.Dir(s.CheckString(1)))
	return 1
}

// @param path string
//
// @return string
func pathBase(s yocki.YockState) int {
	s.PushString(filepath.Base(s.CheckString(1)))
	return 1
}

// @param path string
//
// @return string
func pathClean(s yocki.YockState) int {
	s.PushString(filepath.Clean(s.CheckString(1)))
	return 1
}

// @param path string
//
// @return string
func pathExt(s yocki.YockState) int {
	s.PushString(filepath.Ext(s.CheckString(1)))
	return 1
}

// @param path string
//
// @return string, err
func pathAbs(s yocki.YockState) int {
	abs, err := filepath.Abs(s.CheckString(1))
	s.PushString(abs).PushError(err)
	return 2
}
