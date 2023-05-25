package scheduler

import (
	"sync"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func loadSync(vm *YockScheduler) lua.LValue {
	synclib := &lua.LTable{}
	synclib.RawSetString("new", vm.Interp().NewClosure(
		func(l *lua.LState) int {
			l.Push(luar.New(l, &sync.WaitGroup{}))
			return 1
		},
	))
	return synclib
}
