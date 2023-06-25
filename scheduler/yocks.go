// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"os"
	"path"
	"path/filepath"

	yockd "github.com/ansurfen/yock/daemon/client"
	yockf "github.com/ansurfen/yock/ffi"
	yocki "github.com/ansurfen/yock/interface"
	yockr "github.com/ansurfen/yock/runtime"
	"github.com/ansurfen/yock/util"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

var (
	_ yocki.YockScheduler = (*YockScheduler)(nil)
)

type loader func(yocki.YockScheduler)

// YockScheduler runs and schedules yock scripts in task units.
type YockScheduler struct {
	// Interpreter for yock scripts.
	//
	// Note that concurrency is not safe,
	// so each different task or asynchronous function will call the thread method
	// to derive a new interpreter and isolate execution.
	yockr.YockRuntime

	*yockLoader

	// env carries local environment variables for the program to run.
	// For example, working directory, executable path, passed flags, etc.
	env *yockr.YockLib

	// like env, opt also stores local environmental information.
	// However, opt is provided by the job_option functions declared in the script.
	// When the attributes of env and opt coincide, env prevails.
	// Their relationship is like global and local variables.
	opt *yockr.Table

	// envVar is initialized only when OptionEnableEnvVar is called.
	// Once initialized, the user can manipulate environment variables in the script.
	envVar util.EnvVar

	// it's deprecated in lateset version
	driverManager *yockDriverManager

	// task is the smallest scheduling unit for asynchronous tasks,
	// which is consists of single or multiple jobs.
	// By default, the scheduler assigns coroutines to each task.
	task map[string][]*yockJob

	// goroutines organizes and manages asynchronous functions in scripts.
	// Due to the single-threaded setting of Lua coroutines, the advantages of multi-core CPUs cannot be exploited.
	// Yock exported the Golang's coroutines to provide Lua with true asynchronous capabilities.
	goroutines yocki.GoPool

	// signals manage the signals generated when the script runs.
	// Using the wait and notify methods provided by yock,
	// you can easily implement the synchronization relationship of asynchronous tasks.
	signals yocki.SignalStream

	// yocki is a third-party module extension interface provided by Yock.
	// It is used to enhance yock's scripts, just like yockf.
	yocki yockInterfaces

	// daemon manages and schedules Yock's background tasks.
	// yockd on each computer can be regarded as a node, and
	// different nodes can form clusters to complete parallel build, deployment and etc.
	daemon yockd.YockDaemonClient
}

func New(opts ...YockSchedulerOption) *YockScheduler {
	yocks := &YockScheduler{
		YockRuntime: yockr.New(),
		signals:     NewSingleSignalStream(),
		task:        make(map[string][]*yockJob),
		goroutines:  newChannelPool(10),
		// daemon:      *yockd.New(&yockd.DaemonOption{}),
	}

	for _, opt := range opts {
		if err := opt(yocks); err != nil {
			util.Ycho.Fatal(err.Error())
		}
	}

	if yocks.driverManager != nil {
		if err := util.SafeBatchMkdirs([]string{util.PluginPath, util.DriverPath}); err != nil {
			util.Ycho.Fatal(err.Error())
		}
	}

	yocks.yockLoader = NewYockLoader(yocks.State())
	yocks.env = yocks.CreateLib("env")

	yocks.parseFlags()
	yocks.loadLibs()

	return yocks
}

func (yocks *YockScheduler) EnvVar() util.EnvVar {
	return yocks.envVar
}

func (yocks *YockScheduler) State() *yockr.YockState {
	return yocks.YockRuntime.State()
}

func wrapFunc(fn yocki.YocksFunction, yocks yocki.YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		return fn(yocks, yockr.UpgradeLState(l))
	}
}

func (yocks *YockScheduler) RegYocksFn(funcs yocki.YocksFuncs) {
	for name, fn := range funcs {
		yocks.State().SetGlobal(name,
			yocks.State().LState.NewFunction(wrapFunc(fn, yocks)),
		)
	}
}

func (yocks *YockScheduler) MntYocksFn(lib *yockr.YockLib, funcs yocki.YocksFuncs) {
	for name, fn := range funcs {
		lib.SetFunction(name, wrapFunc(fn, yocks))
	}
}

func (yocks *YockScheduler) Signal() yocki.SignalStream {
	return yocks.signals
}

func (yocks *YockScheduler) getPlugins() *lua.LTable {
	if yocks.driverManager != nil {
		return yocks.driverManager.plugins
	}
	return &lua.LTable{}
}

func (yocks *YockScheduler) getDrivers() *lua.LTable {
	if yocks.driverManager != nil {
		return yocks.driverManager.drivers
	}
	return &lua.LTable{}
}

// parseFlags parses -- the following parameters serve as the flags of script
func (yocks *YockScheduler) parseFlags() {
	idx := 0
	args := &lua.LTable{}
	yocks.env.SetField(map[string]any{
		"args": args,
	})
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

func (yocks *YockScheduler) GetTask(name string) bool {
	if _, ok := yocks.task[name]; ok {
		return true
	}
	return false
}

func (yocks *YockScheduler) AppendTask(name string, job yocki.YockJob) {
	yocks.task[name] = append(yocks.task[name], &yockJob{job.Func()})
}

const (
	yockLibYcho = "ycho"
	yockLibFFI  = "ffi"
)

type yockLib struct {
	name   string
	handle func(*YockScheduler) lua.LValue
}

type YGFunction func(*YockScheduler, *yockr.YockState) int

type luaFuncs map[string]lua.LGFunction

var yockLibs = []yockLib{
	{yockLibYcho, func(yocks *YockScheduler) lua.LValue {
		return luar.New(yocks.State().LState, util.Ycho)
	}},
	{yockLibFFI, func(ys *YockScheduler) lua.LValue {
		return yockf.LoadFFI(ys.State().LState)
	}},
}

type yockFunc func(*YockScheduler) luaFuncs

var yockFuncs = []yockFunc{
	loadPlugin,
	loadDriver,
}

func (yocks *YockScheduler) Do(f func()) {
	yocks.goroutines.Go(f)
}

// loadLibs loads the libraries that go provides to Lua
func (yocks *YockScheduler) loadLibs() {
	for _, fn := range yockFuncs {
		yocks.SetGlobalFn(fn(yocks))
	}

	for _, load := range libgo {
		load(yocks)
	}

	for _, load := range libyock {
		load(yocks)
	}

	var yockGlobalVars = map[string]lua.LValue{
		"env": yocks.env.Meta().LTable,
	}

	if yocks.driverManager != nil {
		yockGlobalVars["plugins"] = yocks.driverManager.plugins
		yockGlobalVars["ldns"] = luar.New(yocks.State().LState, yocks.driverManager.localDNS)
		yockGlobalVars["gdns"] = luar.New(yocks.State().LState, yocks.driverManager.globalDNS)
	}
	yocks.setGlobalVars(yockGlobalVars)

	for _, lib := range yockLibs {
		yocks.SetGlobalVar(lib.name, lib.handle(yocks))
	}

	lib_path := util.Pathf("~/lib")
	files, err := os.ReadDir(lib_path)
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
	for _, file := range files {
		if fn := file.Name(); filepath.Ext(fn) == ".lua" {
			if err := yocks.EvalFile(path.Join(lib_path, fn)); err != nil {
				util.Ycho.Fatal(err.Error())
			}
		}
	}

	// self-boot
	if util.YockBuild != "dev" {
		boot_path := util.Pathf("~/lib/boot")
		files, err = os.ReadDir(boot_path)
		if err != nil {
			util.Ycho.Fatal(err.Error())
		}
		for _, file := range files {
			if fn := file.Name(); filepath.Ext(fn) == ".lua" {
				if err := yocks.EvalFile(path.Join(boot_path, fn)); err != nil {
					util.Ycho.Fatal(err.Error())
				}
			}
		}
	}
}

func (yocks *YockScheduler) setGlobalVars(vars map[string]lua.LValue) {
	for k, v := range vars {
		yocks.SetGlobalVar(k, v)
	}
}

func (yocks *YockScheduler) Opt() *yockr.Table {
	return yocks.opt
}

func (yocks *YockScheduler) SetOpt(o *yockr.Table) {
	yocks.opt = o
}

// LaunchTask executes the corresponding task based on the task name
func (yocks *YockScheduler) LaunchTask(name string) {
	var flags *lua.LTable
	if yocks.opt != nil {
		if tmp, ok := yocks.opt.GetTable("flags"); ok {
			flags = tmp
		}
	}
	for _, job := range yocks.task[name] {
		tmp, cancel := yocks.NewState()
		tbl := yocks.env.Meta().Clone(tmp.LState)
		tbl.SetString("job", name)
		if flags != nil {
			if tmp, ok := flags.RawGetString(name).(*lua.LTable); ok {
				tbl.RawSetString("flags", tmp)
			}
		}
		if err := tmp.Call(yockr.YockFuncInfo{
			Fn: job.fn,
		}, tbl.LTable); err != nil {
			util.Ycho.Warn(err.Error())
		}
		if cancel != nil {
			cancel()
		}
	}
}

// EventLoop periodically takes fn from goroutines and assigns goroutine execution
//
// In a future version, event loop will be based on the AST syntax tree or not
func (yocks *YockScheduler) EventLoop() {
	yocks.goroutines.Run()
}
