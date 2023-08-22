// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// dns, plugin, and driver are all derivatives of the dependency analysis pattern.
// They are now abandoned, see pack/dependency.go for details.
package yocks

import (
	"fmt"
	"regexp"
	"strings"

	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util"
	lua "github.com/yuin/gopher-lua"
)

func loadPlugin(yocks yocki.YockScheduler) {
	yocks.RegYocksFn(yocki.YocksFuncs{
		"parse_plugin": pluginParsePlugin,
		"load_plugin":  pluginLoadPlugin,
		"plugin":       pluginPlugin,
	})
}

// @param plugin string
//
// @return string, table
func pluginParsePlugin(yocks yocki.YockScheduler, state yocki.YockState) int {
	plugin := state.CheckString(1)
	path := ""
	args := &lua.LTable{}
	if strings.Contains(plugin, "@") {
		if p, a, ok := strings.Cut(state.CheckString(1), "@"); ok {
			path = p
			if kvs := strings.Split(a, "&"); len(kvs) > 0 {
				for _, kv := range kvs {
					bind := strings.Split(kv, ":")
					if len(bind) == 2 {
						args.RawSetString(bind[0], lua.LString(bind[1]))
					}
				}
			}
		}
	} else {
		path = plugin
	}
	state.Push(lua.LString(path))
	state.Push(args)
	return 2
}

// @param file string
//
// @return string
func pluginLoadPlugin(yocks yocki.YockScheduler, state yocki.YockState) int {
	file := state.CheckString(1)
	out, err := util.ReadStraemFromFile(file)
	if err != nil {
		panic(err)
	}
	uid := util.RandString(8)
	reg := regexp.MustCompile(`plugin\s*\((.*)\s*\{`)
	yocks.Eval(reg.ReplaceAllString(string(out), fmt.Sprintf(`plugin("%s",{`, uid)))
	state.Push(lua.LString(uid))
	return 1
}

// @param uid string
//
// @param tbl table
func pluginPlugin(yocks yocki.YockScheduler, state yocki.YockState) int {
	// uid := state.CheckString(1)
	// tbl := yocks.(*YockScheduler).getPlugins()
	// cur := &lua.LTable{}
	// tbl.Value().RawSetString(uid, cur)
	// state.CheckLTable(2).ForEach(func(fn, callback lua.LValue) {
	// 	cur.RawSetString(fn.String(), callback)
	// })
	return 0
}
