package scheduler

import (
	"regexp"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func loadRegexp(vm *YockScheduler) lua.LValue {
	reg := &lua.LTable{}
	reg.RawSetString("Compile", vm.Interp().NewClosure(func(l *lua.LState) int {
		r, err := regexp.Compile(l.CheckString(1))
		l.Push(luar.New(l, r))
		handleErr(l, err)
		return 2
	}))
	reg.RawSetString("MustCompile", vm.Interp().NewClosure(func(l *lua.LState) int {
		r := regexp.MustCompile(l.CheckString(1))
		l.Push(luar.New(l, r))
		return 1
	}))
	return reg 
}
