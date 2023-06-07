// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package runtime

import (
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

type Table struct {
	*lua.LTable
}

func NewTable() *Table {
	return &Table{LTable: &lua.LTable{}}
}

func UpgradeTable(tbl *lua.LTable) *Table {
	return &Table{tbl}
}

func (t *Table) ToString(n int) string {
	return t.RawGetInt(n).String()
}

func (t *Table) ToTable(n int) *lua.LTable {
	if tbl, ok := t.RawGetInt(n).(*lua.LTable); ok {
		return tbl
	}
	return nil
}

func (t *Table) ToFunction(n int) *lua.LFunction {
	return t.RawGetInt(n).(*lua.LFunction)
}

func (t *Table) ToFloat32(n int) float32 {
	return float32(t.RawGetInt(n).(lua.LNumber))
}

func (t *Table) ToFloat64(n int) float64 {
	return float64(t.RawGetInt(n).(lua.LNumber))
}

func (t *Table) ToInt(n int) int {
	return int(t.RawGetInt(n).(lua.LNumber))
}

func (t *Table) Bind(v any) error {
	return gluamapper.Map(t.LTable, v)
}

func (t *Table) MustGetTable(key string) *lua.LTable {
	return t.RawGetString(key).(*lua.LTable)
}

func (t *Table) GetTable(key string) (*lua.LTable, bool) {
	tbl := t.RawGetString(key)
	if tbl.Type() == lua.LTTable {
		return tbl.(*lua.LTable), true
	}
	return nil, false
}

func (tbl *Table) Clone(l *lua.LState) *lua.LTable {
	netTable := &lua.LTable{}
	copyTable(l, tbl.LTable, netTable)
	return netTable
}

func copyTable(l *lua.LState, src *lua.LTable, dst *lua.LTable) {
	src.ForEach(func(key lua.LValue, value lua.LValue) {
		if tbl, ok := value.(*lua.LTable); ok {
			newTbl := l.NewTable()
			copyTable(l, tbl, newTbl)
			dst.RawSet(key, newTbl)
		} else {
			dst.RawSet(key, value)
		}
	})
}
