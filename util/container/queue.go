// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

import (
	"container/list"
	"container/ring"
	"errors"
)

type Deque[T any] interface {
	Queue[T]
	Back() (T, bool)
	EnqueueFront(v T) error
	DequeueBack() error
}

type Queue[T any] interface {
	Cap() int
	Enqueue(v T) error
	Dequeue() error
	Front() (T, bool)
	Full() bool
	Empty() bool
}

var (
	_ Queue[nilType] = (*ListQueue[nilType])(nil)
	_ Queue[nilType] = (*ListRing[nilType])(nil)
	_ Queue[nilType] = (*ArrayRing[nilType])(nil)

	_ Deque[nilType] = (*ListDeque[nilType])(nil)
)

type ListRing[T any] struct {
	cap   int
	next  *ring.Ring
	front *ring.Ring
	len   int
}

func NewListRing[T any](cap int) *ListRing[T] {
	r := ring.New(cap)
	if r == nil {
		panic("invalid cap")
	}
	return &ListRing[T]{cap: cap, next: r, front: r, len: 0}
}

func (queue *ListRing[T]) Cap() int {
	return queue.cap
}

func (queue *ListRing[T]) Enqueue(v T) error {
	queue.next.Value = v
	queue.next = queue.next.Next()
	queue.len++
	return nil
}

func (queue *ListRing[T]) Dequeue() error {
	queue.front.Value = nil
	queue.front = queue.front.Next()
	queue.len--
	return nil
}

func (queue *ListRing[T]) Front() (T, bool) {
	if queue.front.Value == nil {
		var v T
		return v, false
	}
	return queue.front.Value.(T), true
}

func (queue *ListRing[T]) Empty() bool {
	return queue.len == 0
}

func (queue *ListRing[T]) Full() bool {
	return queue.len == queue.cap
}

type ArrayRing[T any] struct {
	data  []T
	cap   int
	front int
	rear  int
	len   int
}

func NewArrayRing[T any](cap int) Queue[T] {
	return &ArrayRing[T]{cap: cap, len: 0, data: make([]T, cap)}
}

func (queue *ArrayRing[T]) Enqueue(v T) error {
	if queue.Full() {
		return errors.New("queue is full")
	}
	queue.data[queue.rear] = v
	queue.rear = (queue.rear + 1) % queue.cap
	queue.len++
	return nil
}

func (queue *ArrayRing[T]) Dequeue() error {
	if queue.Empty() {
		return errors.New("index out of range")
	}
	queue.front = (queue.front + 1) % queue.cap
	queue.len--
	return nil
}

func (queue *ArrayRing[T]) Front() (T, bool) {
	if queue.Empty() {
		var v T
		return v, false
	}
	return queue.data[queue.front], true
}

func (queue *ArrayRing[T]) Empty() bool {
	return queue.len == 0
}

func (queue *ArrayRing[T]) Full() bool {
	return queue.len == queue.cap
}

func (queue *ArrayRing[T]) Cap() int {
	return queue.cap
}

type ListQueue[T any] struct {
	cap  int
	data *list.List
}

func NewListQueue[T any](cap ...int) Queue[T] {
	c := -1
	if len(cap) > 0 {
		c = cap[0]
	}
	return &ListQueue[T]{cap: c, data: list.New()}
}

func (queue *ListQueue[T]) Enqueue(v T) error {
	if queue.cap > -1 && queue.Full() {
		return errors.New("index out of range")
	}
	queue.data.PushBack(v)
	return nil
}

func (queue *ListQueue[T]) Dequeue() error {
	if queue.Empty() {
		return errors.New("queue is empty")
	}
	queue.data.Remove(queue.data.Front())
	return nil
}

func (queue *ListQueue[T]) Front() (T, bool) {
	if queue.Empty() {
		var v T
		return v, false
	}
	return queue.data.Front().Value.(T), true
}

func (queue *ListQueue[T]) Empty() bool {
	return queue.data.Len() == 0
}

func (queue *ListQueue[T]) Full() bool {
	return queue.data.Len() == queue.cap
}

func (queue *ListQueue[T]) Cap() int {
	return queue.cap
}

type ListDeque[T any] struct {
	*ListQueue[T]
}

func NewListDeque[T any](cap ...int) Deque[T] {
	return &ListDeque[T]{NewListQueue[T](cap...).(*ListQueue[T])}
}

func (queue *ListDeque[T]) Back() (T, bool) {
	if queue.Empty() {
		var v T
		return v, false
	}
	return queue.data.Back().Value.(T), true
}

func (queue *ListDeque[T]) EnqueueFront(v T) error {
	if queue.Full() {
		return errors.New("index out of range")
	}
	queue.data.PushFront(v)
	return nil
}

func (queue *ListDeque[T]) DequeueBack() error {
	if queue.Empty() {
		return errors.New("queue is empty")
	}
	queue.data.Remove(queue.data.Back())
	return nil
}
