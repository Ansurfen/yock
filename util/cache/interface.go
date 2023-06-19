// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cache

import "github.com/ansurfen/yock/util"

type Entry[T util.Comparable, R any] interface {
	SetKey(k T)
	Key() T
	SetValue(v R)
	Value() R
	Free()
}

type Cache[T util.Comparable, R any] interface {
	Get(k T) *Entry[T, R]
	Put(k T, v R) *Entry[T, R]
	Cap() int
}

type IncEntry[T util.Comparable, R any] interface {
	Entry[T, R]
	
}
