package scheduler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/cushion/utils"
	parser "github.com/ansurfen/yock/pack"
	"github.com/ansurfen/yock/util"
	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
	luar "layeh.com/gopher-luar"
)

type YockScheduler struct {
	runtime.VirtualMachine
	env     *lua.LTable
	plugins *lua.LTable
	drivers *lua.LTable
	opt     *lua.LTable

	envVar utils.EnvVar

	globalDNS *DNS
	localDNS  *DNS

	jobs       map[string][]*yockJob
	goroutines chan func()
	signals    SignalStream
}

func New(opts ...YockSchedulerOption) *YockScheduler {
	scheduler := &YockScheduler{
		VirtualMachine: runtime.NewVirtualMachine().Default(),
		env:            &lua.LTable{},
		drivers:        &lua.LTable{},
		plugins:        &lua.LTable{},
		goroutines:     make(chan func(), 10),
		signals:        NewSingleSignalStream(),
		jobs:           make(map[string][]*yockJob),
		// envVar:         utils.NewEnvVar(),
	}

	for _, opt := range opts {
		if err := opt(scheduler); err != nil {
			util.YchoFatal("", err.Error())
		}
	}

	utils.SafeBatchMkdirs([]string{util.PluginPath, util.DriverPath})

	scheduler.globalDNS = CreateDNS(util.Pathf("@/global.json"))
	scheduler.localDNS = CreateDNS(util.Pathf("@/local.json"))

	scheduler.parseFlags()
	scheduler.injectGlobal()

	return scheduler
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

func (vm *YockScheduler) injectGlobal() {
	vm.setGlobalFns(
		loadNet(vm),
		loadIO(),
		loadOS(),
		loadGoroutine(vm),
		loadStrings(vm),
		loadPlugin(vm),
		loadDriver(vm),
		loadJob(vm),
		loadCtl(vm),
		loadGNU())
	vm.setGlobalVars(map[string]lua.LValue{
		"ldns":    luar.New(vm.Interp(), vm.localDNS),
		"gdns":    luar.New(vm.Interp(), vm.globalDNS),
		"env":     vm.env,
		"plugins": vm.plugins,
	})
	files, err := os.ReadDir(util.Pathf("~/lib"))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if fn := file.Name(); filepath.Ext(fn) == ".lua" {
			if err := vm.EvalFile(util.Pathf("~/lib/") + fn); err != nil {
				panic(err)
			}
		}
	}
	if err := vm.EvalFile(util.Pathf("~/ypm/ypm.lua")); err != nil {
		panic(err)
	}
	vm.loadRandom()
	loadPath(vm)
	loadTime(vm)
	loadSync(vm)
	loadJSON(vm)
	loadPsutil(vm)
	vm.Eval(`Import({"cushion-check", "cushion-vm"})`)
}

func (vm *YockScheduler) setGlobalFns(funcs ...runtime.LuaFuncs) {
	for _, fn := range funcs {
		vm.SetGlobalFn(fn)
	}
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

func (vm *YockScheduler) RunJob(name string) {
	var flags *lua.LTable
	if vm.opt != nil {
		if tmp, ok := vm.opt.RawGetString("flags").(*lua.LTable); ok {
			flags = tmp
		}
	}
	for _, job := range vm.jobs[name] {
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
			panic(err)
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

type CompileOpt struct {
	DisableAnalyse bool
}

func (vm *YockScheduler) Compile(opt CompileOpt, file string) *lua.FunctionProto {
	fp, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	chunk, err := parse.Parse(reader, file)
	if err != nil {
		panic(err)
	}
	// it's depreated
	if false && !opt.DisableAnalyse {
		anlyzer := parser.NewLuaDependencyAnalyzer()
		out, err := utils.ReadStraemFromFile(util.Pathf("@/sdk/yock/deps/stdlib.json"))
		if err != nil {
			panic(err)
		}
		if err = json.Unmarshal(out, anlyzer); err != nil {
			panic(err)
		}
		files, err := os.ReadDir(util.Pathf("~/lib"))
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			if fn := file.Name(); filepath.Ext(fn) == ".lua" {
				anlyzer.Load(util.Pathf("~/lib/") + fn)
			}
		}
		undefines, _ := anlyzer.Tidy(file)
		for _, undefine := range undefines {
			undefine = strings.TrimSuffix(undefine, "()")
			vm.Eval(fmt.Sprintf(`%s = uninit_driver("%s")`, undefine, undefine))
		}
	}
	proto, err := lua.Compile(chunk, file)
	if err != nil {
		panic(err)
	}
	return proto
}

func (vm *YockScheduler) DoCompliedFile(proto *lua.FunctionProto) error {
	lvm := vm.Interp()
	lfunc := lvm.NewFunctionFromProto(proto)
	lvm.Push(lfunc)
	return lvm.PCall(0, lua.MultRet, nil)
}
