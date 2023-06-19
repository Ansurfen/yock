// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocki

import (
	"github.com/ansurfen/cushion/utils"
	yockr "github.com/ansurfen/yock/runtime"
	lua "github.com/yuin/gopher-lua"
)

type YockScheduler interface {
	// yock runtime
	YockLoader
	yockr.YockRuntime
	RegYocksFn(funcs YocksFuncs)
	MntYocksFn(lib *yockr.YockLib, funcs YocksFuncs)

	// yocks field
	EnvVar() utils.EnvVar
	Signal() SignalStream
	Opt() *yockr.Table
	SetOpt(o *yockr.Table)

	// yocks goroutines
	Do(f func())

	GetTask(name string) bool
	AppendTask(name string, job YockJob)
}

type YockJob interface {
	Func() *lua.LFunction
}

type YocksFunction func(yocks YockScheduler, state *yockr.YockState) int

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
	CreateLib(name string) *yockr.YockLib
	// OpenLib opens library to be specified
	// and creates a new library when it isn't exist
	OpenLib(name string) *yockr.YockLib
	RegLuaFn(LuaFuncs)
	RegYockFn(YockFuns)
	State() *yockr.YockState
}

type (
	YocksFuncs map[string]YocksFunction
	YockFuns   map[string]yockr.YGFunction
	LuaFuncs   map[string]lua.LGFunction
)
