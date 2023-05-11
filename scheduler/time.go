package scheduler

import (
	"time"

	lua "github.com/yuin/gopher-lua"
)

func loadTime(vm *YockScheduler) {
	timelib := &lua.LTable{}
	timelib.RawSetString("microsecond", lua.LNumber(time.Microsecond))
	timelib.RawSetString("millisecond", lua.LNumber(time.Millisecond))
	timelib.RawSetString("second", lua.LNumber(time.Second))
	timelib.RawSetString("sleep", vm.Interp().NewClosure(func(l *lua.LState) int {
		time.Sleep(time.Duration(l.CheckNumber(1)))
		return 0
	}))
	vm.SetGlobalVar("time", timelib)
}
