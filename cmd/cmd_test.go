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

func TestEchoConsole(t *testing.T) {
	fmt.Println(Echo(EchoOpt{}, "$GOPATH a"))
}

func TestEchoFile(t *testing.T) {
	fmt.Println(Echo(EchoOpt{Fd: []string{"file.txt"}, Mode: "c|a"}, "$GOPATH a"))
}

func TestExec(t *testing.T) {
	Exec(ExecOpt{
		Redirect: true,
		Caller:   "TestExec",
		Quiet:    false,
	}, "echo a")
}

func TestCurl(t *testing.T) {
	Curl(CurlOpt{
		Method: "GET",
		Save:   true,
		Debug:  true,
		Dir:    ".",
		FilenameHandle: func(s string) string {
			return s
		},
		Caller: "TestHTTP",
	}, []string{"https://github.com"})
}

func TestLs(t *testing.T) {
	fmt.Println(Ls(LsOpt{
		Dir: ".",
	}))
}

func TestCh(t *testing.T) {
	Chmod("", 0777)
}

func TestFind(t *testing.T) {
	fmt.Println(Find(FindOpt{Search: true, Pattern: "sd$", Dir: false, File: true}, "../bin"))
}

func TestWhereis(t *testing.T) {
	fmt.Println(Whereis("go"))
}

func TestExport(t *testing.T) {
	Export(ExportOpt{Expand: true}, "a", "b")
	Export(ExportOpt{}, "a", "c")
}

func TestUnset(t *testing.T) {
	Unset("a")
}

func TestUntar(t *testing.T) {
	fmt.Println(Untar("yock.tar.gz", "aaa"))
}