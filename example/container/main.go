// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/util/container"
)

type SafeQueue[T any] struct {
	queue []T
	cap   int
	head  int64
	tail  int64
}

func NewSafeQueue[T any](cap int) *SafeQueue[T] {
	return &SafeQueue[T]{
		queue: make([]T, cap),
		cap:   cap,
		head:  0,
		tail:  0,
	}
}

func (q *SafeQueue[T]) Cap() int {
	return q.cap
}

func (q *SafeQueue[T]) Enqueue(v T) error {
	next := (atomic.LoadInt64(&q.tail) + 1) % int64(q.cap)
	if atomic.LoadInt64(&q.head) == next {
		return util.ErrOutRange
	}
	q.queue[q.tail] = v
	atomic.StoreInt64(&q.tail, next)
	return nil
}

func (q *SafeQueue[T]) Dequeue() error {
	if q.Empty() {
		return errors.New("queue is empty")
	}
	atomic.StoreInt64(&q.head, (atomic.LoadInt64(&q.head)+1)%int64(q.cap))
	return nil
}

func (q *SafeQueue[T]) Front() (T, bool) {
	if q.Empty() {
		var v T
		return v, false
	}
	return q.queue[q.head], true
}

func (q *SafeQueue[T]) Full() bool {
	return atomic.LoadInt64(&q.head) == (atomic.LoadInt64(&q.tail)+1)%int64(q.cap)
}

func (q *SafeQueue[T]) Empty() bool {
	return atomic.LoadInt64(&q.head) == atomic.LoadInt64(&q.tail)
}

var _ container.Queue[struct{}] = (*SafeQueue[struct{}])(nil)

func main() {
	var arr container.Queue[int] = NewSafeQueue[int](5)
	arr.Enqueue(1)
	arr.Enqueue(2)
	arr.Enqueue(3)
	fmt.Println(arr.Front())
	arr.Dequeue()
	fmt.Println(arr.Front())
	arr.Dequeue()
	fmt.Println(arr.Front())
}
