// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"time"

	lua "github.com/yuin/gopher-lua"
)

func loadTime(yocks *YockScheduler) lua.LValue {
	timelib := yocks.registerLib(luaFuncs{
		"sleep": timeSleep,
	})
	timelib.RawSetString("microsecond", lua.LNumber(time.Microsecond))
	timelib.RawSetString("millisecond", lua.LNumber(time.Millisecond))
	timelib.RawSetString("second", lua.LNumber(time.Second))
	return timelib
}

// Sleep pauses the current goroutine for at least the duration d. A negative or zero duration causes Sleep to return immediately.
//
// @param d number
func timeSleep(l *lua.LState) int {
	time.Sleep(time.Duration(l.CheckNumber(1)))
	return 0
}
