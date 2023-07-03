// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// yockw is short for yock watcher, which is used to monitor system performance.

package util

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

func init() {
	// init yock watch
	yockCpu = newCPU()
	yockMem = newMem()
	yockDisk = newDisk()
	yockHost = newHost()
	yockNet = newNet()

	center = &SSHCenter{}
}

var (
	yockCpu  *psutilCPU
	yockMem  *psutilMem
	yockDisk *psutilDisk
	yockHost *psutilHost
	yockNet  *psutilNet
)

type psutilCPU struct {
	LogicalCore  int
	PhysicalCore int
}

func newCPU() *psutilCPU {
	c := &psutilCPU{}
	var err error
	c.LogicalCore, err = cpu.Counts(true)
	if err != nil {
		panic(err)
	}
	c.PhysicalCore, err = cpu.Counts(false)
	if err != nil {
		panic(err)
	}
	return c
}

func (c *psutilCPU) Percent(interval time.Duration, percpu bool) ([]float64, error) {
	return cpu.Percent(interval, percpu)
}

func (c *psutilCPU) Times(percpu bool) ([]cpu.TimesStat, error) {
	return cpu.Times(percpu)
}

func (c *psutilCPU) Info() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

func CPU() *psutilCPU {
	return yockCpu
}

type psutilMem struct{}

func newMem() *psutilMem {
	return &psutilMem{}
}

func (m *psutilMem) SwapMemory() (*mem.SwapMemoryStat, error) {
	return mem.SwapMemory()
}

func (m *psutilMem) VirtualMemory() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

func Mem() *psutilMem {
	return yockMem
}

type psutilDisk struct{}

func newDisk() *psutilDisk {
	return &psutilDisk{}
}

func (d *psutilDisk) IOCounters(names ...string) (map[string]disk.IOCountersStat, error) {
	return disk.IOCounters(names...)
}

func (d *psutilDisk) Partitions(all bool) ([]disk.PartitionStat, error) {
	return disk.Partitions(all)
}

func (d *psutilDisk) Usage(path string) (*disk.UsageStat, error) {
	return disk.Usage(path)
}

func Disk() *psutilDisk {
	return yockDisk
}

type psutilHost struct{}

func newHost() *psutilHost {
	return &psutilHost{}
}

func (h *psutilHost) BootTime() (uint64, error) {
	return host.BootTime()
}

func (h *psutilHost) PlatformInformation() (string, string, string, error) {
	return host.PlatformInformation()
}

func Host() *psutilHost {
	return yockHost
}

type psutilNet struct{}

func newNet() *psutilNet {
	return &psutilNet{}
}

func (n *psutilNet) Interfaces() (net.InterfaceStatList, error) {
	return net.Interfaces()
}

func (n *psutilNet) IOCounters(pernic bool) ([]net.IOCountersStat, error) {
	return net.IOCounters(pernic)
}

func (n *psutilNet) Connections(kind string) ([]net.ConnectionStat, error) {
	return net.Connections(kind)
}

func Net() *psutilNet {
	return yockNet
}
