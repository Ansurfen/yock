// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package runtime

import (
	"path/filepath"

	lua "github.com/yuin/gopher-lua"
)

// New initialize yock runtime and returns the its pointer
func New(opts ...YockrOption) YockRuntime {
	var yockr YockRuntime = &YockInterp{}

	for _, opt := range opts {
		if err := opt(yockr); err != nil {
			panic(err)
		}
	}

	if yockr.State() == nil {
		yockr.SetState(lua.NewState())
	}

	return yockr
}

var _ YockRuntime = (*YockInterp)(nil)

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
	State() *lua.LState
	// SetState sets interp
	SetState(l *lua.LState)
	// NewState returns new interp
	NewState() *lua.LState
}

type YockrOption func(yockr YockRuntime) error

// YockInterp abstracts lua interpreter
type YockInterp struct {
	state *lua.LState
}

// State returns LState
func (yockr *YockInterp) State() *lua.LState {
	return yockr.state
}

func (yockr *YockInterp) SetState(l *lua.LState) {
	yockr.state = l
}

func (yockr *YockInterp) NewState() *lua.LState {
	ls, _ := yockr.state.NewThread()
	return ls
}

// FastCall to call specify function without arguments and not return value
func (yockr *YockInterp) FastCall(fun string) error {
	return yockr.state.CallByParam(lua.P{
		Fn:      yockr.state.GetGlobal(fun),
		NRet:    0,
		Protect: true,
	})
}

// Call to call specify function without arguments
func (yockr *YockInterp) Call(fun string) ([]lua.LValue, error) {
	ret := []lua.LValue{}
	if err := yockr.state.CallByParam(lua.P{
		Fn:      yockr.state.GetGlobal(fun),
		NRet:    0,
		Protect: true,
	}); err != nil {
		return ret, err
	}
	for i := 1; i <= yockr.state.GetTop(); i++ {
		ret = append(ret, yockr.state.CheckAny(i))
	}
	return ret, nil
}

// SetGlobalFn to set global function
func (yockr *YockInterp) SetGlobalFn(loaders map[string]lua.LGFunction) {
	for name, loader := range loaders {
		yockr.state.SetGlobal(name, yockr.state.NewFunction(loader))
	}
}

// SafeSetGlobalFn to set global function when it isn't exist
func (yockr *YockInterp) SafeSetGlobalFn(loaders map[string]lua.LGFunction) {
	for name, loader := range loaders {
		if value := yockr.state.GetGlobal(name); value.String() == "nil" {
			yockr.state.SetGlobal(name, yockr.state.NewFunction(loader))
		}
	}
}

// Eval to execute string of script
func (yockr *YockInterp) Eval(script string) error {
	return yockr.state.DoString(script)
}

// EvalFile to execute file of script
func (yockr *YockInterp) EvalFile(fullpath string) error {
	if filepath.Ext(fullpath) == ".lua" {
		return yockr.state.DoFile(fullpath)
	}
	return nil
}

// EvalFunc to execute function
func (yockr *YockInterp) EvalFunc(fn lua.LValue, args []lua.LValue) ([]lua.LValue, error) {
	ret := []lua.LValue{}
	if err := yockr.state.CallByParam(lua.P{
		Fn:      fn,
		Protect: true,
	}, args...); err != nil {
		return ret, err
	}
	for i := 1; i <= yockr.state.GetTop(); i++ {
		ret = append(ret, yockr.state.CheckAny(i))
	}
	return ret, nil
}

// FastEvalFunc to execute function and not return value
func (yockr *YockInterp) FastEvalFunc(fn lua.LValue, args []lua.LValue) error {
	return yockr.state.CallByParam(lua.P{
		Fn:      fn,
		Protect: true,
	}, args...)
}

// GetGlobalVar returns global variable
func (yockr *YockInterp) GetGlobalVar(name string) lua.LValue {
	return yockr.state.GetGlobal(name)
}

// SetGlobalVar to set global variable
func (yockr *YockInterp) SetGlobalVar(name string, v lua.LValue) {
	yockr.state.SetGlobal(name, v)
}

// SafeSetGlobalVar to set global variable when variable isn't exist
func (yockr *YockInterp) SafeSetGlobalVar(name string, v lua.LValue) {
	if yockr.state.GetGlobal(name).Type().String() == "nil" {
		yockr.state.SetGlobal(name, v)
	}
}
