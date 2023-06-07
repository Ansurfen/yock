// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"net"

	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/cmd"
	yockr "github.com/ansurfen/yock/runtime"
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
		s := yockr.UpgradeLState(l)
		opt := cmd.HttpOpt{Method: "GET"}
		urls := []string{}
		if s.IsTable(1) {
			tbl := s.CheckTable(1)

			if fn := tbl.RawGetString("filename"); fn.Type() == lua.LTFunction {
				opt.Filename = func(s string) string {
					lvm, _ := yocks.State().NewThread()
					if err := lvm.CallByParam(lua.P{
						NRet: 1,
						Fn:   fn.(*lua.LFunction),
					}, lua.LString(s)); err != nil {
						panic(err)
					}
					return lvm.CheckString(1)
				}
			}

			if err := tbl.Bind(&opt); err != nil {
				s.Throw(err)
				return s.Exit()
			}
			for i := 2; i <= l.GetTop(); i++ {
				urls = append(urls, l.CheckString(i))
			}
		} else {
			for i := 1; i < l.GetTop(); i++ {
				urls = append(urls, l.CheckString(i))
			}
		}
		s.PushError(cmd.HTTP(opt, urls))
		return s.Exit()
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
	s := yockr.UpgradeLState(l)
	url := s.CheckString(1)
	if url == "localhost" {
		return s.PushBool(true).Exit()
	}
	addrs, err := net.LookupHost("localhost")
	if err != nil {
		return s.Throw(err).Exit()
	}
	return s.PushBool(len(addrs) > 1 && addrs[1] == url).Exit()
}
