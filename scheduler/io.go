package scheduler

import (
	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/yock/cmd"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

func loadIO() runtime.Handles {
	return runtime.Handles{
		"rm": func(l *lua.LState) int {
			mode := l.CheckAny(1)
			opt := cmd.RmOpt{Safe: true}
			targets := []string{}
			if mode.Type() == lua.LTTable {
				gluamapper.Map(l.CheckTable(1), &opt)
				for i := 2; i <= l.GetTop(); i++ {
					targets = append(targets, l.CheckString(i))
				}
			} else {
				for i := 1; i < l.GetTop(); i++ {
					targets = append(targets, l.CheckString(i))
				}
			}
			cmd.Rm(opt, targets)
			return 0
		},
	}
}
