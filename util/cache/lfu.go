// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/util/container"
)

type LFU[T util.Comparable, R any] struct {
	cap       int
	newEntity func() Entry[T, R]
	find      container.Heap[Entry[T, R]]
}

func NewLFU[T util.Comparable, R any](cap int, newEntity func() Entry[T, R]) *LFU[T, R] {
	return &LFU[T, R]{
		cap:       cap,
		newEntity: newEntity,
		find: container.NewSliceHeap(container.HeapMin, func(a, b Entry[T, R]) int {
			return 1
		}),
	}
}

// func (lfu *LFU[T, R]) Get(k T) *Entry[T, R] {

// }

// func (lfu *LFU[T, R]) Put(k T, v R) *Entry[T, R] {
// 	if lfu.cap == lfu.find.Len() {
// 		lfu.find.Pop()
// 		lfu.newEntity()
// 	}
// }

func (lfu *LFU[T, R]) Cap() int {
	return lfu.cap
}
