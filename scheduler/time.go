package scheduler

import (
	"time"

	lua "github.com/yuin/gopher-lua"
)

func loadTime(yocks *YockScheduler) lua.LValue {
	timelib := yocks.registerLib(luaFuncs{
		"sleep": timeSleep,
	})
	timelib.RawSetString("microsecond", lua.LNumber(time.Microsecond))
	timelib.RawSetString("millisecond", lua.LNumber(time.Millisecond))
	timelib.RawSetString("second", lua.LNumber(time.Second))
	return timelib
}

func timeSleep(l *lua.LState) int {
	time.Sleep(time.Duration(l.CheckNumber(1)))
	return 0
}
