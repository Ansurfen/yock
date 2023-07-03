package yockp

import (
	"testing"

	"github.com/ansurfen/yock/scheduler"
)

func TestCompile(t *testing.T) {
	yockp := New()
	yockp.Compile(CompileOpt{
		VM: yocks.New(),
	}, "print.lua")
}
