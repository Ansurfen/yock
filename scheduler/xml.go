package scheduler

import (
	"github.com/ansurfen/cushion/runtime"
	"github.com/beevik/etree"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func loadXML(vm *YockScheduler) runtime.Handles {
	return runtime.Handles{
		"xml": func(l *lua.LState) int {
			l.Push(luar.New(l, etree.NewDocument()))
			return 1
		},
	}
}
