// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
	"os"
	"testing"

	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/util/test"
)

func TestIO(t *testing.T) {
	Rm(RmOpt{Safe: false}, "a")
	Rm(RmOpt{Safe: false}, "b")
	util.Mkdirs("a")
	err := Mv(MvOpt{}, "a", "b")
	test.Assert(err == nil)
	_, err = os.Stat("a")
	test.Assert(err != nil)
	err = Cp(CpOpt{Recurse: true}, "b", "a")
	test.Assert(err == nil)
	err = Rm(RmOpt{
		Safe: false,
	}, "a")
	test.Assert(err == nil)
	_, err = os.Stat("a")
	test.Assert(err != nil)
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
		Quiet:    false,
	}, "echo a")
}

func TestCurl(t *testing.T) {
	Curl(CurlOpt{
		Method: "GET",
		Save:   true,
		Dir:    ".",
		FilenameHandle: func(s string) string {
			return s
		},
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
	Unset(UnsetOpt{}, "a", "")
}

func TestUntar(t *testing.T) {
	// Rm(RmOpt{Safe: false}, "tmp")
	// util.Mkdirs("tmp/a/b/c")
	// err := Tar("tmp", "yock.tar.gz")
	// test.Assert(err == nil, err.Error())
	// err = Untar("yock.tar.gz", "aaa")
	// test.Assert(err == nil, "fail to untar")
}

func TestUserUnset(t *testing.T) {
	fmt.Println(util.NewTemplate().OnceParse(unsetExpandWindows, map[string]string{
		"Target": "User",
		"Key":    "demo",
		"Value":  "1",
	}))
}
