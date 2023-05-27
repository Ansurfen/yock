// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"time"

	lua "github.com/yuin/gopher-lua"
)

func goroutineFuncs(yocks *YockScheduler) luaFuncs {
	return luaFuncs{
		"go":     goroutineGo(yocks),
		"wait":   goroutineWait(yocks),
		"waits":  goroutineWaits(yocks),
		"notify": goroutineNotify(yocks),
	}
}

// goroutineGo wraps the callback function of the Lua language into a go callback
// and passes it into the goroutines for unified scheduling.
// @param fn function
func goroutineGo(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		fn := l.CheckFunction(1)
		tmp, cancel := l.NewThread()
		yocks.goroutines <- func() {
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

// @param sig string
func goroutineWait(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		sig := l.CheckString(1)
		if _, ok := yocks.signals.Load(sig); !ok {
			yocks.signals.Store(sig, false)
		}
		cnt := 0
		for {
			if sig, ok := yocks.signals.Load(sig); ok && sig.(bool) {
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

// @param sig ...string
func goroutineWaits(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		sigs := []string{}
		for i := 1; i <= l.GetTop(); i++ {
			sigs = append(sigs, l.CheckString(i))
		}
		cnt := 0
		for {
			flag := true
			for i := 0; i < len(sigs); i++ {
				if sig, ok := yocks.signals.Load(sigs[i]); !ok || (ok && !sig.(bool)) {
					flag = false
					break
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

// @param sig string
func goroutineNotify(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		sig := l.CheckString(1)
		yocks.signals.Store(sig, true)
		return 0
	}
}
