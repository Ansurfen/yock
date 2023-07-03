// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package path

import (
	"io/fs"
	"path/filepath"

	yocki "github.com/ansurfen/yock/interface"
	filepathlib "github.com/ansurfen/yock/lib/go/path/filepath"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
	lua "github.com/yuin/gopher-lua"
)

func LoadPath(yocks yocki.YockScheduler) {
	filepathlib.LoadFilepath(yocks)
	lib := yocks.CreateLib("path")
	lib.SetField(map[string]any{
		"Separator": filepath.Separator,
	})
	lib.SetYFunction(map[string]yocki.YGFunction{
		"exist":    pathExist,
		"filename": pathFilename,
		"join":     pathJoin,
		"dir":      pathDir,
		"base":     pathBase,
		"clean":    pathClean,
		"ext":      pathExt,
		"abs":      pathAbs,
		"walk":     pathWalk,
	})
}

// @param path string
//
// @return bool
func pathExist(l yocki.YockState) int {
	ok := util.IsExist(l.LState().CheckString(1))
	l.PushBool(ok)
	return 1
}

// @param path string
//
// @return string
func pathFilename(l yocki.YockState) int {
	l.PushString(util.Filename(l.LState().CheckString(1)))
	return 1
}

// @param elem ...string
//
// @return string
func pathJoin(l yocki.YockState) int {
	elem := []string{}
	for i := 1; i <= l.Argc(); i++ {
		elem = append(elem, l.LState().CheckString(i))
	}
	l.PushString(filepath.Join(elem...))
	return 1
}

// @param path string
//
// @return string
func pathDir(l yocki.YockState) int {
	l.PushString(filepath.Dir(l.LState().CheckString(1)))
	return 1
}

// @param path string
//
// @return string
func pathBase(l yocki.YockState) int {
	l.PushString(filepath.Base(l.LState().CheckString(1)))
	return 1
}

// @param path string
//
// @return string
func pathClean(l yocki.YockState) int {
	l.PushString(filepath.Clean(l.LState().CheckString(1)))
	return 1
}

// @param path string
//
// @return string
func pathExt(l yocki.YockState) int {
	l.PushString(filepath.Ext(l.LState().CheckString(1)))
	return 1
}

// @param path string
//
// @return string, err
func pathAbs(l yocki.YockState) int {
	abs, err := filepath.Abs(l.LState().CheckString(1))
	l.PushString(abs).PushError(err)
	return 2
}

/*
* @param root string
* @param fn fun(path: string, info: fileinfo, err:err): bool
* @return err
 */
func pathWalk(l yocki.YockState) int {
	fn := l.LState().CheckFunction(2)
	err := filepath.Walk(l.LState().CheckString(1), func(path string, info fs.FileInfo, err error) error {
		e := lua.LNil
		if err != nil {
			e = lua.LString(err.Error())
		}
		if err := l.Call(yocki.YockFuncInfo{
			Fn:   fn,
			NRet: 1,
		}, path, info, e); err != nil {
			ycho.Fatal(err)
		}
		ok := l.LState().CheckBool(-1)
		l.PopTop()
		if ok {
			return nil
		}
		return err
	})
	l.PushError(err)
	return 1
}
