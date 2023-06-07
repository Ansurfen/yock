// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import "errors"

type Stack[T any] struct {
	data []T
	top  int
	cap  int
}

func NewStack[T any](cap int) *Stack[T] {
	return &Stack[T]{
		data: make([]T, cap),
		top:  -1,
		cap:  cap,
	}
}

func (s *Stack[T]) Iter() []T {
	return s.data
}

func (s *Stack[T]) Clear() {
	s.data = make([]T, 0)
	s.top = -1
	s.cap = 1
}

func (s *Stack[T]) Empty() bool {
	return s.top == -1
}

func (s *Stack[T]) Pop() error {
	if s.Empty() {
		return errors.New("index out of range")
	}
	s.top--
	if s.top < -1 {
		s.top = -1
	}
	return nil
}

func (s *Stack[T]) Push(v T) error {
	if s.top >= s.cap {
		return errors.New("index out of range")
	}
	s.top++
	s.data[s.top] = v
	return nil
}

func (s *Stack[T]) Top() (T, error) {
	if s.Empty() {
		return *new(T), errors.New("index out of range")
	}
	return s.data[s.top], nil
}

func (s *Stack[T]) Len() int {
	return s.top + 1
}
