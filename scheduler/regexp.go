// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"regexp"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func loadRegexp(yocks *YockScheduler) lua.LValue {
	return yocks.registerLib(regexpLib)
}

var regexpLib = luaFuncs{
	"Compile":     regexpCompile,
	"MustCompile": regexpMustCompile,
}

// @param expr string
//
// @return userdata (*regexp.Regexp), err
func regexpCompile(l *lua.LState) int {
	r, err := regexp.Compile(l.CheckString(1))
	l.Push(luar.New(l, r))
	handleErr(l, err)
	return 2
}

// @param expr string
//
// @return userdata (*regexp.Regexp)
func regexpMustCompile(l *lua.LState) int {
	r := regexp.MustCompile(l.CheckString(1))
	l.Push(luar.New(l, r))
	return 1
}
