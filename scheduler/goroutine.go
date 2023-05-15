package scheduler

import (
	"time"

	"github.com/ansurfen/cushion/runtime"
	lua "github.com/yuin/gopher-lua"
)

func loadGoroutine(vm *YockScheduler) runtime.Handles {
	return runtime.Handles{
		"go":     goroutineGo(vm),
		"wait":   goroutineWait(vm),
		"waits":  goroutineWaits(vm),
		"notify": goroutineNotify(vm),
	}
}

func goroutineGo(vm *YockScheduler) func(*runtime.LuaInterp) int {
	return func(l *runtime.LuaInterp) int {
		fn := l.CheckFunction(1)
		tmp, cancel := l.NewThread()
		vm.goroutines <- func() {
			tmp.CallByParam(lua.P{
				Fn: fn,
			})
			if cancel != nil {
				cancel()
			}
		}
		return 0
	}
}

func goroutineWait(vm *YockScheduler) func(*runtime.LuaInterp) int {
	return func(l *runtime.LuaInterp) int {
		sig := l.CheckString(1)
		if _, ok := vm.signals.Load(sig); !ok {
			vm.signals.Store(sig, false)
		}
		cnt := 0
		for {
			if sig, ok := vm.signals.Load(sig); ok && sig.(bool) {
				break
			}
			round := 1 + cnt>>2
			if round > 10 {
				round = 10
			}
			time.Sleep(time.Duration(round) * time.Second)
			cnt++
		}
		return 0
	}
}

func goroutineWaits(vm *YockScheduler) func(*runtime.LuaInterp) int {
	return func(l *runtime.LuaInterp) int {
		sigs := []string{}
		for i := 1; i <= l.GetTop(); i++ {
			sigs = append(sigs, l.CheckString(i))
		}
		cnt := 0
		for {
			flag := true
			for i := 0; i < len(sigs); i++ {
				if sig, ok := vm.signals.Load(sigs[i]); !ok || (ok && !sig.(bool)) {
					flag = false
				}
			}
			if flag {
				break
			}
			round := 1 + cnt>>2
			if round > 10 {
				round = 10
			}
			time.Sleep(time.Duration(round) * time.Second)
			cnt++
		}
		return 0
	}
}

func goroutineNotify(vm *YockScheduler) func(*runtime.LuaInterp) int {
	return func(l *runtime.LuaInterp) int {
		sig := l.CheckString(1)
		vm.signals.Store(sig, true)
		return 0
	}
}
