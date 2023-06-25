// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"

	yockc "github.com/ansurfen/yock/cmd"
	yocki "github.com/ansurfen/yock/interface"
	yockr "github.com/ansurfen/yock/runtime"
	"github.com/ansurfen/yock/util"
	lua "github.com/yuin/gopher-lua"
)

func LoadGNU(yocks yocki.YockScheduler) {
	yocks.RegYockFn(yocki.YockFuns{
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
		"mv":     gnuMv,
		"cp":     gnuCp,
		"mkdir":  gnuMkdir,
		"rm":     gnuRm,
	})
}

// @return path string
//
// @return err error
func gnuPwd(l *yockr.YockState) int {
	path, err := os.Getwd()
	l.PushString(path).PushError(err)
	return 2
}

// @return username string
//
// @return err error
func gnuWhoami(l *yockr.YockState) int {
	u, err := user.Current()
	l.PushString(u.Username).PushError(err)
	return 2
}

// @param str string
//
// @return string
func gnuEcho(l *yockr.YockState) int {
	str := l.CheckString(1)
	out, err := yockc.Echo(str)
	if err != nil {
		l.Throw(err)
		return 1
	}
	debug := true
	if l.GetTop() >= 2 && l.IsBool(2) {
		debug = l.CheckBool(2)
	}
	if debug {
		fmt.Println(out)
	}
	l.PushString(out)
	return 1
}

// @param opt table
//
// @return table|string, err
func gnuLs(l *yockr.YockState) int {
	var opt yockc.LsOpt
	err := l.CheckTable(1).Bind(&opt)
	if err != nil {
		l.PushNil().Throw(err)
		return 2
	}
	st, str, err := yockc.Ls(opt)
	if opt.Str {
		l.PushString(str)
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
	l.PushError(err)
	return 2
}

// gnuClear clears the output on the screen
func gnuClear(l *yockr.YockState) int {
	yockc.Clear()
	return 0
}

/*
* @param name string
* @param mode number
* @return err
 */
func gnuChmod(l *yockr.YockState) int {
	mode, err := strconv.ParseInt(strconv.Itoa(int(l.CheckNumber(2))), 8, 64)
	if err != nil {
		l.Throw(err)
		return 1
	}
	err = yockc.Chmod(l.CheckString(1), mode)
	l.PushError(err)
	return 1
}

/*
* @param name string
* @param uid number
* @param gid number
* @return err
 */
func gnuChown(l *yockr.YockState) int {
	err := yockc.Chown(l.CheckString(1), int(l.CheckNumber(2)), int(l.CheckNumber(3)))
	l.PushError(err)
	return 1
}

/*
* @param dir string
* @return err
 */
func gnuCd(l *yockr.YockState) int {
	err := yockc.Cd(l.CheckString(1))
	l.PushError(err)
	return 1
}

// @param file string
//
// @return err
func gnuTouch(l *yockr.YockState) int {
	err := util.SafeWriteFile(l.CheckString(1), nil)
	l.PushError(err)
	return 1
}

// @param file string
//
// @return string, err
func gnuCat(l *yockr.YockState) int {
	out, err := util.ReadStraemFromFile(l.CheckString(1))
	l.Push(lua.LString(string(out)))
	l.PushError(err)
	return 2
}

/*
* @param opt table
* @param src string
* @param dst string
* @return err
 */
func gnuMv(l *yockr.YockState) int {
	err := yockc.Mv(yockc.MvOpt{}, l.CheckString(1), l.CheckString(2))
	l.PushError(err)
	return 1
}

/*
* @param opt table
* @param src string
* @param dst string
* @return err
 */
func gnuCp(l *yockr.YockState) int {
	opt := yockc.CpOpt{Recurse: true}
	paths := []string{}
	var g_err error
	if l.IsTable(1) {
		if err := l.CheckTable(1).Bind(&opt); err != nil {
			l.Throw(err)
			return 1
		}
		if l.IsTable(2) {
			l.CheckTable(2).ForEach(func(src, dst lua.LValue) {
				err := yockc.Cp(opt, src.String(), dst.String())
				if err != nil {
					if opt.Strict {
						// TODO
					} else {
						g_err = util.ErrGeneral
					}
					if opt.Debug {
						util.Ycho.Warn(l.Stacktrace() + err.Error())
					}
				}
			})
			l.PushError(g_err)
			return 1
		} else {
			for i := 2; i <= l.GetTop(); i++ {
				paths = append(paths, l.CheckString(i))
			}
		}
	} else {
		for i := 1; i <= l.GetTop(); i++ {
			paths = append(paths, l.CheckString(i))
		}
	}
	if len(paths) >= 2 {
		err := yockc.Cp(opt, paths[0], paths[1])
		if err != nil {
			g_err = err
			if opt.Debug {
				util.Ycho.Warn(l.Stacktrace() + err.Error())
			}
			if opt.Strict {
				// TODO
			} else {
				g_err = util.ErrGeneral
			}
		}
		l.PushError(g_err)
	} else {
		util.ReadLineFromString(paths[0], func(s string) string {
			if len(s) == 0 {
				return ""
			}
			kv := strings.Split(s, " ")
			if len(kv) == 2 {
				err := yockc.Cp(opt, kv[0], kv[1])
				if err != nil {
					g_err = err
					if opt.Debug {
						util.Ycho.Warn(err.Error())
					}
					if opt.Strict {
						// TODO
					} else {
						g_err = util.ErrGeneral
					}
				}
			}
			return ""
		})

	}
	l.PushError(g_err)
	return 1
}

// @param path string
//
// @return err
func gnuMkdir(l *yockr.YockState) int {
	var g_err error
	for i := 1; i <= l.GetTop(); i++ {
		err := util.SafeMkdirs(l.CheckString(i))
		if err != nil {
			util.Ycho.Warn(err.Error())
			g_err = err
		}
	}
	l.PushError(g_err)
	return 1
}

// @param opt table
//
// @param files ...string
//
// @return err
func gnuRm(l *yockr.YockState) int {
	opt := yockc.RmOpt{Safe: true}
	targets := []string{}
	if l.IsTable(1) {
		if err := l.CheckTable(1).Bind(&opt); err != nil {
			l.Throw(err)
			return 1
		}
		for i := 2; i <= l.GetTop(); i++ {
			targets = append(targets, l.CheckString(i))
		}
	} else {
		for i := 1; i < l.GetTop(); i++ {
			targets = append(targets, l.CheckString(i))
		}
	}
	l.PushError(yockc.Rm(opt, targets))
	return 1
}
