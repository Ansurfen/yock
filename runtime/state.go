// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package runtime

import (
	lua "github.com/yuin/gopher-lua"
)

type YockState struct {
	*lua.LState
	rn int
}

func NewYState() *YockState {
	return &YockState{}
}

func UpgradeLState(s *lua.LState) *YockState {
	return &YockState{LState: s}
}

func (s *YockState) CheckTable(n int) *Table {
	return UpgradeTable(s.LState.CheckTable(n))
}

func (s *YockState) IsTable(n int) bool {
	return s.LState.CheckAny(n).Type() == lua.LTTable
}

func (s *YockState) Throw(err error) *YockState {
	s.LState.Push(lua.LString(err.Error()))
	s.rn++
	return s
}

func (s *YockState) PushError(err error) *YockState {
	if err != nil {
		s.LState.Push(lua.LString(err.Error()))
	} else {
		s.LState.Push(lua.LNil)
	}
	s.rn++
	return s
}

func (s *YockState) PushNil() *YockState {
	s.LState.Push(lua.LNil)
	s.rn++
	return s
}

func (s *YockState) Push(v lua.LValue) *YockState {
	s.LState.Push(v)
	s.rn++
	return s
}

func (s *YockState) PushString(str string) *YockState {
	s.LState.Push(lua.LString(str))
	s.rn++
	return s
}

func (s *YockState) PushBool(b bool) *YockState {
	if b {
		s.LState.Push(lua.LTrue)
	} else {
		s.LState.Push(lua.LFalse)
	}
	s.rn++
	return s
}

func (s *YockState) Exit() int {
	return s.rn
}
