// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"strings"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/cmd"
	"github.com/spf13/cobra"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

var osFuncs = luaFuncs{
	"sh":          osSh,
	"exec":        osExec,
	"cmdf":        osCmdf,
	"new_command": osNewCommand,
}

// @param opt table
func osExec(l *lua.LState) int {
	mode := l.CheckAny(1)
	opt := cmd.ExecOpt{Quiet: true}
	cmds := []string{}
	if mode.Type() == lua.LTTable {
		gluamapper.Map(l.CheckTable(1), &opt)
		for i := 2; i <= l.GetTop(); i++ {
			cmds = append(cmds, l.CheckString(i))
		}
	} else {
		for i := 1; i < l.GetTop(); i++ {
			cmds = append(cmds, l.CheckString(i))
		}
	}
	cmd.Exec(opt, cmds)
	return 0
}

// @param cmds string
func osSh(l *lua.LState) int {
	cmds := l.CheckString(1)
	utils.ReadLineFromString(cmds, func(s string) string {
		if len(s) > 0 {
			cmd.Exec(cmd.ExecOpt{}, []string{s})
		}
		return ""
	})
	return 0
}

// @param cmd ...string
//
// @return string
func osCmdf(l *lua.LState) int {
	tmp := []string{}
	for i := 0; i <= l.GetTop(); i++ {
		switch l.CheckAny(i).Type() {
		case lua.LTNumber:
			tmp = append(tmp, l.CheckNumber(i).String())
		case lua.LTString:
			tmp = append(tmp, l.CheckString(i))
		}
	}
	l.Push(lua.LString(strings.Join(tmp, " ")))
	return 1
}

// @return userdata
func osNewCommand(l *lua.LState) int {
	l.Push(luar.New(l, &cobra.Command{}))
	return 1
}
