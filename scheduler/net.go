package scheduler

import (
	"fmt"
	"net"

	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/cmd"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

func loadNet(vm *YockScheduler) runtime.Handles {
	return runtime.Handles{
		"http": func(l *runtime.LuaInterp) int {
			mode := l.CheckAny(1)
			opt := cmd.HttpOpt{Method: "GET"}
			urls := []string{}
			if mode.Type() == lua.LTTable {
				if fn := l.CheckTable(1).RawGetString("fn"); fn.Type() == lua.LTFunction {
					opt.Filename = func(s string) string {
						lvm, _ := vm.Interp().NewThread()
						if err := lvm.CallByParam(lua.P{
							NRet: 1,
							Fn:   fn.(*lua.LFunction),
						}, lua.LString(s)); err != nil {
							panic(err)
						}
						return lvm.CheckString(1)
					}
				}
				gluamapper.Map(l.CheckTable(1), &opt)
				for i := 2; i <= l.GetTop(); i++ {
					urls = append(urls, l.CheckString(i))
				}
			} else {
				for i := 1; i < l.GetTop(); i++ {
					urls = append(urls, l.CheckString(i))
				}
			}
			cmd.HTTP(opt, urls)
			return 0
		},
		"is_url": func(l *lua.LState) int {
			if utils.IsURL(l.CheckString(1)) {
				l.Push(lua.LTrue)
			} else {
				l.Push(lua.LFalse)
			}
			return 1
		},
		"is_localhost": func(l *lua.LState) int {
			url := l.CheckString(1)
			if url == "localhost" {
				l.Push(lua.LTrue)
				return 1
			}
			addrs, err := net.LookupHost("localhost")
			if err != nil {
				fmt.Println("error:", err)
			}
			if len(addrs) > 1 && addrs[1] == url {
				l.Push(lua.LTrue)
			} else {
				l.Push(lua.LFalse)
			}
			return 1
		},
	}
}
