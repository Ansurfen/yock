// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

import (
	"hash/crc32"
	"math/rand"
	"time"
)

type BloomFilter struct {
	bm    *Bitmap
	round int
	seeds []int
	size  int
}

func NewBloomFilter(size, round int) *BloomFilter {
	filter := &BloomFilter{
		seeds: randomSeed(size),
		round: round,
		bm:    NewBitmap(size),
		size:  size,
	}
	return filter
}

func (filter *BloomFilter) Set(v string) {
	for _, raw := range filter.hash([]byte(v)) {
		filter.bm.Set(raw)
	}
}

func (filter *BloomFilter) Check(v string) bool {
	for _, raw := range filter.hash([]byte(v)) {
		if !filter.bm.Chcek(raw) {
			return false
		}
	}
	return true
}

func (filter *BloomFilter) hash(v []byte) []int {
	idxs := make([]int, filter.round)
	for i := 0; i < filter.round; i++ {
		idxs[i] = (int(crc32.ChecksumIEEE(v)) + filter.seeds[i]) % filter.size
	}
	return idxs
}

func randomSeed(size int) []int {
	seeds := make([]int, size)
	rand.Seed(time.Now().Unix() / 1000)
	for i := 0; i < size; i++ {
		seeds[i] = rand.Intn(size * 8)
	}
	return seeds
}
