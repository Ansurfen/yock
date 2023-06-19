// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockr

import (
	lua "github.com/yuin/gopher-lua"
)

type YockLib struct {
	tbl   *Table
	state *YockState
	name  string
}

// CreateYockLib returns new library whether it exist or not.
func CreateYockLib(state *YockState, name string, tbl ...*Table) *YockLib {
	var t *Table
	if len(tbl) > 0 {
		t = tbl[0]
	} else {
		t = NewTable()
	}
	state.SetGlobal(name, t.LTable)
	return &YockLib{name: name, tbl: t, state: state}
}

// OpenYockLib returns existed library based on name
// and creates new library when lib isn't exist.
func OpenYockLib(state *YockState, name string) *YockLib {
	var t *Table
	if v := state.GetGlobal(name); v.Type() == lua.LTNil {
		t = NewTable()
		state.SetGlobal(name, t.LTable)
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
	return lib.tbl
}

func (lib *YockLib) SetField(v map[string]any) {
	lib.tbl.SetField(lib.state.LState, v)
}

func (lib *YockLib) SetFunction(name string, fn lua.LGFunction) {
	lib.tbl.RawSetString(name, lib.state.NewFunction(fn))
}

func (lib *YockLib) SetFunctions(v map[string]lua.LGFunction) {
	for name, fn := range v {
		lib.tbl.RawSetString(name, lib.state.NewFunction(fn))
	}
}

func wrapYFunction(fn YGFunction) lua.LGFunction {
	return func(l *lua.LState) int {
		return fn(UpgradeLState(l))
	}
}

func (lib *YockLib) SetYFunction(v map[string]YGFunction) {
	for name, fn := range v {
		lib.tbl.RawSetString(name, lib.state.NewFunction(wrapYFunction(fn)))
	}
}

func (lib *YockLib) SetClosure(v map[string]lua.LGFunction) {
	for name, fn := range v {
		lib.tbl.RawSetString(name, lib.state.NewClosure(fn))
	}
}

func (lib *YockLib) Meta() *Table {
	return lib.tbl
}

func (lib *YockLib) SetTable(t *Table) {
	lib.tbl = t
}

func (lib *YockLib) State() *YockState {
	return lib.state
}

func (lib *YockLib) SetState(s *YockState) {
	lib.state = s
}
