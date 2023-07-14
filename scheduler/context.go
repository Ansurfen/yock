// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocks

import (
	"context"
	"fmt"
	"time"

	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/ycho"
	lua "github.com/yuin/gopher-lua"
)

type Context struct {
	s      yocki.YockState
	tbl    yocki.Table
	cancel context.CancelFunc
}

func newContext(name string, job *yockJob, yocks yocki.YockScheduler) *Context {
	s, cancel := yocks.NewState()
	tbl := yocks.Env().Meta().Clone(s.LState())
	tbl.SetString("job", name)
	tbl.SetField(s.LState(), map[string]any{
		"info": func(msg string) {
			dbg, ok := s.Stack(1)
			if ok {
				ycho.Info(fmt.Sprintf("%s:%d %s", dbg.Source, dbg.CurrentLine, msg))
			}
		},
	})
	return &Context{
		s:      s,
		cancel: cancel,
		tbl:    tbl,
	}
}

func (ctx *Context) LValue() lua.LValue {
	return ctx.tbl.Value()
}

func (ctx *Context) Close() {
	if ctx.cancel != nil {
		ctx.cancel()
	}
}

func (ctx *Context) Suspend(deadline ...time.Duration) {

}

func (ctx *Context) Infof(msg string, a ...any) {
	dbg, ok := ctx.s.Stack(1)
	if ok {
		ycho.Infof("%s:%d %s", dbg.Source, dbg.CurrentLine, msg)
	}
}
