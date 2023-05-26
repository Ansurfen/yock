package scheduler

import (
	"github.com/ansurfen/cushion/utils"
	lua "github.com/yuin/gopher-lua"
)

func loadRandom(yocks *YockScheduler) lua.LValue {
	return yocks.registerLib(randomLib)
}

var randomLib = luaFuncs{
	"str": randomStr,
}

func randomStr(l *lua.LState) int {
	l.Push(lua.LString(utils.RandString(int(l.CheckNumber(1)))))
	return 1
}
