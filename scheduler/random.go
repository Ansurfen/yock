package scheduler

import (
	"github.com/ansurfen/cushion/utils"
	lua "github.com/yuin/gopher-lua"
)

func loadRandom(vm *YockScheduler) lua.LValue {
	random := &lua.LTable{}
	random.RawSetString("str", vm.Interp().NewClosure(func(l *lua.LState) int {
		l.Push(lua.LString(utils.RandString(int(l.CheckNumber(1)))))
		return 1
	}))
	return random
}
