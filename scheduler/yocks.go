// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocks

import (
	"os"
	"path"
	"path/filepath"

	"github.com/ansurfen/yock/ctl/conf"
	"github.com/ansurfen/yock/daemon/net"
	yocke "github.com/ansurfen/yock/env"
	yocki "github.com/ansurfen/yock/interface"
	yockr "github.com/ansurfen/yock/runtime"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

var _ yocki.YockScheduler = (*YockScheduler)(nil)

type loader func(yocki.YockScheduler)

// YockScheduler runs and schedules yock scripts in task units.
type YockScheduler struct {
	// Interpreter for yock scripts.
	//
	// Note that concurrency is not safe,
	// so each different task or asynchronous function will call the thread method
	// to derive a new interpreter and isolate execution.
	yocki.YockRuntime

	*yockLoader

	// env carries local environment variables for the program to run.
	// For example, working directory, executable path, passed flags, etc.
	env yocki.YockLib

	// like env, opt also stores local environmental information.
	// However, opt is provided by the job_option functions declared in the script.
	// When the attributes of env and opt coincide, env prevails.
	// Their relationship is like global and local variables.
	opt yocki.Table

	// envVar is initialized only when OptionEnableEnvVar is called.
	// Once initialized, the user can manipulate environment variables in the script.
	envVar yocki.EnvVar

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
	yocki *yockInterfaces

	// daemon manages and schedules Yock's background tasks.
	// yockd on each computer can be regarded as a node, and
	// different nodes can form clusters to complete parallel build, deployment and etc.
	daemon map[string]yocki.YockdClient

	libPath string

	*yocksDB
}

func New(opts ...YockSchedulerOption) *YockScheduler {
	yocks := &YockScheduler{
		YockRuntime: yockr.New(),
		signals:     NewSingleSignalStream(),
		task:        make(map[string][]*yockJob),
		goroutines:  newChannelPool(10),
		yocksDB:     newYocksDB(),
		yocki:       newYockInterface(),
		daemon:      make(map[string]yocki.YockdClient),
	}

	for _, opt := range opts {
		if err := opt(yocks); err != nil {
			ycho.Fatal(err)
		}
	}

	if yocks.driverManager != nil {
		if err := util.SafeBatchMkdirs([]string{util.PluginPath, util.DriverPath}); err != nil {
			ycho.Fatal(err)
		}
	}

	yocks.yockLoader = NewYockLoader(yocks.State())
	yocks.env = yocks.CreateLib("env")

	yocks.parseFlags()

	return yocks
}

func Default(opts ...YockSchedulerOption) *YockScheduler {
	yocks := New(opts...)
	yocks.LoadLibs()
	yocks.LoadYocki()
	yocks.LoadYockd()
	return yocks
}

func (yocks *YockScheduler) EnvVar() yocki.EnvVar {
	return yocks.envVar
}

func (yocks *YockScheduler) State() yocki.YockState {
	return yocks.YockRuntime.State()
}

func wrapFunc(fn yocki.YocksFunction, yocks yocki.YockScheduler) lua.LGFunction {
	return func(l *lua.LState) int {
		return fn(yocks, yockr.UpgradeLState(l))
	}
}

func (yocks *YockScheduler) RegYocksFn(funcs yocki.YocksFuncs) {
	for name, fn := range funcs {
		yocks.State().LState().SetGlobal(name,
			yocks.State().LState().NewFunction(wrapFunc(fn, yocks)),
		)
	}
}

func (yocks *YockScheduler) MntYocksFn(lib yocki.YockLib, funcs yocki.YocksFuncs) {
	for name, fn := range funcs {
		lib.SetFunction(name, wrapFunc(fn, yocks))
	}
}

func (yocks *YockScheduler) Signal() yocki.SignalStream {
	return yocks.signals
}

func (yocks *YockScheduler) getPlugins() yocki.Table {
	if yocks.driverManager != nil {
		return yocks.driverManager.plugins
	}
	return &yockr.Table{}
}

func (yocks *YockScheduler) getDrivers() yocki.Table {
	if yocks.driverManager != nil {
		return yocks.driverManager.drivers
	}
	return &yockr.Table{}
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
	yocks.task[name] = append(yocks.task[name], &yockJob{fn: job.Func(), name: job.Name()})
}

type YGFunction func(*YockScheduler, *yockr.YockState) int

func (yocks *YockScheduler) Do(f func()) {
	yocks.goroutines.Go(f)
}

// deprecated
func (yocks *YockScheduler) LoadLibsV1() {
	loadDriver(yocks)
	loadPlugin(yocks)
}

func (yocks *YockScheduler) LoadYocki() {
	lib := yocks.CreateLib("yocki")
	lib.SetField(map[string]any{
		"connect": func(name, ip string, port int) {
			yocks.yocki.Connect(name, ip, port)
		},
		"close": func(name string) {
			yocks.yocki.Close(name)
		},
		"call": func(name, fn, arg string) (string, string) {
			res, err := yocks.yocki.Call(name, fn, arg)
			if err != nil {
				return res, err.Error()
			}
			return res, ""
		},
		"list": func() *lua.LTable {
			tbl := &lua.LTable{}
			for name := range yocks.yocki.clients {
				tbl.Append(lua.LString(name))
			}
			return tbl
		},
	})
}

// LoadLibs loads the libraries that go provides to Lua
func (yocks *YockScheduler) LoadLibs() {
	for _, load := range libgo {
		load(yocks)
	}

	for _, load := range libyock {
		load(yocks)
	}

	var yockGlobalVars = map[string]lua.LValue{
		"env": yocks.env.Meta().Value(),
	}

	if yocks.driverManager != nil {
		yockGlobalVars["plugins"] = yocks.driverManager.plugins.Value()
		yockGlobalVars["ldns"] = luar.New(yocks.State().LState(), yocks.driverManager.localDNS)
		yockGlobalVars["gdns"] = luar.New(yocks.State().LState(), yocks.driverManager.globalDNS)
	}
	yocks.setGlobalVars(yockGlobalVars)

	files, err := os.ReadDir(yocks.libPath)
	if err != nil {
		ycho.Fatal(err)
	}
	for _, file := range files {
		if fn := file.Name(); filepath.Ext(fn) == ".lua" {
			if err := yocks.EvalFile(path.Join(yocks.libPath, fn)); err != nil {
				ycho.Fatal(err)
			}
		}
	}

	// self-boot
	if util.YockBuild != "dev" {
		boot_path := util.Pathf("~/lib/boot")
		files, err = os.ReadDir(boot_path)
		if err != nil {
			ycho.Fatal(err)
		}
		for _, file := range files {
			if fn := file.Name(); filepath.Ext(fn) == ".lua" {
				if err := yocks.EvalFile(path.Join(boot_path, fn)); err != nil {
					ycho.Fatal(err)
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

func (yocks *YockScheduler) defaultYockd() yocki.YockdClient {
	if d, ok := yocks.daemon["default"]; !ok {
		conf := yocke.GetEnv[*conf.YockConf]().Conf()
		yocks.daemon["default"] = net.NewDirect(&net.YockdClientOption{
			IP:   conf.Yockd.IP,
			Port: conf.Yockd.Port,
		})
		return yocks.daemon["default"]
	} else {
		return d
	}
}

func (yocks *YockScheduler) Opt() yocki.Table {
	return yocks.opt
}

func (yocks *YockScheduler) SetOpt(o yocki.Table) {
	yocks.opt = o
}

func (yocks *YockScheduler) Env() yocki.YockLib {
	return yocks.env
}

// LaunchTask executes the corresponding task based on the task name
func (yocks *YockScheduler) LaunchTask(name string) {
	defer func() {
		msg := recover()
		switch msg.(type) {
		case error:

		}
	}()
	var flags yocki.Table
	if yocks.opt != nil {
		if tmp, ok := yocks.opt.GetTable("flags"); ok {
			flags = tmp
		}
	}
	var (
		inherit bool
		super   *Context
	)
	for _, job := range yocks.task[name] {
		ctx := newContext(name, job, flags, yocks)
		defer ctx.Close()
		if inherit {
			ctx.Extends(super)
			inherit = false
			super = nil
		}
		code := ctx.Call(job.fn)
		ycho.Infof("[%s] exit, %s", ctx.source, code)
		switch code {
		case 0:
			return
		case 1:
			// continue
		case 2:
			inherit = true
			super = ctx
		default:
			ycho.Warnf("unkown status for context")
		}
	}
	ycho.Infof("[%s] exit", name)
}

// EventLoop periodically takes fn from goroutines and assigns goroutine execution
//
// In a future version, event loop will be based on the AST syntax tree or not
func (yocks *YockScheduler) EventLoop() {
	yocks.goroutines.Run()
}
