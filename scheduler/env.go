// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"os"

	"github.com/ansurfen/cushion/utils"
	yocki "github.com/ansurfen/yock/interface"
	yockr "github.com/ansurfen/yock/runtime"
	"github.com/ansurfen/yock/util"
	lua "github.com/yuin/gopher-lua"
)

func loadEnv(yocks yocki.YockScheduler) {
	lib := yocks.OpenLib("env")
	lib.SetField(map[string]any{
		"platform":  utils.CurPlatform,
		"workdir":   util.WorkSpace,
		"yock_path": util.YockPath,
		"conf":      util.Conf(),
	})
	lib.SetFunctions(map[string]lua.LGFunction{
		"environ":  envEnviron,
		"set_args": envSetArgs,
	})
	yocks.MntYocksFn(lib, yocki.YocksFuncs{
		"set_path":  envSetPath,
		"safe_set":  envSafeSet,
		"set":       envSet,
		"unset":     envUnset,
		"setl":      envSetL,
		"safe_setl": envSafeSetL,
		"export":    envExport,
		"print":     envPrint,
	})
}

// @param path string
//
// @return err
func envSetPath(yocks yocki.YockScheduler, s *yockr.YockState) int {
	err := yocks.EnvVar().SetPath(s.CheckString(1))
	s.PushError(err)
	return 1
}

/*
* @param key string
* @param value string
* @return err
 */
func envSafeSet(yocks yocki.YockScheduler, s *yockr.YockState) int {
	err := yocks.EnvVar().SafeSet(s.CheckString(1), envVarTypeCvt(s.CheckAny(2)))
	s.PushError(err)
	return 1
}

/*
* @param key string
* @param value string
* @return err
 */
func envSet(yocks yocki.YockScheduler, s *yockr.YockState) int {
	err := yocks.EnvVar().Set(s.CheckString(1), envVarTypeCvt(s.CheckAny(2)))
	s.PushError(err)
	return 1
}

// @param key string
//
// @return err
func envUnset(yocks yocki.YockScheduler, s *yockr.YockState) int {
	err := yocks.EnvVar().Unset(s.CheckString(1))
	s.PushError(err)
	return 1
}

/*
* @param key string
* @param value string
* @return err
 */
func envSetL(yocks yocki.YockScheduler, s *yockr.YockState) int {
	err := yocks.EnvVar().SetL(s.CheckString(1), s.CheckString(2))
	s.PushError(err)
	return 1
}

/*
* @param key string
* @param value string
* @return err
 */
func envSafeSetL(yocks yocki.YockScheduler, s *yockr.YockState) int {
	err := yocks.EnvVar().SafeSetL(s.CheckString(1), s.CheckString(2))
	s.PushError(err)
	return 1
}

// envExport exports current enviroment string into specify file
//
// @param path string
//
// @return err
func envExport(yocks yocki.YockScheduler, s *yockr.YockState) int {
	err := yocks.EnvVar().Export(s.CheckString(1))
	s.PushError(err)
	return 1
}

// envPrint prints current enviroment variable
func envPrint(yocks yocki.YockScheduler, s *yockr.YockState) int {
	yocks.EnvVar().Print()
	return 0
}

// envEnviron returns table of enviroment variable
//
// @return table
func envEnviron(l *lua.LState) int {
	envs := &lua.LTable{}
	for i, e := range os.Environ() {
		envs.Insert(i+1, lua.LString(e))
	}
	l.Push(envs)
	return 1
}

// envSetArgs resets the value of os.Args
//
// @param args table
func envSetArgs(l *lua.LState) int {
	os.Args = append(os.Args[:0], os.Args[0])
	l.CheckTable(1).ForEach(func(_, s lua.LValue) {
		os.Args = append(os.Args, s.String())
	})
	return 0
}

func envVarTypeCvt(v lua.LValue) any {
	switch vv := v.(type) {
	case lua.LString:
		return vv.String()
	case *lua.LTable:
		str := []string{}
		vv.ForEach(func(_, s lua.LValue) {
			str = append(str, s.String())
		})
		return str
	default:
		return nil
	}
}
