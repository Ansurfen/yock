// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"time"

	"github.com/ansurfen/yock/util"

	lua "github.com/yuin/gopher-lua"
)

func loadWatch(yocks *YockScheduler) lua.LValue {
	cpu := yocks.registerLib(luaFuncs{
		"percent": watchCpuPercent,
		"times":   watchCpuTimes,
	})
	cpu.RawSetString("physics_core", lua.LNumber(util.CPU().PhysicalCore))
	cpu.RawSetString("logical_core", lua.LNumber(util.CPU().LogicalCore))

	var cpuInfo lua.LValue
	if infos, err := util.CPU().Info(); err == nil {
		if len(infos) > 0 {
			if info, err := Decode(yocks.Interp(), []byte(infos[0].String())); err == nil {
				cpuInfo = info
			}
		}
	}
	if cpuInfo == nil {
		cpuInfo = &lua.LTable{}
	}
	cpu.RawSetString("info", cpuInfo)

	mem := yocks.registerLib(luaFuncs{
		"info": watchMemVirtualMemory,
		"swap": watchMemSwapMemory,
	})

	disk := yocks.registerLib(luaFuncs{
		"info":       watchDiskInfo,
		"partitions": watchDiskPartitions,
		"usage":      watchDiskUsage,
	})

	host := yocks.registerLib(luaFuncs{
		"info":      watchHostInfo,
		"boot_time": watchHostBootTime,
	})

	net := yocks.registerLib(luaFuncs{
		"interfaces":  watchNetInterfaces,
		"io":          watchNetIO,
		"connections": watchNetConnections,
	})

	yocks.setGlobalVars(map[string]lua.LValue{
		"cpu":  cpu,
		"mem":  mem,
		"disk": disk,
		"host": host,
		"net":  net,
	})
	return nil
}

/*
* @param interval number
* @param percpu bool
* @retrun table, error
 */
func watchCpuPercent(l *lua.LState) int {
	per, err := util.CPU().Percent(time.Duration(l.CheckInt64(1)), l.CheckBool(2))
	ptbl := &lua.LTable{}
	for i := 0; i < len(per); i++ {
		ptbl.Insert(i+1, lua.LNumber(per[i]))
	}
	l.Push(ptbl)
	handleErr(l, err)
	return 2
}

/*
* @param percpu bool
* @retrun table, error
 */
func watchCpuTimes(l *lua.LState) int {
	stats, err := util.CPU().Times(l.CheckBool(1))
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
func watchHostBootTime(l *lua.LState) int {
	timestamp, _ := util.Host().BootTime()
	t := time.Unix(int64(timestamp), 0)
	l.Push(lua.LString(t.Local().Format("2006-01-02 15:04:05")))
	return 1
}

// @return string, string, string, err
func watchHostInfo(l *lua.LState) int {
	platform, family, version, err := util.Host().PlatformInformation()
	l.Push(lua.LString(platform))
	l.Push(lua.LString(family))
	l.Push(lua.LString(version))
	handleErr(l, err)
	return 4
}

// @return table, err
func watchNetInterfaces(l *lua.LState) int {
	stats, err := util.Net().Interfaces()
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
* @param pernic bool
* @return table, err
*/
func watchNetIO(l *lua.LState) int {
	info := &lua.LTable{}
	stats, err := util.Net().IOCounters(l.CheckBool(1))
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
* @param kind string
* @return table, err
*/
func watchNetConnections(l *lua.LState) int {
	info := &lua.LTable{}
	stats, err := util.Net().Connections(l.CheckString(1))
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