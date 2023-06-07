// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/ansurfen/cushion/utils"
	yockd "github.com/ansurfen/yock/daemon/client"
	yockf "github.com/ansurfen/yock/ffi"
	yockr "github.com/ansurfen/yock/runtime"
	"github.com/ansurfen/yock/util"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

// YockScheduler runs and schedules yock scripts in task units.
type YockScheduler struct {
	// Interpreter for yock scripts.
	//
	// Note that concurrency is not safe,
	// so each different task or asynchronous function will call the thread method
	// to derive a new interpreter and isolate execution.
	yockr.YockRuntime

	// env carries local environment variables for the program to run.
	// For example, working directory, executable path, passed flags, etc.
	env *yockr.Table

	// like env, opt also stores local environmental information.
	// However, opt is provided by the job_option functions declared in the script.
	// When the attributes of env and opt coincide, env prevails.
	// Their relationship is like global and local variables.
	opt *yockr.Table

	// envVar is initialized only when OptionEnableEnvVar is called.
	// Once initialized, the user can manipulate environment variables in the script.
	envVar utils.EnvVar

	// it's deprecated in lateset version
	driverManager *yockDriverManager

	// task is the smallest scheduling unit for asynchronous tasks,
	// which is consists of single or multiple jobs.
	// By default, the scheduler assigns coroutines to each task.
	task map[string][]*yockJob

	// goroutines organizes and manages asynchronous functions in scripts.
	// Due to the single-threaded setting of Lua coroutines, the advantages of multi-core CPUs cannot be exploited.
	// Yock exported the Golang's coroutines to provide Lua with true asynchronous capabilities.
	goroutines chan func()

	// signals manage the signals generated when the script runs.
	// Using the wait and notify methods provided by yock,
	// you can easily implement the synchronization relationship of asynchronous tasks.
	signals SignalStream

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
		env:         yockr.NewTable(),
		goroutines:  make(chan func(), 10),
		signals:     NewSingleSignalStream(),
		task:        make(map[string][]*yockJob),
		// daemon:      *yockd.New(&yockd.DaemonOption{}),
	}

	for _, opt := range opts {
		if err := opt(yocks); err != nil {
			util.Ycho.Fatal(err.Error())
		}
	}

	if yocks.driverManager != nil {
		if err := utils.SafeBatchMkdirs([]string{util.PluginPath, util.DriverPath}); err != nil {
			util.Ycho.Fatal(err.Error())
		}
	}

	yocks.parseFlags()
	yocks.loadLibs()

	return yocks
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
	yocks.env.RawSetString("args", args)
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
	yockLibRandom  = "random"
	yockLibSSH     = "ssh"
	yockLibTime    = "time"
	yockLibJSON    = "json"
	yockLibPath    = "path"
	yockLibRegexp  = "regexp"
	yockLibSync    = "sync"
	yockLibWatch   = "watch"
	yockLibStrings = "strings"
	yockLibYcho    = "ycho"
	yockLibFFI     = "ffi"
)

type yockLib struct {
	name   string
	handle func(*YockScheduler) lua.LValue
}

type luaFuncs map[string]lua.LGFunction

var yockLibs = []yockLib{
	{yockLibRandom, loadRandom},
	{yockLibJSON, loadJSON},
	{yockLibTime, loadTime},
	{yockLibPath, loadPath},
	{yockLibRegexp, loadRegexp},
	{yockLibSync, loadSync},
	{yockLibWatch, loadWatch},
	{yockLibStrings, loadStrings},
	{yockLibYcho, func(yocks *YockScheduler) lua.LValue {
		return luar.New(yocks.State(), util.Ycho)
	}},
	{yockLibFFI, func(ys *YockScheduler) lua.LValue {
		return yockf.LoadFFI(ys.State())
	}},
}

type yockFunc func(*YockScheduler) luaFuncs

func wrapLuaFuns(fs luaFuncs) yockFunc {
	return func(ys *YockScheduler) luaFuncs {
		return fs
	}
}

var yockFuncs = []yockFunc{
	loadEnv,
	netFuncs,
	goroutineFuncs,
	loadPlugin,
	loadDriver,
	taskFuncs,
	loadType,
	loadXML,
	sshFuncs,
	loadCheck,
	wrapLuaFuns(tmplFuncs),
	wrapLuaFuns(osFuncs),
	wrapLuaFuns(gnuFuncs),
	wrapLuaFuns(ioFuncs),
}

// loadLibs loads the libraries that go provides to Lua
func (yocks *YockScheduler) loadLibs() {
	for _, fn := range yockFuncs {
		yocks.SetGlobalFn(fn(yocks))
	}

	var yockGlobalVars = map[string]lua.LValue{
		"env": yocks.env.LTable,
	}

	if yocks.driverManager != nil {
		yockGlobalVars["plugins"] = yocks.driverManager.plugins
		yockGlobalVars["ldns"] = luar.New(yocks.State(), yocks.driverManager.localDNS)
		yockGlobalVars["gdns"] = luar.New(yocks.State(), yocks.driverManager.globalDNS)
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

// LaunchTask executes the corresponding task based on the task name
func (yocks *YockScheduler) LaunchTask(name string) {
	var flags *lua.LTable
	if yocks.opt != nil {
		if tmp, ok := yocks.opt.GetTable("flags"); ok {
			flags = tmp
		}
	}
	for _, job := range yocks.task[name] {
		tmp, cancel := yocks.State().NewThread()
		tbl := yocks.env.Clone(tmp)
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

// EventLoop periodically takes fn from goroutines and assigns goroutine execution
//
// In a future version, event loop will be based on the AST syntax tree or not
func (yocks *YockScheduler) EventLoop() {
	for {
		select {
		case fn := <-yocks.goroutines:
			go fn()
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

// handleErr returns the appropriate value depending on whether the error exists or not.
// Exists, returns error's text information, otherwise returns null.
//
// @return string|nil
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

// registerLib creates an empty table, injects functions into the table, and return the pointer to the table.
func (yocks *YockScheduler) registerLib(funcs luaFuncs) *lua.LTable {
	lib := &lua.LTable{}
	ls := yocks.State()
	for name, fn := range funcs {
		lib.RawSetString(name, ls.NewClosure(fn))
	}
	return lib
}

// mountLib mounts functions to the specified table.
func (yocks *YockScheduler) mountLib(lib *lua.LTable, funcs luaFuncs) *lua.LTable {
	ls := yocks.State()
	for name, fn := range funcs {
		lib.RawSetString(name, ls.NewClosure(fn))
	}
	return lib
}

// stacktrace returns the stack info of function, in form of file:line
func stacktrace(l *lua.LState) string {
	dgb, ok := l.GetStack(1)
	if ok {
		l.GetInfo("S", dgb, &lua.LFunction{})
		l.GetInfo("l", dgb, &lua.LFunction{})
		return fmt.Sprintf("%s:%d\t", dgb.Source, dgb.CurrentLine)
	}
	return ""
}
