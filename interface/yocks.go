// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocki

import (
	"context"

	lua "github.com/yuin/gopher-lua"
)

type YockScheduler interface {
	// yock runtime
	YockLoader
	YockRuntime
	RegYocksFn(funcs YocksFuncs)
	MntYocksFn(lib YockLib, funcs YocksFuncs)

	// yocks field
	EnvVar() EnvVar
	Signal() SignalStream
	Opt() Table
	SetOpt(o Table)
	Env() YockLib

	// yocks goroutines
	Do(f func())

	GetTask(name string) bool
	AppendTask(name string, job YockJob)
}

var Y_MODE YockMode

type YockMode interface {
	Mode() int32
	SetMode(m int32)

	Debug() bool
	Strict() bool
}

type YockJob interface {
	Func() *lua.LFunction
}

// SignalStream is an abstract interface for distributing and updating singals
type SignalStream interface {
	// Load returns the value of the specified singal.
	// If the singal isn't exist, the second parameter returns false, and vice versa.
	Load(sig string) (any, bool)
	// Store settings specify the value of the singal, similar to map's kv storage.
	Store(sig string, v bool)
}

type YockLoader interface {
	// CreateLib returns a new library
	// and overrides old library when it's exist
	CreateLib(name string) YockLib
	// OpenLib opens library to be specified
	// and creates a new library when it isn't exist
	OpenLib(name string) YockLib
	RegLuaFn(LuaFuncs)
	RegYockFn(YockFuns)
	State() YockState
}

type (
	YocksFuncs    map[string]YocksFunction
	YocksFunction func(yocks YockScheduler, state YockState) int
)

// YockRuntime is an interface to abstract single and multiple interpreters
type YockRuntime interface {
	// Call to call specify function without arguments
	Call(string) ([]lua.LValue, error)
	// FastCall to call specify function without arguments and not return value
	FastCall(string) error
	// Call to call specify function with arguments
	// CallByParam(string, []lua.LValue) ([]lua.LValue, error)
	// FastCallByParam to call specify function with arguments and not return value
	// FastCallByParam(string, []lua.LValue) error
	// Eval to execute string of script
	Eval(string) error
	// EvalFile to execute file of script
	EvalFile(string) error
	// EvalFunc to execute function
	EvalFunc(lua.LValue, []lua.LValue) ([]lua.LValue, error)
	// FastEvalFunc to execute function and not return value
	FastEvalFunc(lua.LValue, []lua.LValue) error
	// SetGlobalFn to set global function
	SetGlobalFn(map[string]lua.LGFunction)
	// SafeSetGlobalFn to set global function when it isn't exist
	SafeSetGlobalFn(map[string]lua.LGFunction)
	// GetGlobalVar returns global variable
	GetGlobalVar(string) lua.LValue
	// SetGlobalVar to set global variable
	SetGlobalVar(string, lua.LValue)
	// SafeSetGlobalVar to set global variable when variable isn't exist
	SafeSetGlobalVar(string, lua.LValue)
	// RegisterModule to register modules
	// RegisterModule(map[string]lua.LGFunction)
	// UnregisterModule to unregister specify module
	// UnregisterModule(string)
	// LoadModule to immediately load module to be specified
	// LoadModule(string, lua.LGFunction)
	// State returns LState
	State() YockState
	// SetState sets interp
	SetState(l YockState)
	// NewState returns new interp
	NewState() (YockState, context.CancelFunc)
}
