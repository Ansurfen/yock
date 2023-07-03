// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	yockc "github.com/ansurfen/yock/cmd"
	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
	lua "github.com/yuin/gopher-lua"
)

var aliases map[string]string

func init() {
	aliases = make(map[string]string)
}

func LoadGNU(yocks yocki.YockScheduler) {
	yocks.RegYockFn(yocki.YockFuns{
		"pwd":     gnuPwd,
		"whoami":  gnuWhoami,
		"echo":    gnuEcho,
		"ls":      gnuLs,
		"clear":   gnuClear,
		"chmod":   gnuChmod,
		"chown":   gnuChown,
		"cd":      gnuCd,
		"touch":   gnuTouch,
		"cat":     gnuCat,
		"mv":      gnuMv,
		"cp":      gnuCp,
		"mkdir":   gnuMkdir,
		"rm":      gnuRm,
		"alias":   gnuAlias,
		"unalias": gnuUnalias,
		"sudo":    gnuSudo,
	})
}

func gnuSudo(s yocki.YockState) int {
	if util.CurPlatform.OS == "windows" {
		sudo := filepath.Join(util.YockPath, "bin", "sudo.bat")
		yockc.Exec(yockc.ExecOpt{}, sudo + " " + s.CheckString(1))
	}
	return 0
}

// @param key string
//
// @return string
func gnuAlias(s yocki.YockState) int {
	aliases[s.CheckString(1)] = s.CheckString(2)
	return 0
}

// @param key string
func gnuUnalias(s yocki.YockState) int {
	delete(aliases, s.CheckString(1))
	return 0
}

// @return path string
//
// @return err error
func gnuPwd(s yocki.YockState) int {
	path, err := os.Getwd()
	s.PushString(path).PushError(err)
	return 2
}

// @return username string
//
// @return err error
func gnuWhoami(s yocki.YockState) int {
	u, err := user.Current()
	s.PushString(u.Username).PushError(err)
	return 2
}

// @param str string
//
// @return string
func gnuEcho(s yocki.YockState) int {
	str := s.CheckString(1)
	out, err := yockc.Echo(str)
	if err != nil {
		s.Throw(err)
		return 1
	}
	debug := true
	if s.Argc() >= 2 && s.IsBool(2) {
		debug = s.CheckBool(2)
	}
	if debug {
		fmt.Println(out)
	}
	s.PushString(out)
	return 1
}

// @param opt table
//
// @return table|string, err
func gnuLs(s yocki.YockState) int {
	var opt yockc.LsOpt
	err := s.CheckTable(1).Bind(&opt)
	if err != nil {
		s.PushNil().Throw(err)
		return 2
	}
	st, str, err := yockc.Ls(opt)
	if opt.Str {
		s.PushString(str)
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
		s.Push(fileinfos)
	}
	s.PushError(err)
	return 2
}

// gnuClear clears the output on the screen
func gnuClear(s yocki.YockState) int {
	yockc.Clear()
	return 0
}

/*
* @param name string
* @param mode number
* @return err
 */
func gnuChmod(s yocki.YockState) int {
	mode, err := strconv.ParseInt(strconv.Itoa(s.CheckInt(2)), 8, 64)
	if err != nil {
		s.Throw(err)
		return 1
	}
	err = yockc.Chmod(s.CheckString(1), mode)
	s.PushError(err)
	return 1
}

/*
* @param name string
* @param uid number
* @param gid number
* @return err
 */
func gnuChown(s yocki.YockState) int {
	err := yockc.Chown(s.CheckString(1), s.CheckInt(2), s.CheckInt(3))
	s.PushError(err)
	return 1
}

/*
* @param dir string
* @return err
 */
func gnuCd(s yocki.YockState) int {
	err := yockc.Cd(s.CheckString(1))
	s.PushError(err)
	return 1
}

// @param file string
//
// @return err
func gnuTouch(s yocki.YockState) int {
	err := util.SafeWriteFile(s.CheckString(1), nil)
	s.PushError(err)
	return 1
}

// @param file string
//
// @return string, err
func gnuCat(s yocki.YockState) int {
	out, err := util.ReadStraemFromFile(s.CheckString(1))
	s.PushString(string(out)).PushError(err)
	return 2
}

/*
* @param opt table
* @param src string
* @param dst string
* @return err
 */
func gnuMv(s yocki.YockState) int {
	err := yockc.Mv(yockc.MvOpt{}, s.CheckString(1), s.CheckString(2))
	s.PushError(err)
	return 1
}

/*
* @param opt table
* @param src string
* @param dst string
* @return err
 */
func gnuCp(s yocki.YockState) int {
	opt := yockc.CpOpt{Recurse: true}
	paths := []string{}
	var g_err error
	if s.IsTable(1) {
		if err := s.CheckTable(1).Bind(&opt); err != nil {
			s.Throw(err)
			return 1
		}
		if s.IsTable(2) {
			s.CheckTable(2).Value().ForEach(func(src, dst lua.LValue) {
				err := yockc.Cp(opt, src.String(), dst.String())
				if err != nil {
					if opt.Strict {
						// TODO
					} else {
						g_err = util.ErrGeneral
					}
					if opt.Debug {
						ycho.Warnf(s.Stacktrace() + err.Error())
					}
				}
			})
			s.PushError(g_err)
			return 1
		} else {
			for i := 2; i <= s.Argc(); i++ {
				paths = append(paths, s.CheckString(i))
			}
		}
	} else {
		for i := 1; i <= s.Argc(); i++ {
			paths = append(paths, s.CheckString(i))
		}
	}
	if len(paths) >= 2 {
		err := yockc.Cp(opt, paths[0], paths[1])
		if err != nil {
			g_err = err
			if opt.Debug {
				ycho.Warnf(s.Stacktrace() + err.Error())
			}
			if opt.Strict {
				// TODO
			} else {
				g_err = util.ErrGeneral
			}
		}
		s.PushError(g_err)
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
						ycho.Warn(err)
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
	s.PushError(g_err)
	return 1
}

// @param path string
//
// @return err
func gnuMkdir(s yocki.YockState) int {
	var g_err error
	for i := 1; i <= s.Argc(); i++ {
		err := util.SafeMkdirs(s.CheckString(i))
		if err != nil {
			ycho.Warn(err)
			g_err = err
		}
	}
	s.PushError(g_err)
	return 1
}

// @param opt table
//
// @param files ...string
//
// @return err
func gnuRm(s yocki.YockState) int {
	opt := yockc.RmOpt{Safe: true}
	targets := []string{}
	if s.IsTable(1) {
		if err := s.CheckTable(1).Bind(&opt); err != nil {
			s.Throw(err)
			return 1
		}
		for i := 2; i <= s.Argc(); i++ {
			targets = append(targets, s.CheckString(i))
		}
	} else {
		for i := 1; i < s.Argc(); i++ {
			targets = append(targets, s.CheckString(i))
		}
	}
	s.PushError(yockc.Rm(opt, targets))
	return 1
}
