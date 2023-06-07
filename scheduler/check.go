// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"github.com/ansurfen/cushion/utils"
	lua "github.com/yuin/gopher-lua"
)

func loadCheck(yocks *YockScheduler) luaFuncs {
	return luaFuncs{
		"CheckVersion":  checkVersion,
		"FormatVersion": checkVersionf,
	}
}

func checkVersion(l *lua.LState) int {
	want := utils.NewCheckedVersion(l.CheckString(1))
	got := utils.NewCheckedVersion(l.CheckString(2))
	handleBool(l, want.Compare(got))
	return 1
}

func checkVersionf(l *lua.LState) int {
	rawVersion := l.CheckString(1)
	targetCnt := l.CheckInt(2)
	cnt := 0
	curVersion := ""
	for _, ch := range rawVersion {
		if targetCnt == cnt-1 {
			break
		}
		if ch == '.' {
			cnt++
		}
		curVersion += string(ch)
	}
	if curVersionLen := len(curVersion); curVersion[curVersionLen-1] == '.' {
		curVersion = curVersion[:curVersionLen-1]
	}
	l.Push(lua.LString(curVersion))
	return 1
}
