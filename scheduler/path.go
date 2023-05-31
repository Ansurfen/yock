// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"path/filepath"

	"github.com/ansurfen/cushion/utils"
	lua "github.com/yuin/gopher-lua"
)

func loadPath(yocks *YockScheduler) lua.LValue {
	pathlib := &lua.LTable{}
	pathlib.RawSetString("Separator", lua.LString(filepath.Separator))
	return yocks.mountLib(pathlib, pathLib)
}

var pathLib = luaFuncs{
	"exist":    pathExist,
	"filename": pathFilename,
	"join":     pathJoin,
	"dir":      pathDir,
	"base":     pathBase,
	"clean":    pathClean,
	"ext":      pathExt,
	"abs":      pathAbs,
}

// @param path string
//
// @return bool
func pathExist(l *lua.LState) int {
	ok := utils.IsExist(l.CheckString(1))
	handleBool(l, ok)
	return 1
}

// @param path string
//
// @return string
func pathFilename(l *lua.LState) int {
	l.Push(lua.LString(utils.Filename(l.CheckString(1))))
	return 1
}

// @param elem ...string
//
// @return string
func pathJoin(l *lua.LState) int {
	elem := []string{}
	for i := 1; i <= l.GetTop(); i++ {
		elem = append(elem, l.CheckString(i))
	}
	l.Push(lua.LString(filepath.Join(elem...)))
	return 1
}

// @param path string
//
// @return string
func pathDir(l *lua.LState) int {
	l.Push(lua.LString(filepath.Dir(l.CheckString(1))))
	return 1
}

// @param path string
//
// @return string
func pathBase(l *lua.LState) int {
	l.Push(lua.LString(filepath.Base(l.CheckString(1))))
	return 1
}

// @param path string
//
// @return string
func pathClean(l *lua.LState) int {
	l.Push(lua.LString(filepath.Clean(l.CheckString(1))))
	return 1
}

// @param path string
//
// @return string
func pathExt(l *lua.LState) int {
	l.Push(lua.LString(filepath.Ext(l.CheckString(1))))
	return 1
}

// @param path string
//
// @return string, err
func pathAbs(l *lua.LState) int {
	abs, err := filepath.Abs(l.CheckString(1))
	l.Push(lua.LString(abs))
	handleErr(l, err)
	return 2
}
