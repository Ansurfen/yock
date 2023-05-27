// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// dns, plugin, and driver are all derivatives of the dependency analysis pattern.
// They are now abandoned, see pack/dependency.go for details.
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

// driverSetDriver overloads the implementation of the specified function
//
// @param name string
//
// @param fn function
func driverSetDriver(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		yocks.getDrivers().RawSetString(l.CheckString(1), l.CheckFunction(2))
		return 0
	}
}

/*
* @param driver string
* @param name string
* @return string
 */
func driverDriver(yocks *YockScheduler) lua.LGFunction {
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

// @param name string
//
// @param args ...string
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
