// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package process

import (
	yockc "github.com/ansurfen/yock/cmd"
	"github.com/ansurfen/yock/ycho"
	"github.com/bwmarrin/snowflake"
)

const (
	P_NEW = iota
	P_READY
	P_SUSPEND
	P_RUNNING
	P_WAIT
	P_STOPPED
)

type pstate uint8

func (s pstate) String() string {
	switch s {
	default:
		return "unknwon state"
	}
}

type Process struct {
	pid   int64
	state pstate
	spec  string
	cmd   string
	run   func(*Process)
}

func New() *Process {
	return &Process{
		state: P_NEW,
	}
}

func (p *Process) Pid() int64 {
	return p.pid
}

func (p *Process) Cmd() string {
	return p.cmd
}

func (p *Process) Spec() string {
	return p.spec
}

func (p *Process) State() pstate {
	return p.state
}

func (p *Process) Run() {
	p.run(p)
}

type ProcessManager struct {
	node    *snowflake.Node
	process map[int64]*Process
	pid2cid map[int64]TID
	remove  chan int64
}

type TID interface {
	IsSystem() bool
	Value() int
}

type OSTID int

func (OSTID) IsSystem() bool {
	return true
}

func (id OSTID) Value() int {
	return int(id)
}

type CTID int

func (CTID) IsSystem() bool {
	return false
}

func (id CTID) Value() int {
	return int(id)
}

func NewProcessManager() *ProcessManager {
	node, err := snowflake.NewNode(2)
	if err != nil {
		panic(err)
	}
	return &ProcessManager{
		node:    node,
		process: make(map[int64]*Process),
		pid2cid: make(map[int64]TID),
		remove:  make(chan int64),
	}
}

func (m *ProcessManager) NextPID() int64 {
	return m.node.Generate().Int64()
}

func (m *ProcessManager) CreateProcess(spec, cmd string) *Process {
	p := &Process{
		spec:  spec,
		cmd:   cmd,
		pid:   m.NextPID(),
		state: P_NEW,
		run: func(p *Process) {
			// ycho.Info()
			res, err := yockc.Exec(yockc.ExecOpt{}, cmd)
			if err != nil {
				p.state = P_SUSPEND
				ycho.Errorf("[%d] process abort, err: %s", p.pid, res)
				return
			}
		},
	}
	if _, ok := m.process[p.pid]; ok {
		panic("internal system error")
	}
	m.process[p.pid] = p
	return p
}

func (m *ProcessManager) mapping(pid int64, cid TID) {
	if _, ok := m.pid2cid[pid]; ok {
		panic("%s already exist")
	}
	m.pid2cid[pid] = cid
}

func (m *ProcessManager) kill(pid int64) {
	m.remove <- pid
	if proc, ok := m.process[pid]; ok {
		proc.state = P_STOPPED
	}
}
