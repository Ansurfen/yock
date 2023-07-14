// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// dns, plugin, and driver are all derivatives of the dependency analysis pattern.
// They are now abandoned, see pack/dependency.go for details.
package yocks

import (
	"fmt"
	"path"
	"regexp"

	yocki "github.com/ansurfen/yock/interface"
	yockr "github.com/ansurfen/yock/runtime"
	"github.com/ansurfen/yock/util"

	lua "github.com/yuin/gopher-lua"
)

type yockDriverManager struct {
	plugins yocki.Table
	drivers yocki.Table

	globalDNS *DNS
	localDNS  *DNS
}

func newDriverManager() *yockDriverManager {
	return &yockDriverManager{
		drivers:   yockr.NewTable(),
		plugins:   yockr.NewTable(),
		globalDNS: CreateDNS(util.Pathf("@/global.json")),
		localDNS:  CreateDNS(util.Pathf("@/local.json")),
	}
}

func loadDriver(yocks yocki.YockScheduler) {
	yocks.RegYocksFn(yocki.YocksFuncs{
		"set_driver": driverSetDriver,
		// driver is a callback that works on developer of yock driver
		"driver":      driverDriver,
		"exec_driver": driverExecDriver,
	})
}

// driverSetDriver overloads the implementation of the specified function
//
// @param name string
//
// @param fn function
func driverSetDriver(yocks yocki.YockScheduler, state yocki.YockState) int {
	ys := yocks.(*YockScheduler)
	ys.getDrivers().Value().RawSetString(state.CheckString(1), state.CheckFunction(2))
	return 0
}

/*
* @param driver string
* @param name string
* @return string
 */
func driverDriver(yocks yocki.YockScheduler, state yocki.YockState) int {
	driver := state.CheckString(1)
	name := state.CheckString(2)
	out, err := util.ReadStraemFromFile(path.Join(util.DriverPath, name+".lua"))
	if err != nil {
		panic(err)
	}
	reg := regexp.MustCompile(`driver\s*\((.*)function`)
	did := driver + "_" + name
	yocks.Eval(reg.ReplaceAllString(string(out), fmt.Sprintf(`driver("%s",function`, did)))
	yocks.SetGlobalVar(driver, yocks.(*YockScheduler).getDrivers().Value().RawGetString(did))
	state.Push(lua.LString(did))
	return 1
}

// @param name string
//
// @param args ...string
func driverExecDriver(yocks yocki.YockScheduler, state yocki.YockState) int {
	if lv := yocks.(*YockScheduler).getDrivers().Value().RawGetString(state.CheckString(1)); lv.Type() == lua.LTFunction {
		args := []lua.LValue{}
		for i := 3; i <= state.Argc(); i++ {
			args = append(args, state.CheckLValue(i))
		}
		yocks.EvalFunc(lv.(*lua.LFunction), append([]lua.LValue{state.CheckLTable(2)}, args...))
	}
	return 0
}
