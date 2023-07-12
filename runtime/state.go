// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockr

import (
	"fmt"

	yocki "github.com/ansurfen/yock/interface"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type YockState struct {
	ls *lua.LState
	rn int
}

var _ yocki.YockState = (*YockState)(nil)

func NewYState() *YockState {
	return &YockState{ls: lua.NewState()}
}

func UpgradeLState(s *lua.LState) *YockState {
	return &YockState{ls: s}
}

func (s *YockState) Call(info yocki.YockFuncInfo, args ...any) error {
	lua_args := []lua.LValue{}
	for _, arg := range args {
		lua_args = append(lua_args, luar.New(s.ls, arg))
	}
	return s.ls.CallByParam(lua.P(info), lua_args...)
}

func (s *YockState) PCall() error {
	return nil
}

func (s *YockState) CheckTable(n int) yocki.Table {
	return UpgradeTable(s.ls.CheckTable(n))
}

func (s *YockState) CheckString(n int) string {
	return s.ls.CheckString(n)
}

func (s *YockState) CheckRune(n int) rune {
	r := s.ls.CheckString(n)
	if len(r) > 0 {
		return rune(r[0])
	}
	return 0
}

func (s *YockState) CheckNumber(n int) lua.LNumber {
	return s.ls.CheckNumber(n)
}

func (s *YockState) CheckInt(n int) int {
	return s.ls.CheckInt(n)
}

func (s *YockState) CheckBool(n int) bool {
	return s.ls.CheckBool(n)
}

func (s *YockState) CheckFunction(n int) *lua.LFunction {
	return s.ls.CheckFunction(n)
}

func (s *YockState) CheckAny(n int) any {
	return s.ls.CheckAny(n)
}

func (s *YockState) IsNil(n int) bool {
	return s.ls.CheckAny(n).Type() == lua.LTNil
}

func (s *YockState) IsFunction(n int) bool {
	return s.ls.CheckAny(n).Type() == lua.LTFunction
}

func (s *YockState) IsNumber(n int) bool {
	return s.ls.CheckAny(n).Type() == lua.LTNumber
}

func (s *YockState) IsBool(n int) bool {
	return s.ls.CheckAny(n).Type() == lua.LTBool
}

func (s *YockState) IsTable(n int) bool {
	return s.ls.CheckAny(n).Type() == lua.LTTable
}

func (s *YockState) IsString(n int) bool {
	return s.ls.CheckAny(n).Type() == lua.LTString
}

func (s *YockState) IsUserData(n int) bool {
	return s.ls.CheckAny(n).Type() == lua.LTUserData
}

func (s *YockState) IsThread(n int) bool {
	return s.ls.CheckAny(n).Type() == lua.LTThread
}

func (s *YockState) IsChannel(n int) bool {
	return s.ls.CheckAny(n).Type() == lua.LTChannel
}

func (s *YockState) Throw(err error) yocki.YockState {
	s.ls.Push(lua.LString(err.Error()))
	s.rn++
	return s
}

// PushError returns the appropriate value depending on whether the error exists or not.
// Exists, returns error's text information, otherwise returns null.
//
// @return string|nil
func (s *YockState) PushError(err error) yocki.YockState {
	if err != nil {
		s.ls.Push(lua.LString(err.Error()))
	} else {
		s.ls.Push(lua.LNil)
	}
	s.rn++
	return s
}

func (s *YockState) PushNil() yocki.YockState {
	s.ls.Push(lua.LNil)
	s.rn++
	return s
}

func (s *YockState) Push(v lua.LValue) yocki.YockState {
	s.ls.Push(v)
	s.rn++
	return s
}

func (s *YockState) PushNilTable() yocki.YockState {
	s.ls.Push(&lua.LTable{})
	s.rn++
	return s
}

func (s *YockState) PushString(str string) yocki.YockState {
	s.ls.Push(lua.LString(str))
	s.rn++
	return s
}

func (s *YockState) PushBool(b bool) yocki.YockState {
	if b {
		s.ls.Push(lua.LTrue)
	} else {
		s.ls.Push(lua.LFalse)
	}
	s.rn++
	return s
}

func (s *YockState) PushInt(i int) yocki.YockState {
	s.ls.Push(lua.LNumber(i))
	s.rn++
	return s
}

func (s *YockState) Pusha(val any) yocki.YockState {
	s.Push(luar.New(s.ls, val))
	return s
}

func (s *YockState) PushAll(vals ...any) yocki.YockState {
	for _, v := range vals {
		s.Push(luar.New(s.ls, v))
	}
	return s
}

func (s *YockState) NewLFunction(f lua.LGFunction) *lua.LFunction {
	return s.ls.NewFunction(f)
}

func (s *YockState) NewYFunction(f yocki.YGFunction) *lua.LFunction {
	return s.ls.NewFunction(func(l *lua.LState) int {
		return f(UpgradeLState(l))
	})
}

// Exit returns amount of return value
func (s *YockState) Exit() int {
	return s.rn
}

// stacktrace returns the stack info of function, in form of file:line
func (s *YockState) Stacktrace() string {
	dbg, ok := s.Stack(1)
	if ok {
		return fmt.Sprintf("%s:%d\t", dbg.Source, dbg.CurrentLine)
	}
	return ""
}

func (s *YockState) Stack(i int) (dbg *lua.Debug, ok bool) {
	dbg, ok = s.ls.GetStack(i)
	if ok {
		s.ls.GetInfo("S", dbg, &lua.LFunction{})
		s.ls.GetInfo("l", dbg, &lua.LFunction{})
	}
	return
}

func (s *YockState) LState() *lua.LState {
	return s.ls
}

func (s *YockState) PopTop() {
	s.ls.Pop(s.ls.GetTop())
}

func (s *YockState) Argc() int {
	return s.ls.GetTop()
}

func LuaDoFunc(lvm *lua.LState, fun *lua.LFunction) error {
	lfunc := lvm.NewFunctionFromProto(fun.Proto)
	lvm.Push(lfunc)
	return lvm.PCall(0, lua.MultRet, nil)
}
