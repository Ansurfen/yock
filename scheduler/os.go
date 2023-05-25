package scheduler

import (
	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/cmd"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

var osFuncs = runtime.LuaFuncs{
	"shell": osShell,
	"exec":  osExec,
}

func osExec(l *lua.LState) int {
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
}

func osShell(l *lua.LState) int {
	cmds := l.CheckString(1)
	utils.ReadLineFromString(cmds, func(s string) string {
		if len(s) > 0 {
			cmd.Exec(cmd.ExecOpt{}, []string{s})
		}
		return ""
	})
	return 0
}
