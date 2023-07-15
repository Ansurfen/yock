// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	"time"

	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/ycho"
)

func LoadGoroutine(yocks yocki.YockScheduler) {
	yocks.RegYocksFn(yocki.YocksFuncs{
		"go":     goroutineGo,
		"wait":   goroutineWait,
		"waits":  goroutineWaits,
		"notify": goroutineNotify,
	})
}

// goroutineGo wraps the callback function of the Lua language into a go callback
// and passes it into the goroutines for unified scheduling.
// @param fn function
func goroutineGo(yocks yocki.YockScheduler, state yocki.YockState) int {
	fn := state.CheckFunction(1)
	yocks.Do(func() {
		tmp, cancel := yocks.NewState()
		if cancel != nil {
			defer cancel()
		}
		if err := tmp.Call(yocki.YockFuncInfo{
			Fn: fn,
		}); err != nil {
			ycho.Warn(err)
		}
	})
	return 0
}

// @param sig string
func goroutineWait(yocks yocki.YockScheduler, state yocki.YockState) int {
	sig := state.CheckString(1)
	if _, ok := yocks.Signal().Load(sig); !ok {
		yocks.Signal().Store(sig, false)
	}
	cnt := 0
	for {
		if sig, ok := yocks.Signal().Load(sig); ok && sig.(bool) {
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

// @param sig ...string
func goroutineWaits(yocks yocki.YockScheduler, state yocki.YockState) int {
	sigs := []string{}
	for i := 1; i <= state.Argc(); i++ {
		sigs = append(sigs, state.CheckString(i))
	}
	cnt := 0
	for {
		flag := true
		for i := 0; i < len(sigs); i++ {
			if sig, ok := yocks.Signal().Load(sigs[i]); !ok || (ok && !sig.(bool)) {
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

// @param sig string
func goroutineNotify(yocks yocki.YockScheduler, state yocki.YockState) int {
	sig := state.CheckString(1)
	yocks.Signal().Store(sig, true)
	return 0
}
