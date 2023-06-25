// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockr

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type YGFunction func(*YockState) int

type YockState struct {
	*lua.LState
	rn int
}

type YockFuncInfo lua.P

func NewYState() *YockState {
	return &YockState{LState: lua.NewState()}
}

func UpgradeLState(s *lua.LState) *YockState {
	return &YockState{LState: s}
}

func (s *YockState) Call(info YockFuncInfo, args ...any) error {
	lua_args := []lua.LValue{}
	for _, arg := range args {
		lua_args = append(lua_args, luar.New(s.LState, arg))
	}
	return s.LState.CallByParam(lua.P(info), lua_args...)
}

func (s *YockState) PCall() error {

	return nil
}

func (s *YockState) CheckTable(n int) *Table {
	return UpgradeTable(s.LState.CheckTable(n))
}

func (s *YockState) IsNil(n int) bool {
	return s.LState.CheckAny(n).Type() == lua.LTNil
}

func (s *YockState) IsFunction(n int) bool {
	return s.LState.CheckAny(n).Type() == lua.LTFunction
}

func (s *YockState) IsNumber(n int) bool {
	return s.LState.CheckAny(n).Type() == lua.LTNumber
}

func (s *YockState) IsBool(n int) bool {
	return s.LState.CheckAny(n).Type() == lua.LTBool
}

func (s *YockState) IsTable(n int) bool {
	return s.LState.CheckAny(n).Type() == lua.LTTable
}

func (s *YockState) IsString(n int) bool {
	return s.LState.CheckAny(n).Type() == lua.LTString
}

func (s *YockState) IsUserData(n int) bool {
	return s.LState.CheckAny(n).Type() == lua.LTUserData
}

func (s *YockState) IsThread(n int) bool {
	return s.LState.CheckAny(n).Type() == lua.LTThread
}

func (s *YockState) IsChannel(n int) bool {
	return s.LState.CheckAny(n).Type() == lua.LTChannel
}

func (s *YockState) Throw(err error) *YockState {
	s.LState.Push(lua.LString(err.Error()))
	s.rn++
	return s
}

// PushError returns the appropriate value depending on whether the error exists or not.
// Exists, returns error's text information, otherwise returns null.
//
// @return string|nil
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

func (s *YockState) PushNilTable() *YockState {
	s.LState.Push(&lua.LTable{})
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

func (s *YockState) PushInt(i int) *YockState {
	s.LState.Push(lua.LNumber(i))
	s.rn++
	return s
}

func (s *YockState) Pusha(val any) *YockState {
	s.Push(luar.New(s.LState, val))
	return s
}

func (s *YockState) PushAll(vals ...any) *YockState {
	for _, v := range vals {
		s.Push(luar.New(s.LState, v))
	}
	return s
}

func (s *YockState) NewLFunction(f lua.LGFunction) *lua.LFunction {
	return s.LState.NewFunction(f)
}

func (s *YockState) NewYFunction(f YGFunction) *lua.LFunction {
	return s.LState.NewFunction(func(l *lua.LState) int {
		return f(UpgradeLState(l))
	})
}

// Exit returns amount of return value
func (s *YockState) Exit() int {
	return s.rn
}

// stacktrace returns the stack info of function, in form of file:line
func (s *YockState) Stacktrace() string {
	dgb, ok := s.GetStack(1)
	if ok {
		s.GetInfo("S", dgb, &lua.LFunction{})
		s.GetInfo("l", dgb, &lua.LFunction{})
		return fmt.Sprintf("%s:%d\t", dgb.Source, dgb.CurrentLine)
	}
	return ""
}

func LuaDoFunc(lvm *lua.LState, fun *lua.LFunction) error {
	lfunc := lvm.NewFunctionFromProto(fun.Proto)
	lvm.Push(lfunc)
	return lvm.PCall(0, lua.MultRet, nil)
}
