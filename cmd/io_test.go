package cmd

import "testing"

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
