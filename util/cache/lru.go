// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/util/container"
)

var (
	_ Cache[int, struct{}] = (*LRU[int, struct{}])(nil)
)

type LRU[T util.Comparable, R any] struct {
	cap       int
	size      int
	find      *container.LinkedHashMap[T, Entry[T, R]]
	newEntity func() Entry[T, R]
}

func NewLRU[T util.Comparable, R any](cap int, newEntity func() Entry[T, R]) *LRU[T, R] {
	return &LRU[T, R]{
		cap:  cap,
		size: 0,
		find: container.NewLinkedHashMap[T, Entry[T, R]](func(key T, cap int) T {
			return key
		}),
		newEntity: newEntity,
	}
}

func (lru *LRU[T, R]) Get(k T) *Entry[T, R] {
	iter := lru.find.Iter()
	if v, ok := lru.find.Get(k); ok {
		lru.find.Del(k)
		lru.find.Put(k, v)
		return &v
	} else {
		back := iter.Back()
		if back == nil {
			lru.size++
		}
		defer back.Value().Free()
		lru.find.Del(back.Value().Key())
		entity := lru.newEntity()
		entity.SetKey(k)
		lru.find.Put(k, entity)
		return &entity
	}
}

func (lru *LRU[T, R]) Put(k T, v R) *Entry[T, R] {
	if vv, ok := lru.find.Get(k); ok {
		lru.find.Del(k)
		lru.find.Put(k, vv)
		return &vv
	} else {
		if lru.cap == lru.size {
			back := lru.find.Iter().Back()
			defer back.Value().Free()
			lru.find.Del(back.Value().Key())
		}
		entity := lru.newEntity()
		entity.SetKey(k)
		entity.SetValue(v)
		lru.find.Put(k, entity)
		lru.size++
		return &entity
	}
}

func (lru *LRU[T, R]) Cap() int {
	return lru.cap
}
