package scheduler

import (
	"path/filepath"

	"github.com/ansurfen/cushion/utils"
	lua "github.com/yuin/gopher-lua"
)

func loadPath(vm *YockScheduler) lua.LValue {
	path := &lua.LTable{}
	path.RawSetString("exist", vm.Interp().NewClosure(func(l *lua.LState) int {
		ok := utils.IsExist(l.CheckString(1))
		if ok {
			l.Push(lua.LTrue)
		} else {
			l.Push(lua.LFalse)
		}
		return 1
	}))
	path.RawSetString("filename", vm.Interp().NewClosure(func(l *lua.LState) int {
		l.Push(lua.LString(utils.Filename(l.CheckString(1))))
		return 1
	}))
	path.RawSetString("join", vm.Interp().NewClosure(func(l *lua.LState) int {
		elem := []string{}
		for i := 1; i <= l.GetTop(); i++ {
			elem = append(elem, l.CheckString(i))
		}
		l.Push(lua.LString(filepath.Join(elem...)))
		return 1
	}))
	path.RawSetString("dir", vm.Interp().NewClosure(func(l *lua.LState) int {
		l.Push(lua.LString(filepath.Dir(l.CheckString(1))))
		return 1
	}))
	path.RawSetString("base", vm.Interp().NewClosure(func(l *lua.LState) int {
		l.Push(lua.LString(filepath.Base(l.CheckString(1))))
		return 1
	}))
	path.RawSetString("clean", vm.Interp().NewClosure(func(l *lua.LState) int {
		l.Push(lua.LString(filepath.Clean(l.CheckString(1))))
		return 1
	}))
	path.RawSetString("ext", vm.Interp().NewClosure(func(l *lua.LState) int {
		l.Push(lua.LString(filepath.Ext(l.CheckString(1))))
		return 1
	}))
	path.RawSetString("abs", vm.Interp().NewClosure(func(l *lua.LState) int {
		abs, err := filepath.Abs(l.CheckString(1))
		l.Push(lua.LString(abs))
		if err != nil {
			l.Push(lua.LString(err.Error()))
		} else {
			l.Push(lua.LString(""))
		}
		return 2
	}))
	return path
}
