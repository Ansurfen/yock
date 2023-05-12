package scheduler

import (
	"strings"

	"github.com/ansurfen/cushion/runtime"
	lua "github.com/yuin/gopher-lua"
)

func loadStrings(vm *YockScheduler) runtime.Handles {
	str := &lua.LTable{}
	str.RawSetString("HasPrefix", vm.Interp().NewClosure(func(l *lua.LState) int {
		if strings.HasPrefix(l.CheckString(1), l.CheckString(2)) {
			vm.Interp().Push(lua.LTrue)
		} else {
			vm.Interp().Push(lua.LFalse)
		}
		return 1
	}))
	str.RawSetString("Cut", vm.Interp().NewClosure(func(l *lua.LState) int {
		before, after, ok := strings.Cut(l.CheckString(1), l.CheckString(2))
		l.Push(lua.LString(before))
		l.Push(lua.LString(after))
		if ok {
			l.Push(lua.LTrue)
		} else {
			l.Push(lua.LFalse)
		}
		return 3
	}))
	str.RawSetString("Contains", vm.Interp().NewClosure(func(l *lua.LState) int {
		if strings.Contains(l.CheckString(1), l.CheckString(2)) {
			l.Push(lua.LTrue)
		} else {
			l.Push(lua.LFalse)
		}
		return 1
	}))
	vm.SetGlobalVar("strings", str)
	return runtime.Handles{
		"cmdf": func(l *runtime.LuaInterp) int {
			tmp := []string{}
			for i := 0; i <= l.GetTop(); i++ {
				switch l.CheckAny(i).Type() {
				case lua.LTNumber:
					tmp = append(tmp, l.CheckNumber(i).String())
				case lua.LTString:
					tmp = append(tmp, l.CheckString(i))
				}
			}
			l.Push(lua.LString(strings.Join(tmp, " ")))
			return 1
		},
		"pathf": func(l *lua.LState) int {
			l.Push(lua.LString(Pathf(l.CheckString(1))))
			return 1
		},
	}
}

func Pathf(path string) string {
	if len(path) > 0 && path[0] == '@' {
		path = WorkSpace + path[1:]
	}
	return path
}
