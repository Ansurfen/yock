package scheduler

import (
	"fmt"
	"path"
	"regexp"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/util"

	lua "github.com/yuin/gopher-lua"
)

type yockDriverManager struct {
	plugins *lua.LTable
	drivers *lua.LTable

	globalDNS *DNS
	localDNS  *DNS
}

func newDriverManager() *yockDriverManager {
	return &yockDriverManager{
		drivers:   &lua.LTable{},
		plugins:   &lua.LTable{},
		globalDNS: CreateDNS(util.Pathf("@/global.json")),
		localDNS:  CreateDNS(util.Pathf("@/local.json")),
	}
}

func loadDriver(yocks *YockScheduler) luaFuncs {
	return luaFuncs{
		"set_driver": driverSetDriver(yocks),
		// driver is a callback that works on developer of yock driver
		"driver":      driverDriver(yocks),
		"exec_driver": driverExecDriver(yocks),
	}
}

func driverSetDriver(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		driver := l.CheckString(1)
		name := l.CheckString(2)
		out, err := utils.ReadStraemFromFile(path.Join(util.DriverPath, name+".lua"))
		if err != nil {
			panic(err)
		}
		reg := regexp.MustCompile(`driver\s*\((.*)function`)
		did := driver + "_" + name
		yocks.Eval(reg.ReplaceAllString(string(out), fmt.Sprintf(`driver("%s",function`, did)))
		yocks.SetGlobalVar(driver, yocks.getDrivers().RawGetString(did))
		l.Push(lua.LString(did))
		return 1
	}
}

func driverDriver(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		yocks.getDrivers().RawSetString(l.CheckString(1), l.CheckFunction(2))
		return 0
	}
}

func driverExecDriver(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		if lv := yocks.getDrivers().RawGetString(l.CheckString(1)); lv.Type() == lua.LTFunction {
			args := []lua.LValue{}
			for i := 3; i <= l.GetTop(); i++ {
				args = append(args, l.CheckAny(i))
			}
			yocks.EvalFunc(lv.(*lua.LFunction), append([]lua.LValue{l.CheckTable(2)}, args...))
		}
		return 0
	}
}
