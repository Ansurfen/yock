// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

import (
	"github.com/ansurfen/yock/util"
)

type HashMap[R util.Comparable, T any] interface {
	Put(k R, v T)
	Get(k R) (T, bool)
	Del(k R)
}

var _ HashMap[int, nilType] = (*LinkedHashMap[int, nilType])(nil)

type LinkedHashMap[R util.Comparable, T any] struct {
	size  int
	list  DoubleLinkedList[T]
	index map[R]ListNode[T]
	hash  func(key R, cap int) R
}

func NewLinkedHashMap[R util.Comparable, T any](hash func(key R, cap int) R) *LinkedHashMap[R, T] {
	return &LinkedHashMap[R, T]{
		size:  0,
		index: make(map[R]ListNode[T]),
		hash:  hash,
		list:  VectorOf[T](),
	}
}

func (h *LinkedHashMap[R, T]) Put(k R, v T) {
	hash := h.hash(k, len(h.index))
	h.index[hash] = h.list.PushFront(v)
}

func (h *LinkedHashMap[R, T]) Get(k R) (T, bool) {
	hash := h.hash(k, len(h.index))
	if h.index[hash] == nil {
		var v T
		return v, false
	}
	return h.index[hash].Value(), true
}

func (h *LinkedHashMap[R, T]) Del(k R) {
	hash := h.hash(k, len(h.index))
	if v, ok := h.index[hash]; ok {
		h.list.Remove(v)
		delete(h.index, k)
	}
}

func (h *LinkedHashMap[R, T]) Iter() DoubleLinkedList[T] {
	return h.list
}
