// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

import (
	"errors"
)

type Heap[T any] interface {
	Fix()
	Len() int
	Peek() (T, error)
	Pop()
	Push(vals ...T)
	Empty() bool
}

var _ Heap[nilType] = (*SliceHeap[nilType])(nil)

const (
	HeapMax = true
	HeapMin = false
)

type SliceHeap[T any] struct {
	data     []T
	heapType bool
	compare  func(a, b T) int
}

func NewSliceHeap[T any](heapType bool, cmp func(a, b T) int) *SliceHeap[T] {
	return &SliceHeap[T]{heapType: heapType, compare: cmp}
}

func (heap *SliceHeap[T]) Fix() {
	n := heap.Len()
	if n == 0 {
		return
	}
	for i := n / 2; i >= 0; i-- {
		heap.sink(i, n)
	}
}

func (heap *SliceHeap[T]) hasChild(cur int) bool {
	// assume conditions: [0], len = 1 and [0, 0] len = 2
	return cur*2+1 <= len(heap.data)-1
}

func (heap *SliceHeap[T]) sink(cur, length int) {
	cmpVal := 1
	if !heap.heapType { // HeapMin
		cmpVal = -1
	}
	temp := heap.data[cur]
	for heap.hasChild(cur) {
		child := cur*2 + 1
		if child+1 < length && heap.compare(heap.data[child+1], heap.data[child]) == cmpVal {
			child++
		}
		if v := heap.compare(temp, heap.data[child]); v == 0 || v == cmpVal {
			break
		}
		heap.data[cur] = heap.data[child]
		cur = child
	}
	heap.data[cur] = temp
}

func (heap *SliceHeap[T]) float(cur int) {
	cmpVal := 1
	if !heap.heapType { // HeapMin
		cmpVal = -1
	}
	temp := heap.data[cur]
	parent := (cur - 1) / 2
	// assume conditions: [0, 0], cur = 1, [0, 0, 0], cur = 2
	for parent > 0 && heap.compare(temp, heap.data[parent]) == cmpVal {
		heap.data[cur] = heap.data[parent]
		cur = parent
		parent = (cur - 1) / 2
	}
	heap.data[cur] = temp
}

func (heap *SliceHeap[T]) Push(vals ...T) {
	n := heap.Len()
	heap.data = append(heap.data, vals...)
	if n == 0 {
		heap.Fix()
		return
	}
	for i := 0; i < len(vals); i++ {
		heap.float(n + i)
	}
}

func (heap *SliceHeap[T]) Pop() {
	if heap.Empty() {
		return
	}
	n := heap.Len() - 1
	heap.data[0], heap.data[n] = heap.data[n], heap.data[0]
	heap.data = heap.data[0:n]
	if n == 0 {
		return
	}
	heap.sink(0, n)
}

func (heap *SliceHeap[T]) Peek() (T, error) {
	if heap.Empty() {
		var v T
		return v, errors.New("index out of range")
	}
	return heap.data[0], nil
}

func (heap *SliceHeap[T]) Len() int {
	return len(heap.data)
}

func (heap *SliceHeap[T]) Empty() bool {
	return len(heap.data) == 0
}
