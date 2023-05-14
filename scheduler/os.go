package scheduler

import (
	"os"

	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/yock/cmd"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

func loadOS() runtime.Handles {
	return runtime.Handles{
		"exec": func(l *lua.LState) int {
			mode := l.CheckAny(1)
			opt := cmd.ExecOpt{Quiet: true}
			cmds := []string{}
			if mode.Type() == lua.LTTable {
				gluamapper.Map(l.CheckTable(1), &opt)
				for i := 2; i <= l.GetTop(); i++ {
					cmds = append(cmds, l.CheckString(i))
				}
			} else {
				for i := 1; i < l.GetTop(); i++ {
					cmds = append(cmds, l.CheckString(i))
				}
			}
			cmd.Exec(opt, cmds)
			return 0
		},
		"pwd": func(l *lua.LState) int {
			path, err := os.Getwd()
			l.Push(lua.LString(path))
			if err != nil {
				l.Push(lua.LFalse)
			} else {
				l.Push(lua.LTrue)
			}
			return 2
		},
	}
}
