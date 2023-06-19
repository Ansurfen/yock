// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

import (
	"fmt"
	"testing"
)

// todo, improve testset
const (
	PUSH = iota
	POP
	TOP
)

func TestStack(t *testing.T) {
	s := NewStack[int](10)
	s.Push(10)
	s.Push(20)
	fmt.Println(s.Top())
	s.Pop()
	fmt.Println(s.Top())
	s.Pop()
	s.Pop()
	fmt.Println(s.Top())
}
