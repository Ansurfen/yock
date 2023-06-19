// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewListQueue[int](3)
	testset := []int{1, 2, 3}
	for _, v := range testset {
		if err := q.Enqueue(v); err != nil {
			fmt.Println(err)
		}
	}
	for i := 0; i < len(testset); i++ {
		if v, ok := q.Front(); ok {
			fmt.Println(v)
			q.Dequeue()
		}
	}
}

func TestRing(t *testing.T) {
	q := NewArrayRing[int](3)
	testset := []int{1, 2, 3}
	for _, v := range testset {
		if err := q.Enqueue(v); err != nil {
			fmt.Println(err)
		}
	}
	for i := 0; i < len(testset); i++ {
		if v, ok := q.Front(); ok {
			fmt.Println(v)
			q.Dequeue()
		}
	}
	fmt.Println(q.Empty())
}

func TestDeque(t *testing.T) {
	q := NewListDeque[int](10)
	testset := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range testset {
		if err := q.Enqueue(v); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(q.Back())
	q.Dequeue()
	fmt.Println(q.Front())
	q.DequeueBack()
	fmt.Println(q.Back())
	q.EnqueueFront(10)
	fmt.Println(q.Front())
}

func TestListRing(t *testing.T) {
	q := NewListRing[int](3)
	testset := []int{1, 2, 3}
	for _, v := range testset {
		if err := q.Enqueue(v); err != nil {
			fmt.Println(err)
		}
	}
	for i := 0; i < len(testset); i++ {
		if v, ok := q.Front(); ok {
			fmt.Println(v)
			q.Dequeue()
		}
	}
	fmt.Println(q.Empty())
}
