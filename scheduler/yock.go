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
	yocks := &YockScheduler{
		VirtualMachine: runtime.NewVirtualMachine(),
		env:            &lua.LTable{},
		goroutines:     make(chan func(), 10),
		signals:        NewSingleSignalStream(),
		task:           make(map[string][]*yockJob),
	}

	for _, opt := range opts {
		if err := opt(yocks); err != nil {
			util.Ycho.Fatal(err.Error())
		}
	}

	if err := utils.SafeBatchMkdirs([]string{util.PluginPath, util.DriverPath}); err != nil {
		util.Ycho.Fatal(err.Error())
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
	yockLibPsutil  = "psutil"
	yockLibStrings = "strings"
)

type yockLib struct {
	name   string
	handle func(*YockScheduler) lua.LValue
}

type luaFuncs map[string]lua.LGFunction

var yockLibs = []yockLib{
	{yockLibRandom, loadRandom},
	{yockLibSSH, loadSSH},
	{yockLibJSON, loadJSON},
	{yockLibTime, loadTime},
	{yockLibPath, loadPath},
	{yockLibRegexp, loadRegexp},
	{yockLibSync, loadSync},
	{yockLibPsutil, loadPsutil},
	{yockLibStrings, loadStrings},
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
	wrapLuaFuns(osFuncs),
	wrapLuaFuns(gnuFuncs),
	wrapLuaFuns(ioFuncs),
}

func (yocks *YockScheduler) loadLibs() {
	for _, fn := range yockFuncs {
		yocks.SetGlobalFn(runtime.LuaFuncs(fn(yocks)))
	}

	var yockGlobalVars = map[string]lua.LValue{
		"env": yocks.env,
	}

	if yocks.driverManager != nil {
		yockGlobalVars["plugins"] = yocks.driverManager.plugins
		yockGlobalVars["ldns"] = luar.New(yocks.Interp(), yocks.driverManager.localDNS)
		yockGlobalVars["gdns"] = luar.New(yocks.Interp(), yocks.driverManager.globalDNS)
	}
	yocks.setGlobalVars(yockGlobalVars)

	files, err := os.ReadDir(util.Pathf("~/lib"))
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
	for _, file := range files {
		if fn := file.Name(); filepath.Ext(fn) == ".lua" {
			if err := yocks.EvalFile(util.Pathf("~/lib/") + fn); err != nil {
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
			if err := yocks.EvalFile(util.Pathf("~/lib/boot/") + fn); err != nil {
				util.Ycho.Fatal(err.Error())
			}
		}
	}

	for _, lib := range yockLibs {
		yocks.SetGlobalVar(lib.name, lib.handle(yocks))
	}

	yocks.Interp().PreloadModule("check", runtime.LoadCheck)
}

func (yocks *YockScheduler) setGlobalVars(vars map[string]lua.LValue) {
	for k, v := range vars {
		yocks.SetGlobalVar(k, v)
	}
}

func (yocks *YockScheduler) LaunchTask(name string) {
	var flags *lua.LTable
	if yocks.opt != nil {
		if tmp, ok := yocks.opt.RawGetString("flags").(*lua.LTable); ok {
			flags = tmp
		}
	}
	for _, job := range yocks.task[name] {
		tmp, cancel := yocks.Interp().NewThread()
		tbl := tableDeepCopy(tmp, yocks.env)
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

func (yocks *YockScheduler) registerLib(funcs luaFuncs) *lua.LTable {
	lib := &lua.LTable{}
	ls := yocks.Interp()
	for name, fn := range funcs {
		lib.RawSetString(name, ls.NewClosure(fn))
	}
	return lib
}

func (yocks *YockScheduler) mountLib(lib *lua.LTable, funcs luaFuncs) *lua.LTable {
	ls := yocks.Interp()
	for name, fn := range funcs {
		lib.RawSetString(name, ls.NewClosure(fn))
	}
	return lib
}
