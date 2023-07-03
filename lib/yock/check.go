// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util"
)

func LoadCheck(yocks yocki.YockScheduler) {
	yocks.RegYockFn(yocki.YockFuns{
		"CheckVersion":  checkVersion,
		"FormatVersion": checkVersionf,
	})
}

func checkVersion(s yocki.YockState) int {
	want := util.NewCheckedVersion(s.CheckString(1))
	got := util.NewCheckedVersion(s.CheckString(2))
	s.PushBool(want.Compare(got))
	return 1
}

func checkVersionf(s yocki.YockState) int {
	rawVersion := s.CheckString(1)
	targetCnt := s.CheckInt(2)
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
	s.PushString(curVersion)
	return 1
}
