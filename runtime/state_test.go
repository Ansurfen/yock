// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockr

import (
	"testing"

	yocki "github.com/ansurfen/yock/interface"
)

func TestYockState(t *testing.T) {
	s := NewYState()
	s.LState().DoString("function Echo(x) print(x) end")
	s.Call(yocki.YockFuncInfo{
		Fn: s.LState().GetGlobal("Echo"),
	}, "Hello World")
}
