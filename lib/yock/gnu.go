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
	luar "layeh.com/gopher-luar"
)

func LoadGNU(yocks yocki.YockScheduler) {
	yocks.RegYocksFn(yocki.YocksFuncs{
		"ssh": gnuSSH,
	})
	yocks.RegYockFn(yocki.YockFuns{
		"pwd":      gnuPwd,
		"whoami":   gnuWhoami,
		"echo":     gnuEcho,
		"ls":       gnuLs,
		"clear":    gnuClear,
		"chmod":    gnuChmod,
		"chown":    gnuChown,
		"cd":       gnuCd,
		"touch":    gnuTouch,
		"cat":      gnuCat,
		"mv":       gnuMv,
		"cp":       gnuCp,
		"mkdir":    gnuMkdir,
		"rm":       gnuRm,
		"alias":    gnuAlias,
		"unalias":  gnuUnalias,
		"sudo":     gnuSudo,
		"find":     gnuFind,
		"nohup":    gnuNohup,
		"ps":       gnuPS,
		"whereis":  gnuWhereis,
		"export":   gnuExport,
		"unset":    gnuUnset,
		"kill":     gnuKill,
		"pgrep":    gnuPGrep,
		"ifconfig": gnuIfconfig,
	})
	systemCtl := yocks.CreateLib("systemctl")
	systemCtl.SetYFunction(map[string]yocki.YGFunction{
		"list":    gnuSystemCtlList,
		"status":  gnuSystemCtlStatus,
		"stop":    gnuSystemCtlStop,
		"delete":  gnuSystemCtlDelete,
		"start":   gnuSystemCtlStart,
		"enable":  gnuSystemCtlEnable,
		"disable": gnuSystemCtlDisable,
		"create":  gnuSystemCtlCreate,
	})
	iptables := yocks.CreateLib("iptables")
	iptables.SetYFunction(map[string]yocki.YGFunction{
		"list": gnuIPTablesList,
		"add":  gnuIPTablesAdd,
		"del":  gnuIPTablesDel,
	})
}

// @param opt table
//
// @param cb function(*SSHClient)
//
// @return userdata (*SSHClient), err
func gnuSSH(yocks yocki.YockScheduler, state yocki.YockState) int {
	opt := yockc.SSHOpt{}
	if state.IsTable(1) {
		state.CheckTable(1).Bind(&opt)
		cli, err := yockc.NewSSHClient(opt)
		if err != nil {
			state.PushNil().Throw(err)
			return 2
		}
		if state.Argc() >= 2 && state.IsFunction(2) {
			fn := state.CheckFunction(2)
			if err := state.Call(yocki.YockFuncInfo{
				Fn: fn,
			}, cli); err != nil {
				ycho.Fatal(err)
			}
		}
		state.Pusha(cli)
	}
	state.PushNil()
	return 2
}

func gnuIPTablesList(s yocki.YockState) int {
	opt := yockc.IPTablesListOpt{}
	if err := s.CheckTable(1).Bind(&opt); err != nil {
		s.PushNilTable().Throw(err)
		return 2
	}
	rules, err := yockc.IPTablesList(opt)
	tbl := &lua.LTable{}
	for _, rule := range rules {
		tbl.Append(luar.New(s.LState(), rule))
	}
	s.Push(tbl).PushError(err)
	return 2
}

func gnuIPTablesAdd(s yocki.YockState) int {
	opt := yockc.IPTablesOpOpt{Op: yockc.IPTablesAdd}
	if err := s.CheckTable(1).Bind(&opt); err != nil {
		s.Throw(err)
		return 1
	}
	if err := yockc.IPTablesOp(opt); err != nil {
		s.Throw(err)
	}
	s.PushNil()
	return 1
}

func gnuIPTablesDel(s yocki.YockState) int {
	opt := yockc.IPTablesOpOpt{Op: yockc.IPTablesDel}
	if err := s.CheckTable(1).Bind(&opt); err != nil {
		s.Throw(err)
		return 1
	}
	if err := yockc.IPTablesOp(opt); err != nil {
		s.Throw(err)
	}
	s.PushNil()
	return 1
}

func gnuSystemCtlCreate(s yocki.YockState) int {
	opt := yockc.SCCreateOpt{}
	name := s.CheckString(1)
	if err := s.CheckTable(2).Bind(&opt); err != nil {
		s.Throw(err)
		return 1
	}
	err := yockc.SystemCtlCreate(name, opt)
	s.PushError(err)
	return 1
}

func gnuSystemCtlList(s yocki.YockState) int {
	var (
		optType   string
		optStatus string
	)
	switch s.Argc() {
	case 2:
		optStatus = s.CheckString(2)
		fallthrough
	case 1:
		optType = s.CheckString(1)
	}
	infos, err := yockc.SystemCtlStatus(yockc.SystemCtlStatusOpt{
		Name:   "",
		Type:   optType,
		Status: optStatus,
	})
	tbl := &lua.LTable{}
	for _, info := range infos {
		tbl.Append(luar.New(s.LState(), info))
	}
	s.Push(tbl).PushError(err)
	return 2
}

func gnuSystemCtlStatus(s yocki.YockState) int {
	infos, err := yockc.SystemCtlStatus(yockc.SystemCtlStatusOpt{
		Name: s.CheckString(1),
	})
	if len(infos) == 0 {
		s.PushNil().PushError(err)
	} else {
		s.Pusha(infos[0]).PushError(err)
	}
	return 2
}

func gnuSystemCtlStop(s yocki.YockState) int {
	s.PushError(yockc.SystemCtlStop(s.CheckString(1)))
	return 1
}

func gnuSystemCtlDelete(s yocki.YockState) int {
	s.PushError(yockc.SystemCtlDelete(s.CheckString(1)))
	return 1
}

func gnuSystemCtlStart(s yocki.YockState) int {
	s.PushError(yockc.SystemCtlStart(s.CheckString(1)))
	return 1
}

func gnuSystemCtlEnable(s yocki.YockState) int {
	s.PushBool(yockc.SystemCtlIsEnable(s.CheckString(1)))
	return 1
}

func gnuSystemCtlDisable(s yocki.YockState) int {
	s.PushError(yockc.SystemCtlDisable(s.CheckString(1)))
	return 1
}

func gnuIfconfig(s yocki.YockState) int {
	stats, err := util.Net().Interfaces()
	if err != nil {
		s.PushNilTable().Throw(err)
		return 2
	}
	if v, err := Decode(s.LState(), []byte(stats.String())); err == nil {
		s.Push(v).PushNil()
	} else {
		s.PushNilTable().Throw(err)
	}
	return 2
}

func gnuPGrep(s yocki.YockState) int {
	s.Pusha(yockc.PGrep(s.CheckString(1)))
	return 1
}

func gnuKill(s yocki.YockState) int {
	if s.IsString(1) {
		yockc.KillByName(s.CheckString(1))
	} else {
		yockc.KillByPid(int32(s.CheckInt(1)))
	}
	return 1
}

func gnuUnset(s yocki.YockState) int {
	err := yockc.Unset(s.CheckString(1))
	s.PushError(err)
	return 1
}

func gnuExport(s yocki.YockState) int {
	if s.Argc() > 1 {
		yockc.Export(yockc.ExportOpt{}, s.CheckString(1), s.CheckString(2))
	} else {
		kv := strings.SplitN(s.CheckString(1), ":", 2)
		if len(kv) == 2 {
			yockc.Export(yockc.ExportOpt{Expand: true}, kv[0], kv[1])
		} else {
			s.PushError(fmt.Errorf("invalid command"))
			return 1
		}
	}
	s.PushNil()
	return 1
}

// @return string, err
func gnuWhereis(s yocki.YockState) int {
	str, err := yockc.Whereis(s.CheckString(1))
	s.PushString(str).PushError(err)
	return 2
}

func gnuPS(s yocki.YockState) int {
	opt := yockc.PSOpt{}
	if err := s.CheckTable(1).Bind(&opt); err != nil {
		s.PushNil().Throw(err)
		return 2
	}
	info, err := yockc.PS(opt)
	s.Pusha(info).PushError(err)
	return 2
}

func gnuNohup(s yocki.YockState) int {
	s.PushError(yockc.Nohup(s.CheckString(1)))
	return 1
}

func gnuFind(s yocki.YockState) int {
	opt := yockc.FindOpt{Dir: true, File: true, Search: true}
	if s.IsTable(1) {
		err := s.CheckTable(1).Bind(&opt)
		if err != nil {
			s.PushNil().Throw(err)
			return 2
		}
		res, err := yockc.Find(opt, s.CheckString(2))
		if err != nil {
			s.PushNil().Throw(err)
			return 2
		}
		tbl := &lua.LTable{}
		for _, str := range res {
			tbl.Append(lua.LString(str))
		}
		s.Push(tbl).PushNil()
		return 2
	} else {
		opt.Search = false
		if _, err := yockc.Find(opt, s.CheckString(1)); err != nil {
			s.PushBool(false)
			return 1
		}
		s.PushBool(true)
		return 1
	}
}

func gnuSudo(s yocki.YockState) int {
	sudo := "sudo"
	if util.CurPlatform.OS == "windows" {
		sudo = filepath.Join(util.YockPath, "bin", "sudo.bat")
	}
	yockc.Exec(yockc.ExecOpt{Quiet: true}, sudo+" "+s.CheckString(1))
	return 0
}

// @param key string
//
// @return string
func gnuAlias(s yocki.YockState) int {
	yockc.Alias(s.CheckString(1), s.CheckString(2))
	return 0
}

// @param key string
func gnuUnalias(s yocki.YockState) int {
	for i := 1; i <= s.Argc(); i++ {
		yockc.Unalias(s.CheckString(i))
	}
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

// @param opt table|string
//
// @varag string
//
// @return string
func gnuEcho(s yocki.YockState) int {
	opt := yockc.EchoOpt{}
	if s.IsTable(1) {
		if err := s.CheckTable(1).Bind(&opt); err != nil {
			s.PushNil().Throw(err)
			return 2
		}
	} else {
		opt.Fd = []string{"stdout"}
	}
	tbl := &lua.LTable{}
	for i := 2; i <= s.Argc(); i++ {
		out, err := yockc.Echo(opt, s.CheckString(i))
		tbl.Append(lua.LString(out))
		if err != nil {
			s.Push(tbl).Throw(err)
			return 2
		}
	}
	s.Push(tbl).PushNil()
	return 2
}

// @param opt table
//
// @return table|string, err
func gnuLs(s yocki.YockState) int {
	opt := yockc.LsOpt{}
	if s.IsTable(1) {
		err := s.CheckTable(1).Bind(&opt)
		if err != nil {
			s.PushNil().Throw(err)
			return 2
		}
	} else {
		opt.Dir = s.CheckString(1)
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
	opt := yockc.CpOpt{Recurse: true, Info: func(name, args string) {
		if yocki.Y_MODE.Debug() {
			ycho.Infof("%s%s %s", s.Stacktrace(), name, args)
		}
	}}
	paths := []string{}
	var g_err error
	if s.IsTable(1) {
		if err := s.CheckTable(1).Bind(&opt); err != nil {
			s.Throw(err)
			return 1
		}
		if s.IsTable(2) {
			s.CheckLTable(2).ForEach(func(src, dst lua.LValue) {
				err := yockc.Cp(opt, src.String(), dst.String())
				if err != nil {
					if yocki.Y_MODE.Strict() {
						s.Throw(err)
					} else {
						g_err = util.ErrGeneral
					}
					if yocki.Y_MODE.Debug() {
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
			if yocki.Y_MODE.Debug() {
				ycho.Warnf(s.Stacktrace() + err.Error())
			}
			if yocki.Y_MODE.Strict() {
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
					if yocki.Y_MODE.Debug() {
						ycho.Warn(err)
					}
					if yocki.Y_MODE.Strict() {
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
		path := s.CheckString(i)
		if yocki.Y_MODE.Debug() {
			ycho.Infof("%smkdir %s", s.Stacktrace(), path)
		}
		err := util.SafeMkdirs(path)
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
	opt := yockc.RmOpt{Safe: true, Info: func(path string) {
		if yocki.Y_MODE.Debug() {
			ycho.Infof("%s%s", s.Stacktrace(), path)
		}
	}, Error: func(err error) error {
		if yocki.Y_MODE.Debug() {
			ycho.Warnf("%s%s", s.Stacktrace(), err)
		}
		return nil
	}}
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
		for i := 1; i <= s.Argc(); i++ {
			targets = append(targets, s.CheckString(i))
		}
	}
	s.PushError(yockc.Rm(opt, targets))
	return 1
}
