package cmd

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
	}, []string{"echo a"})
}

func TestHTTP(t *testing.T) {
	HTTP(HttpOpt{
		Method: "GET",
		Save:   true,
		Debug:  true,
		Dir:    ".",
		Filename: func(s string) string {
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
