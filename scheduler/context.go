// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocks

import (
	"context"
	"fmt"
	"time"

	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type Context struct {
	s      yocki.YockState
	tbl    yocki.Table
	cancel context.CancelFunc
	source string
}

type contextExitCode int

func (c contextExitCode) String() string {
	switch c {
	case 0:
		return "abort all peer jobs"
	case 1:
		return "continue to run peer jobs"
	case 2:
		return "continue to run peer jobs with inherit"
	default:
		return "unknown"
	}
}

func newContext(name string, job yocki.YockJob, flags yocki.Table, yocks yocki.YockScheduler) *Context {
	s, cancel := yocks.NewState()
	// s = s.Clone()
	tbl := yocks.Env().Meta().Clone(s.LState())
	tbl.SetString("task", name)
	tbl.SetFields(s.LState(), map[string]any{
		"throw": func(msg ...string) {
			if len(msg) > 0 {
				s.LState().RaiseError("[%s] %s", name, s.Stacktrace())
			} else {
				s.LState().RaiseError("[%s] %s %s", name, msg[0], s.Stacktrace())
			}
		},
		"yield": func(timeout ...int) {
			if len(timeout) > 0 {
				time.Sleep(time.Duration(timeout[0]))
			} else {
				// TODO: wait for
			}
		},
		"resume": func() {

		},
		"put": func(k string, v lua.LValue) {
			yocks.Put(k, v)
		},
		"get": func(k string) any {
			return yocks.Get(k)
		},
		"assert": func(ok lua.LValue, msg ...string) {
			if ok.Type() == lua.LTBool {
				if len(msg) > 0 {
					s.LState().RaiseError(msg[0])
				} else {
					s.LState().RaiseError("assert failed!")
				}
			}
		},
		"exit": func(code ...int) {
			if len(code) > 0 {
				panic(contextExitCode(code[0]))
			} else {
				panic(contextExitCode(0))
			}
		},
		"set_os": func(os string) {
			platform := tbl.Value().RawGetString("platform").(*lua.LUserData).Value.(util.Platform)
			platform.OS = os
			tbl.Value().RawSetString("platform", luar.New(s.LState(), platform))
		},
	})
	if flags != nil {
		if tmp, ok := flags.GetLTable(name); ok {
			tbl.SetLTable("flags", tmp)
		}
	}
	source := ""
	if name == job.Name() {
		source = name
	} else {
		source = name + ":" + job.Name()
	}
	return &Context{
		s:      s,
		cancel: cancel,
		tbl:    tbl,
		source: source,
	}
}

func (ctx *Context) Call(fn *lua.LFunction) (c contextExitCode) {
	c = contextExitCode(1)
	defer func() {
		msg := recover()
		switch v := msg.(type) {
		case error:
			fmt.Print(v)
			panic(v)
		case contextExitCode:
			c = v
		}
	}()
	err := ctx.s.Call(yocki.YockFuncInfo{
		Fn: fn,
	}, ctx.tbl.Value())
	if err != nil {
		c = contextExitCode(0)
	}
	return
}

func (ctx *Context) Extends(super *Context) {
	platform := super.tbl.Value().RawGetString("platform").(*lua.LUserData).Value.(util.Platform)
	ctx.tbl.Value().RawSetString("platform", luar.New(ctx.s.LState(), platform))
	ctx.tbl.Value().RawSetString("flags", super.tbl.Value().RawGetString("flags"))
}

func (ctx *Context) Close() {
	if ctx.cancel != nil {
		ctx.cancel()
	}
}
