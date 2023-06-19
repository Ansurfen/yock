// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package process

const (
	ProcessNew = iota
	ProcessReady
	ProcessSuspend
	ProcessRunning
	ProcessWait
)

type pstate uint8

func (s pstate) String() string {
	switch s {
	default:
		return "unknwon state"
	}
}

type Process struct {
	pid     int
	state   pstate
	program func() error
}

func New(program func() error) *Process {
	return &Process{state: ProcessNew, program: program}
}

func (p *Process) UpgradeCronTask(loop string) *Process {
	return p
}

func (p *Process) Pid() int {
	return p.pid
}

func (p *Process) State() pstate {
	return p.state
}

func (p *Process) Run() {

}
