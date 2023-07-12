// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

type DoubleLinkedList[T any] interface {
	PushFront(v T) ListNode[T]
	PushBack(v T) ListNode[T]
	PopFront() ListNode[T]
	PopBack() ListNode[T]
	Find(n int) ListNode[T]
	Remove(node ListNode[T])
	Front() ListNode[T]
	Back() ListNode[T]
	Size() int
}

type ListNode[T any] interface {
	Prev() ListNode[T]
	Next() ListNode[T]
	Value() T
}

var (
	_ ListNode[nilType]         = (*VectorNode[nilType])(nil)
	_ DoubleLinkedList[nilType] = (*Vector[nilType])(nil)
)

type VectorNode[T any] struct {
	data T
	prev *VectorNode[T]
	next *VectorNode[T]
}

func newVectorNode[T any](v T) *VectorNode[T] {
	return &VectorNode[T]{data: v, prev: nil, next: nil}
}

func (node *VectorNode[T]) Value() T {
	if node == nil {
		var v T
		return v
	}
	return node.data
}

func (node *VectorNode[T]) Prev() ListNode[T] {
	if node == nil {
		return nil
	}
	return node.prev
}

func (node *VectorNode[T]) Next() ListNode[T] {
	if node == nil {
		return nil
	}
	return node.next
}

type Vector[T any] struct {
	cap  int
	size int
	head *VectorNode[T]
	tail *VectorNode[T]
}

func VectorOf[T any](cap ...int) *Vector[T] {
	c := -1
	if len(cap) > 0 {
		c = cap[0]
	}
	return &Vector[T]{cap: c, size: 0, head: nil, tail: nil}
}

func (vec *Vector[T]) PushFront(v T) ListNode[T] {
	if vec.isEnd() {
		return nil
	}
	node := newVectorNode(v)
	if vec.head == nil {
		vec.head = node
		vec.tail = node
		vec.head.prev = nil
		vec.tail.next = nil
	} else {
		vec.head.prev = node
		node.next = vec.head
		vec.head = node
		vec.head.prev = nil
	}
	vec.size++
	return node
}

func (vec *Vector[T]) PushBack(v T) ListNode[T] {
	if vec.isEnd() {
		return nil
	}
	node := newVectorNode(v)
	if vec.tail == nil {
		vec.tail = node
		vec.head = node
		vec.head.prev = nil
		vec.tail.next = nil
	} else {
		vec.tail.next = node
		node.prev = vec.tail
		vec.tail = node
		vec.tail.next = nil
	}
	vec.size++
	return node
}

func (vec *Vector[T]) PopFront() ListNode[T] {
	if vec.head == nil {
		return nil
	}
	node := vec.head
	if node.next != nil {
		vec.head = node.next
		vec.head.prev = nil
	} else {
		vec.head = nil
		vec.tail = nil
	}
	vec.size--
	return node
}

func (vec *Vector[T]) PopBack() ListNode[T] {
	if vec.tail == nil {
		return nil
	}
	node := vec.tail
	if node.prev != nil {
		vec.tail = node.prev
		vec.tail.next = nil
	} else {
		vec.head = nil
		vec.tail = nil
	}
	vec.size--
	return node
}

func (vec *Vector[T]) Remove(node ListNode[T]) {
	n := node.(*VectorNode[T])
	if n == nil {
		return
	}
	if n == vec.head {
		vec.PopFront()
	} else if n == vec.tail {
		vec.PopBack()
	} else {
		n.next.prev = n.prev
		n.prev.next = n.next
	}
	vec.size--
}

func (vec *Vector[T]) Find(n int) ListNode[T] {
	if n >= vec.size {
		return nil
	}
	var p ListNode[T] = vec.head
	for i := 0; p != nil && i < n; i++ {
		p = p.Next()
	}
	return p
}

func (vec *Vector[T]) Front() ListNode[T] {
	if vec.head == nil {
		return nil
	}
	return vec.head
}

func (vec *Vector[T]) Back() ListNode[T] {
	if vec.tail == nil {
		return nil
	}
	return vec.tail
}

func (vec *Vector[T]) isEnd() bool {
	if vec.cap < 0 {
		return false
	}
	return vec.size >= vec.cap
}

func (vec *Vector[T]) Size() int {
	return vec.size
}
