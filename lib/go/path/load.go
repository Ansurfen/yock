// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package path

import (
	"io/fs"
	"path/filepath"

	"github.com/ansurfen/cushion/utils"
	yocki "github.com/ansurfen/yock/interface"
	yockr "github.com/ansurfen/yock/runtime"
	"github.com/ansurfen/yock/util"
	lua "github.com/yuin/gopher-lua"
)

func LoadPath(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("path")
	lib.SetField(map[string]any{
		"Separator": filepath.Separator,
	})
	lib.SetYFunction(map[string]yockr.YGFunction{
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
func pathExist(l *yockr.YockState) int {
	ok := utils.IsExist(l.CheckString(1))
	l.PushBool(ok)
	return 1
}

// @param path string
//
// @return string
func pathFilename(l *yockr.YockState) int {
	l.PushString(utils.Filename(l.CheckString(1)))
	return 1
}

// @param elem ...string
//
// @return string
func pathJoin(l *yockr.YockState) int {
	elem := []string{}
	for i := 1; i <= l.GetTop(); i++ {
		elem = append(elem, l.CheckString(i))
	}
	l.PushString(filepath.Join(elem...))
	return 1
}

// @param path string
//
// @return string
func pathDir(l *yockr.YockState) int {
	l.PushString(filepath.Dir(l.CheckString(1)))
	return 1
}

// @param path string
//
// @return string
func pathBase(l *yockr.YockState) int {
	l.PushString(filepath.Base(l.CheckString(1)))
	return 1
}

// @param path string
//
// @return string
func pathClean(l *yockr.YockState) int {
	l.PushString(filepath.Clean(l.CheckString(1)))
	return 1
}

// @param path string
//
// @return string
func pathExt(l *yockr.YockState) int {
	l.PushString(filepath.Ext(l.CheckString(1)))
	return 1
}

// @param path string
//
// @return string, err
func pathAbs(l *yockr.YockState) int {
	abs, err := filepath.Abs(l.CheckString(1))
	l.PushString(abs).PushError(err)
	return 2
}

/*
* @param root string
* @param fn fun(path: string, info: fileinfo, err:err): bool
* @return err
 */
func pathWalk(l *yockr.YockState) int {
	fn := l.CheckFunction(2)
	err := filepath.Walk(l.CheckString(1), func(path string, info fs.FileInfo, err error) error {
		e := lua.LNil
		if err != nil {
			e = lua.LString(err.Error())
		}
		if err := l.Call(yockr.YockFuncInfo{
			Fn:   fn,
			NRet: 1,
		}, path, info, e); err != nil {
			util.Ycho.Fatal(err.Error())
		}
		ok := l.CheckBool(-1)
		l.Pop(l.GetTop())
		if ok {
			return nil
		}
		return err
	})
	l.PushError(err)
	return 1
}
