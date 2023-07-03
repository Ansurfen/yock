// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package atomiclib

import (
	"sync/atomic"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadAtomic(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("atomic")
	lib.SetField(map[string]any{
		// functions
		"LoadUint32":            atomic.LoadUint32,
		"CompareAndSwapUint64":  atomic.CompareAndSwapUint64,
		"CompareAndSwapInt64":   atomic.CompareAndSwapInt64,
		"CompareAndSwapUintptr": atomic.CompareAndSwapUintptr,
		"StoreUintptr":          atomic.StoreUintptr,
		"SwapUint64":            atomic.SwapUint64,
		"CompareAndSwapInt32":   atomic.CompareAndSwapInt32,
		"CompareAndSwapPointer": atomic.CompareAndSwapPointer,
		"SwapPointer":           atomic.SwapPointer,
		"LoadInt64":             atomic.LoadInt64,
		"StorePointer":          atomic.StorePointer,
		"LoadPointer":           atomic.LoadPointer,
		"SwapInt64":             atomic.SwapInt64,
		"LoadUintptr":           atomic.LoadUintptr,
		"SwapInt32":             atomic.SwapInt32,
		"StoreUint32":           atomic.StoreUint32,
		"AddUintptr":            atomic.AddUintptr,
		"StoreUint64":           atomic.StoreUint64,
		"AddUint64":             atomic.AddUint64,
		"StoreInt64":            atomic.StoreInt64,
		"SwapUintptr":           atomic.SwapUintptr,
		"LoadInt32":             atomic.LoadInt32,
		"LoadUint64":            atomic.LoadUint64,
		"AddInt32":              atomic.AddInt32,
		"SwapUint32":            atomic.SwapUint32,
		"AddUint32":             atomic.AddUint32,
		"AddInt64":              atomic.AddInt64,
		"StoreInt32":            atomic.StoreInt32,
		"CompareAndSwapUint32":  atomic.CompareAndSwapUint32,
		// constants
		// variable
	})
}
