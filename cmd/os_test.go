package cmd

import "testing"

func TestExec(t *testing.T) {
	Exec(ExecOpt{
		Redirect: true,
		Caller:   "TestExec",
		Quiet:    false,
	}, []string{"echo a"})
}
