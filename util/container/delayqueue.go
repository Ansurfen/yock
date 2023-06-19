// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

type entry[T any] struct {
	value     T
	expration time.Time
}

type DelayQueue[T any] struct {
	heap     Heap[*entry[T]]
	mut      sync.Mutex
	sleeping int32
	wakeup   chan nilType
}

func NewDelayQueue[T any]() *DelayQueue[T] {
	return &DelayQueue[T]{
		heap: NewSliceHeap(HeapMin, func(a, b *entry[T]) int {
			if a.expration.Before(b.expration) {
				return 1
			}
			return -1
		}),
	}
}

func (queue *DelayQueue[T]) Push(value T, delay time.Duration) {
	queue.mut.Lock()
	defer queue.mut.Unlock()
	entry := &entry[T]{
		value:     value,
		expration: time.Now().Add(delay),
	}
	queue.heap.Push(entry)
	if peek, err := queue.heap.Peek(); err == nil && peek == entry {
		if atomic.CompareAndSwapInt32(&queue.sleeping, 1, 0) {
			queue.wakeup <- nilType{}
		}
	}
}

func (queue *DelayQueue[T]) Take(ctx context.Context) (T, bool) {
	for {
		var timer *time.Timer
		queue.mut.Lock()
		if !queue.heap.Empty() {
			entry, _ := queue.heap.Peek()
			now := time.Now()
			if now.After(entry.expration) {
				queue.heap.Pop()
				queue.mut.Unlock()
				return entry.value, true
			}
			timer = time.NewTimer(entry.expration.Sub(now))
		}
		atomic.StoreInt32(&queue.sleeping, 1)
		queue.mut.Unlock()
		if timer != nil {
			select {
			case <-queue.wakeup:
				timer.Stop()
			case <-timer.C:
				if atomic.SwapInt32(&queue.sleeping, 0) == 0 {
					<-queue.wakeup
				}
			case <-ctx.Done():
				timer.Stop()
				var t T
				return t, false
			}
		} else {
			select {
			case <-queue.wakeup:
			case <-ctx.Done():
				var t T
				return t, false
			}
		}
	}
}

func (queue *DelayQueue[T]) Channel(ctx context.Context, size int) <-chan T {
	out := make(chan T, size)
	go func() {
		for {
			entry, ok := queue.Take(ctx)
			if !ok {
				close(out)
				return
			}
			out <- entry
		}
	}()
	return out
}
