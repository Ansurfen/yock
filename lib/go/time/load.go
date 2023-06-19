// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package libtime

import (
	"time"

	yocki "github.com/ansurfen/yock/interface"
	lua "github.com/yuin/gopher-lua"
)

func LoadTime(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("time")
	lib.SetField(map[string]any{
		"nanosecond":  time.Nanosecond,
		"microsecond": time.Microsecond,
		"millisecond": time.Millisecond,
		"second":      time.Second,
		"minute":      time.Minute,
		"hour":        time.Hour,
	})
	lib.SetFunctions(map[string]lua.LGFunction{
		"sleep": timeSleep,
	})

}

// Sleep pauses the current goroutine for at least the duration d. A negative or zero duration causes Sleep to return immediately.
//
// @param d number
func timeSleep(l *lua.LState) int {
	time.Sleep(time.Duration(l.CheckNumber(1)))
	return 0
}
