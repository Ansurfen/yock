// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	"time"

	yocki "github.com/ansurfen/yock/interface"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func LoadType(yocks yocki.YockScheduler) {
	yocks.RegLuaFn(yocki.LuaFuncs{
		"Boolean":     typeBoolean,
		"String":      typeString,
		"StringArray": typeStringArray,
		"Ptr":         typePtr,
		"map2Table": func(l *lua.LState) int {
			if v, ok := l.CheckUserData(1).Value.(map[string]any); ok {
				l.Push(mapToTable(v))
				return 1
			}
			l.Push(&lua.LTable{})
			return 1
		},
	})
}

func mapToTable(m map[string]interface{}) *lua.LTable {
	resultTable := &lua.LTable{}
	for key, element := range m {
		switch el := element.(type) {
		case float64:
			resultTable.RawSetString(key, lua.LNumber(el))
		case int64:
			resultTable.RawSetString(key, lua.LNumber(el))
		case string:
			resultTable.RawSetString(key, lua.LString(el))
		case bool:
			resultTable.RawSetString(key, lua.LBool(el))
		case []byte:
			resultTable.RawSetString(key, lua.LString(string(el)))
		case map[string]interface{}:
			tble := mapToTable(el)
			resultTable.RawSetString(key, tble)
		case map[string][]interface{}:
			globalTable := &lua.LTable{}
			for k, v := range element.(map[string][]interface{}) {
				// Create slice table
				sliceTable := &lua.LTable{}
				// Loop interface slice
				for _, s := range v {
					// Switch interface type
					switch sv := s.(type) {
					case map[string]interface{}:
						// Convert map to table
						t := mapToTable(sv)
						// Append result
						sliceTable.Append(t)
					case float64:
						// Append result as number
						sliceTable.Append(lua.LNumber(sv))
					case string:
						// Append result as string
						sliceTable.Append(lua.LString(sv))
					case bool:
						// Append result as bool
						sliceTable.Append(lua.LBool(sv))
					}
				}
				// Append to main table
				globalTable.RawSetString(k, sliceTable)
			}
			resultTable.RawSetString(key, globalTable)
		case map[string][]string:
			globalTable := &lua.LTable{}
			for k, v := range element.(map[string][]string) {
				// Create slice table
				sliceTable := &lua.LTable{}
				// Loop interface slice
				for _, s := range v {
					sliceTable.Append(lua.LString(s))
				}
				// Append to main table
				globalTable.RawSetString(k, sliceTable)
			}
			resultTable.RawSetString(key, globalTable)
		case time.Time:
			resultTable.RawSetString(key, lua.LNumber(el.Unix()))
		case []map[string]interface{}:
			// Create slice table
			sliceTable := &lua.LTable{}
			// Loop element
			for _, s := range el {
				// Get table from map
				tble := mapToTable(s)
				sliceTable.Append(tble)
			}
			// Set slice table
			resultTable.RawSetString(key, sliceTable)
		case []interface{}:
			// Create slice table
			sliceTable := &lua.LTable{}
			// Loop interface slice
			for _, s := range element.([]interface{}) {
				// Switch interface type
				switch sv := s.(type) {
				case map[string]interface{}:
					// Convert map to table
					t := mapToTable(sv)
					// Append result
					sliceTable.Append(t)
				case float64:
					// Append result as number
					sliceTable.Append(lua.LNumber(sv))
				case string:
					// Append result as string
					sliceTable.Append(lua.LString(sv))
				case bool:
					// Append result as bool
					sliceTable.Append(lua.LBool(sv))
				}
			}
			// Append to main table
			resultTable.RawSetString(key, sliceTable)
		default:
		}
	}
	return resultTable
}

func typePtr(l *lua.LState) int {
	l.Push(luar.New(l, &l.CheckUserData(1).Value))
	return 1
}

// @param b bool
//
// @return userdata (Boolean)
func typeBoolean(l *lua.LState) int {
	b := l.CheckBool(1)
	l.Push(luar.New(l, &Boolean{v: &b}))
	return 1
}

// @param s string
//
// @return userdata (String)
func typeString(l *lua.LState) int {
	s := l.CheckString(1)
	l.Push(luar.New(l, &String{v: &s}))
	return 1
}

// @param s ...string
//
// @return userdata (StringArray)
func typeStringArray(l *lua.LState) int {
	s := []string{}
	for i := 1; i <= l.GetTop(); i++ {
		s = append(s, l.CheckString(i))
	}
	l.Push(luar.New(l, &StringArray{v: &s}))
	return 1
}

type GoBindingType interface {
	Var() any
	Ptr() *any
}

type Boolean struct {
	v *bool
}

func (b *Boolean) Ptr() *bool {
	return b.v
}

func (b *Boolean) Var() bool {
	return *b.v
}

type String struct {
	v *string
}

func (s *String) Ptr() *string {
	return s.v
}

func (s *String) Var() string {
	return *s.v
}

type StringArray struct {
	v *[]string
}

func (arr *StringArray) Ptr() *[]string {
	return arr.v
}

func (arr *StringArray) Var() *lua.LTable {
	res := &lua.LTable{}
	for i, v := range *arr.v {
		res.Insert(i+1, lua.LString(v))
	}
	return res
}
