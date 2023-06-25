// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util"
	lua "github.com/yuin/gopher-lua"
)

func LoadRandom(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("random")
	lib.SetFunctions(map[string]lua.LGFunction{
		"str": randomStr,
	})
}

// @param n number
//
// @return string
func randomStr(l *lua.LState) int {
	l.Push(lua.LString(util.RandString(int(l.CheckNumber(1)))))
	return 1
}
