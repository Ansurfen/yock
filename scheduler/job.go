package scheduler

import (
	"fmt"
	"os"

	lua "github.com/yuin/gopher-lua"
)

type yockJob struct {
	fn *lua.LFunction
}

func taskFuncs(yocks *YockScheduler) luaFuncs {
	return luaFuncs{
		"job":        taskJob(yocks),
		"jobs":       taskJobs(yocks),
		"job_option": taskJobOption(yocks),
	}
}

func taskJob(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		jobName := l.CheckString(1)
		jobFn := l.CheckFunction(2)
		if _, ok := yocks.task[jobName]; ok {
			fmt.Println("dumplicate job name")
			os.Exit(1)
		} else {
			yocks.task[jobName] = append(yocks.task[jobName], &yockJob{
				fn: jobFn,
			})
		}
		return 0
	}
}

func taskJobs(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		groups := []string{}
		for i := 1; i <= l.GetTop(); i++ {
			groups = append(groups, l.CheckString(i))
		}
		if len(groups) <= 1 {
			return 0
		}
		name := groups[0]
		if _, ok := yocks.task[name]; ok {
			fmt.Println("dumplicate job name")
			os.Exit(1)
		}
		for _, n := range groups[1:] {
			if job, ok := yocks.task[n]; ok {
				yocks.task[name] = append(yocks.task[name], job...)
			}
		}
		return 0
	}
}

func taskJobOption(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		yocks.opt = l.CheckTable(1)
		return 0
	}
}
