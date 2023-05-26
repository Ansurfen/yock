package scheduler

import (
	"github.com/beevik/etree"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func loadXML(yocks *YockScheduler) luaFuncs {
	return luaFuncs{
		"xml": func(l *lua.LState) int {
			l.Push(luar.New(l, etree.NewDocument()))
			return 1
		},
	}
}
