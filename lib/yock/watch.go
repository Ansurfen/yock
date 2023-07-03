// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	"time"

	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util"
	lua "github.com/yuin/gopher-lua"
)

func LoadWatch(yocks yocki.YockScheduler) {
	cpu := yocks.CreateLib("cpu")
	cpu.SetYFunction(map[string]yocki.YGFunction{
		"percent": watchCpuPercent,
		"times":   watchCpuTimes,
	})
	cpu.SetField(map[string]any{
		"physics_core": util.CPU().PhysicalCore,
		"logical_core": util.CPU().LogicalCore,
	})

	var cpuInfo lua.LValue
	if infos, err := util.CPU().Info(); err == nil {
		if len(infos) > 0 {
			if info, err := Decode(yocks.State().LState(), []byte(infos[0].String())); err == nil {
				cpuInfo = info
			}
		}
	}
	if cpuInfo == nil {
		cpuInfo = &lua.LTable{}
	}
	cpu.Meta().Value().RawSetString("info", cpuInfo)

	yocks.CreateLib("mem").SetFunctions(map[string]lua.LGFunction{
		"info": watchMemVirtualMemory,
		"swap": watchMemSwapMemory,
	})

	yocks.CreateLib("disk").SetFunctions(map[string]lua.LGFunction{
		"info":       watchDiskInfo,
		"partitions": watchDiskPartitions,
		"usage":      watchDiskUsage,
	})

	yocks.CreateLib("host").SetYFunction(map[string]yocki.YGFunction{
		"info":      watchHostInfo,
		"boot_time": watchHostBootTime,
	})

	yocks.CreateLib("net").SetYFunction(map[string]yocki.YGFunction{
		"interfaces":  watchNetInterfaces,
		"io":          watchNetIO,
		"connections": watchNetConnections,
	})
}

/*
* @param interval number
* @param percpu bool
* @retrun table, error
 */
func watchCpuPercent(l yocki.YockState) int {
	per, err := util.CPU().Percent(time.Duration(l.LState().CheckInt64(1)), l.LState().CheckBool(2))
	ptbl := &lua.LTable{}
	for i := 0; i < len(per); i++ {
		ptbl.Insert(i+1, lua.LNumber(per[i]))
	}
	l.Push(ptbl).PushError(err)
	return 2
}

/*
* @param percpu bool
* @retrun table, error
 */
func watchCpuTimes(l yocki.YockState) int {
	stats, err := util.CPU().Times(l.LState().CheckBool(1))
	pstat := &lua.LTable{}
	for idx, stat := range stats {
		if info, err := Decode(l.LState(), []byte(stat.String())); err == nil {
			pstat.Insert(idx+1, info)
		} else {
			pstat.Insert(idx+1, &lua.LTable{})
		}
	}
	l.Push(pstat).PushError(err)
	return 2
}

// @retrun table, error
func watchMemSwapMemory(l *lua.LState) int {
	stats, err := util.Mem().SwapMemory()
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
}

// @retrun table, error
func watchMemVirtualMemory(l *lua.LState) int {
	stats, err := util.Mem().VirtualMemory()
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
}

/*
* @param names ...string
* @retrun table, error
 */
func watchDiskInfo(l *lua.LState) int {
	names := []string{}
	for i := 1; i <= l.GetTop(); i++ {
		names = append(names, l.CheckString(i))
	}
	info := &lua.LTable{}
	stats, err := util.Disk().IOCounters(names...)
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
}

/*
* @param all bool
* @retrun table, error
 */
func watchDiskPartitions(l *lua.LState) int {
	info := &lua.LTable{}
	stats, err := util.Disk().Partitions(l.CheckBool(1))
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
}

/*
* @param path string
* @retrun table, error
 */
func watchDiskUsage(l *lua.LState) int {
	stats, err := util.Disk().Usage(l.CheckString(1))
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
}

// @retrun string
func watchHostBootTime(l yocki.YockState) int {
	timestamp, _ := util.Host().BootTime()
	t := time.Unix(int64(timestamp), 0)
	l.Push(lua.LString(t.Local().Format("2006-01-02 15:04:05")))
	return 1
}

// @return string, string, string, err
func watchHostInfo(l yocki.YockState) int {
	platform, family, version, err := util.Host().PlatformInformation()
	l.PushString(platform)
	l.PushString(family)
	l.PushString(version)
	l.PushError(err)
	return 4
}

// @return table, err
func watchNetInterfaces(l yocki.YockState) int {
	stats, err := util.Net().Interfaces()
	if err != nil {
		l.PushNilTable().Throw(err)
		return 2
	}
	if v, err := Decode(l.LState(), []byte(stats.String())); err == nil {
		l.Push(v).PushNil()
	} else {
		l.PushNilTable().Throw(err)
	}
	return 2
}

/*
* @param pernic bool
* @return table, err
 */
func watchNetIO(l yocki.YockState) int {
	info := &lua.LTable{}
	stats, err := util.Net().IOCounters(l.LState().CheckBool(1))
	if err != nil {
		l.Push(info).Throw(err)
		return 2
	}
	for idx, stat := range stats {
		if s, err := Decode(l.LState(), []byte(stat.String())); err == nil {
			info.Insert(idx+1, s)
		}
	}
	l.Push(info).PushNil()
	return 2
}

/*
* @param kind string
* @return table, err
 */
func watchNetConnections(l yocki.YockState) int {
	info := &lua.LTable{}
	stats, err := util.Net().Connections(l.LState().CheckString(1))
	if err != nil {
		l.Push(info)
		l.Push(lua.LString(err.Error()))
		return 2
	}
	for idx, stat := range stats {
		if s, err := Decode(l.LState(), []byte(stat.String())); err == nil {
			info.Insert(idx+1, s)
		}
	}
	l.Push(info).PushNil()
	return 2
}
