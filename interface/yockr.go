// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocki

import lua "github.com/yuin/gopher-lua"

type (
	YockFuncInfo lua.P
	YGFunction   func(YockState) int
	YockFuns     map[string]YGFunction
	LuaFuncs     map[string]lua.LGFunction
)

type Table interface {
	TableAt
	TableSetter
	TableGetter
	SetDo(k string, v func(YockState) lua.LValue, env ...YockState)
	Bind(v any) error
	Clone(l *lua.LState) Table
	Value() *lua.LTable
}

type TableSetter interface {
	SetString(k, v string)
	SetBool(k string, v bool)
	SetNil(k string)
	SetInt(k string, v int)
	SetTable(k string, v Table)
	SetLTable(k string, v *lua.LTable)
	SetField(l *lua.LState, k string, v any)
	SetFields(l *lua.LState, v map[string]any)
}

type TableGetter interface {
	GetBool(k string) (bool, bool)
	GetString(k string) (string, bool)
	GetInt(k string) (int, bool)
	GetFloat(k string) (float64, bool)
	GetTable(k string) (Table, bool)
	GetLTable(k string) (*lua.LTable, bool)
	MustGetTable(key string) Table
}

type TableAt interface {
	ToString(n int) string
	ToTable(n int) Table
	ToFunctionByString(k string) *lua.LFunction
	ToFloat32ByString(k string) float32
	ToFloat64ByString(k string) float64
	ToIntByString(k string) int
	ToFunction(n int) *lua.LFunction
	ToFloat32(n int) float32
	ToFloat64(n int) float64
	ToInt(n int) int
}

type YockState interface {
	Call(info YockFuncInfo, args ...any) error

	YockStateIs
	YockStateCheck
	YockStatePush

	NewLFunction(f lua.LGFunction) *lua.LFunction
	NewYFunction(f YGFunction) *lua.LFunction
	Exit() int

	Stack(i int) (*lua.Debug, bool)
	Stacktrace() string

	LState() *lua.LState
	Argc() int
	PopAll()
	Clone() YockState
}

type YockStateIs interface {
	IsNil(n int) bool
	IsFunction(n int) bool
	IsNumber(n int) bool
	IsBool(n int) bool
	IsTable(n int) bool
	IsString(n int) bool
	IsUserData(n int) bool
	IsThread(n int) bool
	IsChannel(n int) bool
}

type YockStateCheck interface {
	CheckTable(n int) Table
	CheckString(n int) string
	CheckRune(n int) rune
	CheckNumber(n int) lua.LNumber
	CheckInt(n int) int
	CheckBool(n int) bool
	CheckFunction(n int) *lua.LFunction
	CheckAny(n int) any
	CheckLValue(n int) lua.LValue
	CheckLTable(n int) *lua.LTable
}

type YockStatePush interface {
	Throw(err error) YockState
	PushError(err error) YockState
	PushNil() YockState
	Push(v lua.LValue) YockState
	PushNilTable() YockState
	PushString(str string) YockState
	PushBool(b bool) YockState
	PushInt(i int) YockState
	Pusha(val any) YockState
	PushAll(vals ...any) YockState
}

type YockLib interface {
	Name() string
	Value() lua.LValue
	SetField(v map[string]any)
	SetFunction(name string, fn lua.LGFunction)
	SetFunctions(v map[string]lua.LGFunction)
	SetYFunction(v map[string]YGFunction)
	SetClosure(v map[string]lua.LGFunction)
	Meta() Table
	SetTable(t Table)
	State() YockState
	SetState(s YockState)
}
