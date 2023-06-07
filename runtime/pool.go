// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package runtime

import (
	"sync"

	"github.com/ansurfen/yock/util"
	lua "github.com/yuin/gopher-lua"
)

var _ YockRuntime = (*YockInterpPool)(nil)

type YockInterpPool struct {
	m       sync.Mutex
	interps util.Stack[YockRuntime]
	idle    YockRuntime
}

func UpgradeInterpPool(yockr YockRuntime) YockRuntime {
	return &YockInterpPool{
		idle: yockr,
	}
}

// Call to call specify function without arguments
func (yockr *YockInterpPool) Call(name string) ([]lua.LValue, error) {
	r := yockr.Get()
	return r.Call(name)
}

// FastCall to call specify function without arguments and not return value
func (yockr *YockInterpPool) FastCall(string) error {
	return nil
}

func (yockr *YockInterpPool) Eval(string) error { return nil }

// EvalFile to execute file of script
func (yockr *YockInterpPool) EvalFile(string) error { return nil }

// EvalFunc to execute function
func (yockr *YockInterpPool) EvalFunc(lua.LValue, []lua.LValue) ([]lua.LValue, error) {
	return nil, nil
}

// FastEvalFunc to execute function and not return value
func (yockr *YockInterpPool) FastEvalFunc(lua.LValue, []lua.LValue) error { return nil }

// SetGlobalFn to set global function
func (yockr *YockInterpPool) SetGlobalFn(map[string]lua.LGFunction) {}

// SafeSetGlobalFn to set global function when it isn't exist
func (yockr *YockInterpPool) SafeSetGlobalFn(map[string]lua.LGFunction) {}

// GetGlobalVar returns global variable
func (yockr *YockInterpPool) GetGlobalVar(string) lua.LValue { return nil }

// SetGlobalVar to set global variable
func (yockr *YockInterpPool) SetGlobalVar(string, lua.LValue) {}

// SafeSetGlobalVar to set global variable when variable isn't exist
func (yockr *YockInterpPool) SafeSetGlobalVar(string, lua.LValue) {}

// RegisterModule to register modules
// RegisterModule(map[string]lua.LGFunction)
// UnregisterModule to unregister specify module
// UnregisterModule(string)
// LoadModule to immediately load module to be specified
// LoadModule(string, lua.LGFunction)
// State returns LState
func (yockr *YockInterpPool) State() *lua.LState { return nil }

func (yockr *YockInterpPool) SetState(l *lua.LState) {}

func (yockr *YockInterpPool) NewState() *lua.LState {
	return yockr.idle.NewState()
}

func (yockr *YockInterpPool) Get() YockRuntime {
	yockr.m.Lock()
	defer yockr.m.Unlock()
	n := yockr.interps.Len()
	if n == 0 {
		return yockr.New()
	}
	interp, err := yockr.interps.Top()
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
	yockr.interps.Pop()
	return interp
}

func (yockr *YockInterpPool) Put(interp YockRuntime) error {
	yockr.m.Lock()
	defer yockr.m.Unlock()
	return yockr.interps.Push(interp)
}

func (yockr *YockInterpPool) New() YockRuntime {
	return New()
}
