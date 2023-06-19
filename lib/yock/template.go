// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	"github.com/ansurfen/cushion/utils/build"
	yocki "github.com/ansurfen/yock/interface"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func LoadTemplate(yocks yocki.YockScheduler) {
	yocks.RegLuaFn(yocki.LuaFuncs{
		"tmpl": func(l *lua.LState) int {
			tmpl := build.NewTemplate()
			l.Push(luar.New(l, tmpl))
			return 1
		},
	})
}
