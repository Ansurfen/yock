package scheduler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/yuin/gopher-lua/parse"

	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/parser"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

var (
	WorkSpace  string
	PluginPath string
	DriverPath string
)

func init() {
	WorkSpace = filepath.ToSlash(path.Join(utils.GetEnv().Workdir(), ".yock"))
	PluginPath = path.Join(WorkSpace, "plugin")
	DriverPath = path.Join(WorkSpace, "driver")
}

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
	signals    *sync.Map
}

func New() *YockScheduler {
	vm := &YockScheduler{
		VirtualMachine: runtime.NewVirtualMachine().Default(),
		env:            &lua.LTable{},
		drivers:        &lua.LTable{},
		plugins:        &lua.LTable{},
		goroutines:     make(chan func(), 10),
		signals:        &sync.Map{},
		jobs:           make(map[string][]*yockJob),
		envVar:         utils.NewEnvVar(),
	}

	utils.SafeBatchMkdirs([]string{PluginPath, DriverPath})

	vm.globalDNS = CreateDNS(Pathf("@/global.json"))
	vm.localDNS = CreateDNS(Pathf("@/local.json"))

	vm.parseFlags()
	vm.injectGlobal()

	return vm
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
	vm.env.RawSetString("workdir", lua.LString(WorkSpace))
	vm.env.RawSetString("set_path", vm.Interp().NewClosure(func(l *lua.LState) int {
		err := vm.envVar.SetPath(l.CheckString(1))
		if err != nil {
			l.Push(lua.LString(err.Error()))
		} else {
			l.Push(lua.LString(""))
		}
		return 1
	}))
	vm.env.RawSetString("safe_set", vm.Interp().NewClosure(func(l *lua.LState) int {
		err := vm.envVar.SafeSet(l.CheckString(1), envVarTypeCvt(l.CheckAny(2)))
		if err != nil {
			l.Push(lua.LString(err.Error()))
		} else {
			l.Push(lua.LString(""))
		}
		return 1
	}))
	vm.env.RawSetString("set", vm.Interp().NewClosure(func(l *lua.LState) int {
		err := vm.envVar.Set(l.CheckString(1), envVarTypeCvt(l.CheckAny(2)))
		if err != nil {
			l.Push(lua.LString(err.Error()))
		} else {
			l.Push(lua.LString(""))
		}
		return 1
	}))
	vm.env.RawSetString("unset", vm.Interp().NewClosure(func(l *lua.LState) int {
		err := vm.envVar.Unset(l.CheckString(1))
		if err != nil {
			l.Push(lua.LString(err.Error()))
		} else {
			l.Push(lua.LString(""))
		}
		return 1
	}))
	vm.env.RawSetString("setl", vm.Interp().NewClosure(func(l *lua.LState) int {
		err := vm.envVar.SetL(l.CheckString(1), l.CheckString(2))
		if err != nil {
			l.Push(lua.LString(err.Error()))
		} else {
			l.Push(lua.LString(""))
		}
		return 1
	}))
	vm.env.RawSetString("safe_setl", vm.Interp().NewClosure(func(l *lua.LState) int {
		err := vm.envVar.SafeSetL(l.CheckString(1), l.CheckString(2))
		if err != nil {
			l.Push(lua.LString(err.Error()))
		} else {
			l.Push(lua.LString(""))
		}
		return 1
	}))
	vm.env.RawSetString("export", vm.Interp().NewClosure(func(l *lua.LState) int {
		err := vm.envVar.Export(l.CheckString(1))
		if err != nil {
			l.Push(lua.LString(err.Error()))
		} else {
			l.Push(lua.LString(""))
		}
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
		loadJob(vm))
	vm.setGlobalVars(map[string]lua.LValue{
		"ldns":    luar.New(vm.Interp(), vm.localDNS),
		"gdns":    luar.New(vm.Interp(), vm.globalDNS),
		"env":     vm.env,
		"plugins": vm.plugins,
	})
	files, err := os.ReadDir(Pathf("@/sdk/yock"))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.Name() != "yock.lua" {
			vm.EvalFile(Pathf("@/sdk/yock/") + file.Name())
		}
	}
	vm.loadRandom()
	loadPath(vm)
	loadTime(vm)
	loadSync(vm)
	loadJSON(vm)
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
	if !opt.DisableAnalyse {
		anlyzer := parser.NewLuaDependencyAnalyzer()
		out, err := utils.ReadStraemFromFile(Pathf("@/sdk/yock/deps/stdlib.json"))
		if err != nil {
			panic(err)
		}
		if err = json.Unmarshal(out, anlyzer); err != nil {
			panic(err)
		}
		files, err := os.ReadDir(Pathf("@/sdk/yock"))
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			if filepath.Ext(file.Name()) == ".lua" {
				anlyzer.Load(Pathf("@/sdk/yock/") + file.Name())
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
