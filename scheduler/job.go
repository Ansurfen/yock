package scheduler

import (
	"fmt"
	"os"

	"github.com/ansurfen/cushion/runtime"
	lua "github.com/yuin/gopher-lua"
)

type yockJob struct {
	fn *lua.LFunction
}

func loadJob(vm *YockScheduler) runtime.Handles {
	return runtime.Handles{
		"job": func(l *runtime.LuaInterp) int {
			jobName := l.CheckString(1)
			jobFn := l.CheckFunction(2)
			if _, ok := vm.jobs[jobName]; ok {
				fmt.Println("dumplicate job name")
				os.Exit(1)
			} else {
				vm.jobs[jobName] = append(vm.jobs[jobName], &yockJob{
					fn: jobFn,
				})
			}
			return 0
		},
		"jobs": func(l *runtime.LuaInterp) int {
			groups := []string{}
			for i := 1; i <= l.GetTop(); i++ {
				groups = append(groups, l.CheckString(i))
			}
			if len(groups) <= 1 {
				return 0
			}
			name := groups[0]
			if _, ok := vm.jobs[name]; ok {
				fmt.Println("dumplicate job name")
				os.Exit(1)
			}
			for _, n := range groups[1:] {
				if job, ok := vm.jobs[n]; ok {
					vm.jobs[name] = append(vm.jobs[name], job...)
				}
			}
			return 0
		},
		"job_option": func(l *lua.LState) int {
			vm.opt = l.CheckTable(1)
			return 0
		},
	}
}
