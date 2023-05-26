package scheduler

import (
	"sync"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func loadSync(yocks *YockScheduler) lua.LValue {
	return yocks.registerLib(syncLib)
}

var syncLib = luaFuncs{
	"new": syncNew,
}

func syncNew(l *lua.LState) int {
	l.Push(luar.New(l, &sync.WaitGroup{}))
	return 1
}
