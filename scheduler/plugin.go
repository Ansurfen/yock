package scheduler

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/cushion/utils"
	lua "github.com/yuin/gopher-lua"
)

func loadPlugin(vm *YockScheduler) runtime.Handles {
	return runtime.Handles{
		"parse_plugin": func(l *runtime.LuaInterp) int {
			plugin := l.CheckString(1)
			path := ""
			args := &lua.LTable{}
			if strings.Contains(plugin, "@") {
				if p, a, ok := strings.Cut(l.CheckString(1), "@"); ok {
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
			l.Push(lua.LString(path))
			l.Push(args)
			return 2
		},
		"load_plugin": func(l *lua.LState) int {
			file := l.CheckString(1)
			out, err := utils.ReadStraemFromFile(file)
			if err != nil {
				panic(err)
			}
			uid := utils.RandString(8)
			reg := regexp.MustCompile(`plugin\s*\((.*)\s*\{`)
			vm.Eval(reg.ReplaceAllString(string(out), fmt.Sprintf(`plugin("%s",{`, uid)))
			l.Push(lua.LString(uid))
			return 1
		},
		"plugin": func(l *lua.LState) int {
			uid := l.CheckString(1)
			tbl := vm.getPlugins()
			cur := &lua.LTable{}
			tbl.RawSetString(uid, cur)
			l.CheckTable(2).ForEach(func(fn, callback lua.LValue) {
				cur.RawSetString(fn.String(), callback)
			})
			return 0
		},
	}
}
