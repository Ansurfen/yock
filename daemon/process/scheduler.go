// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package process

import (
	"strings"
	"time"

	"github.com/ansurfen/yock/ycho"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	prom        *ProcessManager
	timingwheel *cron.Cron
	oschan      *OSNotify
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		prom:        NewProcessManager(),
		timingwheel: cron.New(),
		oschan:      NewOSNotify(),
	}
}

func (s *Scheduler) FindByCmd(cmd string) (ret []*Process) {
	for _, p := range s.prom.process {
		if strings.Contains(p.Cmd(), cmd) {
			ret = append(ret, p)
		}
	}
	return
}

func (s *Scheduler) FindByPID(pid int64) *Process {
	return s.prom.process[pid]
}

func (s *Scheduler) Kill(pid int64) {
	s.prom.kill(pid)
}

func (s *Scheduler) ProcessState() map[int64]*Process {
	return s.prom.process
}

func (s *Scheduler) Run() {
	go func() {
		for {
			select {
			case pid := <-s.prom.remove:
				if id, ok := s.prom.pid2cid[pid]; ok {
					if !id.IsSystem() {
						s.timingwheel.Remove(cron.EntryID(id.Value()))
					} else {
						s.oschan.Remove(id.Value())
					}
				}
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()
	go s.timingwheel.Start()
	s.oschan.Listen()
}

func (s *Scheduler) CreateCronTask(cron, cmd string) (pid int64, err error) {
	defer func() {
		err := recover()
		switch v := err.(type) {
		case error:
			ycho.Error(v)
			err = v
			pid = -1
		case string:
			ycho.Errorf(v)
			err = v
			pid = -1
		}
	}()
	p := s.prom.CreateProcess(cron, cmd)
	id, err := s.timingwheel.AddFunc(cron, func() {
		if p.state == P_STOPPED {
			return
		}
		p.state = P_RUNNING
		p.run(p)
		p.state = P_READY
	})
	if err != nil {
		panic(err)
	}
	s.prom.mapping(p.pid, CTID(id))
	return p.pid, nil
}

func (s *Scheduler) CreateImmediateCronTask(cron, cmd string) (pid int64, err error) {
	defer func() {
		err := recover()
		switch v := err.(type) {
		case error:
			ycho.Error(v)
			err = v
			pid = -1
		case string:
			ycho.Errorf(v)
			err = v
			pid = -1
		}
	}()
	p := s.prom.CreateProcess(cron, cmd)
	id, err := s.timingwheel.AddFunc(cron, func() {
		if p.state == P_STOPPED {
			return
		}
		p.state = P_RUNNING
		p.run(p)
		s.prom.kill(p.pid)
	})
	if err != nil {
		panic(err)
	}
	s.prom.mapping(p.pid, CTID(id))
	pid = p.pid
	return
}

func (s *Scheduler) CreateTimingImmediateCronTask(cron, cmd string) {}

func (s *Scheduler) CreateFSListenTask(paths []string, cmd string) (int64, error) {
	p := s.prom.CreateProcess(strings.Join(paths, ";"), cmd)
	id := s.oschan.AddFunc(paths, func() {
		if p.state == P_STOPPED {
			return
		}
		p.state = P_RUNNING
		p.run(p)
		p.state = P_WAIT
	})
	s.prom.mapping(p.pid, OSTID(id))
	return p.pid, nil
}
