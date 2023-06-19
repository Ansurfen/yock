// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

import (
	"container/list"
	"errors"
)

type Stack[T any] interface {
	Clear()
	Empty() bool
	Full() bool
	Iter() []T
	Len() int
	Cap() int
	Pop() error
	Push(v T) error
	Top() (T, error)
}

var (
	_ Stack[nilType] = (*ArrayStack[nilType])(nil)
)

type ListStack[T any] struct {
	data *list.List
}

func NewListStack[T any]() *ListStack[T] {
	return &ListStack[T]{data: list.New()}
}

type ArrayStack[T any] struct {
	data []T
	top  int
	cap  int
}

func NewStack[T any](cap int) *ArrayStack[T] {
	return &ArrayStack[T]{
		data: make([]T, cap),
		top:  -1,
		cap:  cap,
	}
}

func (s *ArrayStack[T]) Cap() int {
	return s.cap
}

func (s *ArrayStack[T]) Iter() []T {
	return s.data
}

func (s *ArrayStack[T]) Clear() {
	s.data = make([]T, 0)
	s.top = -1
	s.cap = 1
}

func (s *ArrayStack[T]) Full() bool {
	return s.top >= s.cap
}

func (s *ArrayStack[T]) Empty() bool {
	return s.top == -1
}

func (s *ArrayStack[T]) Pop() error {
	if s.Empty() {
		return errors.New("index out of range")
	}
	s.top--
	if s.top < -1 {
		s.top = -1
	}
	return nil
}

func (s *ArrayStack[T]) Push(v T) error {
	if s.Full() {
		return errors.New("index out of range")
	}
	s.top++
	s.data[s.top] = v
	return nil
}

func (s *ArrayStack[T]) Top() (T, error) {
	if s.Empty() {
		return *new(T), errors.New("index out of range")
	}
	return s.data[s.top], nil
}

func (s *ArrayStack[T]) Len() int {
	return s.top + 1
}
