// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"sync"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func loadSync(yocks *YockScheduler) lua.LValue {
	return yocks.registerLib(syncLib)
}

var syncLib = luaFuncs{
	"new": syncNew,
}

// @return userdata
func syncNew(l *lua.LState) int {
	l.Push(luar.New(l, &sync.WaitGroup{}))
	return 1
}
