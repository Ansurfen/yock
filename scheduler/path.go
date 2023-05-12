package scheduler

import (
	"path/filepath"

	"github.com/ansurfen/cushion/utils"
	lua "github.com/yuin/gopher-lua"
)

func loadPath(vm *YockScheduler) {
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
	vm.SetGlobalVar("path", path)
}
