// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/util"
	lua "github.com/yuin/gopher-lua"
)

var ioFuncs = luaFuncs{
	"safe_write": safe_write,
	"zip":        zip,
	"unzip":      unzip,
	"write_file": write_file,
	"is_exist":   is_exist,
	"printf":     printf,
	"pathf":      ioPathf,
}

/*
* @param file string
* @param data string
* @return err
 */
func safe_write(l *lua.LState) int {
	err := utils.SafeWriteFile(l.CheckString(1), []byte(l.CheckString(2)))
	handleErr(l, err)
	return 1
}

/*
* @param src string
* @param dst string
* @return err
 */
func zip(l *lua.LState) int {
	zipPath := l.CheckString(1)
	paths := []string{}
	for i := 2; i <= l.GetTop(); i++ {
		paths = append(paths, l.CheckString(i))
	}
	err := utils.Zip(zipPath, paths...)
	handleErr(l, err)
	return 1
}

/*
* @param src string
* @param dst string
* @return err
 */
func unzip(l *lua.LState) int {
	err := utils.Unzip(l.CheckString(1), l.CheckString(2))
	handleErr(l, err)
	return 1
}

/*
* @param file string
* @param data string
* @return err
 */
func write_file(l *lua.LState) int {
	err := utils.WriteFile(l.CheckString(1), []byte(l.CheckString(2)))
	handleErr(l, err)
	return 1
}

// @param path string
//
// @return bool
func is_exist(l *lua.LState) int {
	ok := utils.IsExist(l.CheckString(1))
	handleBool(l, ok)
	return 1
}

// @param title string
//
// @param rows string[][]
func printf(l *lua.LState) int {
	title := []string{}
	rows := [][]string{}
	l.CheckTable(1).ForEach(func(idx, el lua.LValue) {
		title = append(title, el.String())
	})
	l.CheckTable(2).ForEach(func(ri, row lua.LValue) {
		tmp := []string{}
		row.(*lua.LTable).ForEach(func(fi, field lua.LValue) {
			tmp = append(tmp, field.String())
		})
		rows = append(rows, tmp)
	})
	utils.Prinf(utils.PrintfOpt{MaxLen: 30}, title, rows)
	return 0
}

// ioPathf formats path
//
// @/abc => {WorkSpace}/abc (WorkSpace = UserHome + .yock)
//
// ~/abc => {YockPath}/abc (YockPath = executable file path)
/*
* @param path string
* @return string
 */
func ioPathf(l *lua.LState) int {
	l.Push(lua.LString(util.Pathf(l.CheckString(1))))
	return 1
}
