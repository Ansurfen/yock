// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocks

import (
	"time"

	yockr "github.com/ansurfen/yock/runtime"
	"github.com/ansurfen/yock/ycho"
)

type Context struct {
	s *yockr.YockState
}

func (ctx *Context) Suspend(deadline ...time.Duration) {
}

func (ctx *Context) Infof(msg string, a ...any) {
	dbg, ok := ctx.s.Stack(1)
	if ok {
		ycho.Infof("%s:%d %s", dbg.Source, dbg.CurrentLine, msg)
	}
}
