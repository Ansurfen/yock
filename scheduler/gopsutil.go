package scheduler

import (
	"time"

	"github.com/shirou/gopsutil/v3/net"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	lua "github.com/yuin/gopher-lua"
)

func loadPsutil(vm *YockScheduler) {
	psutilCpu := &lua.LTable{}
	physicalCnt, err := cpu.Counts(false)
	if err != nil {
		physicalCnt = 0
	}
	psutilCpu.RawSetString("physics_core", lua.LNumber(physicalCnt))
	logicalCnt, err := cpu.Counts(true)
	if err != nil {
		logicalCnt = 0
	}
	psutilCpu.RawSetString("logical_core", lua.LNumber(logicalCnt))
	psutilCpu.RawSetString("percent", vm.Interp().NewClosure(func(l *lua.LState) int {
		per, err := cpu.Percent(time.Duration(l.CheckInt64(1)), l.CheckBool(2))
		ptbl := &lua.LTable{}
		for i := 0; i < len(per); i++ {
			ptbl.Insert(i+1, lua.LNumber(per[i]))
		}
		l.Push(ptbl)
		handleErr(l, err)
		return 2
	}))
	psutilCpu.RawSetString("times", vm.Interp().NewClosure(func(l *lua.LState) int {
		stats, err := cpu.Times(l.CheckBool(1))
		pstat := &lua.LTable{}
		for idx, stat := range stats {
			if info, err := Decode(l, []byte(stat.String())); err == nil {
				pstat.Insert(idx+1, info)
			} else {
				pstat.Insert(idx+1, &lua.LTable{})
			}
		}
		l.Push(pstat)
		handleErr(l, err)
		return 2
	}))
	if infos, err := cpu.Info(); err == nil {
		if len(infos) > 0 {
			if info, err := Decode(vm.Interp(), []byte(infos[0].String())); err == nil {
				psutilCpu.RawSetString("info", info)
			} else {
				psutilCpu.RawSetString("info", &lua.LTable{})
			}
		}
	} else {
		psutilCpu.RawSetString("info", &lua.LTable{})
	}
	cpu.Times(false)
	psutilMem := &lua.LTable{}
	psutilDisk := &lua.LTable{}
	psutilDisk.RawSetString("info", vm.Interp().NewClosure(func(l *lua.LState) int {
		names := []string{}
		for i := 1; i <= l.GetTop(); i++ {
			names = append(names, l.CheckString(i))
		}
		info := &lua.LTable{}
		stats, err := disk.IOCounters(names...)
		if err != nil {
			l.Push(info)
			l.Push(lua.LString(err.Error()))
			return 2
		}
		for name, stat := range stats {
			if s, err := Decode(l, []byte(stat.String())); err == nil {
				info.RawSetString(name, s)
			}
		}
		l.Push(info)
		l.Push(lua.LNil)
		return 2
	}))
	psutilDisk.RawSetString("partitions", vm.Interp().NewClosure(func(l *lua.LState) int {
		info := &lua.LTable{}
		stats, err := disk.Partitions(l.CheckBool(1))
		if err != nil {
			l.Push(info)
			l.Push(lua.LString(err.Error()))
			return 2
		}
		for idx, stat := range stats {
			if s, err := Decode(l, []byte(stat.String())); err == nil {
				info.Insert(idx+1, s)
			}
		}
		l.Push(info)
		l.Push(lua.LNil)
		return 2
	}))
	psutilDisk.RawSetString("usage", vm.Interp().NewClosure(func(l *lua.LState) int {
		stats, err := disk.Usage(l.CheckString(1))
		if err != nil {
			l.Push(&lua.LTable{})
			l.Push(lua.LString(err.Error()))
			return 2
		}
		if v, err := Decode(l, []byte(stats.String())); err == nil {
			l.Push(v)
			l.Push(lua.LNil)
		} else {
			l.Push(&lua.LTable{})
			l.Push(lua.LString(err.Error()))
		}
		return 2
	}))
	psutilMem.RawSetString("swap", vm.Interp().NewClosure(func(l *lua.LState) int {
		stats, err := mem.SwapMemory()
		if err != nil {
			l.Push(&lua.LTable{})
			l.Push(lua.LString(err.Error()))
			return 2
		}
		if v, err := Decode(l, []byte(stats.String())); err == nil {
			l.Push(v)
			l.Push(lua.LNil)
		} else {
			l.Push(&lua.LTable{})
			l.Push(lua.LString(err.Error()))
		}
		return 2
	}))
	psutilMem.RawSetString("info", vm.Interp().NewClosure(func(l *lua.LState) int {
		stats, err := mem.VirtualMemory()
		if err != nil {
			l.Push(&lua.LTable{})
			l.Push(lua.LString(err.Error()))
			return 2
		}
		if v, err := Decode(l, []byte(stats.String())); err == nil {
			l.Push(v)
			l.Push(lua.LNil)
		} else {
			l.Push(&lua.LTable{})
			l.Push(lua.LString(err.Error()))
		}
		return 2
	}))
	psutilHost := &lua.LTable{}
	psutilHost.RawSetString("boot_time", vm.Interp().NewClosure(func(l *lua.LState) int {
		timestamp, _ := host.BootTime()
		t := time.Unix(int64(timestamp), 0)
		l.Push(lua.LString(t.Local().Format("2006-01-02 15:04:05")))
		return 1
	}))
	psutilHost.RawSetString("info", vm.Interp().NewClosure(func(l *lua.LState) int {
		platform, family, version, err := host.PlatformInformation()
		l.Push(lua.LString(platform))
		l.Push(lua.LString(family))
		l.Push(lua.LString(version))
		handleErr(l, err)
		return 4
	}))
	psutilNet := &lua.LTable{}
	psutilNet.RawSetString("interfaces", vm.Interp().NewClosure(func(l *lua.LState) int {
		stats, err := net.Interfaces()
		if err != nil {
			l.Push(&lua.LTable{})
			l.Push(lua.LString(err.Error()))
			return 2
		}
		if v, err := Decode(l, []byte(stats.String())); err == nil {
			l.Push(v)
			l.Push(lua.LNil)
		} else {
			l.Push(&lua.LTable{})
			l.Push(lua.LString(err.Error()))
		}
		return 2
	}))
	psutilNet.RawSetString("io", vm.Interp().NewClosure(func(l *lua.LState) int {
		info := &lua.LTable{}
		stats, err := net.IOCounters(l.CheckBool(1))
		if err != nil {
			l.Push(info)
			l.Push(lua.LString(err.Error()))
			return 2
		}
		for idx, stat := range stats {
			if s, err := Decode(l, []byte(stat.String())); err == nil {
				info.Insert(idx+1, s)
			}
		}
		l.Push(info)
		l.Push(lua.LNil)
		return 2
	}))
	psutilNet.RawSetString("connections", vm.Interp().NewClosure(func(l *lua.LState) int {
		info := &lua.LTable{}
		stats, err := net.Connections(l.CheckString(1))
		if err != nil {
			l.Push(info)
			l.Push(lua.LString(err.Error()))
			return 2
		}
		for idx, stat := range stats {
			if s, err := Decode(l, []byte(stat.String())); err == nil {
				info.Insert(idx+1, s)
			}
		}
		l.Push(info)
		l.Push(lua.LNil)
		return 2
	}))
	vm.setGlobalVars(map[string]lua.LValue{
		"cpu":  psutilCpu,
		"mem":  psutilMem,
		"disk": psutilDisk,
		"host": psutilHost,
		"net":  psutilNet,
	})
}
