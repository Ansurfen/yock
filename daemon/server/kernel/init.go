// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kernel

import (
	"github.com/ansurfen/yock/daemon/server/fs"
	"github.com/ansurfen/yock/daemon/server/mem"
	"github.com/ansurfen/yock/daemon/server/net"
	"github.com/ansurfen/yock/daemon/server/process"
	"github.com/ansurfen/yock/daemon/server/user"
)

type YockKernel struct {
	*fs.FileSystem
	*mem.MemWatch
	*ynet.NetworkManager
	*user.UserGroup
	*process.Scheduler
	*SingalStream
}

func NewKernel() *YockKernel {
	return &YockKernel{
		FileSystem:   fs.NewFileSystem(),
		SingalStream: newSingalStream(),
	}
}

func (yockk *YockKernel) Setup() {
	p := process.New(func() error {
		return yockk.MemWatch.GabargeCollect()
	})
	yockk.Scheduler.Put(p.UpgradeCronTask("0 0 * * * ?"))
}

func (yockk *YockKernel) Main() {
	yockk.Scheduler.Run()
}
