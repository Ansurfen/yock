// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package libsync

import (
	"sync"

	yocki "github.com/ansurfen/yock/interface"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func LoadSync(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("sync")
	lib.SetFunctions(map[string]lua.LGFunction{
		"new": syncNew,
	})
}

// @return userdata
func syncNew(l *lua.LState) int {
	l.Push(luar.New(l, &sync.WaitGroup{}))
	return 1
}
