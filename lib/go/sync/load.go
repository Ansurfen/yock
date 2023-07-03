// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package synclib

import (
	"sync"

	yocki "github.com/ansurfen/yock/interface"
	atomiclib "github.com/ansurfen/yock/lib/go/sync/atomic"
)

func LoadSync(yocks yocki.YockScheduler) {
	atomiclib.LoadAtomic(yocks)
	lib := yocks.CreateLib("sync")
	lib.SetField(map[string]any{
		// functions
		"NewCond": sync.NewCond,
		// constants
		// variable
	})
	lib.SetYFunction(map[string]yocki.YGFunction{
		"new":   syncNew,
		"mutex": syncMutex,
	})
}

// @return userdata
func syncNew(l yocki.YockState) int {
	l.Pusha(&sync.WaitGroup{})
	return 1
}

// @return userdata
func syncMutex(l yocki.YockState) int {
	l.Pusha(&sync.Mutex{})
	return 1
}
