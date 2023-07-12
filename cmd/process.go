// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

type PSOpt struct {
	User bool
	CPU  bool
	Mem  bool
	Time bool

	String bool
}

type ProcessInfo struct {
	User  string
	CPU   float64
	Cmd   string
	Mem   *process.MemoryInfoStat
	Start int64
}

func PS(opt PSOpt) (info map[int32]*ProcessInfo, err error) {
	info = make(map[int32]*ProcessInfo)
	processes, _ := process.Processes()
	for _, p := range processes {
		pid := p.Pid
		cmd, _ := p.Cmdline()
		pi := &ProcessInfo{
			Cmd: cmd,
		}
		info[pid] = pi
		if opt.CPU {
			c, _ := p.CPUPercent()
			pi.CPU = c
		}
		if opt.User {
			u, _ := p.Username()
			pi.User = u
		}
		if opt.Mem {
			m, _ := p.MemoryInfo()
			pi.Mem = m
		}
		if opt.Time {
			t, _ := p.CreateTime()
			pi.Start = t
		}
	}
	return
}

type PGrepResult struct {
	Pid  int32
	Name string
}

func PGrep(name string) []PGrepResult {
	res := []PGrepResult{}
	processes, _ := process.Processes()
	for _, p := range processes {
		pn, err := p.Name()
		if err != nil {
			continue
		}
		if strings.Contains(pn, name) {
			res = append(res, PGrepResult{Pid: p.Pid, Name: name})
		}
	}
	return res
}

func KillByPid(pid int32) error {
	proc, err := process.NewProcess(pid)
	if err != nil {
		return err
	}
	err = proc.Kill()
	if err != nil {
		return err
	}
	return nil
}

func KillByName(name string) error {
	res := PGrep(name)
	for _, r := range res {
		if err := KillByPid(r.Pid); err != nil {
			return err
		}
	}
	return nil
}
