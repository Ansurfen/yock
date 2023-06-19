// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package api

import (
	yockc "github.com/ansurfen/yock/cmd"
	"github.com/ansurfen/yock/daemon/server/process"
)

func (daemon *YockDaemon) PS() {}

func (daemon *YockDaemon) Kill() {}

func (daemon *YockDaemon) Getpid() {}

func (daemon *YockDaemon) Spwan(cron, cmd string) {
	p := process.New(func() error {
		yockc.WindowsTerm(cmd).Exec(&yockc.ExecOpt{})
		return nil
	})
	if len(cron) > 0 {
		p = p.UpgradeCronTask(cron)
	}
	daemon.kernel.Scheduler.Put(p)
}
