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
		"lsof":     gnuLsof,
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

func gnuLsof(s yocki.YockState) int {
	infos, err := yockc.Lsof()
	if err != nil {
		s.PushNilTable()
		return 1
	}
	port := ""
	if s.Argc() >= 1 {
		port = strconv.Itoa(int(s.CheckNumber(1)))
	}
	tbl := &lua.LTable{}
	if n := len(infos); n == 1 {
		tbl.RawSetString("pid", lua.LNumber(infos[0].Pid))
		tbl.RawSetString("state", lua.LString(infos[0].State))
		tbl.RawSetString("proto", lua.LString(infos[0].Protocal))
		tbl.RawSetString("Local", lua.LString(infos[0].Local))
		tbl.RawSetString("foreign", lua.LString(infos[0].Foreign))
	} else if n > 1 {
		for _, info := range infos {
			if len(port) != 0 {
				if !strings.HasSuffix(info.Local, port) {
					continue
				}
			}
			linfo := &lua.LTable{}
			tbl.Append(linfo)
			linfo.RawSetString("pid", lua.LNumber(info.Pid))
			linfo.RawSetString("state", lua.LString(info.State))
			linfo.RawSetString("proto", lua.LString(info.Protocal))
			linfo.RawSetString("Local", lua.LString(info.Local))
			linfo.RawSetString("foreign", lua.LString(info.Foreign))
		}
	}
	s.Push(tbl)
	return 1
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
	if err != nil {
		s.PushNilTable().Throw(err)
		return 2
	}
	tbl := &lua.LTable{}
	if len(rules) > 1 {
		for _, rule := range rules {
			r := &lua.LTable{}
			r.RawSetString("name", lua.LString(rule.Name()))
			r.RawSetString("proto", lua.LString(rule.Proto()))
			r.RawSetString("src", lua.LString(rule.Src()))
			r.RawSetString("dst", lua.LString(rule.Dst()))
			r.RawSetString("action", lua.LString(rule.Action()))
			tbl.Append(r)
		}
	} else {
		tbl.RawSetString("name", lua.LString(rules[0].Name()))
		tbl.RawSetString("proto", lua.LString(rules[0].Proto()))
		tbl.RawSetString("src", lua.LString(rules[0].Src()))
		tbl.RawSetString("dst", lua.LString(rules[0].Dst()))
		tbl.RawSetString("action", lua.LString(rules[0].Action()))
	}
	s.Push(tbl).PushError(err)
	return 2
}

func gnuIPTablesAdd(s yocki.YockState) int {
	opt := yockc.IPTablesOpOpt{Op: yockc.IPTablesAdd}
	if err := s.CheckTable(1).Bind(&opt); err != nil {
		ychoLogger(err, "%siptables add", s.Stacktrace())
		s.Throw(err)
		return 1
	}
	err := yockc.IPTablesOp(opt)
	ychoLogger(err, "%siptables add", s.Stacktrace())
	if err != nil {
		s.Throw(err)
	} else {
		s.PushNil()
	}
	return 1
}

func gnuIPTablesDel(s yocki.YockState) int {
	opt := yockc.IPTablesOpOpt{Op: yockc.IPTablesDel}
	if err := s.CheckTable(1).Bind(&opt); err != nil {
		ychoLogger(err, "%siptables delete", s.Stacktrace())
		s.Throw(err)
		return 1
	}
	err := yockc.IPTablesOp(opt)
	ychoLogger(err, "%siptables delete", s.Stacktrace())
	if err != nil {
		s.Throw(err)
	} else {
		s.PushNil()
	}
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
		sinfo := &lua.LTable{}
		sinfo.RawSetString("pid", lua.LNumber(info.PID()))
		sinfo.RawSetString("name", lua.LString(info.Name()))
		sinfo.RawSetString("status", lua.LString(info.Status()))
		tbl.Append(sinfo)
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
		tbl := &lua.LTable{}
		tbl.RawSetString("pid", lua.LNumber(infos[0].PID()))
		tbl.RawSetString("name", lua.LString(infos[0].Name()))
		tbl.RawSetString("status", lua.LString(infos[0].Status()))
		s.Push(tbl).PushError(err)
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
	process := yockc.PGrep(s.CheckString(1))
	tbl := &lua.LTable{}
	for i := 0; i < len(process); i++ {
		pinfo := &lua.LTable{}
		pinfo.RawSetString("pid", lua.LNumber(process[i].Pid))
		pinfo.RawSetString("name", lua.LString(process[i].Name))
		tbl.Append(pinfo)
	}
	s.Push(tbl)
	return 1
}

func gnuKill(s yocki.YockState) int {
	if s.IsString(1) {
		name := s.CheckString(1)
		err := yockc.KillByName(name)
		ychoLogger(err, "%skill %s", s.Stacktrace(), name)
		s.PushError(err)
	} else {
		id := s.CheckInt(1)
		err := yockc.KillByPid(int32(id))
		ychoLogger(err, "%skill %d", s.Stacktrace(), id)
		s.PushError(err)
	}
	return 1
}

func gnuUnset(s yocki.YockState) int {
	name := s.CheckString(1)
	err := yockc.Unset(name)
	ychoLogger(err, "%sunset %s", s.Stacktrace(), name)
	s.PushError(err)
	return 1
}

func gnuExport(s yocki.YockState) int {
	if s.Argc() > 1 {
		k := s.CheckString(1)
		v := s.CheckString(2)
		err := yockc.Export(yockc.ExportOpt{}, k, v)
		ychoLogger(err, "%sexport %s=%s", s.Stacktrace(), k, v)
	} else {
		kv := strings.SplitN(s.CheckString(1), ":", 2)
		if len(kv) == 2 {
			err := yockc.Export(yockc.ExportOpt{Expand: true}, kv[0], kv[1])
			ychoLogger(err, "%sexport %s=$%s:%s", s.Stacktrace(), kv[0], kv[0], kv[1])
		} else {
			err := fmt.Errorf("invalid command")
			ychoLogger(err, "export")
			s.PushError(err)
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
	p := -1
	cmd := ""
	if s.Argc() >= 1 {
		if s.IsTable(1) {
			if err := s.CheckTable(1).Bind(&opt); err != nil {
				s.PushNil().Throw(err)
				return 2
			}
		} else if s.IsNumber(1) {
			p = int(s.CheckNumber(1))
		} else if s.IsString(1) {
			cmd = s.CheckString(1)
		}
	}
	infos, err := yockc.PS(opt)
	tbl := &lua.LTable{}
	for pid, info := range infos {
		if p != -1 && p != int(pid) {
			continue
		}
		if len(cmd) != 0 && !strings.Contains(info.Cmd, cmd) {
			continue
		}
		pinfo := &lua.LTable{}
		pinfo.RawSetString("cmd", lua.LString(info.Cmd))
		pinfo.RawSetString("name", lua.LString(info.Name))
		if opt.User {
			pinfo.RawSetString("user", lua.LString(info.User))
		}
		if opt.CPU {
			pinfo.RawSetString("cpu", lua.LNumber(info.CPU))
		}
		if opt.Time {
			pinfo.RawSetString("start", lua.LNumber(info.Start))
		}
		tbl.RawSet(lua.LNumber(pid), pinfo)
	}
	s.Push(tbl).PushError(err)
	return 2
}

func gnuNohup(s yocki.YockState) int {
	cmd := s.CheckString(1)
	err := yockc.Nohup(cmd)
	ychoLogger(err, "%snohup %s", s.Stacktrace(), cmd)
	s.PushError(err)
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
	cmd := s.CheckString(1)
	_, err := yockc.Exec(yockc.ExecOpt{Quiet: true}, sudo+" "+cmd)
	ychoLogger(err, "%ssudo %s", s.Stacktrace(), cmd)
	return 0
}

// @param key string
//
// @return string
func gnuAlias(s yocki.YockState) int {
	k := s.CheckString(1)
	v := s.CheckString(2)
	ychoLogger(nil, "alias %s %s", k, v)
	yockc.Alias(k, v)
	return 0
}

// @param key string
func gnuUnalias(s yocki.YockState) int {
	for i := 1; i <= s.Argc(); i++ {
		k := s.CheckString(i)
		ychoLogger(nil, "unalias %s", k)
		yockc.Unalias(k)
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
	start := 1
	if s.IsTable(1) {
		if err := s.CheckTable(1).Bind(&opt); err != nil {
			s.PushNil().Throw(err)
			return 2
		}
		start++
	} else {
		opt.Fd = []string{"stdout"}
	}
	tbl := &lua.LTable{}
	for i := start; i <= s.Argc(); i++ {
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

// @param dir string
//
// @return err
func gnuCd(s yocki.YockState) int {
	wd := s.CheckString(1)
	err := yockc.Cd(wd)
	ychoLogger(err, "%scd %s", s.Stacktrace(), wd)
	s.PushError(err)
	return 1
}

func ychoLogger(err error, format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	if yocki.Y_MODE.Strict() && err != nil {
		ycho.Errorf(msg)
		panic(err)
	} else if err != nil {
		ycho.Warnf("%s\n%s", msg, err.Error())
	} else {
		ycho.Info(msg)
	}
}

// @param file string
//
// @return err
func gnuTouch(s yocki.YockState) int {
	file := s.CheckString(1)
	err := util.SafeWriteFile(file, nil)
	ychoLogger(err, "%stouch %s", s.Stacktrace(), file)
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
	src := s.CheckString(1)
	dst := s.CheckString(2)
	err := yockc.Mv(yockc.MvOpt{}, src, dst)
	ychoLogger(err, "%smv %s %s", s.Stacktrace(), src, dst)
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
			s.CheckLTable(2).ForEach(func(src, dst lua.LValue) {
				err := yockc.Cp(opt, src.String(), dst.String())
				ychoLogger(err, fmt.Sprintf("%scp %s %s", s.Stacktrace(), src.String(), dst.String()))
				if err != nil {
					g_err = err
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
		g_err = yockc.Cp(opt, paths[0], paths[1])
		ychoLogger(g_err, fmt.Sprintf("%scp %s %s", s.Stacktrace(), paths[0], paths[1]))
	} else {
		util.ReadLineFromString(paths[0], func(str string) string {
			if len(str) == 0 {
				return ""
			}
			kv := strings.Split(str, " ")
			if len(kv) == 2 {
				err := yockc.Cp(opt, kv[0], kv[1])
				ychoLogger(err, fmt.Sprintf("%scp %s %s", s.Stacktrace(), kv[0], kv[1]))
				if err != nil {
					g_err = err
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
	var err error
	for i := 1; i <= s.Argc(); i++ {
		path := s.CheckString(i)
		err = util.SafeMkdirs(path)
		ychoLogger(err, "%smkdir %s", s.Stacktrace(), path)
	}
	s.PushError(err)
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
		for i := 1; i <= s.Argc(); i++ {
			targets = append(targets, s.CheckString(i))
		}
	}
	var err error
	for _, t := range targets {
		err = yockc.Rm(opt, t)
		ychoLogger(err, "%srm %s", s.Stacktrace(), t)
	}
	s.PushError(err)
	return 1
}
