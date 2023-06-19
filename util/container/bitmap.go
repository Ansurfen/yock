// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

type Bitmap struct {
	data []byte
}

func NewBitmap(cap int, expand ...bool) *Bitmap {
	return &Bitmap{data: make([]byte, cap)}
}

func (bm *Bitmap) Set(n int) {
	if bm.overflow(n) {
		return
	}
	bm.data[n/8] |= 1 << (n % 8)
}

func (bm *Bitmap) Unset(n int) {
	if bm.overflow(n) {
		return
	}
	bm.data[n/8] &^= 1 << (n % 8)
}

func (bm *Bitmap) Chcek(n int) bool {
	if bm.overflow(n) {
		return false
	}
	return bm.data[n/8]&(1<<(n%8)) != 0
}

func (bm *Bitmap) overflow(n int) bool {
	return n/8 >= len(bm.data)
}
