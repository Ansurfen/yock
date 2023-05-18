package scheduler

import (
	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/cushion/utils"
	lua "github.com/yuin/gopher-lua"
)

func loadIO() runtime.Handles {
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
	}
}
