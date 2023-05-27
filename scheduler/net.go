// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"fmt"
	"net"

	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/cmd"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

func netFuncs(yocks *YockScheduler) luaFuncs {
	return luaFuncs{
		"http":         netHTTP(yocks),
		"is_url":       netIsURL,
		"is_localhost": netIsLocalhost,
	}
}

// netHTTP is capable of sending HTTP requests, which defaults to the GET method
/*
* @param opt table
* @param urls ...string
* @retrun err
 */
func netHTTP(yocks *YockScheduler) lua.LGFunction {
	return func(l *runtime.LuaInterp) int {
		mode := l.CheckAny(1)
		opt := cmd.HttpOpt{Method: "GET"}
		urls := []string{}
		if mode.Type() == lua.LTTable {
			if fn := l.CheckTable(1).RawGetString("filename"); fn.Type() == lua.LTFunction {
				opt.Filename = func(s string) string {
					lvm, _ := yocks.Interp().NewThread()
					if err := lvm.CallByParam(lua.P{
						NRet: 1,
						Fn:   fn.(*lua.LFunction),
					}, lua.LString(s)); err != nil {
						panic(err)
					}
					return lvm.CheckString(1)
				}
			}
			gluamapper.Map(l.CheckTable(1), &opt)
			for i := 2; i <= l.GetTop(); i++ {
				urls = append(urls, l.CheckString(i))
			}
		} else {
			for i := 1; i < l.GetTop(); i++ {
				urls = append(urls, l.CheckString(i))
			}
		}
		handleErr(l, cmd.HTTP(opt, urls))
		return 1
	}
}

// @param url string
//
// @return bool
func netIsURL(l *lua.LState) int {
	handleBool(l, utils.IsURL(l.CheckString(1)))
	return 1
}

// @param url string
//
// @return bool
func netIsLocalhost(l *lua.LState) int {
	url := l.CheckString(1)
	if url == "localhost" {
		l.Push(lua.LTrue)
		return 1
	}
	addrs, err := net.LookupHost("localhost")
	if err != nil {
		fmt.Println("error:", err)
	}
	if len(addrs) > 1 && addrs[1] == url {
		l.Push(lua.LTrue)
	} else {
		l.Push(lua.LFalse)
	}
	return 1
}
