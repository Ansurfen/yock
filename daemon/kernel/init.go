// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kernel

import (
	"github.com/ansurfen/yock/daemon/fs"
	"github.com/ansurfen/yock/daemon/mem"
	"github.com/ansurfen/yock/daemon/net"
	"github.com/ansurfen/yock/daemon/process"
	"github.com/ansurfen/yock/daemon/user"
	"github.com/ansurfen/yock/ycho"
)

type YockKernel struct {
	*fs.FileSystem
	*mem.MemWatch
	*net.NetworkManager
	*user.UserGroup
	*process.Scheduler
	*SignalStream
}

func NewKernel() *YockKernel {
	ycho.Info("kernel init")
	return &YockKernel{
		FileSystem:     fs.NewFileSystem(),
		SignalStream:   newSingalStream(),
		NetworkManager: net.NewNetworkManager(),
		Scheduler:      process.NewScheduler(),
	}
}

func (k *YockKernel) Init() {
	k.Scheduler.Run()
}
