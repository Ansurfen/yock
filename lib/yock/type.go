// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	yocki "github.com/ansurfen/yock/interface"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func LoadType(yocks yocki.YockScheduler) {
	yocks.RegLuaFn(yocki.LuaFuncs{
		"Boolean":     typeBoolean,
		"String":      typeString,
		"StringArray": typeStringArray,
	})
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
