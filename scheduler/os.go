// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"strings"

	"github.com/ansurfen/cushion/utils"
	yockc "github.com/ansurfen/yock/cmd"
	yockr "github.com/ansurfen/yock/runtime"
	"github.com/ansurfen/yock/util"
	"github.com/spf13/cobra"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

var osFuncs = luaFuncs{
	"sh":          osSh,
	"cmdf":        osCmdf,
	"strf":        osStrf,
	"new_command": osNewCommand,
}

/*
* @param str string
* @param charset string
* @return string
 */
func osStrf(l *lua.LState) int {
	out := utils.ConvertByte2String([]byte(l.CheckString(1)), utils.Charset(l.CheckString(2)))
	l.Push(lua.LString(out))
	return 1
}

/*
* @param opt? table
* @param cmds string
* @return table, err
 */
func osSh(l *lua.LState) int {
	s := yockr.UpgradeLState(l)
	cmds := []string{}
	opt := yockc.ExecOpt{Quiet: true}
	if s.IsTable(1) {
		tbl := s.CheckTable(1)
		if err := tbl.Bind(&opt); err != nil {
			return s.PushNil().Throw(err).Exit()
		}
		for i := 2; i <= l.GetTop(); i++ {
			cmds = append(cmds, l.CheckString(i))
		}
	} else {
		opt.Redirect = true
		for i := 1; i <= l.GetTop(); i++ {
			cmds = append(cmds, l.CheckString(i))
		}
	}
	outs := &lua.LTable{}
	var g_err error
	for _, cmd := range cmds {
		utils.ReadLineFromString(cmd, func(s string) string {
			if len(s) > 0 {
				out, err := yockc.Exec(opt, s)
				outs.Append(lua.LString(out))
				if err != nil {
					if opt.Debug {
						util.Ycho.Warn(err.Error())
					}
					if opt.Strict {
						return ""
					} else {
						g_err = util.ErrGeneral
					}
				}
			}
			return ""
		})
	}

	return s.Push(outs).PushError(g_err).Exit()
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
