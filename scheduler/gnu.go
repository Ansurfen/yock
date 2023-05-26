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

func gnuPwd(l *lua.LState) int {
	path, err := os.Getwd()
	l.Push(lua.LString(path))
	handleErr(l, err)
	return 2
}

func gnuWhoami(l *lua.LState) int {
	u, err := user.Current()
	l.Push(lua.LString(u.Username))
	handleErr(l, err)
	return 2
}

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

func gnuClear(l *lua.LState) int {
	cmd.Clear()
	return 0
}

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

func gnuChown(l *lua.LState) int {
	err := cmd.Chown(l.CheckString(1), int(l.CheckNumber(2)), int(l.CheckNumber(3)))
	handleErr(l, err)
	return 1
}

func gnuCd(l *lua.LState) int {
	err := cmd.Cd(l.CheckString(1))
	handleErr(l, err)
	return 1
}

func gnuTouch(l *lua.LState) int {
	err := utils.SafeWriteFile(l.CheckString(1), nil)
	handleErr(l, err)
	return 1
}

func gnuCat(l *lua.LState) int {
	out, err := utils.ReadStraemFromFile(l.CheckString(1))
	l.Push(lua.LString(string(out)))
	handleErr(l, err)
	return 2
}

func gnuCmd(l *lua.LState) int {
	out, err := cmd.Cmd(cmd.ExecOpt{Redirect: false, Quiet: true}, l.CheckString(1))
	l.Push(lua.LString(out))
	handleErr(l, err)
	return 2
}

func gnuMv(l *lua.LState) int {
	err := cmd.Mv(cmd.MvOpt{}, l.CheckString(1), l.CheckString(2))
	handleErr(l, err)
	return 1
}

func gnuCp(l *lua.LState) int {
	err := cmd.Cp(cmd.CpOpt{
		Recurse: true,
	}, l.CheckString(1), l.CheckString(2))
	handleErr(l, err)
	return 1
}

func gnuMkdir(l *lua.LState) int {
	err := utils.SafeMkdirs(l.CheckString(1))
	handleErr(l, err)
	return 1
}

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
	cmd.Rm(opt, targets)
	return 0
}
