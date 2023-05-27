// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"fmt"
	"os"
	"os/user"
	"strconv"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/cmd"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

// gnuFuncs provides some simple GNU functions.
// Scripters can use cross-platform GNU commands in the form of global functions in Lua.
// For specific parameters and functions, see docs.
var gnuFuncs = luaFuncs{
	"pwd":    gnuPwd,
	"whoami": gnuWhoami,
	"echo":   gnuEcho,
	"ls":     gnuLs,
	"clear":  gnuClear,
	"chmod":  gnuChmod,
	"chown":  gnuChown,
	"cd":     gnuCd,
	"touch":  gnuTouch,
	"cat":    gnuCat,
	"cmd":    gnuCmd,
	"mv":     gnuMv,
	"cp":     gnuCp,
	"mkdir":  gnuMkdir,
	"rm":     gnuRm,
}

// @return path string
//
// @return err error
func gnuPwd(l *lua.LState) int {
	path, err := os.Getwd()
	l.Push(lua.LString(path))
	handleErr(l, err)
	return 2
}

// @return username string
//
// @return err error
func gnuWhoami(l *lua.LState) int {
	u, err := user.Current()
	l.Push(lua.LString(u.Username))
	handleErr(l, err)
	return 2
}

// @param str string
//
// @return string
func gnuEcho(l *lua.LState) int {
	str := l.CheckString(1)
	out, err := cmd.Echo(str)
	if err != nil {
		l.Push(lua.LString(""))
		return 1
	}
	debug := true
	if l.GetTop() >= 2 && l.CheckAny(2).Type() == lua.LTBool {
		debug = l.CheckBool(2)
	}
	if debug {
		fmt.Println(out)
	}
	l.Push(lua.LString(out))
	return 1
}

// @param opt table
//
// @return table|string, err
func gnuLs(l *lua.LState) int {
	var opt cmd.LsOpt
	gluamapper.Map(l.CheckTable(1), &opt)
	st, str, err := cmd.Ls(opt)
	if opt.Str {
		l.Push(lua.LString(str))
	} else {
		fileinfos := &lua.LTable{}
		for idx, info := range st {
			linfo := &lua.LTable{}
			linfo.Insert(1, lua.LString(info.Perm))
			linfo.Insert(2, lua.LNumber(info.Size))
			linfo.Insert(3, lua.LString(info.ModTime))
			linfo.Insert(4, lua.LString(info.Filename))
			fileinfos.Insert(idx+1, linfo)
		}
		l.Push(fileinfos)
	}
	handleErr(l, err)
	return 2
}

// gnuClear clears the output on the screen
func gnuClear(l *lua.LState) int {
	cmd.Clear()
	return 0
}

/*
* @param name string
* @param mode number
* @return err
 */
func gnuChmod(l *lua.LState) int {
	mode, err := strconv.ParseInt(strconv.Itoa(int(l.CheckNumber(2))), 8, 64)
	if err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}
	err = cmd.Chmod(l.CheckString(1), mode)
	handleErr(l, err)
	return 1
}

/*
* @param name string
* @param uid number
* @param gid number
* @return err
 */
func gnuChown(l *lua.LState) int {
	err := cmd.Chown(l.CheckString(1), int(l.CheckNumber(2)), int(l.CheckNumber(3)))
	handleErr(l, err)
	return 1
}

/*
* @param dir string
* @return err
 */
func gnuCd(l *lua.LState) int {
	err := cmd.Cd(l.CheckString(1))
	handleErr(l, err)
	return 1
}

// @param file string
//
// @return err
func gnuTouch(l *lua.LState) int {
	err := utils.SafeWriteFile(l.CheckString(1), nil)
	handleErr(l, err)
	return 1
}

// @param file string
//
// @return string, err
func gnuCat(l *lua.LState) int {
	out, err := utils.ReadStraemFromFile(l.CheckString(1))
	l.Push(lua.LString(string(out)))
	handleErr(l, err)
	return 2
}

// @param opt string
//
// @param cmd string
//
// @return string, err
func gnuCmd(l *lua.LState) int {
	out, err := cmd.Cmd(cmd.ExecOpt{Redirect: false, Quiet: true}, l.CheckString(1))
	l.Push(lua.LString(out))
	handleErr(l, err)
	return 2
}

/*
* @param opt table
* @param src string
* @param dst string
* @return err
 */
func gnuMv(l *lua.LState) int {
	err := cmd.Mv(cmd.MvOpt{}, l.CheckString(1), l.CheckString(2))
	handleErr(l, err)
	return 1
}

/*
* @param opt table
* @param src string
* @param dst string
* @return err
 */
func gnuCp(l *lua.LState) int {
	err := cmd.Cp(cmd.CpOpt{
		Recurse: true,
	}, l.CheckString(1), l.CheckString(2))
	handleErr(l, err)
	return 1
}

// @param path string
//
// @return err
func gnuMkdir(l *lua.LState) int {
	err := utils.SafeMkdirs(l.CheckString(1))
	handleErr(l, err)
	return 1
}

// @param opt table
//
// @param files ...string
//
// @return err
func gnuRm(l *lua.LState) int {
	mode := l.CheckAny(1)
	opt := cmd.RmOpt{Safe: true}
	targets := []string{}
	if mode.Type() == lua.LTTable {
		gluamapper.Map(l.CheckTable(1), &opt)
		for i := 2; i <= l.GetTop(); i++ {
			targets = append(targets, l.CheckString(i))
		}
	} else {
		for i := 1; i < l.GetTop(); i++ {
			targets = append(targets, l.CheckString(i))
		}
	}
	handleErr(l, cmd.Rm(opt, targets))
	return 1
}
