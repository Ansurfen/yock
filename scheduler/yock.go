package scheduler

import (
	"os"
	"path/filepath"
	"time"

	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/util"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type YockScheduler struct {
	runtime.VirtualMachine
	env *lua.LTable
	opt *lua.LTable

	envVar utils.EnvVar

	// it's deprecated in lateset version
	driverManager *yockDriverManager

	task       map[string][]*yockJob
	goroutines chan func()
	signals    SignalStream
}

func New(opts ...YockSchedulerOption) *YockScheduler {
	scheduler := &YockScheduler{
		VirtualMachine: runtime.NewVirtualMachine().Default(),
		env:            &lua.LTable{},
		goroutines:     make(chan func(), 10),
		signals:        NewSingleSignalStream(),
		task:           make(map[string][]*yockJob),
	}

	for _, opt := range opts {
		if err := opt(scheduler); err != nil {
			util.Ycho.Fatal(err.Error())
		}
	}

	if err := utils.SafeBatchMkdirs([]string{util.PluginPath, util.DriverPath}); err != nil {
		util.Ycho.Fatal(err.Error())
	}

	scheduler.parseFlags()
	scheduler.loadLibs()

	return scheduler
}

func (s *YockScheduler) getPlugins() *lua.LTable {
	if s.driverManager != nil {
		return s.driverManager.plugins
	}
	return &lua.LTable{}
}

func (s *YockScheduler) getDrivers() *lua.LTable {
	if s.driverManager != nil {
		return s.driverManager.drivers
	}
	return &lua.LTable{}
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

func (vm *YockScheduler) parseFlags() {
	idx := 0
	args := &lua.LTable{}
	vm.env.RawSetString("args", args)
	vm.env.RawSetString("platform", luar.New(vm.Interp(), utils.CurPlatform))
	vm.env.RawSetString("workdir", lua.LString(util.WorkSpace))
	vm.env.RawSetString("set_path", vm.Interp().NewClosure(func(l *lua.LState) int {
		err := vm.envVar.SetPath(l.CheckString(1))
		handleErr(l, err)
		return 1
	}))
	vm.env.RawSetString("safe_set", vm.Interp().NewClosure(func(l *lua.LState) int {
		err := vm.envVar.SafeSet(l.CheckString(1), envVarTypeCvt(l.CheckAny(2)))
		handleErr(l, err)
		return 1
	}))
	vm.env.RawSetString("set", vm.Interp().NewClosure(func(l *lua.LState) int {
		err := vm.envVar.Set(l.CheckString(1), envVarTypeCvt(l.CheckAny(2)))
		handleErr(l, err)
		return 1
	}))
	vm.env.RawSetString("unset", vm.Interp().NewClosure(func(l *lua.LState) int {
		err := vm.envVar.Unset(l.CheckString(1))
		handleErr(l, err)
		return 1
	}))
	vm.env.RawSetString("setl", vm.Interp().NewClosure(func(l *lua.LState) int {
		err := vm.envVar.SetL(l.CheckString(1), l.CheckString(2))
		handleErr(l, err)
		return 1
	}))
	vm.env.RawSetString("safe_setl", vm.Interp().NewClosure(func(l *lua.LState) int {
		err := vm.envVar.SafeSetL(l.CheckString(1), l.CheckString(2))
		handleErr(l, err)
		return 1
	}))
	vm.env.RawSetString("export", vm.Interp().NewClosure(func(l *lua.LState) int {
		err := vm.envVar.Export(l.CheckString(1))
		handleErr(l, err)
		return 1
	}))
	vm.env.RawSetString("print", vm.Interp().NewClosure(func(l *lua.LState) int {
		vm.envVar.Print()
		return 0
	}))
	vm.env.RawSetString("get_all", vm.Interp().NewClosure(func(l *lua.LState) int {
		envs := &lua.LTable{}
		for i, e := range os.Environ() {
			envs.Insert(i+1, lua.LString(e))
		}
		l.Push(envs)
		return 1
	}))
	vm.env.RawSetString("set_args", vm.Interp().NewClosure(func(l *lua.LState) int {
		os.Args = append(os.Args[:0], os.Args[0])
		l.CheckTable(1).ForEach(func(_, s lua.LValue) {
			os.Args = append(os.Args, s.String())
		})
		return 0
	}))
	vm.env.RawSetString("yock_path", lua.LString(util.YockPath))
	for i, j := 0, 1; i < len(os.Args); i++ {
		if os.Args[i] == "--" {
			idx = i
			continue
		}
		if idx > 0 {
			args.Insert(j, lua.LString(os.Args[i]))
			j++
		}
	}
	if idx > 0 {
		os.Args = os.Args[:idx]
	}
}

const (
	yockLibRandom = "random"
	yockLibSSH    = "ssh"
	yockLibTime   = "time"
	yockLibJSON   = "json"
	yockLibPath   = "path"
	yockLibRegexp = "regexp"
	yockLibSync   = "sync"
	yockLibPsutil = "psutil"
)

type yockLib struct {
	name   string
	handle func(*YockScheduler) lua.LValue
}

var yockLibs = []yockLib{
	{yockLibRandom, loadRandom},
	{yockLibSSH, loadSSH},
	{yockLibJSON, loadJSON},
	{yockLibTime, loadTime},
	{yockLibPath, loadPath},
	{yockLibRegexp, loadRegexp},
	{yockLibSync, loadSync},
	{yockLibPsutil, loadPsutil},
}

type yockFunc func(*YockScheduler) runtime.LuaFuncs

func wrapLuaFuns(fs runtime.LuaFuncs) yockFunc {
	return func(ys *YockScheduler) runtime.LuaFuncs {
		return fs
	}
}

var yockFuncs = []yockFunc{
	netFuncs,
	goroutineFuncs,
	loadStrings,
	loadPlugin,
	loadDriver,
	taskFuncs,
	loadCtl,
	loadXML,
	wrapLuaFuns(osFuncs),
	wrapLuaFuns(gnuFuncs),
	wrapLuaFuns(ioFuncs()),
}

func (s *YockScheduler) loadLibs() {
	for _, fn := range yockFuncs {
		s.SetGlobalFn(fn(s))
	}

	var yockGlobalVars = map[string]lua.LValue{
		"env": s.env,
	}

	if s.driverManager != nil {
		yockGlobalVars["plugins"] = s.driverManager.plugins
		yockGlobalVars["ldns"] = luar.New(s.Interp(), s.driverManager.localDNS)
		yockGlobalVars["gdns"] = luar.New(s.Interp(), s.driverManager.globalDNS)
	}
	s.setGlobalVars(yockGlobalVars)

	files, err := os.ReadDir(util.Pathf("~/lib"))
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
	for _, file := range files {
		if fn := file.Name(); filepath.Ext(fn) == ".lua" {
			if err := s.EvalFile(util.Pathf("~/lib/") + fn); err != nil {
				util.Ycho.Fatal(err.Error())
			}
		}
	}

	// self-boot
	files, err = os.ReadDir(util.Pathf("~/lib/boot"))
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
	for _, file := range files {
		if fn := file.Name(); filepath.Ext(fn) == ".lua" {
			if err := s.EvalFile(util.Pathf("~/lib/") + fn); err != nil {
				util.Ycho.Fatal(err.Error())
			}
		}
	}

	for _, lib := range yockLibs {
		s.SetGlobalVar(lib.name, lib.handle(s))
	}
	// s.Eval(`Import({"cushion-check", "cushion-vm"})`)
}

func (vm *YockScheduler) setGlobalVars(vars map[string]lua.LValue) {
	for k, v := range vars {
		vm.SetGlobalVar(k, v)
	}
}

func DeepCopy(L *lua.LState, tbl *lua.LTable) *lua.LTable {
	table := tbl
	newTable := &lua.LTable{}
	copyTable(L, table, newTable)
	return newTable
}

func copyTable(L *lua.LState, srcTable *lua.LTable, dstTable *lua.LTable) {
	srcTable.ForEach(func(key lua.LValue, value lua.LValue) {
		if tbl, ok := value.(*lua.LTable); ok {
			newTbl := L.NewTable()
			copyTable(L, tbl, newTbl)
			dstTable.RawSet(key, newTbl)
		} else {
			dstTable.RawSet(key, value)
		}
	})
}

func (vm *YockScheduler) LaunchTask(name string) {
	var flags *lua.LTable
	if vm.opt != nil {
		if tmp, ok := vm.opt.RawGetString("flags").(*lua.LTable); ok {
			flags = tmp
		}
	}
	for _, job := range vm.task[name] {
		tmp, cancel := vm.Interp().NewThread()
		tbl := DeepCopy(tmp, vm.env)
		tbl.RawSetString("job", lua.LString(name))
		if flags != nil {
			if tmp, ok := flags.RawGetString(name).(*lua.LTable); ok {
				tbl.RawSetString("flags", tmp)
			}
		}
		if err := tmp.CallByParam(lua.P{
			Fn: job.fn,
		}, tbl); err != nil {
			util.Ycho.Warn(err.Error())
		}
		if cancel != nil {
			cancel()
		}
	}
}

// scan ast to determine whether it enable
func (vm *YockScheduler) EventLoop() {
	for {
		select {
		case fn := <-vm.goroutines:
			go fn()
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func (vm *YockScheduler) DoCompliedFile(proto *lua.FunctionProto) error {
	lvm := vm.Interp()
	lfunc := lvm.NewFunctionFromProto(proto)
	lvm.Push(lfunc)
	return lvm.PCall(0, lua.MultRet, nil)
}

func handleErr(l *lua.LState, err error) {
	if err != nil {
		l.Push(lua.LString(err.Error()))
	} else {
		l.Push(lua.LNil)
	}
}

func handleBool(l *lua.LState, b bool) {
	if b {
		l.Push(lua.LTrue)
	} else {
		l.Push(lua.LFalse)
	}
}
