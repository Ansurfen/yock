// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cache

import (
	"fmt"
	"testing"
)

type cacheUint struct {
	k string
	v int
}

func (cache *cacheUint) SetKey(k string) {
	cache.k = k
}

func (cache *cacheUint) Key() string {
	return cache.k
}

func (cache *cacheUint) SetValue(v int) {
	cache.v = v
}

func (cache *cacheUint) Value() int {
	return cache.v
}

func (cache *cacheUint) Free() {
	fmt.Println(cache.k, "free")
}

func TestLRU(t *testing.T) {
	lru := NewLRU(1, func() Entry[string, int] { return &cacheUint{} })
	lru.Put("a", 10)
	lru.Put("b", 20)
	fmt.Println((*lru.Get("b")).Value())
	fmt.Println((*lru.Get("a")).Value())
}
