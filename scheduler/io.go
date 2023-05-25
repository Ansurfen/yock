package scheduler

import (
	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/cushion/utils"
	lua "github.com/yuin/gopher-lua"
)

func ioFuncs() runtime.Handles {
	return runtime.Handles{
		"safe_write": func(l *lua.LState) int {
			err := utils.SafeWriteFile(l.CheckString(1), []byte(l.CheckString(2)))
			handleErr(l, err)
			return 1
		},
		"zip": func(l *lua.LState) int {
			zipPath := l.CheckString(1)
			paths := []string{}
			for i := 2; i <= l.GetTop(); i++ {
				paths = append(paths, l.CheckString(i))
			}
			err := utils.Zip(zipPath, paths...)
			handleErr(l, err)
			return 1
		},
		"unzip": func(l *lua.LState) int {
			err := utils.Unzip(l.CheckString(1), l.CheckString(2))
			handleErr(l, err)
			return 1
		},
		"write_file": func(l *lua.LState) int {
			err := utils.WriteFile(l.CheckString(1), []byte(l.CheckString(2)))
			handleErr(l, err)
			return 1
		},
		"is_exist": func(l *lua.LState) int {
			ok := utils.IsExist(l.CheckString(1))
			handleBool(l, ok)
			return 1
		},
		"printf": func(l *lua.LState) int {
			title := []string{}
			rows := [][]string{}
			l.CheckTable(1).ForEach(func(idx, el lua.LValue) {
				title = append(title, el.String())
			})
			l.CheckTable(2).ForEach(func(ri, row lua.LValue) {
				tmp := []string{}
				row.(*lua.LTable).ForEach(func(fi, field lua.LValue) {
					tmp = append(tmp, field.String())
				})
				rows = append(rows, tmp)
			})
			utils.Prinf(utils.PrintfOpt{MaxLen: 30}, title, rows)
			return 0
		},
	}
}
