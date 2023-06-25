// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
	"testing"
)

func TestRm(t *testing.T) {
	Rm(RmOpt{
		Safe:   true,
		Caller: "TestRm",
	}, []string{"test"})
}

func TestMv(t *testing.T) {
	Mv(MvOpt{
		Caller: "TestMv",
	}, "a", "b")
}

func TestCp(t *testing.T) {
	Cp(CpOpt{
		Caller: "TestCp",
	}, "a", "b")
}

func TestEcho(t *testing.T) {
	fmt.Println(Echo("$GOPATH a"))
}

func TestExec(t *testing.T) {
	Exec(ExecOpt{
		Redirect: true,
		Caller:   "TestExec",
		Quiet:    false,
	}, "echo a")
}

func TestHTTP(t *testing.T) {
	HTTP(HttpOpt{
		Method: "GET",
		Save:   true,
		Debug:  true,
		Dir:    ".",
		FilenameHandle: func(s string) string {
			return s
		},
		Caller: "TestHTTP",
	}, []string{"https://www.github.com"})
}

func TestLs(t *testing.T) {
	fmt.Println(Ls(LsOpt{
		Dir: ".",
	}))
}

func TestCh(t *testing.T) {
	Chmod("", 0777)
}
