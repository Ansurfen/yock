package scheduler

import (
	"github.com/ansurfen/cushion/utils"
	lua "github.com/yuin/gopher-lua"
)

func (vm *YockScheduler) loadRandom() {
	random := &lua.LTable{}
	random.RawSetString("str", vm.Interp().NewClosure(func(l *lua.LState) int {
		l.Push(lua.LString(utils.RandString(8)))
		return 1
	}))
	vm.SetGlobalVar("random", random)
}
