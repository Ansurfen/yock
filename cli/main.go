package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/cushion/utils"
	_ "github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock"
	"github.com/gocolly/colly"
	lua "github.com/yuin/gopher-lua"
)

var (
	file            = flag.String("f", "", "")
	enableLuaScript = flag.Bool("l", false, "")
)

func main() {
	flag.Parse()
	if len(*file) != 0 {
		if *enableLuaScript {
			vm := runtime.NewVirtualMachine().Default()
			vm.RegisterModule(runtime.LuaFuncs{
				"yock-io": func(l *lua.LState) int {
					return runtime.LuaModuleLoader(
						vm.Interp(), runtime.LuaFuncs{
							"move": func(l *lua.LState) int {
								mv := yock.NewMoveCmd()
								mv.Exec(fmt.Sprintf("%s %s", l.CheckString(1), l.CheckString(2)))
								return 0
							},
							"copy": func(l *lua.LState) int {
								cp := yock.NewCpCmd()
								cp.Exec(fmt.Sprintf("-r %s %s", l.CheckString(1), l.CheckString(2)))
								return 0
							},
							"curl": func(l *lua.LState) int {
								c := colly.NewCollector()
								c.OnResponse(func(r *colly.Response) {
									utils.WriteFile(l.CheckString(2), r.Body)
								})
								c.Visit(l.CheckString(1))
								return 0
							},
							"file_replace": func(l *lua.LState) int {
								src, err := utils.ReadStraemFromFile(l.CheckString(1))
								if err != nil {
									panic(err)
								}
								dst := strings.ReplaceAll(string(src), l.CheckString(2), l.CheckString(3))
								utils.WriteFile(l.CheckString(1), []byte(dst))
								return 0
							},
							"exec": func(l *lua.LState) int {
								l.CheckTable(1).ForEach(func(_, cmd lua.LValue) {
									if _, err := utils.ExecStr(cmd.String()); err != nil {
										return
									}
								})
								return 0
							},
							"compose": func(l *lua.LState) int { return 0 },
						},
					)
				},
			})
			vm.EvalFile(*file)
		} else {
			yock.LoadBySh(*file)
		}
	} else {
		out, err := yock.LoadByStr(strings.Join(os.Args[1:], " "))
		if err != nil {
			panic(err)
		}
		if len(out) > 0 {
			fmt.Println(string(out))
		}
	}
}
