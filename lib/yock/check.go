// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	"github.com/ansurfen/cushion/utils"
	yocki "github.com/ansurfen/yock/interface"
	yockr "github.com/ansurfen/yock/runtime"
	lua "github.com/yuin/gopher-lua"
)

func LoadCheck(yocks yocki.YockScheduler) {
	yocks.RegYockFn(yocki.YockFuns{
		"CheckVersion":  checkVersion,
		"FormatVersion": checkVersionf,
	})
}

func checkVersion(l *yockr.YockState) int {
	want := utils.NewCheckedVersion(l.CheckString(1))
	got := utils.NewCheckedVersion(l.CheckString(2))
	l.PushBool(want.Compare(got))
	return 1
}

func checkVersionf(l *yockr.YockState) int {
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
