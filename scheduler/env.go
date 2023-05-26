package scheduler

import (
	"os"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/util"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func loadEnv(yocks *YockScheduler) luaFuncs {
	yocks.env.RawSetString("platform", luar.New(yocks.Interp(), utils.CurPlatform))
	yocks.env.RawSetString("workdir", lua.LString(util.WorkSpace))
	yocks.env.RawSetString("yock_path", lua.LString(util.YockPath))
	yocks.mountLib(yocks.env, luaFuncs{
		"set_path":  envSetPath(yocks),
		"safe_set":  envSafeSet(yocks),
		"set":       envSet(yocks),
		"unset":     envUnset(yocks),
		"setl":      envSetL(yocks),
		"safe_setl": envSafeSetL(yocks),
		"export":    envExport(yocks),
		"print":     envPrint(yocks),
		"get_all":   envGetAll,
		"set_args":  envSetArgs,
	})
	return nil
}

func envSetPath(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		err := yocks.envVar.SetPath(l.CheckString(1))
		handleErr(l, err)
		return 1
	}
}

func envSafeSet(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		err := yocks.envVar.SafeSet(l.CheckString(1), envVarTypeCvt(l.CheckAny(2)))
		handleErr(l, err)
		return 1
	}
}

func envSet(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		err := yocks.envVar.Set(l.CheckString(1), envVarTypeCvt(l.CheckAny(2)))
		handleErr(l, err)
		return 1
	}
}

func envUnset(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		err := yocks.envVar.Unset(l.CheckString(1))
		handleErr(l, err)
		return 1
	}
}

func envSetL(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		err := yocks.envVar.SetL(l.CheckString(1), l.CheckString(2))
		handleErr(l, err)
		return 1
	}
}

func envSafeSetL(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		err := yocks.envVar.SafeSetL(l.CheckString(1), l.CheckString(2))
		handleErr(l, err)
		return 1
	}
}

func envExport(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		err := yocks.envVar.Export(l.CheckString(1))
		handleErr(l, err)
		return 1
	}
}

func envPrint(yocks *YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		yocks.envVar.Print()
		return 0
	}
}

func envGetAll(l *lua.LState) int {
	envs := &lua.LTable{}
	for i, e := range os.Environ() {
		envs.Insert(i+1, lua.LString(e))
	}
	l.Push(envs)
	return 1
}

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
