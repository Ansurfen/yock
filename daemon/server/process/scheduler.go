// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package process

import "sync"

type Scheduler struct {
	mut        sync.Mutex
	processes  []*Process
	ps         map[int]*Process
	readyQueue chan func()
}

func (s *Scheduler) Put(p *Process) {
	s.mut.Lock()
	defer s.mut.Unlock()
	s.processes = append(s.processes, p)
}

func (s *Scheduler) Run() {
	go func() {
		for {
			f := <-s.readyQueue
			go f()
		}
	}()
}
