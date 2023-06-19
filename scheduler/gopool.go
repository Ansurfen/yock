// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"time"

	yocki "github.com/ansurfen/yock/interface"
)

var _ yocki.GoPool = (*ChannelPool)(nil)

type ChannelPool struct {
	goroutines chan func()
}

func newChannelPool(cap int) *ChannelPool {
	return &ChannelPool{goroutines: make(chan func(), cap)}
}

func (pool *ChannelPool) Go(f func()) {
	pool.goroutines <- f
}

func (pool *ChannelPool) Run() {
	sleeping := time.Microsecond * 500
	for {
		timer := time.NewTimer(sleeping)
		select {
		case f := <-pool.goroutines:
			f()
		default:
			// time.Sleep(1 * time.Second)
			<-timer.C
			timer.Reset(sleeping)
		}
	}
}
