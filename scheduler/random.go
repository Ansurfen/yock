// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"github.com/ansurfen/cushion/utils"
	lua "github.com/yuin/gopher-lua"
)

func loadRandom(yocks *YockScheduler) lua.LValue {
	return yocks.registerLib(randomLib)
}

var randomLib = luaFuncs{
	"str": randomStr,
}

// @param n number
//
// @return string
func randomStr(l *lua.LState) int {
	l.Push(lua.LString(utils.RandString(int(l.CheckNumber(1)))))
	return 1
}
