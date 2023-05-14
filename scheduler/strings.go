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
	str.RawSetString("HasSuffix", vm.Interp().NewClosure(func(l *lua.LState) int {
		if strings.HasSuffix(l.CheckString(1), l.CheckString(2)) {
			vm.Interp().Push(lua.LTrue)
		} else {
			vm.Interp().Push(lua.LFalse)
		}
		return 1
	}))
	str.RawSetString("Contains", vm.Interp().NewClosure(func(l *lua.LState) int {
		if strings.Contains(l.CheckString(1), l.CheckString(2)) {
			l.Push(lua.LTrue)
		} else {
			l.Push(lua.LFalse)
		}
		return 1
	}))
	str.RawSetString("Join", vm.Interp().NewClosure(func(l *lua.LState) int {

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
	str.RawSetString("CutSuffix", vm.Interp().NewClosure(func(l *lua.LState) int {
		before, found := strings.CutSuffix(l.CheckString(1), l.CheckString(2))
		l.Push(lua.LString(before))
		if found {
			l.Push(lua.LTrue)
		} else {
			l.Push(lua.LFalse)
		}
		return 2
	}))
	str.RawSetString("CutPrefix", vm.Interp().NewClosure(func(l *lua.LState) int {
		after, found := strings.CutPrefix(l.CheckString(1), l.CheckString(2))
		l.Push(lua.LString(after))
		if found {
			l.Push(lua.LTrue)
		} else {
			l.Push(lua.LFalse)
		}
		return 2
	}))
	str.RawSetString("Clone", vm.Interp().NewClosure(func(l *lua.LState) int {
		l.Push(lua.LString(strings.Clone(l.CheckString(1))))
		return 1
	}))
	str.RawSetString("Compare", vm.Interp().NewClosure(func(l *lua.LState) int {
		l.Push(lua.LNumber(strings.Compare(l.CheckString(1), l.CheckString(2))))
		return 1
	}))
	str.RawSetString("ContainsAny", vm.Interp().NewClosure(func(l *lua.LState) int {
		if strings.ContainsAny(l.CheckString(1), l.CheckString(2)) {
			l.Push(lua.LTrue)
		} else {
			l.Push(lua.LFalse)
		}
		return 1
	}))
	str.RawSetString("ContainsRune", vm.Interp().NewClosure(func(l *lua.LState) int {
		if strings.ContainsRune(l.CheckString(1), rune(l.CheckString(2)[0])) {
			l.Push(lua.LTrue)
		} else {
			l.Push(lua.LFalse)
		}
		return 1
	}))
	str.RawSetString("Count", vm.Interp().NewClosure(func(l *lua.LState) int {
		l.Push(lua.LNumber(strings.Count(l.CheckString(1), l.CheckString(2))))
		return 1
	}))
	str.RawSetString("Replace", vm.Interp().NewClosure(func(l *lua.LState) int {
		l.Push(lua.LString(strings.Replace(l.CheckString(1), l.CheckString(2), l.CheckString(3), l.CheckInt(4))))
		return 1
	}))
	str.RawSetString("ReplaceAll", vm.Interp().NewClosure(func(l *lua.LState) int {
		l.Push(lua.LString(strings.ReplaceAll(l.CheckString(1), l.CheckString(2), l.CheckString(3))))
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
