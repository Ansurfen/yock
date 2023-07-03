// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockr

import (
	yocki "github.com/ansurfen/yock/interface"
	lua "github.com/yuin/gopher-lua"
)

type YockLib struct {
	tbl   yocki.Table
	state yocki.YockState
	name  string
}

var _ yocki.YockLib = (*YockLib)(nil)

// CreateYockLib returns new library whether it exist or not.
func CreateYockLib(state yocki.YockState, name string, tbl ...*Table) *YockLib {
	var t *Table
	if len(tbl) > 0 {
		t = tbl[0]
	} else {
		t = NewTable()
	}
	state.LState().SetGlobal(name, t.LTable)
	return &YockLib{name: name, tbl: t, state: state}
}

// OpenYockLib returns existed library based on name
// and creates new library when lib isn't exist.
func OpenYockLib(state yocki.YockState, name string) *YockLib {
	var t yocki.Table
	if v := state.LState().GetGlobal(name); v.Type() == lua.LTNil {
		t = NewTable()
		state.LState().SetGlobal(name, t.Value())
	} else if v.Type() == lua.LTTable {
		t = UpgradeTable(v.(*lua.LTable))
	} else {
		t = NewTable()
	}
	return &YockLib{name: name, tbl: t, state: state}
}

func (lib *YockLib) Name() string {
	return lib.name
}

func (lib *YockLib) Value() lua.LValue {
	return lib.tbl.Value()
}

func (lib *YockLib) SetField(v map[string]any) {
	lib.tbl.SetField(lib.state.LState(), v)
}

func (lib *YockLib) SetFunction(name string, fn lua.LGFunction) {
	lib.tbl.Value().RawSetString(name, lib.state.LState().NewFunction(fn))
}

func (lib *YockLib) SetFunctions(v map[string]lua.LGFunction) {
	for name, fn := range v {
		lib.tbl.Value().RawSetString(name, lib.state.LState().NewFunction(fn))
	}
}

func wrapYFunction(fn yocki.YGFunction) lua.LGFunction {
	return func(l *lua.LState) int {
		return fn(UpgradeLState(l))
	}
}

func (lib *YockLib) SetYFunction(v map[string]yocki.YGFunction) {
	for name, fn := range v {
		lib.tbl.Value().RawSetString(name, lib.state.LState().NewFunction(wrapYFunction(fn)))
	}
}

func (lib *YockLib) SetClosure(v map[string]lua.LGFunction) {
	for name, fn := range v {
		lib.tbl.Value().RawSetString(name, lib.state.LState().NewClosure(fn))
	}
}

func (lib *YockLib) Meta() yocki.Table {
	return lib.tbl
}

func (lib *YockLib) SetTable(t yocki.Table) {
	lib.tbl = t
}

func (lib *YockLib) State() yocki.YockState {
	return lib.state
}

func (lib *YockLib) SetState(s yocki.YockState) {
	lib.state = s
}
