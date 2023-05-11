package scheduler

import (
	"sync"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func loadSync(vm *YockScheduler) {
	synclib := &lua.LTable{}
	synclib.RawSetString("new", vm.Interp().NewClosure(
		func(l *lua.LState) int {
			l.Push(luar.New(l, &sync.WaitGroup{}))
			return 1
		},
	))
	vm.SetGlobalVar("sync", synclib)
}
