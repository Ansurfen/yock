// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockr

import (
	"testing"

	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util/test"
	lua "github.com/yuin/gopher-lua"
)

func TestYockState(t *testing.T) {
	s := NewYState()
	s.LState().DoString("function Echo(x) print(x) end")
	s.Call(yocki.YockFuncInfo{
		Fn: s.LState().GetGlobal("Echo"),
	}, "Hello World")
}

func TestUpgrade(t *testing.T) {
	l := lua.NewState()
	s := UpgradeLState(l)
	test.Assert(s.LState() == l)
}

func TestYStateCheck(t *testing.T) {
	s := NewYState()

	s.Push(lua.LString("1"))
	s.PushString("2")
	s.PushInt(3)

	test.Assert(s.Exit() == 3)
	test.Assert(s.Argc() == 3)
	test.Assert(s.IsNumber(3))
	test.Assert(s.CheckAny(2).(lua.LString).Type() == lua.LTString)
	test.Assert(s.CheckLValue(1).Type() == lua.LTString)

	s.PopAll()
	s.Pusha(1).PushAll("2", true, nil, 3.14)

	test.Assert(s.Argc() == 5)
	test.Assert(s.IsNumber(1))
	test.Assert(s.IsString(2))
	test.Assert(s.IsBool(3))
	test.Assert(s.IsNil(4))
	test.Assert(s.IsNumber(5))

	new := s.Clone()
	test.Assert(new != s)
}
