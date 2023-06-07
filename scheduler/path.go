// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"io/fs"
	"path/filepath"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/util"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
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
	"walk":     pathWalk,
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

/*
* @param root string
* @param fn fun(path: string, info: fileinfo, err:err): bool
* @return err
 */
func pathWalk(l *lua.LState) int {
	fn := l.CheckFunction(2)
	err := filepath.Walk(l.CheckString(1), func(path string, info fs.FileInfo, err error) error {
		e := lua.LNil
		if err != nil {
			e = lua.LString(err.Error())
		}
		if err := l.CallByParam(lua.P{
			Fn:   fn,
			NRet: 1,
		}, lua.LString(path), luar.New(l, info), e); err != nil {
			util.Ycho.Fatal(err.Error())
		}
		ok := l.CheckBool(-1)
		l.Pop(l.GetTop())
		if ok {
			return nil
		}
		return err
	})
	handleErr(l, err)
	return 1
}
