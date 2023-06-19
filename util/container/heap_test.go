// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

import (
	"fmt"
	"testing"
)

func upCompare(a, b int) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

func TestMaxHeap(t *testing.T) {
	h := NewSliceHeap(HeapMax, upCompare)
	h.Push(0, 3, 1, 2)
	fmt.Println(h.data)
	h.Pop()
	fmt.Println(h.data)
	h.Pop()
	fmt.Println(h.data)
}

func TestMinHeap(t *testing.T) {
	h := NewSliceHeap(HeapMin, upCompare)
	h.Push(0, 3, 1, 2)
	fmt.Println(h.data)
	h.Pop()
	fmt.Println(h.data)
	h.Pop()
	fmt.Println(h.data)
}
