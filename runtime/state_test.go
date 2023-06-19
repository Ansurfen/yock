// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockr

import "testing"

func TestYockState(t *testing.T) {
	s := NewYState()
	s.DoString("function Echo(x) print(x) end")
	s.Call(YockFuncInfo{
		Fn: s.GetGlobal("Echo"),
	}, "Hello World")
}
