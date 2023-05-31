// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"github.com/ansurfen/cushion/utils/build"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

var tmplFuncs = luaFuncs{
	"tmpl": func(l *lua.LState) int {
		tmpl := build.NewTemplate()
		l.Push(luar.New(l, tmpl))
		return 1
	},
}
