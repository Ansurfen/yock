// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockr

import (
	"context"
	"path/filepath"

	yocki "github.com/ansurfen/yock/interface"
	lua "github.com/yuin/gopher-lua"
)

// New initialize yock runtime and returns the its pointer
func New(opts ...YockrOption) yocki.YockRuntime {
	var yockr yocki.YockRuntime = &YockInterp{state: NewYState()}

	for _, opt := range opts {
		if err := opt(yockr); err != nil {
			panic(err)
		}
	}

	if yockr.State() == nil {
		s, cancel := yockr.NewState()
		if cancel != nil {
			defer cancel()
		}
		yockr.SetState(s)
	}

	return yockr
}

var _ yocki.YockRuntime = (*YockInterp)(nil)

type YockrOption func(yockr yocki.YockRuntime) error

// YockInterp abstracts lua interpreter
type YockInterp struct {
	state yocki.YockState
}

// State returns LState
func (yockr *YockInterp) State() yocki.YockState {
	return yockr.state
}

func (yockr *YockInterp) SetState(l yocki.YockState) {
	yockr.state = l
}

func (yockr *YockInterp) NewState() (yocki.YockState, context.CancelFunc) {
	ls, cancel := yockr.state.LState().NewThread()
	return UpgradeLState(ls), cancel
}

// FastCall to call specify function without arguments and not return value
func (yockr *YockInterp) FastCall(fun string) error {
	return yockr.state.LState().CallByParam(lua.P{
		Fn:      yockr.state.LState().GetGlobal(fun),
		NRet:    0,
		Protect: true,
	})
}

// Call to call specify function without arguments
func (yockr *YockInterp) Call(fun string) ([]lua.LValue, error) {
	ret := []lua.LValue{}
	if err := yockr.state.LState().CallByParam(lua.P{
		Fn:      yockr.state.LState().GetGlobal(fun),
		NRet:    0,
		Protect: true,
	}); err != nil {
		return ret, err
	}
	for i := 1; i <= yockr.state.LState().GetTop(); i++ {
		ret = append(ret, yockr.state.LState().CheckAny(i))
	}
	return ret, nil
}

// SetGlobalFn to set global function
func (yockr *YockInterp) SetGlobalFn(loaders map[string]lua.LGFunction) {
	for name, loader := range loaders {
		yockr.state.LState().SetGlobal(name, yockr.state.LState().NewFunction(loader))
	}
}

// SafeSetGlobalFn to set global function when it isn't exist
func (yockr *YockInterp) SafeSetGlobalFn(loaders map[string]lua.LGFunction) {
	for name, loader := range loaders {
		if value := yockr.state.LState().GetGlobal(name); value.String() == "nil" {
			yockr.state.LState().SetGlobal(name, yockr.state.LState().NewFunction(loader))
		}
	}
}

// Eval to execute string of script
func (yockr *YockInterp) Eval(script string) error {
	return yockr.state.LState().DoString(script)
}

// EvalFile to execute file of script
func (yockr *YockInterp) EvalFile(fullpath string) error {
	if filepath.Ext(fullpath) == ".lua" {
		return yockr.state.LState().DoFile(fullpath)
	}
	return nil
}

// EvalFunc to execute function
func (yockr *YockInterp) EvalFunc(fn lua.LValue, args []lua.LValue) ([]lua.LValue, error) {
	ret := []lua.LValue{}
	if err := yockr.state.LState().CallByParam(lua.P{
		Fn:      fn,
		Protect: true,
	}, args...); err != nil {
		return ret, err
	}
	for i := 1; i <= yockr.state.LState().GetTop(); i++ {
		ret = append(ret, yockr.state.LState().CheckAny(i))
	}
	return ret, nil
}

// FastEvalFunc to execute function and not return value
func (yockr *YockInterp) FastEvalFunc(fn lua.LValue, args []lua.LValue) error {
	return yockr.state.LState().CallByParam(lua.P{
		Fn:      fn,
		Protect: true,
	}, args...)
}

// GetGlobalVar returns global variable
func (yockr *YockInterp) GetGlobalVar(name string) lua.LValue {
	return yockr.state.LState().GetGlobal(name)
}

// SetGlobalVar to set global variable
func (yockr *YockInterp) SetGlobalVar(name string, v lua.LValue) {
	yockr.state.LState().SetGlobal(name, v)
}

// SafeSetGlobalVar to set global variable when variable isn't exist
func (yockr *YockInterp) SafeSetGlobalVar(name string, v lua.LValue) {
	if yockr.state.LState().GetGlobal(name).Type().String() == "nil" {
		yockr.state.LState().SetGlobal(name, v)
	}
}
