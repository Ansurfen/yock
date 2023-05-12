package scheduler

import (
	"fmt"

	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/cmd"
	yock "github.com/ansurfen/yock/internal/deprecated"
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
		"mkdir": func(l *lua.LState) int {
			utils.SafeMkdirs(l.CheckString(1))
			return 0
		},
		"mv": func(l *lua.LState) int {
			mv := yock.NewMoveCmd()
			mv.Exec(fmt.Sprintf("%s %s", l.CheckString(1), l.CheckString(2)))
			return 0
		},
		"cp": func(l *lua.LState) int {
			cp := yock.NewCpCmd()
			cp.Exec(fmt.Sprintf("-r %s %s", l.CheckString(1), l.CheckString(2)))
			return 0
		},
	}
}
